package crd

import (
	"github.com/JiaoDean/crd-controller/pkg/apis/crd/v1beta1"
	crdClient "github.com/JiaoDean/crd-controller/pkg/client/clientset/versioned"
	externalversions "github.com/JiaoDean/crd-controller/pkg/client/informers/externalversions/crd/v1beta1"
	listers "github.com/JiaoDean/crd-controller/pkg/client/listers/crd/v1beta1"
	"github.com/JiaoDean/crd-controller/pkg/fsm"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

const CrdProtectionFinalizers = "kubernetes.io/crd-protection"

/*var (
	initialSet = hashset.New(fsm.Pending, fsm.Creating, fsm.Initialization)
	workSet    = hashset.New(fsm.Running)
	fatalSet   = hashset.New(fsm.Terminating)
)*/

type CrdController struct {
	CrdClient  *crdClient.Clientset
	KubeClient *kubernetes.Clientset
	vspSynced  cache.InformerSynced

	crdLister    listers.CrdLister
	synced       cache.InformerSynced
	crdStatusMap map[string]statusDesc
	crdMap       map[string]*v1beta1.Crd
}

func NewCrdController(crdClient *crdClient.Clientset, kubeClient *kubernetes.Clientset, crdInformer externalversions.CrdInformer) *CrdController {
	c := CrdController{
		CrdClient:  crdClient,
		KubeClient: kubeClient,
		crdLister:  crdInformer.Lister(),
		synced:     crdInformer.Informer().HasSynced,
	}

	crdInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.handleCrdAdd,
		DeleteFunc: c.handleCrdDelete,
		UpdateFunc: c.handleCrdUpdate,
	})
	return &c
}

func (c *CrdController) handleCrdAdd(obj interface{}) {
	crdObj, ok := obj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer crd %+v is failed", crdObj)
		return
	}
	c.init(crdObj)
}

func (c *CrdController) handleCrdDelete(obj interface{}) {
	crdObj, ok := obj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer crd %+v is failed", crdObj)
		return
	}
}

func (c *CrdController) handleCrdUpdate(oldObj interface{}, newObj interface{}) {
	newCrdObj, ok := newObj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer new crd %+v is failed", newCrdObj)
		return
	}
	oldCrdObj, ok := oldObj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer old crd %+v is failed", oldCrdObj)
		return
	}
}

func (c *CrdController) init(crdObj *v1beta1.Crd) {
	fsm.NewFSM()
}
