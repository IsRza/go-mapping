package mapping

import (
	"context"
	"fmt"
	"reflect"
)

func Update[T, S any](target *T, source S, optionsArr ...Option) error {
	return UpdateContext(nil, target, source, optionsArr...)
}

func UpdateContext[T, S any](ctx context.Context, target *T, source S, optionsArr ...Option) error {
	options := optionsArrToMap(ctx, optionsArr)

	var skeleton T
	targetTypes := reflect.TypeOf(skeleton)
	if targetTypes.Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct not '%s'", targetTypes.Kind())
	}

	sourceTypes := reflect.TypeOf(source)
	if sourceTypes.Kind() != reflect.Struct {
		return fmt.Errorf("source must be a struct not '%s'", sourceTypes.Kind())
	}

	sourceValues := reflect.ValueOf(source)
	targetValues := reflect.ValueOf(target)

	mapping(ctx, targetTypes, targetValues, sourceValues, options)
	return nil
}
