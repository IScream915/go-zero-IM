package implement

import "context"

// InsertSingleRecord 插入单条记录
func (d *DbToolHelper[T]) InsertSingleRecord(ctx context.Context, model *T) error {
	if err := d.DB.
		WithContext(ctx).
		Create(&model).Error; err != nil {
		return err
	}

	return nil
}

// InsertSingleRecordAndReturn 插入单条记录并返回该记录
func (d *DbToolHelper[T]) InsertSingleRecordAndReturn(ctx context.Context, model *T) (*T, error) {
	if err := d.DB.
		WithContext(ctx).
		Create(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
