/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/JiaoDean/crd-controller/pkg/apis/crd/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CrdLister helps list Crds.
type CrdLister interface {
	// List lists all Crds in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.Crd, err error)
	// Get retrieves the Crd from the index for a given name.
	Get(name string) (*v1beta1.Crd, error)
	CrdListerExpansion
}

// crdLister implements the CrdLister interface.
type crdLister struct {
	indexer cache.Indexer
}

// NewCrdLister returns a new CrdLister.
func NewCrdLister(indexer cache.Indexer) CrdLister {
	return &crdLister{indexer: indexer}
}

// List lists all Crds in the indexer.
func (s *crdLister) List(selector labels.Selector) (ret []*v1beta1.Crd, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Crd))
	})
	return ret, err
}

// Get retrieves the Crd from the index for a given name.
func (s *crdLister) Get(name string) (*v1beta1.Crd, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("crd"), name)
	}
	return obj.(*v1beta1.Crd), nil
}
