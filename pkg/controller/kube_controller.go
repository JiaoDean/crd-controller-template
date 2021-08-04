package controller

import (
	corev1 "k8s.io/api/core/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	storagelisters "k8s.io/client-go/listers/storage/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
)

type KubeController struct {
	// kubeclientset is a standard kubernetes clientset
	clientset kubernetes.Interface

	pvLister corelisters.PersistentVolumeLister
	pvSynced cache.InformerSynced

	pvcLister corelisters.PersistentVolumeClaimLister
	pvcSynced cache.InformerSynced

	scLister storagelisters.StorageClassLister
	scSynced cache.InformerSynced

	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder

	Asow *ActualStorageOfWorld
}

type ActualStorageOfWorld struct {
	PvList  map[string]*corev1.PersistentVolume
	PvcList map[string]*corev1.PersistentVolumeClaim
}

func NewKubeController(kubeClient kubernetes.Interface,
	pvInformer coreinformers.PersistentVolumeInformer,
	pvcInformer coreinformers.PersistentVolumeClaimInformer,
	recorder record.EventRecorder) *KubeController {

	asow := &ActualStorageOfWorld{}
	pvList := map[string]*corev1.PersistentVolume{}
	pvcList := map[string]*corev1.PersistentVolumeClaim{}
	asow.PvList = pvList
	asow.PvcList = pvcList

	controller := &KubeController{
		clientset: kubeClient,
		pvcLister: pvcInformer.Lister(),
		pvcSynced: pvcInformer.Informer().HasSynced,
		pvLister:  pvInformer.Lister(),
		pvSynced:  pvInformer.Informer().HasSynced,
		recorder:  recorder,
		Asow:      asow,
	}

	pvInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handlePvAdd,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*corev1.PersistentVolume)
			oldDepl := old.(*corev1.PersistentVolume)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			controller.handlePvUpdate(new)
		},
		DeleteFunc: controller.handlePvDelete,
	})

	pvcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handlePvcAdd,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*corev1.PersistentVolumeClaim)
			oldDepl := old.(*corev1.PersistentVolumeClaim)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			controller.handlePvcUpdate(new)
		},
		DeleteFunc: controller.handlePvcDelete,
	})
	return controller
}
