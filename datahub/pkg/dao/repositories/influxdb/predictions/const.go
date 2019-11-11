package predictions

import (
	"github.com/containers-ai/alameda/internal/pkg/database/influxdb"
)

const (
	Node        influxdb.Measurement = "node"
	Container   influxdb.Measurement = "container"
	Controller  influxdb.Measurement = "controller"
	Application influxdb.Measurement = "application"
	Namespace   influxdb.Measurement = "namespace"
	Cluster     influxdb.Measurement = "cluster"
)
