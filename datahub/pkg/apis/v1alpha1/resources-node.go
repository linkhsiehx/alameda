package v1alpha1

import (
	DaoClusterStatus "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus"
	DaoClusterStatusInflux "github.com/containers-ai/alameda/datahub/pkg/dao/interfaces/clusterstatus/influxdb"
	AlamedaUtils "github.com/containers-ai/alameda/pkg/utils"
	ApiResources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"golang.org/x/net/context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/genproto/googleapis/rpc/status"
)

// CreateAlamedaNodes add node information to database
func (s *ServiceV1alpha1) CreateNodes(ctx context.Context, in *ApiResources.CreateNodesRequest) (*status.Status, error) {
	scope.Debug("Request received from CreateAlamedaNodes grpc function: " + AlamedaUtils.InterfaceToString(in))

	var nodeDAO DaoClusterStatus.NodeOperation = &DaoClusterStatusInflux.Node{
		InfluxDBConfig: *s.Config.InfluxDB,
	}
	if err := nodeDAO.RegisterAlamedaNodes(in.GetNodes()); err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}

/*
// ListAlamedaNodes list nodes in cluster
func (s *ServiceV1alpha1) ListAlamedaNodes(ctx context.Context, in *ApiResources.ListAlamedaNodesRequest) (*ApiResources.ListNodesResponse, error) {
	scope.Debug("Request received from ListAlamedaNodes grpc function: " + AlamedaUtils.InterfaceToString(in))

	var nodeDAO DaoClusterStatus.NodeOperation = &DaoClusterStatusInflux.Node{
		InfluxDBConfig: *s.Config.InfluxDB,
	}

	timeRange := in.GetTimeRange()

	if alamedaNodes, err := nodeDAO.ListAlamedaNodes(timeRange); err != nil {
		scope.Error(err.Error())
		return &ApiResources.ListNodesResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	} else {
		return &ApiResources.ListNodesResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Nodes: alamedaNodes,
		}, nil
	}
}
*/

func (s *ServiceV1alpha1) ListNodes(ctx context.Context, in *ApiResources.ListNodesRequest) (*ApiResources.ListNodesResponse, error) {
	scope.Debug("Request received from ListNodes grpc function: " + AlamedaUtils.InterfaceToString(in))

	var nodeDAO DaoClusterStatus.NodeOperation = &DaoClusterStatusInflux.Node{
		InfluxDBConfig: *s.Config.InfluxDB,
	}

	nodeNames := make([]string, 0)
	for _, objectMeta := range in.GetObjectMeta() {
		nodeNames = append(nodeNames, objectMeta.Name)
	}

	req := DaoClusterStatus.ListNodesRequest{
		NodeNames: nodeNames,
		InCluster: true,
	}
	if nodes, err := nodeDAO.ListNodes(req); err != nil {
		scope.Error(err.Error())
		return &ApiResources.ListNodesResponse{
			Status: &status.Status{
				Code:    int32(code.Code_INTERNAL),
				Message: err.Error(),
			},
		}, nil
	} else {
		return &ApiResources.ListNodesResponse{
			Status: &status.Status{
				Code: int32(code.Code_OK),
			},
			Nodes: nodes,
		}, nil
	}
}

/*
// DeleteAlamedaNodes remove node information to database
func (s *ServiceV1alpha1) DeleteAlamedaNodes(ctx context.Context, in *ApiResources.DeleteAlamedaNodesRequest) (*status.Status, error) {
	scope.Debug("Request received from DeleteAlamedaNodes grpc function: " + AlamedaUtils.InterfaceToString(in))

	var nodeDAO DaoClusterStatus.NodeOperation = &DaoClusterStatusInflux.Node{
		InfluxDBConfig: *s.Config.InfluxDB,
	}
	alamedaNodeList := []*ApiResources.Node{}
	for _, alamedaNode := range in.GetAlamedaNodes() {
		alamedaNodeList = append(alamedaNodeList, &ApiResources.Node{
			Name: alamedaNode.GetName(),
		})
	}
	if err := nodeDAO.DeregisterAlamedaNodes(alamedaNodeList); err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}
*/

// DeleteAlamedaNodes remove node information to database
func (s *ServiceV1alpha1) DeleteNodes(ctx context.Context, in *ApiResources.DeleteNodesRequest) (*status.Status, error) {
	scope.Debug("Request received from DeleteAlamedaNodes grpc function: " + AlamedaUtils.InterfaceToString(in))

	var nodeDAO DaoClusterStatus.NodeOperation = &DaoClusterStatusInflux.Node{
		InfluxDBConfig: *s.Config.InfluxDB,
	}
	nodeList := make([]*ApiResources.Node, 0)
	for _, node := range in.GetNodes() {
		nodeList = append(nodeList, &ApiResources.Node{
			ObjectMeta: &ApiResources.ObjectMeta{
				Name: node.GetObjectMeta().GetName(),
			},
		})
	}
	if err := nodeDAO.DeregisterAlamedaNodes(nodeList); err != nil {
		scope.Error(err.Error())
		return &status.Status{
			Code:    int32(code.Code_INTERNAL),
			Message: err.Error(),
		}, nil
	}

	return &status.Status{
		Code: int32(code.Code_OK),
	}, nil
}