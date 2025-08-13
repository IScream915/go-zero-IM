package implement

import "context"

func (d *DbToolHelper[T]) SearchCount(ctx context.Context, whereCondition string) (int64, error) {
	var count int64

	model := new(T)
	db := d.DB.WithContext(ctx).Model(model)
	if whereCondition != "" {
		db = db.Where(whereCondition)
	}

	if err := db.
		Count(&count).
		Error; err != nil {
		return count, err
	}

	return count, nil
}
