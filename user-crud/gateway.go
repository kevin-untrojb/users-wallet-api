package user_crud

import (
	"context"

)

//go:generate mockgen -destination=mock_gateway.go -package=user_crud -source=gateway.go Gateway

type Gateway interface {
	Create(context.Context, user) error

	Get(context.Context, string) error

}

type gateway struct {
	dao
}
