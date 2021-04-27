/*
Copyright 2021 Kong, Inc.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/kong/kubernetes-ingress-controller/railgun/apis/configuration/v1"
	scheme "github.com/kong/kubernetes-ingress-controller/railgun/pkg/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KongClusterPluginsGetter has a method to return a KongClusterPluginInterface.
// A group's client should implement this interface.
type KongClusterPluginsGetter interface {
	KongClusterPlugins(namespace string) KongClusterPluginInterface
}

// KongClusterPluginInterface has methods to work with KongClusterPlugin resources.
type KongClusterPluginInterface interface {
	Create(ctx context.Context, kongClusterPlugin *v1.KongClusterPlugin, opts metav1.CreateOptions) (*v1.KongClusterPlugin, error)
	Update(ctx context.Context, kongClusterPlugin *v1.KongClusterPlugin, opts metav1.UpdateOptions) (*v1.KongClusterPlugin, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.KongClusterPlugin, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.KongClusterPluginList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KongClusterPlugin, err error)
	KongClusterPluginExpansion
}

// kongClusterPlugins implements KongClusterPluginInterface
type kongClusterPlugins struct {
	client rest.Interface
	ns     string
}

// newKongClusterPlugins returns a KongClusterPlugins
func newKongClusterPlugins(c *ConfigurationV1Client, namespace string) *kongClusterPlugins {
	return &kongClusterPlugins{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kongClusterPlugin, and returns the corresponding kongClusterPlugin object, and an error if there is any.
func (c *kongClusterPlugins) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.KongClusterPlugin, err error) {
	result = &v1.KongClusterPlugin{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongclusterplugins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KongClusterPlugins that match those selectors.
func (c *kongClusterPlugins) List(ctx context.Context, opts metav1.ListOptions) (result *v1.KongClusterPluginList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.KongClusterPluginList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kongclusterplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kongClusterPlugins.
func (c *kongClusterPlugins) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kongclusterplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kongClusterPlugin and creates it.  Returns the server's representation of the kongClusterPlugin, and an error, if there is any.
func (c *kongClusterPlugins) Create(ctx context.Context, kongClusterPlugin *v1.KongClusterPlugin, opts metav1.CreateOptions) (result *v1.KongClusterPlugin, err error) {
	result = &v1.KongClusterPlugin{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kongclusterplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kongClusterPlugin).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kongClusterPlugin and updates it. Returns the server's representation of the kongClusterPlugin, and an error, if there is any.
func (c *kongClusterPlugins) Update(ctx context.Context, kongClusterPlugin *v1.KongClusterPlugin, opts metav1.UpdateOptions) (result *v1.KongClusterPlugin, err error) {
	result = &v1.KongClusterPlugin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kongclusterplugins").
		Name(kongClusterPlugin.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kongClusterPlugin).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kongClusterPlugin and deletes it. Returns an error if one occurs.
func (c *kongClusterPlugins) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongclusterplugins").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kongClusterPlugins) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kongclusterplugins").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kongClusterPlugin.
func (c *kongClusterPlugins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KongClusterPlugin, err error) {
	result = &v1.KongClusterPlugin{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kongclusterplugins").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}