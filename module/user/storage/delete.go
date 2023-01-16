package storageuser

import (
	"be-music/module/user/model"
	"context"
)

func (s *sqlStore) Delete(
	ctx context.Context,
	id int) error { // *model.User khong dung con tro o day thi tra ve struct rong ton bo nho

	if err := s.db.Table(model.User{}.
		TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
