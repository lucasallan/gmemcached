package parser

type getValidator struct {
	cmd []string
}

var _ commandValidator = &getValidator{}

func (v *getValidator) Validate(str []string) error {
	return validateNumberOfArguments(str, 2, 0)
}
