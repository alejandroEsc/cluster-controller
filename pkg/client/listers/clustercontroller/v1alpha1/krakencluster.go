/*
Copyright 2018 Samsung CNCT.

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

package v1alpha1

import (
	v1alpha1 "github.com/samsung-cnct/cluster-controller/pkg/apis/clustercontroller/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KrakenClusterLister helps list KrakenClusters.
type KrakenClusterLister interface {
	// List lists all KrakenClusters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.KrakenCluster, err error)
	// KrakenClusters returns an object that can list and get KrakenClusters.
	KrakenClusters(namespace string) KrakenClusterNamespaceLister
	KrakenClusterListerExpansion
}

// krakenClusterLister implements the KrakenClusterLister interface.
type krakenClusterLister struct {
	indexer cache.Indexer
}

// NewKrakenClusterLister returns a new KrakenClusterLister.
func NewKrakenClusterLister(indexer cache.Indexer) KrakenClusterLister {
	return &krakenClusterLister{indexer: indexer}
}

// List lists all KrakenClusters in the indexer.
func (s *krakenClusterLister) List(selector labels.Selector) (ret []*v1alpha1.KrakenCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KrakenCluster))
	})
	return ret, err
}

// KrakenClusters returns an object that can list and get KrakenClusters.
func (s *krakenClusterLister) KrakenClusters(namespace string) KrakenClusterNamespaceLister {
	return krakenClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KrakenClusterNamespaceLister helps list and get KrakenClusters.
type KrakenClusterNamespaceLister interface {
	// List lists all KrakenClusters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.KrakenCluster, err error)
	// Get retrieves the KrakenCluster from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.KrakenCluster, error)
	KrakenClusterNamespaceListerExpansion
}

// krakenClusterNamespaceLister implements the KrakenClusterNamespaceLister
// interface.
type krakenClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all KrakenClusters in the indexer for a given namespace.
func (s krakenClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.KrakenCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.KrakenCluster))
	})
	return ret, err
}

// Get retrieves the KrakenCluster from the indexer for a given namespace and name.
func (s krakenClusterNamespaceLister) Get(name string) (*v1alpha1.KrakenCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("krakencluster"), name)
	}
	return obj.(*v1alpha1.KrakenCluster), nil
}
