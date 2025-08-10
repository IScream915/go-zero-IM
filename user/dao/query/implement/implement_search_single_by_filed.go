package implement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (d *DbToolHelper[T]) SearchSingleByField(ctx context.Context, field string, keyword interface{}, preload ...string) (*T, error) {
	var result T

	db := d.DB.WithContext(ctx)
	if len(preload) > 0 {
		for i := 0; i < len(preload); i++ {
			db = db.Preload(preload[i])
		}
	}

	if field != "" && keyword != nil {
		db = db.Where(fmt.Sprintf("%s = ?", field), keyword)
	}
	if err := db.
		First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}
