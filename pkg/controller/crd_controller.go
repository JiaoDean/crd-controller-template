package controller

import (
	crdClient "github.com/JiaoDean/crd-controller/pkg/client/clientset/versioned"
	externalversions "github.com/JiaoDean/crd-controller/pkg/client/informers/externalversions/crd/v1beta1"
	listers "github.com/JiaoDean/crd-controller/pkg/client/listers/crd/v1beta1"
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

	crdLister listers.CrdLister
	synced    cache.InformerSynced
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
