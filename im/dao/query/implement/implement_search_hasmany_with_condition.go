package implement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (d *DbToolHelper[T]) SearchHasManyRecordsWithCondition(ctx context.Context, field, keyword string) ([]*T, error) {
	var results []*T

	db := d.DB.WithContext(ctx)
	if field != "" && keyword != "" {
		db = db.Where(fmt.Sprintf("`%s` LIKE '%s'", field, "%"+keyword+"%"))
	}

	if err := db.
		Find(&results).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*T{}, nil
		}

		return []*T{}, err
	}

	return results, nil
}
