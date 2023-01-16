package userbiz

import (
	"be-music/module/user/model"
	"context"
)

type CreateUserStore interface {
	CreateUser(ctx context.Context, user *model.UserCreate) error
}
type createUserBiz struct {
	store CreateUserStore
}

func NewCreateUserBiz(store CreateUserStore) *createUserBiz {
	return &createUserBiz{store: store}
}
func (biz *createUserBiz) CreateUser(ctx context.Context, user *model.UserCreate) error {
	if err := biz.store.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}
