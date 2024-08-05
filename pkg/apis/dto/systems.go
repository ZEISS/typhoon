package dto

import (
	"github.com/zeiss/pkg/cast"
	"github.com/zeiss/typhoon/internal/api/controllers"
	"github.com/zeiss/typhoon/internal/models"
	openapi "github.com/zeiss/typhoon/pkg/apis"
)

// FromCreateSystemRequest ...
func FromCreateSystemRequest(req openapi.CreateSystemRequestObject) controllers.CreateSystemCommand {
	return controllers.CreateSystemCommand{
		Name:        req.Body.Name,
		Description: cast.Value(req.Body.Description),
		OperatorID:  req.Body.OperatorId,
	}
}

// ToCreateSystemResponse ...
func ToCreateSystemResponse(system models.System) openapi.CreateSystem201JSONResponse {
	res := openapi.CreateSystem201JSONResponse{}
	res.Id = cast.Ptr(system.ID)
	res.Name = system.Name
	res.Description = cast.Ptr(system.Description)

	for _, cluster := range system.Clusters {
		res.Clusters = append(res.Clusters, openapi.Cluster{
			Id:          cast.Ptr(cluster.ID),
			Name:        cluster.Name,
			Description: cast.Ptr(cluster.Description),
			CreatedAt:   cast.Ptr(cluster.CreatedAt),
			UpdatedAt:   cast.Ptr(cluster.UpdatedAt),
			DeletedAt:   cast.Ptr(cluster.DeletedAt.Time),
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
	res.Id = cast.Ptr(system.ID)
	res.Name = system.Name
	res.Description = cast.Ptr(system.Description)
	res.CreatedAt = cast.Ptr(system.CreatedAt)
	res.UpdatedAt = cast.Ptr(system.UpdatedAt)
	res.DeletedAt = cast.Ptr(system.DeletedAt.Time)

	return res
}

// FromListSystemsRequest ...
func FromListSystemsRequest(req openapi.ListSystemsRequestObject) controllers.ListSystemsQuery {
	return controllers.ListSystemsQuery{
		Offset: cast.Value(req.Params.Offset),
		Limit:  cast.Value(req.Params.Limit),
	}
}

// ToListSystemsResponse ...
func ToListSystemsResponse(sys models.Pagination[models.System]) openapi.ListSystems200JSONResponse {
	res := openapi.ListSystems200JSONResponse{}
	res.Limit = cast.Ptr(sys.Limit)
	res.Offset = cast.Ptr(sys.Offset)
	res.Total = cast.Ptr(sys.TotalRows)

	systems := make([]openapi.System, 0, len(sys.Rows))
	for _, sys := range sys.Rows {
		systems = append(systems, openapi.System{
			Id:          cast.Ptr(sys.ID),
			Name:        sys.Name,
			Description: cast.Ptr(sys.Description),
			CreatedAt:   cast.Ptr(sys.CreatedAt),
			UpdatedAt:   cast.Ptr(sys.UpdatedAt),
			DeletedAt:   cast.Ptr(sys.DeletedAt.Time),
		})
	}
	res.Results = &systems

	return res
}
