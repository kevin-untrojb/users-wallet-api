package user_crud

import "context"

type Dao interface {
	InsertUser(context.Context, user) (int64, error)

	GetUser(context.Context, string) (user,error)
}