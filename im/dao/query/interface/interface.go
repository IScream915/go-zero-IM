package _interface

import (
	"context"

	"gorm.io/gorm"
)

type PreloadFunc func(*gorm.DB) *gorm.DB

type CommonQuery[T any] interface {
	// SearchFuzzy 按字段模糊搜索并返回该字段列表
	SearchFuzzy(ctx context.Context, field, keyword string) ([]string, error)

	// SearchFuzzyGroupByField 按字段模糊搜索并返回该字段列表(group by)
	SearchFuzzyGroupByField(ctx context.Context, field, keyword string) ([]string, error)

	// SearchCount 统计数量
	SearchCount(ctx context.Context, whereCondition string) (int64, error)

	// SearchHasManyWithChildPreloadAndFuzzy 一对多关联查询，同时支持子表preload + 子表指定字段模糊搜索
	SearchHasManyWithChildPreloadAndFuzzy(ctx context.Context, preloadFiled, field, keyword string) ([]*T, error)

	// SearchHasManyRecordsWithCondition 一对多查询，支持指定条件
	SearchHasManyRecordsWithCondition(ctx context.Context, field, keyword string) ([]*T, error)

	// SearchSingleByField 按指定字段查询并返回单条记录
	SearchSingleByField(ctx context.Context, field string, keyword interface{}, preload ...string) (*T, error)

	// SearchRecordsWithPaginationAndMultiConditions 一对多+分页+多条件组合查询并返回所有记录
	SearchRecordsWithPaginationAndMultiConditions(ctx context.Context, pageIndex, pageSize int, preloads []PreloadFunc, queryConditions map[string]interface{}, orderBy ...string) ([]*T, int64, error)

	// SearchManyToManyRecordsWithPaginationAndMultiConditions 多对多+分页+多条件组合查询并返回所有记录
	SearchManyToManyRecordsWithPaginationAndMultiConditions(ctx context.Context, pageIndex, pageSize int, preloads []PreloadFunc, queryConditions map[string]interface{}, joinsCondition, whereCondition string, orderBy ...string) ([]*T, int64, error)

	// InsertSingleRecord 插入单条记录
	InsertSingleRecord(ctx context.Context, model *T) error

	// InsertBatchRecords 插入多条记录
	InsertBatchRecords(ctx context.Context, models []T) error

	// InsertSingleRecordAndReturn 插入单条记录并返回该记录
	InsertSingleRecordAndReturn(ctx context.Context, model *T) (*T, error)

	// InsertOrUpdate 不存在则插入存在则更新
	InsertOrUpdate(ctx context.Context, model, model1, model2 T) error

	// UpdateOneOrMultiFields 更新一个或多个字段
	UpdateOneOrMultiFields(ctx context.Context, field string, keyword interface{}, updates map[string]interface{}) error

	// Delete 删除记录
	Delete(ctx context.Context, condition interface{}, args ...interface{}) error
}
