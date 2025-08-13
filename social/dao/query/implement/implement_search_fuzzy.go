package implement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// SearchFuzzy 按字段模糊搜索
func (d *DbToolHelper[T]) SearchFuzzy(ctx context.Context, field, keyword string) ([]string, error) {
	var result []string

	model := new(T)
	db := d.DB.WithContext(ctx).Model(model)
	if field != "" && keyword != "" {
		db = db.Where(fmt.Sprintf("`%s` LIKE '%s'", field, "%"+keyword+"%"))
	}

	if err := db.
		Select(field).
		Pluck(field, &result).
		Order(fmt.Sprintf("%s DESC", field)).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []string{}, nil
		}

		return []string{}, nil
	}

	return result, nil
}
