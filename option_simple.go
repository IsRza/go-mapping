package mapping

type SimpleOption struct {
	Target string
	Source string
}

func (o SimpleOption) target() string {
	return o.Target
}

func (o SimpleOption) source() string {
	return o.Source
}

func (o SimpleOption) ignore() bool {
	return false
}

func (o SimpleOption) qualifier() func(any) (any, error) {
	return nil
}
