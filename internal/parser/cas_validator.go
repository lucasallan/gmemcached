package parser

type casValidator struct {
}

var _ commandValidator = &casValidator{}

func (v *casValidator) Validate(str []string) error {
	return validateNumberOfArguments(str, 2, 0)
}
