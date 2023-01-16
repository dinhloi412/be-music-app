package storageuser

import (
	"be-music/module/user/model"
	"context"
)

func (s *sqlStore) FindUserWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	morekeys ...string) (*model.User, error) { // *model.User khong dung con tro o day thi tra ve struct rong ton bo nho
	var user model.User
	if err := s.db.Where(condition).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
