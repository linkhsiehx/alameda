package metric

import (
	"errors"
	"fmt"
	"time"

	"github.com/containers-ai/alameda/datahub/pkg/entity/prometheus/containerMemoryUsageBytes"
	"github.com/containers-ai/alameda/datahub/pkg/repository/prometheus"
)

// PodContainerMemoryUsageBytesRepository comment
type PodContainerMemoryUsageBytesRepository struct {
	PrometheusConfig prometheus.Config
}

func NewPodContainerMemoryUsageBytesRepositoryWithConfig(cfg prometheus.Config) PodContainerMemoryUsageBytesRepository {
	return PodContainerMemoryUsageBytesRepository{PrometheusConfig: cfg}
}

// ListMetricsByPodNamespacedName comment
func (c PodContainerMemoryUsageBytesRepository) ListMetricsByPodNamespacedName(namespace string, podName string, startTime, endTime time.Time) ([]prometheus.Entity, error) {

	var (
		err error

		prometheusClient *prometheus.Prometheus

		metricName        string
		queryLabelsString string
		queryExpression   string

		response prometheus.Response

		entities []prometheus.Entity
	)

	prometheusClient, err = prometheus.New(c.PrometheusConfig)
	if err != nil {
		return entities, errors.New("QueryRangeByPodNamespacedName failed: " + err.Error())
	}

	metricName = containerMemoryUsageBytes.MetricName
	queryLabelsString = c.buildQueryLabelsStringByNamespaceAndPodName(namespace, podName)

	if queryLabelsString != "" {
		queryExpression = fmt.Sprintf("%s{%s}", metricName, queryLabelsString)
	} else {
		queryExpression = fmt.Sprintf("%s", metricName)
	}

	response, err = prometheusClient.QueryRange(queryExpression, startTime, endTime)
	if err != nil {
		return entities, errors.New("QueryRangeByPodNamespacedName failed: " + err.Error())
	} else if response.Status != prometheus.SuccessStatus {
		return entities, errors.New("QueryRangeByPodNamespacedName failed: receive error response from prometheus: " + response.Error)
	}

	entities, err = response.GetEntitis()
	if err != nil {
		return entities, errors.New("ListMetricsByPodNamespacedName failed: " + err.Error())
	}

	return entities, nil
}

func (c PodContainerMemoryUsageBytesRepository) buildDefaultQueryLabelsString() string {

	var queryLabelsString = ""

	queryLabelsString += fmt.Sprintf(`%s != "" ,`, containerMemoryUsageBytes.PodLabelName)
	queryLabelsString += fmt.Sprintf(`%s != "" ,`, containerMemoryUsageBytes.ContainerLabel)
	queryLabelsString += fmt.Sprintf(`%s != "POD"`, containerMemoryUsageBytes.ContainerLabel)

	return queryLabelsString
}

func (c PodContainerMemoryUsageBytesRepository) buildQueryLabelsStringByNamespaceAndPodName(namespace string, podName string) string {

	var (
		queryLabelsString = c.buildDefaultQueryLabelsString()
	)

	if namespace != "" {
		queryLabelsString += fmt.Sprintf(`,%s = "%s"`, containerMemoryUsageBytes.NamespaceLabel, namespace)
	}

	if podName != "" {
		queryLabelsString += fmt.Sprintf(`,%s = "%s"`, containerMemoryUsageBytes.PodLabelName, podName)
	}

	return queryLabelsString
}
