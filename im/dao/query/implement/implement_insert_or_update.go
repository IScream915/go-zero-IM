package implement

import (
	"context"
)

// InsertOrUpdate 不存在则插入存在则更新
func (d *DbToolHelper[T]) InsertOrUpdate(ctx context.Context, model, model1, model2 T) error {
	if err := d.DB.
		WithContext(ctx).
		Where(model1).
		Assign(model2).
		FirstOrCreate(&model).Error; err != nil {
		return err
	}

	return nil
}
