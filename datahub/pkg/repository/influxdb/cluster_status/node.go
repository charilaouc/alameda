package clusterstatus

import (
	"fmt"
	"strings"
	"time"

	cluster_status_dao "github.com/containers-ai/alameda/datahub/pkg/dao/cluster_status"
	cluster_status_entity "github.com/containers-ai/alameda/datahub/pkg/entity/influxdb/cluster_status"
	"github.com/containers-ai/alameda/datahub/pkg/repository/influxdb"
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/pkg/utils/log"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	influxdb_client "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
)

var (
	scope = log.RegisterScope("influxdb_repo_node_measurement", "InfluxDB repository node measurement", 0)
)

type NodeRepository struct {
	influxDB *influxdb.InfluxDBRepository
}

func (nodeRepository *NodeRepository) IsTag(column string) bool {
	for _, tag := range cluster_status_entity.NodeTags {
		if column == string(tag) {
			return true
		}
	}
	return false
}

func NewNodeRepository(influxDBCfg *influxdb.Config) *NodeRepository {
	return &NodeRepository{
		influxDB: &influxdb.InfluxDBRepository{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

// AddAlamedaNodes add node information to database
func (nodeRepository *NodeRepository) AddAlamedaNodes(alamedaNodes []*datahub_v1alpha1.Node) error {
	points := []*influxdb_client.Point{}
	for _, alamedaNode := range alamedaNodes {
		isInCluster := true
		entity := cluster_status_entity.NodeEntity{
			Time:        influxdb.ZeroTime,
			Name:        &alamedaNode.Name,
			IsInCluster: &isInCluster,
		}
		if nodeCapacity := alamedaNode.GetCapacity(); nodeCapacity != nil {
			entity.CPUCores = &nodeCapacity.CpuCores
			entity.MemoryBytes = &nodeCapacity.MemoryBytes
		}
		if pt, err := entity.InfluxDBPoint(string(Node)); err == nil {
			points = append(points, pt)
		} else {
			scope.Error(err.Error())
		}
	}
	err := nodeRepository.influxDB.WritePoints(points, influxdb_client.BatchPointsConfig{
		Database: string(influxdb.ClusterStatus),
	})
	if err != nil {
		return errors.Wrapf(err, "add alameda nodes failed: %s", err.Error())
	}
	return nil
}

func (nodeRepository *NodeRepository) RemoveAlamedaNodes(alamedaNodes []*datahub_v1alpha1.Node) error {
	points := []*influxdb_client.Point{}
	for _, alamedaNode := range alamedaNodes {
		// SELECT * FROM node WHERE "name"='%s' AND in_cluster=true ORDER BY time ASC LIMIT 1
		cmd := fmt.Sprintf("SELECT * FROM %s WHERE \"%s\"='%s' AND \"%s\"=%s ORDER BY time ASC LIMIT 1",
			string(Node), string(cluster_status_entity.NodeName), alamedaNode.Name,
			string(cluster_status_entity.NodeInCluster), "true")
		if results, _ := nodeRepository.influxDB.QueryDB(cmd, string(influxdb.ClusterStatus)); len(results) == 1 && len(results[0].Series) == 1 {
			newFields := map[string]interface{}{}
			newTags := map[string]string{}
			originTime := time.Now()
			for columnIdx, column := range results[0].Series[0].Columns {
				colVal := results[0].Series[0].Values[0][columnIdx]
				if column == influxdb.Time && colVal != nil {
					originTime, _ = utils.ParseTime(colVal.(string))
				} else if nodeRepository.IsTag(column) && column != influxdb.Time && colVal != nil {
					newTags[column] = colVal.(string)
				} else if !nodeRepository.IsTag(column) {
					if column == string(cluster_status_entity.NodeInCluster) {
						newFields[column] = false
					} else {
						newFields[column] = results[0].Series[0].Values[0][columnIdx]
					}
				}
			}

			if pt, err := influxdb_client.NewPoint(string(Node), newTags, newFields, originTime); err == nil {
				points = append(points, pt)
			} else {
				scope.Error(err.Error())
			}
		}
	}
	nodeRepository.influxDB.WritePoints(points, influxdb_client.BatchPointsConfig{
		Database: string(influxdb.ClusterStatus),
	})
	return nil
}

func (nodeRepository *NodeRepository) ListAlamedaNodes() ([]*cluster_status_entity.NodeEntity, error) {

	nodeEntities := []*cluster_status_entity.NodeEntity{}
	cmd := fmt.Sprintf("SELECT * FROM %s WHERE \"%s\"=%s", string(Node), string(cluster_status_entity.NodeInCluster), "true")
	scope.Infof(fmt.Sprintf("Query nodes in cluster: %s", cmd))
	results, err := nodeRepository.influxDB.QueryDB(cmd, string(influxdb.ClusterStatus))
	if err != nil {
		return nodeEntities, errors.Wrap(err, "list alameda nodes from influxdb failed")
	}

	if len(results) == 1 && len(results[0].Series) > 0 {

		influxdbRows := influxdb.PackMap(results)
		for _, influxdbRow := range influxdbRows {
			for _, data := range influxdbRow.Data {
				nodeEntity := cluster_status_entity.NewNodeEntityFromMap(data)
				nodeEntities = append(nodeEntities, &nodeEntity)
			}
		}
	}

	return nodeEntities, nil
}

func (nodeRepository *NodeRepository) ListNodes(request cluster_status_dao.ListNodesRequest) ([]*cluster_status_entity.NodeEntity, error) {

	nodeEntities := []*cluster_status_entity.NodeEntity{}

	whereClause := nodeRepository.buildInfluxQLWhereClauseFromRequest(request)
	cmd := fmt.Sprintf("SELECT * FROM %s %s", string(Node), whereClause)
	scope.Debug(fmt.Sprintf("Query nodes in cluster: %s", cmd))
	results, err := nodeRepository.influxDB.QueryDB(cmd, string(influxdb.ClusterStatus))
	if err != nil {
		return nodeEntities, errors.Wrap(err, "list nodes from influxdb failed")
	}

	influxdbRows := influxdb.PackMap(results)
	for _, influxdbRow := range influxdbRows {
		for _, data := range influxdbRow.Data {
			nodeEntity := cluster_status_entity.NewNodeEntityFromMap(data)
			nodeEntities = append(nodeEntities, &nodeEntity)
		}
	}

	return nodeEntities, nil
}

func (nodeRepository *NodeRepository) buildInfluxQLWhereClauseFromRequest(request cluster_status_dao.ListNodesRequest) string {

	var (
		whereClause string
		conditions  string
	)

	conditions += fmt.Sprintf("\"%s\" = %t", cluster_status_entity.NodeInCluster, request.InCluster)

	statementFilteringNodes := ""
	for _, nodeName := range request.NodeNames {
		statementFilteringNodes += fmt.Sprintf(`"%s" = '%s' OR `, cluster_status_entity.NodeName, nodeName)
	}
	statementFilteringNodes = strings.TrimSuffix(statementFilteringNodes, "OR ")
	if statementFilteringNodes != "" {
		conditions = fmt.Sprintf("(%s) AND (%s)", conditions, statementFilteringNodes)
	}

	whereClause = fmt.Sprintf("WHERE %s", conditions)

	return whereClause
}