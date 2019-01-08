// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	imageregistryv1 "github.com/openshift/cluster-image-registry-operator/pkg/apis/imageregistry/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageRegistries implements ImageRegistryInterface
type FakeImageRegistries struct {
	Fake *FakeImageregistryV1
}

var imageregistriesResource = schema.GroupVersionResource{Group: "imageregistry.operator.openshift.io", Version: "v1", Resource: "imageregistries"}

var imageregistriesKind = schema.GroupVersionKind{Group: "imageregistry.operator.openshift.io", Version: "v1", Kind: "ImageRegistry"}

// Get takes name of the imageRegistry, and returns the corresponding imageRegistry object, and an error if there is any.
func (c *FakeImageRegistries) Get(name string, options v1.GetOptions) (result *imageregistryv1.ImageRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(imageregistriesResource, name), &imageregistryv1.ImageRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.ImageRegistry), err
}

// List takes label and field selectors, and returns the list of ImageRegistries that match those selectors.
func (c *FakeImageRegistries) List(opts v1.ListOptions) (result *imageregistryv1.ImageRegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(imageregistriesResource, imageregistriesKind, opts), &imageregistryv1.ImageRegistryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &imageregistryv1.ImageRegistryList{ListMeta: obj.(*imageregistryv1.ImageRegistryList).ListMeta}
	for _, item := range obj.(*imageregistryv1.ImageRegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageRegistries.
func (c *FakeImageRegistries) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(imageregistriesResource, opts))
}

// Create takes the representation of a imageRegistry and creates it.  Returns the server's representation of the imageRegistry, and an error, if there is any.
func (c *FakeImageRegistries) Create(imageRegistry *imageregistryv1.ImageRegistry) (result *imageregistryv1.ImageRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(imageregistriesResource, imageRegistry), &imageregistryv1.ImageRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.ImageRegistry), err
}

// Update takes the representation of a imageRegistry and updates it. Returns the server's representation of the imageRegistry, and an error, if there is any.
func (c *FakeImageRegistries) Update(imageRegistry *imageregistryv1.ImageRegistry) (result *imageregistryv1.ImageRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(imageregistriesResource, imageRegistry), &imageregistryv1.ImageRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.ImageRegistry), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageRegistries) UpdateStatus(imageRegistry *imageregistryv1.ImageRegistry) (*imageregistryv1.ImageRegistry, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(imageregistriesResource, "status", imageRegistry), &imageregistryv1.ImageRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.ImageRegistry), err
}

// Delete takes name of the imageRegistry and deletes it. Returns an error if one occurs.
func (c *FakeImageRegistries) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(imageregistriesResource, name), &imageregistryv1.ImageRegistry{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageRegistries) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(imageregistriesResource, listOptions)

	_, err := c.Fake.Invokes(action, &imageregistryv1.ImageRegistryList{})
	return err
}

// Patch applies the patch and returns the patched imageRegistry.
func (c *FakeImageRegistries) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *imageregistryv1.ImageRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(imageregistriesResource, name, data, subresources...), &imageregistryv1.ImageRegistry{})
	if obj == nil {
		return nil, err
	}
	return obj.(*imageregistryv1.ImageRegistry), err
}