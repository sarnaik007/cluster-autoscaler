/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1alpha1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/poc.autoscaling.k8s.io/v1alpha1"
	scheme "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

// VerticalPodAutoscalersGetter has a method to return a VerticalPodAutoscalerInterface.
// A group's client should implement this interface.
type VerticalPodAutoscalersGetter interface {
	VerticalPodAutoscalers(namespace string) VerticalPodAutoscalerInterface
}

// VerticalPodAutoscalerInterface has methods to work with VerticalPodAutoscaler resources.
type VerticalPodAutoscalerInterface interface {
	Create(*v1alpha1.VerticalPodAutoscaler) (*v1alpha1.VerticalPodAutoscaler, error)
	Update(*v1alpha1.VerticalPodAutoscaler) (*v1alpha1.VerticalPodAutoscaler, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VerticalPodAutoscaler, error)
	List(opts v1.ListOptions) (*v1alpha1.VerticalPodAutoscalerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VerticalPodAutoscaler, err error)
	VerticalPodAutoscalerExpansion
}

// verticalPodAutoscalers implements VerticalPodAutoscalerInterface
type verticalPodAutoscalers struct {
	client rest.Interface
	ns     string
}

// newVerticalPodAutoscalers returns a VerticalPodAutoscalers
func newVerticalPodAutoscalers(c *PocV1alpha1Client, namespace string) *verticalPodAutoscalers {
	return &verticalPodAutoscalers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the verticalPodAutoscaler, and returns the corresponding verticalPodAutoscaler object, and an error if there is any.
func (c *verticalPodAutoscalers) Get(name string, options v1.GetOptions) (result *v1alpha1.VerticalPodAutoscaler, err error) {
	result = &v1alpha1.VerticalPodAutoscaler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VerticalPodAutoscalers that match those selectors.
func (c *verticalPodAutoscalers) List(opts v1.ListOptions) (result *v1alpha1.VerticalPodAutoscalerList, err error) {
	result = &v1alpha1.VerticalPodAutoscalerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested verticalPodAutoscalers.
func (c *verticalPodAutoscalers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a verticalPodAutoscaler and creates it.  Returns the server's representation of the verticalPodAutoscaler, and an error, if there is any.
func (c *verticalPodAutoscalers) Create(verticalPodAutoscaler *v1alpha1.VerticalPodAutoscaler) (result *v1alpha1.VerticalPodAutoscaler, err error) {
	result = &v1alpha1.VerticalPodAutoscaler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		Body(verticalPodAutoscaler).
		Do().
		Into(result)
	return
}

// Update takes the representation of a verticalPodAutoscaler and updates it. Returns the server's representation of the verticalPodAutoscaler, and an error, if there is any.
func (c *verticalPodAutoscalers) Update(verticalPodAutoscaler *v1alpha1.VerticalPodAutoscaler) (result *v1alpha1.VerticalPodAutoscaler, err error) {
	result = &v1alpha1.VerticalPodAutoscaler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		Name(verticalPodAutoscaler.Name).
		Body(verticalPodAutoscaler).
		Do().
		Into(result)
	return
}

// Delete takes name of the verticalPodAutoscaler and deletes it. Returns an error if one occurs.
func (c *verticalPodAutoscalers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *verticalPodAutoscalers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched verticalPodAutoscaler.
func (c *verticalPodAutoscalers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VerticalPodAutoscaler, err error) {
	result = &v1alpha1.VerticalPodAutoscaler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("verticalpodautoscalers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
