/*
Copyright 2019 VMware

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

package v1alpha1

import (
	"time"

	scheme "github.com/projectcontour/contour/apis/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/projectcontour/contour/apis/projectcontour/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TLSCertificateDelegationsGetter has a method to return a TLSCertificateDelegationInterface.
// A group's client should implement this interface.
type TLSCertificateDelegationsGetter interface {
	TLSCertificateDelegations(namespace string) TLSCertificateDelegationInterface
}

// TLSCertificateDelegationInterface has methods to work with TLSCertificateDelegation resources.
type TLSCertificateDelegationInterface interface {
	Create(*v1alpha1.TLSCertificateDelegation) (*v1alpha1.TLSCertificateDelegation, error)
	Update(*v1alpha1.TLSCertificateDelegation) (*v1alpha1.TLSCertificateDelegation, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TLSCertificateDelegation, error)
	List(opts v1.ListOptions) (*v1alpha1.TLSCertificateDelegationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TLSCertificateDelegation, err error)
	TLSCertificateDelegationExpansion
}

// tLSCertificateDelegations implements TLSCertificateDelegationInterface
type tLSCertificateDelegations struct {
	client rest.Interface
	ns     string
}

// newTLSCertificateDelegations returns a TLSCertificateDelegations
func newTLSCertificateDelegations(c *ProjectcontourV1alpha1Client, namespace string) *tLSCertificateDelegations {
	return &tLSCertificateDelegations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tLSCertificateDelegation, and returns the corresponding tLSCertificateDelegation object, and an error if there is any.
func (c *tLSCertificateDelegations) Get(name string, options v1.GetOptions) (result *v1alpha1.TLSCertificateDelegation, err error) {
	result = &v1alpha1.TLSCertificateDelegation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TLSCertificateDelegations that match those selectors.
func (c *tLSCertificateDelegations) List(opts v1.ListOptions) (result *v1alpha1.TLSCertificateDelegationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TLSCertificateDelegationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tLSCertificateDelegations.
func (c *tLSCertificateDelegations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a tLSCertificateDelegation and creates it.  Returns the server's representation of the tLSCertificateDelegation, and an error, if there is any.
func (c *tLSCertificateDelegations) Create(tLSCertificateDelegation *v1alpha1.TLSCertificateDelegation) (result *v1alpha1.TLSCertificateDelegation, err error) {
	result = &v1alpha1.TLSCertificateDelegation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		Body(tLSCertificateDelegation).
		Do().
		Into(result)
	return
}

// Update takes the representation of a tLSCertificateDelegation and updates it. Returns the server's representation of the tLSCertificateDelegation, and an error, if there is any.
func (c *tLSCertificateDelegations) Update(tLSCertificateDelegation *v1alpha1.TLSCertificateDelegation) (result *v1alpha1.TLSCertificateDelegation, err error) {
	result = &v1alpha1.TLSCertificateDelegation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		Name(tLSCertificateDelegation.Name).
		Body(tLSCertificateDelegation).
		Do().
		Into(result)
	return
}

// Delete takes name of the tLSCertificateDelegation and deletes it. Returns an error if one occurs.
func (c *tLSCertificateDelegations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tLSCertificateDelegations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched tLSCertificateDelegation.
func (c *tLSCertificateDelegations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TLSCertificateDelegation, err error) {
	result = &v1alpha1.TLSCertificateDelegation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tlscertificatedelegations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
