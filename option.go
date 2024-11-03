package mapping

import "context"

type Option interface {
	target() string
	source() string
	ignore() bool
	qualifier() func(any) (any, error)
}

func optionsArrToMap(ctx context.Context, optionArr []Option) map[string]Option {
	optionsMap := map[string]Option{}
	for _, option := range optionArr {
		if _, ok := optionsMap[option.target()]; ok {
			getLogger(ctx).Warnf("Overriding option for target '%s'", option.target())
		}
		optionsMap[option.target()] = option
	}
	return optionsMap
}
