package parser

import "testing"
import "github.com/stretchr/testify/assert"

func TestValidate(t *testing.T) {
	cmds := []string{"set", "add", "replace", "append", "prepend", "cas", "get", "gets"}

	for _, cmd := range cmds {
		t.Run(cmd, func(t *testing.T) {
			memCmd := cmd + " foo"

			err := Validate(memCmd)
			assert.Nil(t, err)

			err = Validate(cmd)
			assert.NotNil(t, err)
		})
	}

	t.Run("empty command", func(t *testing.T) {
		err := Validate("")

		assert.NotNil(t, err)
	})
}
