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

package fake

import (
	"context"

	configurationv1 "github.com/kong/kubernetes-ingress-controller/v2/pkg/apis/configuration/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKongIngresses implements KongIngressInterface
type FakeKongIngresses struct {
	Fake *FakeConfigurationV1
	ns   string
}

var kongingressesResource = schema.GroupVersionResource{Group: "configuration", Version: "v1", Resource: "kongingresses"}

var kongingressesKind = schema.GroupVersionKind{Group: "configuration", Version: "v1", Kind: "KongIngress"}

// Get takes name of the kongIngress, and returns the corresponding kongIngress object, and an error if there is any.
func (c *FakeKongIngresses) Get(ctx context.Context, name string, options v1.GetOptions) (result *configurationv1.KongIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kongingressesResource, c.ns, name), &configurationv1.KongIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongIngress), err
}

// List takes label and field selectors, and returns the list of KongIngresses that match those selectors.
func (c *FakeKongIngresses) List(ctx context.Context, opts v1.ListOptions) (result *configurationv1.KongIngressList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kongingressesResource, kongingressesKind, c.ns, opts), &configurationv1.KongIngressList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &configurationv1.KongIngressList{ListMeta: obj.(*configurationv1.KongIngressList).ListMeta}
	for _, item := range obj.(*configurationv1.KongIngressList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kongIngresses.
func (c *FakeKongIngresses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kongingressesResource, c.ns, opts))

}

// Create takes the representation of a kongIngress and creates it.  Returns the server's representation of the kongIngress, and an error, if there is any.
func (c *FakeKongIngresses) Create(ctx context.Context, kongIngress *configurationv1.KongIngress, opts v1.CreateOptions) (result *configurationv1.KongIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kongingressesResource, c.ns, kongIngress), &configurationv1.KongIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongIngress), err
}

// Update takes the representation of a kongIngress and updates it. Returns the server's representation of the kongIngress, and an error, if there is any.
func (c *FakeKongIngresses) Update(ctx context.Context, kongIngress *configurationv1.KongIngress, opts v1.UpdateOptions) (result *configurationv1.KongIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kongingressesResource, c.ns, kongIngress), &configurationv1.KongIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongIngress), err
}

// Delete takes name of the kongIngress and deletes it. Returns an error if one occurs.
func (c *FakeKongIngresses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kongingressesResource, c.ns, name, opts), &configurationv1.KongIngress{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKongIngresses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kongingressesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &configurationv1.KongIngressList{})
	return err
}

// Patch applies the patch and returns the patched kongIngress.
func (c *FakeKongIngresses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *configurationv1.KongIngress, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kongingressesResource, c.ns, name, pt, data, subresources...), &configurationv1.KongIngress{})

	if obj == nil {
		return nil, err
	}
	return obj.(*configurationv1.KongIngress), err
}
