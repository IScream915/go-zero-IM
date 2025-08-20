package implement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (d *DbToolHelper[T]) SearchHasManyWithChildPreloadAndFuzzy(ctx context.Context, preloadFiled, field, keyword string) ([]*T, error) {
	var results []*T
	if err := d.DB.
		WithContext(ctx).
		Preload(preloadFiled, func(db *gorm.DB) *gorm.DB {
			if field != "" && keyword != "" {
				return db.Where(fmt.Sprintf("`%s` LIKE '%s'", field, "%"+keyword+"%"))
			}

			return db
		}).
		Find(&results).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*T{}, nil
		}

		return []*T{}, err
	}

	return results, nil
}
