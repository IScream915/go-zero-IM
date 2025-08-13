package implement

import (
	"context"
)

// Delete 删除记录
func (d *DbToolHelper[T]) Delete(ctx context.Context, condition interface{}, args ...interface{}) error {
	if err := d.DB.
		WithContext(ctx).
		Where(condition, args...).
		Delete(new(T)).Error; err != nil {
		return err
	}

	return nil
}
