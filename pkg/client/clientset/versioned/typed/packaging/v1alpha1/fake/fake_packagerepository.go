// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/vmware-tanzu/carvel-kapp-controller/pkg/apis/packaging/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePackageRepositories implements PackageRepositoryInterface
type FakePackageRepositories struct {
	Fake *FakePackagingV1alpha1
	ns   string
}

var packagerepositoriesResource = schema.GroupVersionResource{Group: "packaging.carvel.dev", Version: "v1alpha1", Resource: "packagerepositories"}

var packagerepositoriesKind = schema.GroupVersionKind{Group: "packaging.carvel.dev", Version: "v1alpha1", Kind: "PackageRepository"}

// Get takes name of the packageRepository, and returns the corresponding packageRepository object, and an error if there is any.
func (c *FakePackageRepositories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PackageRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(packagerepositoriesResource, c.ns, name), &v1alpha1.PackageRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRepository), err
}

// List takes label and field selectors, and returns the list of PackageRepositories that match those selectors.
func (c *FakePackageRepositories) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PackageRepositoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(packagerepositoriesResource, packagerepositoriesKind, c.ns, opts), &v1alpha1.PackageRepositoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PackageRepositoryList{ListMeta: obj.(*v1alpha1.PackageRepositoryList).ListMeta}
	for _, item := range obj.(*v1alpha1.PackageRepositoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested packageRepositories.
func (c *FakePackageRepositories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(packagerepositoriesResource, c.ns, opts))

}

// Create takes the representation of a packageRepository and creates it.  Returns the server's representation of the packageRepository, and an error, if there is any.
func (c *FakePackageRepositories) Create(ctx context.Context, packageRepository *v1alpha1.PackageRepository, opts v1.CreateOptions) (result *v1alpha1.PackageRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(packagerepositoriesResource, c.ns, packageRepository), &v1alpha1.PackageRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRepository), err
}

// Update takes the representation of a packageRepository and updates it. Returns the server's representation of the packageRepository, and an error, if there is any.
func (c *FakePackageRepositories) Update(ctx context.Context, packageRepository *v1alpha1.PackageRepository, opts v1.UpdateOptions) (result *v1alpha1.PackageRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(packagerepositoriesResource, c.ns, packageRepository), &v1alpha1.PackageRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRepository), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePackageRepositories) UpdateStatus(ctx context.Context, packageRepository *v1alpha1.PackageRepository, opts v1.UpdateOptions) (*v1alpha1.PackageRepository, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(packagerepositoriesResource, "status", c.ns, packageRepository), &v1alpha1.PackageRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRepository), err
}

// Delete takes name of the packageRepository and deletes it. Returns an error if one occurs.
func (c *FakePackageRepositories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(packagerepositoriesResource, c.ns, name), &v1alpha1.PackageRepository{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePackageRepositories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(packagerepositoriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PackageRepositoryList{})
	return err
}

// Patch applies the patch and returns the patched packageRepository.
func (c *FakePackageRepositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PackageRepository, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(packagerepositoriesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PackageRepository{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PackageRepository), err
}
