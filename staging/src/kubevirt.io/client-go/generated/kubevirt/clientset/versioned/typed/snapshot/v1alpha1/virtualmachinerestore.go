/*
Copyright 2020 The KubeVirt Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "kubevirt.io/client-go/apis/snapshot/v1alpha1"
	scheme "kubevirt.io/client-go/generated/kubevirt/clientset/versioned/scheme"
)

// VirtualMachineRestoresGetter has a method to return a VirtualMachineRestoreInterface.
// A group's client should implement this interface.
type VirtualMachineRestoresGetter interface {
	VirtualMachineRestores(namespace string) VirtualMachineRestoreInterface
}

// VirtualMachineRestoreInterface has methods to work with VirtualMachineRestore resources.
type VirtualMachineRestoreInterface interface {
	Create(*v1alpha1.VirtualMachineRestore) (*v1alpha1.VirtualMachineRestore, error)
	Update(*v1alpha1.VirtualMachineRestore) (*v1alpha1.VirtualMachineRestore, error)
	UpdateStatus(*v1alpha1.VirtualMachineRestore) (*v1alpha1.VirtualMachineRestore, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualMachineRestore, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualMachineRestoreList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineRestore, err error)
	VirtualMachineRestoreExpansion
}

// virtualMachineRestores implements VirtualMachineRestoreInterface
type virtualMachineRestores struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineRestores returns a VirtualMachineRestores
func newVirtualMachineRestores(c *SnapshotV1alpha1Client, namespace string) *virtualMachineRestores {
	return &virtualMachineRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineRestore, and returns the corresponding virtualMachineRestore object, and an error if there is any.
func (c *virtualMachineRestores) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachineRestore, err error) {
	result = &v1alpha1.VirtualMachineRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineRestores that match those selectors.
func (c *virtualMachineRestores) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineRestoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualMachineRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineRestores.
func (c *virtualMachineRestores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualMachineRestore and creates it.  Returns the server's representation of the virtualMachineRestore, and an error, if there is any.
func (c *virtualMachineRestores) Create(virtualMachineRestore *v1alpha1.VirtualMachineRestore) (result *v1alpha1.VirtualMachineRestore, err error) {
	result = &v1alpha1.VirtualMachineRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Body(virtualMachineRestore).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualMachineRestore and updates it. Returns the server's representation of the virtualMachineRestore, and an error, if there is any.
func (c *virtualMachineRestores) Update(virtualMachineRestore *v1alpha1.VirtualMachineRestore) (result *v1alpha1.VirtualMachineRestore, err error) {
	result = &v1alpha1.VirtualMachineRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(virtualMachineRestore.Name).
		Body(virtualMachineRestore).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualMachineRestores) UpdateStatus(virtualMachineRestore *v1alpha1.VirtualMachineRestore) (result *v1alpha1.VirtualMachineRestore, err error) {
	result = &v1alpha1.VirtualMachineRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(virtualMachineRestore.Name).
		SubResource("status").
		Body(virtualMachineRestore).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualMachineRestore and deletes it. Returns an error if one occurs.
func (c *virtualMachineRestores) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineRestores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualMachineRestore.
func (c *virtualMachineRestores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachineRestore, err error) {
	result = &v1alpha1.VirtualMachineRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachinerestores").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
