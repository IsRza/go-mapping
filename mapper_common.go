package mapping

import (
	"context"
	"fmt"
	"reflect"
)

func mapping(
	ctx context.Context,
	targetTypes reflect.Type,
	targetValues reflect.Value,
	sourceValues reflect.Value,
	options map[string]Option,
) {
	for i := 0; i < targetTypes.NumField(); i++ {
		field := targetTypes.Field(i).Name

		option, hasOption := options[field]
		if hasOption && option.ignore() {
			continue
		}

		var sourceValue reflect.Value
		var usingCustomSourceName = false
		if hasOption && option.source() != "" {
			sourceValue = sourceValues.FieldByName(option.source())
			usingCustomSourceName = true
		} else {
			sourceValue = sourceValues.FieldByName(field)
		}

		if !sourceValue.IsValid() {
			message := fmt.Sprintf("Source value is not found for field: %s", field)
			if usingCustomSourceName {
				getLogger(ctx).Error(message)
			} else {
				getLogger(ctx).Warn(message)
			}
			continue
		}

		targetValue := targetValues.Elem().FieldByName(field)

		if hasOption && option.qualifier() != nil {
			qualified, err := option.qualifier()(sourceValue.Interface())
			if err != nil {
				getLogger(ctx).Errorf("Casting error: %v", err)
				continue
			}
			sourceValue = reflect.ValueOf(qualified)
		}

		if sourceValue.Type().AssignableTo(targetValue.Type()) {
			targetValue.Set(sourceValue)
		} else {
			getLogger(ctx).Errorf(
				"Cannot update field %s(%s) with type %s, consider using a qualifier",
				field, targetValue.Type(), sourceValue.Type(),
			)
		}
	}
}
