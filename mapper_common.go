package mapping

import (
	"context"
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
		if !hasOption {
			option = Simple{Target: field, Source: field}
		}

		if option.ignore() {
			continue
		}

		var sourceValue reflect.Value

		if option.source() == "" && option.defVal() != nil {
			sourceValue = reflect.ValueOf(option.defVal())
		} else {
			sourceValue = sourceValues.FieldByName(option.source())
			if !sourceValue.IsValid() {
				getLogger(ctx).Warnf("Source value is not found for field: %s", field)
				continue
			}
		}

		if option.qualifier() != nil {
			qualified, err := option.qualifier()(sourceValue.Interface())
			if err != nil {
				getLogger(ctx).Errorf("Casting error: %v", err)
				continue
			}
			sourceValue = reflect.ValueOf(qualified)
		}

		if sourceValue.IsZero() && option.defVal() != nil {
			sourceValue = reflect.ValueOf(option.defVal())
		}

		targetValue := targetValues.Elem().FieldByName(field)

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
