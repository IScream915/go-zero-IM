package implement

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	_interface "go-zero-IM/social/dao/query/interface"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var PreloadAssociations = []_interface.PreloadFunc{func(db *gorm.DB) *gorm.DB {
	return db.Preload(clause.Associations)
}}

func (d *DbToolHelper[T]) SearchRecordsWithPaginationAndMultiConditions(ctx context.Context, pageIndex, pageSize int, preloads []_interface.PreloadFunc, queryConditions map[string]interface{}, orderBy ...string) ([]*T, int64, error) {
	var (
		results []*T
		count   int64
	)

	model := new(T)
	db := d.DB.WithContext(ctx).Model(model)

	var startTimestamp, endTimestamp int64
	for key, value := range queryConditions {
		v := reflect.ValueOf(value)
		switch v.Kind() {
		case reflect.String:
			if v.String() != "" {
				db = db.Where(fmt.Sprintf("%s LIKE ?", key), "%"+v.String()+"%")
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if v.Uint() > 0 {
				db = db.Where(fmt.Sprintf("%s = ?", key), v.Uint())
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			if v.Int() > 0 {
				db = db.Where(fmt.Sprintf("%s = ?", key), v.Int())
			}
		case reflect.Int64:
			if v.Int() > 0 {
				if strings.HasSuffix(key, "start") {
					startTimestamp = v.Int()
				} else if strings.HasSuffix(key, "end") {
					endTimestamp = v.Int()
				} else {
					db = db.Where(fmt.Sprintf("%s = ?", key), v.Int())
				}
			}
		default:
		}
	}

	s := time.Unix(startTimestamp, 0).Format(time.DateTime + ".000")
	e := time.Unix(endTimestamp, 0).Format(time.DateTime + ".999")
	if startTimestamp > 0 && endTimestamp > 0 {
		db = db.Where("created_at BETWEEN ? AND ?", s, e)
	} else if startTimestamp > 0 && endTimestamp == 0 {
		db = db.Where("created_at >= ?", s)
	} else if startTimestamp == 0 && endTimestamp > 0 {
		db = db.Where("created_at <= ?", e)
	}

	if len(preloads) > 0 {
		for _, preload := range preloads {
			db = preload(db)
		}
	}

	if len(orderBy) > 0 {
		db = db.Order(orderBy[0])
	}

	if err := db.
		Count(&count).
		Offset((pageIndex - 1) * pageSize).
		Limit(pageSize).
		Find(&results).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []*T{}, count, nil
		}

		return []*T{}, count, err
	}

	return results, count, nil
}
