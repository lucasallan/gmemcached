package parser

type appendValidator struct {
}

var _ commandValidator = &appendValidator{}

func (v *appendValidator) Validate(str []string) error {
	return validateNumberOfArguments(str, 2, 0)
}
