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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	gatewayapiapisv1 "sigs.k8s.io/gateway-api/apis/v1"
	versioned "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
	internalinterfaces "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions/internalinterfaces"
	apisv1 "sigs.k8s.io/gateway-api/pkg/client/listers/apis/v1"
)

// HTTPRouteInformer provides access to a shared informer and lister for
// HTTPRoutes.
type HTTPRouteInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() apisv1.HTTPRouteLister
}

type hTTPRouteInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHTTPRouteInformer constructs a new informer for HTTPRoute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHTTPRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHTTPRouteInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHTTPRouteInformer constructs a new informer for HTTPRoute type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHTTPRouteInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1().HTTPRoutes(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1().HTTPRoutes(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1().HTTPRoutes(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GatewayV1().HTTPRoutes(namespace).Watch(ctx, options)
			},
		},
		&gatewayapiapisv1.HTTPRoute{},
		resyncPeriod,
		indexers,
	)
}

func (f *hTTPRouteInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHTTPRouteInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hTTPRouteInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&gatewayapiapisv1.HTTPRoute{}, f.defaultInformer)
}

func (f *hTTPRouteInformer) Lister() apisv1.HTTPRouteLister {
	return apisv1.NewHTTPRouteLister(f.Informer().GetIndexer())
}
