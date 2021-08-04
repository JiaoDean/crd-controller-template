package main

import (
	"flag"
	"github.com/JiaoDean/crd-controller/pkg/client/clientset/versioned"
	"github.com/JiaoDean/crd-controller/pkg/client/informers/externalversions"
	"github.com/JiaoDean/crd-controller/pkg/controller"
	"github.com/JiaoDean/crd-controller/pkg/task"
	"github.com/JiaoDean/crd-controller/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	logfilePrefix = "/var/log/crd-controller/"
	mbsize        = 1024 * 1024
)

var onlyOneSignalHandler = make(chan struct{})
var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

var (
	logLevel = flag.String("log-level", "Info", "Set Log Level")
)

func startHealthzServer(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Infof("start healthz server and listen on %s", addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("healthz server cause a error: %+v", err)
	}
}

func main() {
	flag.Parse()
	// set log config
	setLogAttribute("crd.controller.log")
	stopCh := setupSignalHandler()

	var masterURL, kubeconfig string
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Get KubeConfig is failed from master, err: %s", err.Error())
	}
	// new crdClient and kubeClient
	crdClient := versioned.NewForConfigOrDie(cfg)
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Create KubeClient is failed, err: %s", err.Error())
	}
	kubeInformerFactory := informers.NewSharedInformerFactory(kubeClient, time.Second*30)
	crdInformerFactory := externalversions.NewSharedInformerFactory(crdClient, time.Second*30)

	recorder := utils.NewEventRecorder()

	kubeController := controller.NewKubeController(kubeClient,
		kubeInformerFactory.Core().V1().PersistentVolumes(),
		kubeInformerFactory.Core().V1().PersistentVolumeClaims(),
		recorder)

	crdController := controller.NewCrdController(
		crdClient, kubeClient,
		crdInformerFactory.Storage().V1beta1().Crds())

	kubeInformerFactory.Start(stopCh)
	crdInformerFactory.Start(stopCh)
	kubeInformerFactory.WaitForCacheSync(stopCh)
	crdInformerFactory.WaitForCacheSync(stopCh)

	time.Sleep(10 * time.Second)

	go startHealthzServer(":10254")

	task.Run(kubeController, crdController, kubeClient, stopCh)
}

// rotate log file by 2M bytes
// default print log to stdout and file both.
func setLogAttribute(logName string) {
	logType := os.Getenv("LOG_TYPE")
	logType = strings.ToLower(logType)
	if logType != "stdout" && logType != "host" {
		logType = "both"
	}
	if logType == "stdout" {
		return
	}

	os.MkdirAll(logfilePrefix, os.FileMode(0755))
	logFile := logfilePrefix + logName + ".log"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(1)
	}

	// rotate the log file if too large
	if fi, err := f.Stat(); err == nil && fi.Size() > 2*mbsize {
		f.Close()
		timeStr := time.Now().Format("-2006-01-02-15:04:05")
		timedLogfile := logfilePrefix + logName + timeStr + ".log"
		os.Rename(logFile, timedLogfile)
		f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			os.Exit(1)
		}
	}
	if logType == "both" {
		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
	} else {
		log.SetOutput(f)
	}

	logLevelLow := strings.ToLower(*logLevel)
	if logLevelLow == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if logLevelLow == "warning" {
		log.SetLevel(log.WarnLevel)
	}
	log.Infof("Set Log level to %s...", logLevelLow)
}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func setupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
