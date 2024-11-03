package mapping

type Ignore struct {
	Target string
}

func (o Ignore) target() string {
	return o.Target
}

func (o Ignore) source() string {
	return ""
}

func (o Ignore) ignore() bool {
	return true
}

func (o Ignore) defVal() any {
	return nil
}

func (o Ignore) qualifier() func(any) (any, error) {
	return nil
}
