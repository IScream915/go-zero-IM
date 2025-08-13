package implement

import (
	"context"
	"fmt"
)

func (d *DbToolHelper[T]) UpdateOneOrMultiFields(ctx context.Context, field string, keyword interface{}, updates map[string]interface{}) error {
	model := new(T)
	if err := d.DB.
		WithContext(ctx).
		Model(model).
		Where(fmt.Sprintf("%s = ?", field), keyword).
		Updates(updates).
		Error; err != nil {
		return err
	}

	return nil
}
