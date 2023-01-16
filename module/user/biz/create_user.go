package userbiz

import (
	"be-music/common"
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
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}

// validate
// func (user *model.UserCreate) validate() error {
// 	user.UserName = strings.TrimSpace(user.UserName)

// 	if user.UserName != "" {
// 		return common.ErrInvalidRequest(err)
// 	}
// 	return nil
// }
