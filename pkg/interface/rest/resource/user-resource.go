package resource

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/pangud/pangud/pkg/shared/logger"
	"go.uber.org/zap"
)

type UserResource struct {
	log *zap.Logger
}

func NewUserResource() *UserResource {
	return &UserResource{
		log: logger.New(),
	}
}
func (ur *UserResource) WebService() *restful.WebService {
	ws := &restful.WebService{}
	ws.Path("/users")
	ws.Route(ws.GET("").To(ur.listAll))
	return ws
}
func (ur *UserResource) listAll(req *restful.Request, res *restful.Response) {
	ur.log.Debug("list all")
	ur.log.Info("list all")

	res.WriteAsJson([]string{"users"})
}
