package app

import "flag"

var (
	nodeName    = flag.String("nodeName", "node_default", "Node name in cluster")
	serviceName = flag.String("serviceName", "", "Name of service")
	groupName   = flag.String("groupName", "", "Name of application")
)
