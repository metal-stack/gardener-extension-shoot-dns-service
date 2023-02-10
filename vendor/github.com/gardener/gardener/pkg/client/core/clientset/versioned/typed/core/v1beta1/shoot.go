/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/authentication/v1alpha1"
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	scheme "github.com/gardener/gardener/pkg/client/core/clientset/versioned/scheme"
)

// ShootsGetter has a method to return a ShootInterface.
// A group's client should implement this interface.
type ShootsGetter interface {
	Shoots(namespace string) ShootInterface
}

// ShootInterface has methods to work with Shoot resources.
type ShootInterface interface {
	Create(ctx context.Context, shoot *v1beta1.Shoot, opts v1.CreateOptions) (*v1beta1.Shoot, error)
	Update(ctx context.Context, shoot *v1beta1.Shoot, opts v1.UpdateOptions) (*v1beta1.Shoot, error)
	UpdateStatus(ctx context.Context, shoot *v1beta1.Shoot, opts v1.UpdateOptions) (*v1beta1.Shoot, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.Shoot, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ShootList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Shoot, err error)
	CreateAdminKubeconfigRequest(ctx context.Context, shootName string, adminKubeconfigRequest *v1alpha1.AdminKubeconfigRequest, opts v1.CreateOptions) (*v1alpha1.AdminKubeconfigRequest, error)
	UpdateBinding(ctx context.Context, shootName string, shoot *v1beta1.Shoot, opts v1.UpdateOptions) (*v1beta1.Shoot, error)

	ShootExpansion
}

// shoots implements ShootInterface
type shoots struct {
	client rest.Interface
	ns     string
}

// newShoots returns a Shoots
func newShoots(c *CoreV1beta1Client, namespace string) *shoots {
	return &shoots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the shoot, and returns the corresponding shoot object, and an error if there is any.
func (c *shoots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Shoot, err error) {
	result = &v1beta1.Shoot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shoots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Shoots that match those selectors.
func (c *shoots) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ShootList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ShootList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shoots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested shoots.
func (c *shoots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("shoots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a shoot and creates it.  Returns the server's representation of the shoot, and an error, if there is any.
func (c *shoots) Create(ctx context.Context, shoot *v1beta1.Shoot, opts v1.CreateOptions) (result *v1beta1.Shoot, err error) {
	result = &v1beta1.Shoot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("shoots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shoot).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a shoot and updates it. Returns the server's representation of the shoot, and an error, if there is any.
func (c *shoots) Update(ctx context.Context, shoot *v1beta1.Shoot, opts v1.UpdateOptions) (result *v1beta1.Shoot, err error) {
	result = &v1beta1.Shoot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shoots").
		Name(shoot.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shoot).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *shoots) UpdateStatus(ctx context.Context, shoot *v1beta1.Shoot, opts v1.UpdateOptions) (result *v1beta1.Shoot, err error) {
	result = &v1beta1.Shoot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shoots").
		Name(shoot.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shoot).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the shoot and deletes it. Returns an error if one occurs.
func (c *shoots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shoots").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *shoots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shoots").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched shoot.
func (c *shoots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Shoot, err error) {
	result = &v1beta1.Shoot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("shoots").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// CreateAdminKubeconfigRequest takes the representation of a adminKubeconfigRequest and creates it.  Returns the server's representation of the adminKubeconfigRequest, and an error, if there is any.
func (c *shoots) CreateAdminKubeconfigRequest(ctx context.Context, shootName string, adminKubeconfigRequest *v1alpha1.AdminKubeconfigRequest, opts v1.CreateOptions) (result *v1alpha1.AdminKubeconfigRequest, err error) {
	result = &v1alpha1.AdminKubeconfigRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("shoots").
		Name(shootName).
		SubResource("adminkubeconfig").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(adminKubeconfigRequest).
		Do(ctx).
		Into(result)
	return
}

// UpdateBinding takes the top resource name and the representation of a shoot and updates it. Returns the server's representation of the shoot, and an error, if there is any.
func (c *shoots) UpdateBinding(ctx context.Context, shootName string, shoot *v1beta1.Shoot, opts v1.UpdateOptions) (result *v1beta1.Shoot, err error) {
	result = &v1beta1.Shoot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shoots").
		Name(shootName).
		SubResource("binding").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(shoot).
		Do(ctx).
		Into(result)
	return
}
