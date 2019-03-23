/*
Copyright 2019 The Pharmer Authors.

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
	cloudv1 "github.com/pharmer/cloud/pkg/apis/cloud/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	testing "k8s.io/client-go/testing"
)

// FakeRegions implements RegionInterface
type FakeRegions struct {
	Fake *FakeCloudV1
}

var regionsResource = schema.GroupVersionResource{Group: "cloud.pharmer.io", Version: "v1", Resource: "regions"}

var regionsKind = schema.GroupVersionKind{Group: "cloud.pharmer.io", Version: "v1", Kind: "Region"}

// Get takes name of the region, and returns the corresponding region object, and an error if there is any.
func (c *FakeRegions) Get(name string, options v1.GetOptions) (result *cloudv1.Region, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(regionsResource, name), &cloudv1.Region{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.Region), err
}

// List takes label and field selectors, and returns the list of Regions that match those selectors.
func (c *FakeRegions) List(opts v1.ListOptions) (result *cloudv1.RegionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(regionsResource, regionsKind, opts), &cloudv1.RegionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &cloudv1.RegionList{ListMeta: obj.(*cloudv1.RegionList).ListMeta}
	for _, item := range obj.(*cloudv1.RegionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Create takes the representation of a region and creates it.  Returns the server's representation of the region, and an error, if there is any.
func (c *FakeRegions) Create(region *cloudv1.Region) (result *cloudv1.Region, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(regionsResource, region), &cloudv1.Region{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.Region), err
}

// Update takes the representation of a region and updates it. Returns the server's representation of the region, and an error, if there is any.
func (c *FakeRegions) Update(region *cloudv1.Region) (result *cloudv1.Region, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(regionsResource, region), &cloudv1.Region{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.Region), err
}

// Delete takes name of the region and deletes it. Returns an error if one occurs.
func (c *FakeRegions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(regionsResource, name), &cloudv1.Region{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(regionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &cloudv1.RegionList{})
	return err
}

// Patch applies the patch and returns the patched region.
func (c *FakeRegions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *cloudv1.Region, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(regionsResource, name, pt, data, subresources...), &cloudv1.Region{})
	if obj == nil {
		return nil, err
	}
	return obj.(*cloudv1.Region), err
}