package parser

type setValidator struct {
}

var _ commandValidator = &setValidator{}

func (v *setValidator) Validate(str []string) error {
	return validateNumberOfArguments(str, 2, 0)
}
