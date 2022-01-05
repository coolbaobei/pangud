package query

import "github.com/pangud/pangud/pkg/applicaiton/query/dto"

type UserQueryService interface {
	listAll() (dto.UserDTO, error)
}
