package test

import (
	controllers "github.com/liqotech/liqo/internal/liqonet"
	"time"
)

const (
	Namespace            = "test"
	NattedNamespace      = Namespace + "-" + HomeClusterId
	HostName             = "testHost"
	NodeName             = "testNode"
	AdvName              = "advertisement-" + ForeignClusterId
	TepName              = controllers.TunEndpointNamePrefix + ForeignClusterId
	EndpointsName        = "testEndpoints"
	HomeClusterId        = "homeClusterID"
	ForeignClusterId     = "foreignClusterID"
	LocalRemappedPodCIDR = "100.200.0.0/16"
	Timeout              = 10 * time.Second
)
