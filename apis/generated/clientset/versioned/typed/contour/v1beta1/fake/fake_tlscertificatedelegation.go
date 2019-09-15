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

package fake

import (
	v1beta1 "github.com/projectcontour/contour/apis/contour/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTLSCertificateDelegations implements TLSCertificateDelegationInterface
type FakeTLSCertificateDelegations struct {
	Fake *FakeContourV1beta1
	ns   string
}

var tlscertificatedelegationsResource = schema.GroupVersionResource{Group: "contour.heptio.com", Version: "v1beta1", Resource: "tlscertificatedelegations"}

var tlscertificatedelegationsKind = schema.GroupVersionKind{Group: "contour.heptio.com", Version: "v1beta1", Kind: "TLSCertificateDelegation"}

// Get takes name of the tLSCertificateDelegation, and returns the corresponding tLSCertificateDelegation object, and an error if there is any.
func (c *FakeTLSCertificateDelegations) Get(name string, options v1.GetOptions) (result *v1beta1.TLSCertificateDelegation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tlscertificatedelegationsResource, c.ns, name), &v1beta1.TLSCertificateDelegation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TLSCertificateDelegation), err
}

// List takes label and field selectors, and returns the list of TLSCertificateDelegations that match those selectors.
func (c *FakeTLSCertificateDelegations) List(opts v1.ListOptions) (result *v1beta1.TLSCertificateDelegationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tlscertificatedelegationsResource, tlscertificatedelegationsKind, c.ns, opts), &v1beta1.TLSCertificateDelegationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.TLSCertificateDelegationList{ListMeta: obj.(*v1beta1.TLSCertificateDelegationList).ListMeta}
	for _, item := range obj.(*v1beta1.TLSCertificateDelegationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tLSCertificateDelegations.
func (c *FakeTLSCertificateDelegations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tlscertificatedelegationsResource, c.ns, opts))

}

// Create takes the representation of a tLSCertificateDelegation and creates it.  Returns the server's representation of the tLSCertificateDelegation, and an error, if there is any.
func (c *FakeTLSCertificateDelegations) Create(tLSCertificateDelegation *v1beta1.TLSCertificateDelegation) (result *v1beta1.TLSCertificateDelegation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tlscertificatedelegationsResource, c.ns, tLSCertificateDelegation), &v1beta1.TLSCertificateDelegation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TLSCertificateDelegation), err
}

// Update takes the representation of a tLSCertificateDelegation and updates it. Returns the server's representation of the tLSCertificateDelegation, and an error, if there is any.
func (c *FakeTLSCertificateDelegations) Update(tLSCertificateDelegation *v1beta1.TLSCertificateDelegation) (result *v1beta1.TLSCertificateDelegation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tlscertificatedelegationsResource, c.ns, tLSCertificateDelegation), &v1beta1.TLSCertificateDelegation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TLSCertificateDelegation), err
}

// Delete takes name of the tLSCertificateDelegation and deletes it. Returns an error if one occurs.
func (c *FakeTLSCertificateDelegations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tlscertificatedelegationsResource, c.ns, name), &v1beta1.TLSCertificateDelegation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTLSCertificateDelegations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tlscertificatedelegationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.TLSCertificateDelegationList{})
	return err
}

// Patch applies the patch and returns the patched tLSCertificateDelegation.
func (c *FakeTLSCertificateDelegations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.TLSCertificateDelegation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tlscertificatedelegationsResource, c.ns, name, pt, data, subresources...), &v1beta1.TLSCertificateDelegation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.TLSCertificateDelegation), err
}
