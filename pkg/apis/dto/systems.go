package dto

import (
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/api/models"
	"github.com/zeiss/typhoon/internal/utils"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateSystemRequest ...
func FromCreateSystemRequest(req openapi.CreateSystemRequestObject) controllers.CreateSystemCommand {
	return controllers.CreateSystemCommand{
		Name:        req.Body.Name,
		Description: utils.PtrStr(req.Body.Description),
		OperatorID:  req.Body.OperatorId,
	}
}

// ToCreateSystemResponse ...
func ToCreateSystemResponse(system models.System) openapi.CreateSystem201JSONResponse {
	res := openapi.CreateSystem201JSONResponse{}
	res.Id = utils.PtrUUID(system.ID)
	res.Name = system.Name
	res.Description = utils.StrPtr(system.Description)

	for _, cluster := range system.Clusters {
		res.Clusters = append(res.Clusters, openapi.Cluster{
			Id:          utils.PtrUUID(cluster.ID),
			Name:        cluster.Name,
			Description: utils.StrPtr(cluster.Description),
			CreatedAt:   utils.PtrTime(cluster.CreatedAt),
			UpdatedAt:   utils.PtrTime(cluster.UpdatedAt),
			DeletedAt:   utils.PtrTime(cluster.DeletedAt.Time),
		})
	}

	return res
}

// FromDeleteSystemRequest ...
func FromDeleteSystemRequest(req openapi.DeleteSystemRequestObject) controllers.DeleteSystemCommand {
	return controllers.DeleteSystemCommand{
		ID: req.SystemId,
	}
}

// ToDeleteSystemResponse ...
func ToDeleteSystemResponse() openapi.DeleteSystem204Response {
	return openapi.DeleteSystem204Response{}
}

// FromGetSystemRequest ...
func FromGetSystemRequest(req openapi.GetSystemRequestObject) controllers.GetSystemQuery {
	return controllers.GetSystemQuery{
		ID: req.SystemId,
	}
}

// ToGetSystemResponse ...
func ToGetSystemResponse(system models.System) openapi.GetSystem200JSONResponse {
	res := openapi.GetSystem200JSONResponse{}
	res.Id = utils.PtrUUID(system.ID)
	res.Name = system.Name
	res.Description = utils.StrPtr(system.Description)
	res.CreatedAt = utils.PtrTime(system.CreatedAt)
	res.UpdatedAt = utils.PtrTime(system.UpdatedAt)
	res.DeletedAt = utils.PtrTime(system.DeletedAt.Time)

	return res
}

// FromListSystemsRequest ...
func FromListSystemsRequest(req openapi.ListSystemsRequestObject) controllers.ListSystemsQuery {
	return controllers.ListSystemsQuery{
		Offset: utils.IntPtr(req.Params.Offset),
		Limit:  utils.IntPtr(req.Params.Limit),
	}
}

// ToListSystemsResponse ...
func ToListSystemsResponse(sys models.Pagination[models.System]) openapi.ListSystems200JSONResponse {
	res := openapi.ListSystems200JSONResponse{}
	res.Limit = utils.PtrInt(sys.Limit)
	res.Offset = utils.PtrInt(sys.Offset)
	res.Total = utils.PtrInt(sys.TotalRows)

	systems := make([]openapi.System, 0, len(sys.Rows))
	for _, sys := range sys.Rows {
		systems = append(systems, openapi.System{
			Id:          utils.PtrUUID(sys.ID),
			Name:        sys.Name,
			Description: utils.StrPtr(sys.Description),
			CreatedAt:   utils.PtrTime(sys.CreatedAt),
			UpdatedAt:   utils.PtrTime(sys.UpdatedAt),
			DeletedAt:   utils.PtrTime(sys.DeletedAt.Time),
		})
	}
	res.Results = &systems

	return res
}
