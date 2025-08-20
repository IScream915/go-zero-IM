package implement

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// SearchFuzzyGroupByField 按字段模糊搜索并返回该字段列表(group by)
func (d *DbToolHelper[T]) SearchFuzzyGroupByField(ctx context.Context, field, keyword string) ([]string, error) {
	var result []string

	model := new(T)
	db := d.DB.WithContext(ctx).Model(model)
	if field != "" && keyword != "" {
		db = db.Where(fmt.Sprintf("`%s` LIKE '%s'", field, "%"+keyword+"%"))
	}

	if err := db.
		Select(field).
		Group(field).
		Pluck(field, &result).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []string{}, nil
		}

		return []string{}, nil
	}

	return result, nil
}
