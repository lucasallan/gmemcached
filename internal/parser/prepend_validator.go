package parser

type prependValidator struct {
}

var _ commandValidator = &prependValidator{}

func (v *prependValidator) Validate(str []string) error {
	return validateNumberOfArguments(str, 2, 0)
}
