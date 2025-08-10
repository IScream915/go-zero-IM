package implement

import "context"

// InsertBatchRecords 插入多条记录
func (d *DbToolHelper[T]) InsertBatchRecords(ctx context.Context, models []T) error {
	if err := d.DB.
		WithContext(ctx).
		Create(&models).Error; err != nil {
		return err
	}

	return nil
}
