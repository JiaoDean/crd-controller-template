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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/JiaoDean/crd-controller/pkg/apis/crd/v1beta1"
	scheme "github.com/JiaoDean/crd-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CrdsGetter has a method to return a CrdInterface.
// A group's client should implement this interface.
type CrdsGetter interface {
	Crds() CrdInterface
}

// CrdInterface has methods to work with Crd resources.
type CrdInterface interface {
	Create(ctx context.Context, crd *v1beta1.Crd, opts v1.CreateOptions) (*v1beta1.Crd, error)
	Update(ctx context.Context, crd *v1beta1.Crd, opts v1.UpdateOptions) (*v1beta1.Crd, error)
	UpdateStatus(ctx context.Context, crd *v1beta1.Crd, opts v1.UpdateOptions) (*v1beta1.Crd, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Crd, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.CrdList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Crd, err error)
	CrdExpansion
}

// crds implements CrdInterface
type crds struct {
	client rest.Interface
}

// newCrds returns a Crds
func newCrds(c *StorageV1beta1Client) *crds {
	return &crds{
		client: c.RESTClient(),
	}
}

// Get takes name of the crd, and returns the corresponding crd object, and an error if there is any.
func (c *crds) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Crd, err error) {
	result = &v1beta1.Crd{}
	err = c.client.Get().
		Resource("crds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Crds that match those selectors.
func (c *crds) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CrdList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.CrdList{}
	err = c.client.Get().
		Resource("crds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested crds.
func (c *crds) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("crds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a crd and creates it.  Returns the server's representation of the crd, and an error, if there is any.
func (c *crds) Create(ctx context.Context, crd *v1beta1.Crd, opts v1.CreateOptions) (result *v1beta1.Crd, err error) {
	result = &v1beta1.Crd{}
	err = c.client.Post().
		Resource("crds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(crd).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a crd and updates it. Returns the server's representation of the crd, and an error, if there is any.
func (c *crds) Update(ctx context.Context, crd *v1beta1.Crd, opts v1.UpdateOptions) (result *v1beta1.Crd, err error) {
	result = &v1beta1.Crd{}
	err = c.client.Put().
		Resource("crds").
		Name(crd.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(crd).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *crds) UpdateStatus(ctx context.Context, crd *v1beta1.Crd, opts v1.UpdateOptions) (result *v1beta1.Crd, err error) {
	result = &v1beta1.Crd{}
	err = c.client.Put().
		Resource("crds").
		Name(crd.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(crd).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the crd and deletes it. Returns an error if one occurs.
func (c *crds) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("crds").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *crds) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("crds").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched crd.
func (c *crds) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Crd, err error) {
	result = &v1beta1.Crd{}
	err = c.client.Patch(pt).
		Resource("crds").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}