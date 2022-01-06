package query

import "github.com/pangud/pangud/pkg/application/query/dto"

type UserQueryService interface {
	listAll() (dto.UserDTO, error)
}
