package task

import (
	"fmt"
	"github.com/JiaoDean/crd-controller/pkg/controller/crd"
	"github.com/JiaoDean/crd-controller/pkg/controller/kube"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"time"
)

var handlers = make(map[string]handler)
var syncPeriod = make(map[string]time.Duration)

type handler func(kubeController *kube.KubeController, crdController *crd.CrdController, kubeClient kubernetes.Interface) (Task, error)

func register(name string, handler handler, duration time.Duration) {
	if handler == nil {
		panic("Register job is failed, job is nil poniter.")
	}
	if _, dup := handlers[name]; dup {
		msg := fmt.Sprintf("Register job %s is failed, err: register a duplicate job", name)
		panic(msg)
	}
	handlers[name] = handler
	syncPeriod[name] = duration
}

type Task interface {
	Run()
	doTask(key, value interface{}) bool
}

func Run(csiController *kube.KubeController, crdController *crd.CrdController, kubeClient kubernetes.Interface, stopCh <-chan struct{}) {
	for name, jobHandler := range handlers {
		job, err := jobHandler(csiController, crdController, kubeClient)
		if err != nil {
			log.Errorf("Handler job %s is failed, err: %s", name, err.Error())
		}
		go wait.Until(job.Run, syncPeriod[name], stopCh)
	}
	<-stopCh
}
