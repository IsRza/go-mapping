package mapping

import "fmt"

type QualifierOption[I, O any] struct {
	Target    string
	Source    string
	Qualifier func(I) O
}

func (o QualifierOption[I, O]) target() string {
	return o.Target
}

func (o QualifierOption[I, O]) source() string {
	return o.Source
}

func (o QualifierOption[I, O]) ignore() bool {
	return false
}

func (o QualifierOption[I, O]) qualifier() func(any) (any, error) {
	if o.Qualifier == nil {
		return nil
	}
	return func(arg any) (any, error) {
		casted, ok := arg.(I)
		if !ok {
			return nil, fmt.Errorf(
				"input type of Qualifier(Target: %s, Source: %s) does not match with source value",
				o.Target, o.Source,
			)
		}

		return o.Qualifier(casted), nil
	}
}
