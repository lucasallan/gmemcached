package parser

import (
	"fmt"
	"strings"
)

const genericError = "ERROR\r\n"
const clientError = "CLIENT_ERROR %s\r\n"
const serverError = "SERVER_ERROR %s\r\n"

var commandList map[string]commandValidator

func init() {
	commandList = make(map[string]commandValidator)

	commandList["get"] = &getValidator{}
	commandList["gets"] = &getValidator{}
	commandList["set"] = &setValidator{}
	commandList["add"] = &addValidator{}
	commandList["replace"] = &replaceValidator{}
	commandList["append"] = &appendValidator{}
	commandList["prepend"] = &prependValidator{}
	commandList["cas"] = &casValidator{}
}

type commandValidator interface {
	Validate(str []string) error
}

func validateNumberOfArguments(args []string, minExpectedCount, maxExpectedCount int) error {
	if len(args) < minExpectedCount {
		return fmt.Errorf(clientError, "invalid number of arguments")
	}

	if maxExpectedCount != 0 && len(args) > maxExpectedCount {
		return fmt.Errorf(clientError, "too many arguments")
	}

	return nil
}

// Validate validate a command input
func Validate(command string) error {
	args := strings.Split(command, " ")

	if len(args) < 2 {
		return fmt.Errorf(clientError, "invalid number of arguments")
	}

	action := args[0]
	validator, exist := commandList[action]

	if !exist {
		return fmt.Errorf(clientError, "invalid command: "+action)
	}

	return validator.Validate(args)
}
