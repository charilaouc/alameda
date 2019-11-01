package recommendations

import (
	EntityInfluxRecommend "github.com/containers-ai/alameda/datahub/pkg/dao/entities/influxdb/recommendations"
	RepoInflux "github.com/containers-ai/alameda/datahub/pkg/dao/repositories/influxdb"
	DBCommon "github.com/containers-ai/alameda/internal/pkg/database/common"
	InternalInflux "github.com/containers-ai/alameda/internal/pkg/database/influxdb"
	ApiRecommendations "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/recommendations"
	ApiResources "github.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	InfluxClient "github.com/influxdata/influxdb/client/v2"
	"strconv"
	"time"
)

type ControllerRepository struct {
	influxDB *InternalInflux.InfluxClient
}

func NewControllerRepository(influxDBCfg *InternalInflux.Config) *ControllerRepository {
	return &ControllerRepository{
		influxDB: &InternalInflux.InfluxClient{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

func (c *ControllerRepository) CreateControllerRecommendations(controllerRecommendations []*ApiRecommendations.ControllerRecommendation) error {
	points := make([]*InfluxClient.Point, 0)
	for _, conrollerRecommendation := range controllerRecommendations {
		recommendedType := conrollerRecommendation.GetRecommendedType()

		if recommendedType == ApiRecommendations.ControllerRecommendedType_CRT_Primitive {
			recommendedSpec := conrollerRecommendation.GetRecommendedSpec()

			tags := map[string]string{
				EntityInfluxRecommend.ControllerNamespace: recommendedSpec.GetNamespacedName().GetNamespace(),
				EntityInfluxRecommend.ControllerName:      recommendedSpec.GetNamespacedName().GetName(),
				EntityInfluxRecommend.ControllerType:      ApiRecommendations.ControllerRecommendedType_CRT_Primitive.String(),
			}

			fields := map[string]interface{}{
				EntityInfluxRecommend.ControllerCurrentReplicas: recommendedSpec.GetCurrentReplicas(),
				EntityInfluxRecommend.ControllerDesiredReplicas: recommendedSpec.GetDesiredReplicas(),
				EntityInfluxRecommend.ControllerCreateTime:      recommendedSpec.GetCreateTime().GetSeconds(),
				EntityInfluxRecommend.ControllerKind:            recommendedSpec.GetKind().String(),

				EntityInfluxRecommend.ControllerCurrentCPURequest: recommendedSpec.GetCurrentCpuRequests(),
				EntityInfluxRecommend.ControllerCurrentMEMRequest: recommendedSpec.GetCurrentMemRequests(),
				EntityInfluxRecommend.ControllerCurrentCPULimit:   recommendedSpec.GetCurrentCpuLimits(),
				EntityInfluxRecommend.ControllerCurrentMEMLimit:   recommendedSpec.GetCurrentMemLimits(),
				EntityInfluxRecommend.ControllerDesiredCPULimit:   recommendedSpec.GetDesiredCpuLimits(),
				EntityInfluxRecommend.ControllerDesiredMEMLimit:   recommendedSpec.GetDesiredMemLimits(),
				EntityInfluxRecommend.ControllerTotalCost:         recommendedSpec.GetTotalCost(),
			}

			pt, err := InfluxClient.NewPoint(string(Controller), tags, fields, time.Unix(recommendedSpec.GetTime().GetSeconds(), 0))
			if err != nil {
				scope.Error(err.Error())
			}

			points = append(points, pt)

		} else if recommendedType == ApiRecommendations.ControllerRecommendedType_CRT_K8s {
			recommendedSpec := conrollerRecommendation.GetRecommendedSpecK8S()

			tags := map[string]string{
				EntityInfluxRecommend.ControllerNamespace: recommendedSpec.GetNamespacedName().GetNamespace(),
				EntityInfluxRecommend.ControllerName:      recommendedSpec.GetNamespacedName().GetName(),
				EntityInfluxRecommend.ControllerType:      ApiRecommendations.ControllerRecommendedType_CRT_K8s.String(),
			}

			fields := map[string]interface{}{
				EntityInfluxRecommend.ControllerCurrentReplicas: recommendedSpec.GetCurrentReplicas(),
				EntityInfluxRecommend.ControllerDesiredReplicas: recommendedSpec.GetDesiredReplicas(),
				EntityInfluxRecommend.ControllerCreateTime:      recommendedSpec.GetCreateTime().GetSeconds(),
				EntityInfluxRecommend.ControllerKind:            recommendedSpec.GetKind().String(),
			}

			pt, err := InfluxClient.NewPoint(string(Controller), tags, fields, time.Unix(recommendedSpec.GetTime().GetSeconds(), 0))
			if err != nil {
				scope.Error(err.Error())
			}

			points = append(points, pt)
		}
	}

	err := c.influxDB.WritePoints(points, InfluxClient.BatchPointsConfig{
		Database: string(RepoInflux.Recommendation),
	})

	if err != nil {
		scope.Error(err.Error())
		return err
	}

	return nil
}

func (c *ControllerRepository) ListControllerRecommendations(in *ApiRecommendations.ListControllerRecommendationsRequest) ([]*ApiRecommendations.ControllerRecommendation, error) {
	namespace := in.GetNamespacedName().GetNamespace()
	name := in.GetNamespacedName().GetName()
	recommendationType := in.GetRecommendedType()

	influxdbStatement := InternalInflux.Statement{
		Measurement:    Controller,
		QueryCondition: DBCommon.BuildQueryConditionV1(in.GetQueryCondition()),
	}

	influxdbStatement.AppendWhereClause(EntityInfluxRecommend.ControllerNamespace, "=", namespace)
	influxdbStatement.AppendWhereClause(EntityInfluxRecommend.ControllerName, "=", name)
	influxdbStatement.AppendWhereClauseFromTimeCondition()
	influxdbStatement.SetOrderClauseFromQueryCondition()
	influxdbStatement.SetLimitClauseFromQueryCondition()

	if recommendationType != ApiRecommendations.ControllerRecommendedType_CRT_Undefined {
		influxdbStatement.AppendWhereClause(EntityInfluxRecommend.ControllerType, "=", recommendationType.String())
	}

	cmd := influxdbStatement.BuildQueryCmd()

	results, err := c.influxDB.QueryDB(cmd, string(RepoInflux.Recommendation))
	if err != nil {
		return make([]*ApiRecommendations.ControllerRecommendation, 0), err
	}

	influxdbRows := InternalInflux.PackMap(results)
	recommendations := c.getControllersRecommendationsFromInfluxRows(influxdbRows)

	return recommendations, nil
}

func (c *ControllerRepository) getControllersRecommendationsFromInfluxRows(rows []*InternalInflux.InfluxRow) []*ApiRecommendations.ControllerRecommendation {
	recommendations := make([]*ApiRecommendations.ControllerRecommendation, 0)
	for _, influxdbRow := range rows {
		for _, data := range influxdbRow.Data {
			currentReplicas, _ := strconv.ParseInt(data[EntityInfluxRecommend.ControllerCurrentReplicas], 10, 64)
			desiredReplicas, _ := strconv.ParseInt(data[EntityInfluxRecommend.ControllerDesiredReplicas], 10, 64)
			createTime, _ := strconv.ParseInt(data[EntityInfluxRecommend.ControllerCreateTime], 10, 64)

			t, _ := time.Parse(time.RFC3339, data[EntityInfluxRecommend.ControllerTime])
			tempTime, _ := ptypes.TimestampProto(t)

			currentCpuRequests, _ := strconv.ParseFloat(data[EntityInfluxRecommend.ControllerCurrentCPURequest], 64)
			currentMemRequests, _ := strconv.ParseFloat(data[EntityInfluxRecommend.ControllerCurrentMEMRequest], 64)
			currentCpuLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.ControllerCurrentCPULimit], 64)
			currentMemLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.ControllerCurrentMEMLimit], 64)
			desiredCpuLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.ControllerDesiredCPULimit], 64)
			desiredMemLimits, _ := strconv.ParseFloat(data[EntityInfluxRecommend.ControllerDesiredMEMLimit], 64)
			totalCost, _ := strconv.ParseFloat(data[EntityInfluxRecommend.ControllerTotalCost], 64)

			var commendationType ApiRecommendations.ControllerRecommendedType
			if tempType, exist := data[EntityInfluxRecommend.ControllerType]; exist {
				if value, ok := ApiRecommendations.ControllerRecommendedType_value[tempType]; ok {
					commendationType = ApiRecommendations.ControllerRecommendedType(value)
				}
			}

			var commendationKind ApiResources.Kind
			if tempKind, exist := data[EntityInfluxRecommend.ControllerKind]; exist {
				if value, ok := ApiResources.Kind_value[tempKind]; ok {
					commendationKind = ApiResources.Kind(value)
				}
			}

			if commendationType == ApiRecommendations.ControllerRecommendedType_CRT_Primitive {
				tempRecommendation := &ApiRecommendations.ControllerRecommendation{
					RecommendedType: commendationType,
					RecommendedSpec: &ApiRecommendations.ControllerRecommendedSpec{
						NamespacedName: &ApiResources.NamespacedName{
							Namespace: data[string(EntityInfluxRecommend.ControllerNamespace)],
							Name:      data[string(EntityInfluxRecommend.ControllerName)],
						},
						CurrentReplicas: int32(currentReplicas),
						DesiredReplicas: int32(desiredReplicas),
						Time:            tempTime,
						CreateTime: &timestamp.Timestamp{
							Seconds: createTime,
						},
						Kind:               commendationKind,
						CurrentCpuRequests: currentCpuRequests,
						CurrentMemRequests: currentMemRequests,
						CurrentCpuLimits:   currentCpuLimits,
						CurrentMemLimits:   currentMemLimits,
						DesiredCpuLimits:   desiredCpuLimits,
						DesiredMemLimits:   desiredMemLimits,
						TotalCost:          totalCost,
					},
				}

				recommendations = append(recommendations, tempRecommendation)

			} else if commendationType == ApiRecommendations.ControllerRecommendedType_CRT_K8s {
				tempRecommendation := &ApiRecommendations.ControllerRecommendation{
					RecommendedType: commendationType,
					RecommendedSpecK8S: &ApiRecommendations.ControllerRecommendedSpecK8S{
						NamespacedName: &ApiResources.NamespacedName{
							Namespace: data[string(EntityInfluxRecommend.ControllerNamespace)],
							Name:      data[string(EntityInfluxRecommend.ControllerName)],
						},
						CurrentReplicas: int32(currentReplicas),
						DesiredReplicas: int32(desiredReplicas),
						Time:            tempTime,
						CreateTime: &timestamp.Timestamp{
							Seconds: createTime,
						},
						Kind: commendationKind,
					},
				}

				recommendations = append(recommendations, tempRecommendation)
			}
		}
	}

	return recommendations
}