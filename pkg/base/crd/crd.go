package crd

import (
	"bytes"
	"context"
	"github.com/JiaoDean/crd-controller/pkg/apis/crd/v1beta1"
	crdClient "github.com/JiaoDean/crd-controller/pkg/client/clientset/versioned"
	log "github.com/sirupsen/logrus"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
	"sync"
)

var crdMutex sync.RWMutex

func GetCrd(crdClient *crdClient.Clientset, name string) (*v1beta1.Crd, error) {
	crdMutex.RLock()
	defer crdMutex.RUnlock()
	cnfs, err := crdClient.KubernetesV1beta1().Crds().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		log.Errorf("Get crd %s is failed, err:%s", name, err)
		return nil, err
	}
	return cnfs, nil
}


func ParseCrd(yamlPath string) (*v1beta1.Crd, error) {
	r, err := utils.ReadFile(yamlPath)
	if err != nil {
		log.Errorf("Parse cnfs is failed, yamlPath:%s, err:%s", yamlPath, err.Error())
		return nil, err
	}
	cnfs := v1beta1.ContainerNetworkFileSystem{}
	err = yaml.NewYAMLOrJSONDecoder(bytes.NewReader(r), defaultBufferSize).Decode(&cnfs)
	if err != nil {
		log.Errorf("Parse cnfs yaml decoded is failed, yamlPath:%s, err:%s", yamlPath, err.Error())
		return nil, err
	}
	return &cnfs, nil
}