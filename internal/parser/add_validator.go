package parser

type addValidator struct {
}

var _ commandValidator = &addValidator{}

func (v *addValidator) Validate(str []string) error {
	return validateNumberOfArguments(str, 2, 0)
}
