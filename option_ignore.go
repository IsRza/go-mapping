package mapping

type IgnoreOption struct {
	Target string
}

func (o IgnoreOption) target() string {
	return o.Target
}

func (o IgnoreOption) source() string {
	return ""
}

func (o IgnoreOption) ignore() bool {
	return true
}

func (o IgnoreOption) qualifier() func(any) (any, error) {
	return nil
}
