package services

import (
	"context"
	"net/http"

	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// ListSystems ...
func (a *ApiHandlers) ListSystems(ctx context.Context, req openapi.ListSystemsRequestObject) (openapi.ListSystemsResponseObject, error) {
	pagination := models.Pagination[models.System]{}

	result, err := a.systems.ListSystems(ctx, pagination)
	if err != nil {
		return nil, err
	}

	res := openapi.ListSystems200JSONResponse{
		Limit:  utils.PtrInt(result.Limit),
		Offset: utils.PtrInt(result.Offset),
		Total:  utils.PtrInt(result.TotalRows),
	}

	systems := []openapi.System{}
	for _, system := range result.Rows {
		sys := openapi.System{
			Id:          &system.ID,
			Name:        system.Name,
			Description: utils.StrPtr(system.Description),
			CreatedAt:   &system.CreatedAt,
			UpdatedAt:   &system.UpdatedAt,
			DeletedAt:   &system.DeletedAt.Time,
		}

		for _, cluster := range system.Clusters {
			sys.Clusters = append(sys.Clusters, openapi.Cluster{
				Name:        cluster.Name,
				Description: &cluster.Description,
				ServerURL:   cluster.ServerURL,
				CreatedAt:   utils.PtrTime(cluster.CreatedAt),
				DeletedAt:   utils.PtrTime(cluster.DeletedAt.Time),
				UpdatedAt:   utils.PtrTime(cluster.UpdatedAt),
			})
		}

		systems = append(systems, sys)
	}
	res.Results = &systems

	return openapi.ListSystems200JSONResponse(res), nil
}

// CreateSystem ...
func (a *ApiHandlers) CreateSystem(ctx context.Context, req openapi.CreateSystemRequestObject) (openapi.CreateSystemResponseObject, error) {
	system := &models.System{}
	system.Name = req.Body.Name
	system.OperatorID = utils.PtrUUID(req.Body.OperatorId)
	system.Clusters = []models.Cluster{}

	for _, cluster := range req.Body.Clusters {
		system.Clusters = append(system.Clusters, models.Cluster{
			Name:        cluster.Name,
			Description: utils.PtrStr(cluster.Description),
			ServerURL:   cluster.ServerURL,
		})
	}

	system, err := a.systems.CreateSystem(ctx, system)
	if err != nil {
		return nil, err
	}

	res := openapi.System{
		Id:          utils.PtrUUID(system.ID),
		Name:        system.Name,
		Description: utils.StrPtr(system.Description),
	}

	return openapi.CreateSystem201JSONResponse(res), nil
}

// GetSystemOperator ...
func (a *ApiHandlers) GetSystemOperator(ctx context.Context, req openapi.GetSystemOperatorRequestObject) (openapi.GetSystemOperatorResponseObject, error) {
	system, err := a.systems.GetSystem(ctx, req.SystemId)
	if err != nil {
		return nil, err
	}

	if system.OperatorID == nil {
		return openapi.GetSystemOperatordefaultJSONResponse(openapi.GetGroupdefaultJSONResponse{StatusCode: http.StatusNotFound, Body: openapi.ErrorNotFound("could not find an operator")}), nil
	}

	return openapi.GetSystemOperator200JSONResponse(openapi.Operator{Id: utils.PtrUUID(system.Operator.ID), Name: system.Operator.Name}), nil
}
