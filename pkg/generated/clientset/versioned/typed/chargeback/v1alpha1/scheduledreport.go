package v1alpha1

import (
	v1alpha1 "github.com/coreos-inc/kube-chargeback/pkg/apis/chargeback/v1alpha1"
	scheme "github.com/coreos-inc/kube-chargeback/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScheduledReportsGetter has a method to return a ScheduledReportInterface.
// A group's client should implement this interface.
type ScheduledReportsGetter interface {
	ScheduledReports(namespace string) ScheduledReportInterface
}

// ScheduledReportInterface has methods to work with ScheduledReport resources.
type ScheduledReportInterface interface {
	Create(*v1alpha1.ScheduledReport) (*v1alpha1.ScheduledReport, error)
	Update(*v1alpha1.ScheduledReport) (*v1alpha1.ScheduledReport, error)
	UpdateStatus(*v1alpha1.ScheduledReport) (*v1alpha1.ScheduledReport, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ScheduledReport, error)
	List(opts v1.ListOptions) (*v1alpha1.ScheduledReportList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScheduledReport, err error)
	ScheduledReportExpansion
}

// scheduledReports implements ScheduledReportInterface
type scheduledReports struct {
	client rest.Interface
	ns     string
}

// newScheduledReports returns a ScheduledReports
func newScheduledReports(c *ChargebackV1alpha1Client, namespace string) *scheduledReports {
	return &scheduledReports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scheduledReport, and returns the corresponding scheduledReport object, and an error if there is any.
func (c *scheduledReports) Get(name string, options v1.GetOptions) (result *v1alpha1.ScheduledReport, err error) {
	result = &v1alpha1.ScheduledReport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScheduledReports that match those selectors.
func (c *scheduledReports) List(opts v1.ListOptions) (result *v1alpha1.ScheduledReportList, err error) {
	result = &v1alpha1.ScheduledReportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scheduledreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scheduledReports.
func (c *scheduledReports) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scheduledreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a scheduledReport and creates it.  Returns the server's representation of the scheduledReport, and an error, if there is any.
func (c *scheduledReports) Create(scheduledReport *v1alpha1.ScheduledReport) (result *v1alpha1.ScheduledReport, err error) {
	result = &v1alpha1.ScheduledReport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scheduledreports").
		Body(scheduledReport).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scheduledReport and updates it. Returns the server's representation of the scheduledReport, and an error, if there is any.
func (c *scheduledReports) Update(scheduledReport *v1alpha1.ScheduledReport) (result *v1alpha1.ScheduledReport, err error) {
	result = &v1alpha1.ScheduledReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledreports").
		Name(scheduledReport.Name).
		Body(scheduledReport).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *scheduledReports) UpdateStatus(scheduledReport *v1alpha1.ScheduledReport) (result *v1alpha1.ScheduledReport, err error) {
	result = &v1alpha1.ScheduledReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scheduledreports").
		Name(scheduledReport.Name).
		SubResource("status").
		Body(scheduledReport).
		Do().
		Into(result)
	return
}

// Delete takes name of the scheduledReport and deletes it. Returns an error if one occurs.
func (c *scheduledReports) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledreports").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scheduledReports) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scheduledreports").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched scheduledReport.
func (c *scheduledReports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScheduledReport, err error) {
	result = &v1alpha1.ScheduledReport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scheduledreports").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
