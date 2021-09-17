package kube

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
)

func (c *KubeController) handlePvcAdd(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Error decoding object tombstone, invalid type"))
			return
		}
		log.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}

	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	pvcIndex := c.GetPvcIndex(pvcObj)
	c.SetPvcMap(pvcIndex, pvcObj)
}

func (c *KubeController) handlePvcDelete(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Error decoding object tombstone, invalid type"))
			return
		}
		log.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}

	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	pvcIndex := c.GetPvcIndex(pvcObj)
	delete(c.GetPvcMap(), pvcIndex)
}

func (c *KubeController) handlePvcUpdate(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("Error decoding object tombstone, invalid type"))
			return
		}
		log.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}

	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	pvcIndex := c.GetPvcIndex(pvcObj)
	c.SetPvcMap(pvcIndex, pvcObj)
}

func (c *KubeController) GetPvcIndex(pvcObj *corev1.PersistentVolumeClaim) string {
	pvcIndex := pvcObj.Namespace + "/" + pvcObj.Name
	return pvcIndex
}
