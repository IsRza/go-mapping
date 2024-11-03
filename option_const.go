package mapping

type Constant struct {
	Target string
	Value  any
}

func (o Constant) target() string {
	return o.Target
}

func (o Constant) source() string {
	return ""
}

func (o Constant) ignore() bool {
	return false
}

func (o Constant) defVal() any {
	return o.Value
}

func (o Constant) qualifier() func(any) (any, error) {
	return nil
}
