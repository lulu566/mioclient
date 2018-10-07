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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/hidevopsio/mioclient/pkg/apis/mio/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuildConfigs implements BuildConfigInterface
type FakeBuildConfigs struct {
	Fake *FakeMioV1alpha1
	ns   string
}

var buildconfigsResource = schema.GroupVersionResource{Group: "mio.k8s.io", Version: "v1alpha1", Resource: "buildconfigs"}

var buildconfigsKind = schema.GroupVersionKind{Group: "mio.k8s.io", Version: "v1alpha1", Kind: "BuildConfig"}

// Get takes name of the buildConfig, and returns the corresponding buildConfig object, and an error if there is any.
func (c *FakeBuildConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildconfigsResource, c.ns, name), &v1alpha1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BuildConfig), err
}

// List takes label and field selectors, and returns the list of BuildConfigs that match those selectors.
func (c *FakeBuildConfigs) List(opts v1.ListOptions) (result *v1alpha1.BuildConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildconfigsResource, buildconfigsKind, c.ns, opts), &v1alpha1.BuildConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BuildConfigList{}
	for _, item := range obj.(*v1alpha1.BuildConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buildConfigs.
func (c *FakeBuildConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildconfigsResource, c.ns, opts))

}

// Create takes the representation of a buildConfig and creates it.  Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *FakeBuildConfigs) Create(buildConfig *v1alpha1.BuildConfig) (result *v1alpha1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildconfigsResource, c.ns, buildConfig), &v1alpha1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BuildConfig), err
}

// Update takes the representation of a buildConfig and updates it. Returns the server's representation of the buildConfig, and an error, if there is any.
func (c *FakeBuildConfigs) Update(buildConfig *v1alpha1.BuildConfig) (result *v1alpha1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildconfigsResource, c.ns, buildConfig), &v1alpha1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BuildConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuildConfigs) UpdateStatus(buildConfig *v1alpha1.BuildConfig) (*v1alpha1.BuildConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildconfigsResource, "status", c.ns, buildConfig), &v1alpha1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BuildConfig), err
}

// Delete takes name of the buildConfig and deletes it. Returns an error if one occurs.
func (c *FakeBuildConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(buildconfigsResource, c.ns, name), &v1alpha1.BuildConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuildConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.BuildConfigList{})
	return err
}

// Patch applies the patch and returns the patched buildConfig.
func (c *FakeBuildConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BuildConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildconfigsResource, c.ns, name, data, subresources...), &v1alpha1.BuildConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BuildConfig), err
}
