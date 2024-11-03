package mapping

import (
	"context"
	"fmt"
	"reflect"
)

func Create[T, S any](source S, optionsArr ...Option) (T, error) {
	return CreateContext[T, S](nil, source, optionsArr...)
}

func CreateContext[T, S any](ctx context.Context, source S, optionsArr ...Option) (T, error) {
	options := optionsArrToMap(ctx, optionsArr)

	var target T
	targetTypes := reflect.TypeOf(target)
	if targetTypes.Kind() != reflect.Struct {
		return target, fmt.Errorf("target type must be a struct not '%s'", targetTypes.Kind())
	}

	sourceTypes := reflect.TypeOf(source)
	if sourceTypes.Kind() != reflect.Struct {
		return target, fmt.Errorf("source must be a struct not '%s'", sourceTypes.Kind())
	}

	sourceValues := reflect.ValueOf(source)
	targetValues := reflect.ValueOf(&target)

	mapping(ctx, targetTypes, targetValues, sourceValues, options)

	return target, nil
}
