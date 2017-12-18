package parser

import (
	"fmt"
	"strings"
)

const genericError = "ERROR\r\n"
const clientError = "CLIENT_ERROR %s\r\n"
const serverError = "SERVER_ERROR %s\r\n"

const (
	storageCommand = iota
	retrivalCommand
)

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

func validateNumberOfArguments(cmds []string, minExpectedCount, maxExpectedCount int) error {
	if len(cmds) < minExpectedCount {
		return fmt.Errorf(clientError, "invalid number of arguments")
	}

	if maxExpectedCount != 0 && len(cmds) > maxExpectedCount {
		return fmt.Errorf(clientError, "too many arguments")
	}

	return nil
}

// Validate validate a command input
func Validate(command string) error {
	cmds := strings.Split(command, " ")

	if len(cmds) < 2 {
		return fmt.Errorf(clientError, "invalid number of arguments")
	}

	action := cmds[0]
	validator, exist := commandList[action]

	if !exist {
		return fmt.Errorf(clientError, "invalid command: "+action)
	}

	return validator.Validate(cmds)
}
