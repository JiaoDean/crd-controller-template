package controller

import (
	"github.com/JiaoDean/crd-controller/pkg/apis/crd/v1beta1"
	log "github.com/sirupsen/logrus"
)

func (c *CrdController) handleCrdAdd(obj interface{}) {
	crdObj, ok := obj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer crd %+v is failed", crdObj)
		return
	}
}

func (c *CrdController) handleCrdDelete(obj interface{}) {
	crd, ok := obj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer crd %+v is failed", crd)
		return
	}
}

func (c *CrdController) handleCrdUpdate(oldObj interface{}, newObj interface{}) {
	newCrd, ok := newObj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer new crd %+v is failed", newCrd)
		return
	}
	oldCrd, ok := oldObj.(*v1beta1.Crd)
	if !ok {
		log.Errorf("Transfer old crd %+v is failed", oldCrd)
		return
	}
}
