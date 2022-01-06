package endpoint

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/google/wire"
	"github.com/pangud/pangud/pkg/infrastructure/mapper"
	"github.com/pangud/pangud/pkg/interface/rest/resource"
	"github.com/pangud/pangud/pkg/shared/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type RestfulEndpoint struct {
	DB           *gorm.DB
	UserResource *resource.UserResource
}

func NewRestfulEndpoint(db *gorm.DB, userResource *resource.UserResource) *RestfulEndpoint {
	return &RestfulEndpoint{
		DB:           db,
		UserResource: userResource,
	}
}

var RestfulEndpointSet = wire.NewSet(NewRestfulEndpoint, resource.ResourceSet, mapper.MapperSet)

func (e *RestfulEndpoint) Serve() {
	log := logger.New()
	defer log.Sync()
	restful.Add(e.UserResource.WebService())

	log.Fatal("serve error", zap.Error(http.ListenAndServe(":9000", restful.DefaultContainer)))
}
