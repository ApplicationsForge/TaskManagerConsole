package taskterminal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPluralize(t *testing.T) {
	assert := assert.New(t)
	assert.Equal("task", pluralize(1, "task", "tasks"))
	assert.Equal("tasks", pluralize(2, "task", "tasks"))
}
