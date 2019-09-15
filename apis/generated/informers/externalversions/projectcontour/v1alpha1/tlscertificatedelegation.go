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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	versioned "github.com/projectcontour/contour/apis/generated/clientset/versioned"
	internalinterfaces "github.com/projectcontour/contour/apis/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/projectcontour/contour/apis/generated/listers/projectcontour/v1alpha1"
	projectcontourv1alpha1 "github.com/projectcontour/contour/apis/projectcontour/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TLSCertificateDelegationInformer provides access to a shared informer and lister for
// TLSCertificateDelegations.
type TLSCertificateDelegationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TLSCertificateDelegationLister
}

type tLSCertificateDelegationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTLSCertificateDelegationInformer constructs a new informer for TLSCertificateDelegation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTLSCertificateDelegationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTLSCertificateDelegationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTLSCertificateDelegationInformer constructs a new informer for TLSCertificateDelegation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTLSCertificateDelegationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcontourV1alpha1().TLSCertificateDelegations(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcontourV1alpha1().TLSCertificateDelegations(namespace).Watch(options)
			},
		},
		&projectcontourv1alpha1.TLSCertificateDelegation{},
		resyncPeriod,
		indexers,
	)
}

func (f *tLSCertificateDelegationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTLSCertificateDelegationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tLSCertificateDelegationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&projectcontourv1alpha1.TLSCertificateDelegation{}, f.defaultInformer)
}

func (f *tLSCertificateDelegationInformer) Lister() v1alpha1.TLSCertificateDelegationLister {
	return v1alpha1.NewTLSCertificateDelegationLister(f.Informer().GetIndexer())
}
