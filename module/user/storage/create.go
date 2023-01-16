package storageuser

import (
	"be-music/common"
	"be-music/module/user/model"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, user *model.UserCreate) error {
	if err := s.db.Create(&user).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
