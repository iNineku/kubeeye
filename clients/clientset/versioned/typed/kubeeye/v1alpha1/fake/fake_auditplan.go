/*
Copyright 2022.

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

	v1alpha1 "github.com/kubesphere/kubeeye/api/kubeeye/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAuditPlans implements AuditPlanInterface
type FakeAuditPlans struct {
	Fake *FakeKubeeyeV1alpha1
	ns   string
}

var auditplansResource = schema.GroupVersionResource{Group: "kubeeye.kubesphere.io", Version: "v1alpha1", Resource: "auditplans"}

var auditplansKind = schema.GroupVersionKind{Group: "kubeeye.kubesphere.io", Version: "v1alpha1", Kind: "AuditPlan"}

// Get takes name of the auditPlan, and returns the corresponding auditPlan object, and an error if there is any.
func (c *FakeAuditPlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AuditPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(auditplansResource, c.ns, name), &v1alpha1.AuditPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditPlan), err
}

// List takes label and field selectors, and returns the list of AuditPlans that match those selectors.
func (c *FakeAuditPlans) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AuditPlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(auditplansResource, auditplansKind, c.ns, opts), &v1alpha1.AuditPlanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AuditPlanList{ListMeta: obj.(*v1alpha1.AuditPlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.AuditPlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested auditPlans.
func (c *FakeAuditPlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(auditplansResource, c.ns, opts))

}

// Create takes the representation of a auditPlan and creates it.  Returns the server's representation of the auditPlan, and an error, if there is any.
func (c *FakeAuditPlans) Create(ctx context.Context, auditPlan *v1alpha1.AuditPlan, opts v1.CreateOptions) (result *v1alpha1.AuditPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(auditplansResource, c.ns, auditPlan), &v1alpha1.AuditPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditPlan), err
}

// Update takes the representation of a auditPlan and updates it. Returns the server's representation of the auditPlan, and an error, if there is any.
func (c *FakeAuditPlans) Update(ctx context.Context, auditPlan *v1alpha1.AuditPlan, opts v1.UpdateOptions) (result *v1alpha1.AuditPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(auditplansResource, c.ns, auditPlan), &v1alpha1.AuditPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditPlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAuditPlans) UpdateStatus(ctx context.Context, auditPlan *v1alpha1.AuditPlan, opts v1.UpdateOptions) (*v1alpha1.AuditPlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(auditplansResource, "status", c.ns, auditPlan), &v1alpha1.AuditPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditPlan), err
}

// Delete takes name of the auditPlan and deletes it. Returns an error if one occurs.
func (c *FakeAuditPlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(auditplansResource, c.ns, name, opts), &v1alpha1.AuditPlan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAuditPlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(auditplansResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AuditPlanList{})
	return err
}

// Patch applies the patch and returns the patched auditPlan.
func (c *FakeAuditPlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AuditPlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(auditplansResource, c.ns, name, pt, data, subresources...), &v1alpha1.AuditPlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AuditPlan), err
}