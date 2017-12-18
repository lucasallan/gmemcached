package parser

type replaceValidator struct {
}

var _ commandValidator = &replaceValidator{}

func (v *replaceValidator) Validate(str []string) error {
	return validateNumberOfArguments(str, 2, 0)
}
