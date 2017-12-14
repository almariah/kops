/*
Copyright 2017 The Kubernetes Authors.

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
package v1alpha2

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha2 "k8s.io/kops/pkg/apis/kops/v1alpha2"
	scheme "k8s.io/kops/pkg/client/clientset_generated/internalclientset/scheme"
)

// SSHCredentialsGetter has a method to return a SSHCredentialInterface.
// A group's client should implement this interface.
type SSHCredentialsGetter interface {
	SSHCredentials(namespace string) SSHCredentialInterface
}

// SSHCredentialInterface has methods to work with SSHCredential resources.
type SSHCredentialInterface interface {
	Create(*v1alpha2.SSHCredential) (*v1alpha2.SSHCredential, error)
	Update(*v1alpha2.SSHCredential) (*v1alpha2.SSHCredential, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.SSHCredential, error)
	List(opts v1.ListOptions) (*v1alpha2.SSHCredentialList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.SSHCredential, err error)
	SSHCredentialExpansion
}

// sSHCredentials implements SSHCredentialInterface
type sSHCredentials struct {
	client rest.Interface
	ns     string
}

// newSSHCredentials returns a SSHCredentials
func newSSHCredentials(c *KopsV1alpha2Client, namespace string) *sSHCredentials {
	return &sSHCredentials{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sSHCredential, and returns the corresponding sSHCredential object, and an error if there is any.
func (c *sSHCredentials) Get(name string, options v1.GetOptions) (result *v1alpha2.SSHCredential, err error) {
	result = &v1alpha2.SSHCredential{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshcredentials").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SSHCredentials that match those selectors.
func (c *sSHCredentials) List(opts v1.ListOptions) (result *v1alpha2.SSHCredentialList, err error) {
	result = &v1alpha2.SSHCredentialList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshcredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sSHCredentials.
func (c *sSHCredentials) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sshcredentials").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a sSHCredential and creates it.  Returns the server's representation of the sSHCredential, and an error, if there is any.
func (c *sSHCredentials) Create(sSHCredential *v1alpha2.SSHCredential) (result *v1alpha2.SSHCredential, err error) {
	result = &v1alpha2.SSHCredential{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sshcredentials").
		Body(sSHCredential).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sSHCredential and updates it. Returns the server's representation of the sSHCredential, and an error, if there is any.
func (c *sSHCredentials) Update(sSHCredential *v1alpha2.SSHCredential) (result *v1alpha2.SSHCredential, err error) {
	result = &v1alpha2.SSHCredential{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sshcredentials").
		Name(sSHCredential.Name).
		Body(sSHCredential).
		Do().
		Into(result)
	return
}

// Delete takes name of the sSHCredential and deletes it. Returns an error if one occurs.
func (c *sSHCredentials) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshcredentials").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sSHCredentials) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshcredentials").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sSHCredential.
func (c *sSHCredentials) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.SSHCredential, err error) {
	result = &v1alpha2.SSHCredential{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sshcredentials").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
