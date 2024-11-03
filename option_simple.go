package mapping

type Simple struct {
	Target  string
	Source  string
	Default any
}

func (o Simple) target() string {
	return o.Target
}

func (o Simple) source() string {
	return o.Source
}

func (o Simple) ignore() bool {
	return false
}

func (o Simple) qualifier() func(any) (any, error) {
	return nil
}

func (o Simple) defVal() any {
	return o.Default
}

func (o Simple) constant() bool {
	return false
}
