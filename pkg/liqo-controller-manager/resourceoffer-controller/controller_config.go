package resourceoffercontroller

import (
	"time"

	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/liqotech/liqo/pkg/clusterid"
	"github.com/liqotech/liqo/pkg/vkMachinery/forge"
)

// NewResourceOfferController creates and returns a new reconciler for the ResourceOffers.
func NewResourceOfferController(
	mgr manager.Manager, clusterID clusterid.ClusterID,
	resyncPeriod time.Duration, liqoNamespace string,
	virtualKubeletOpts *forge.VirtualKubeletOpts) *ResourceOfferReconciler {
	return &ResourceOfferReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),

		eventsRecorder: mgr.GetEventRecorderFor("ResourceOffer"),
		clusterID:      clusterID,

		liqoNamespace: liqoNamespace,

		virtualKubeletOpts: virtualKubeletOpts,

		resyncPeriod: resyncPeriod,
	}
}
