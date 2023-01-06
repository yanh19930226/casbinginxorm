package proc

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessName(t *testing.T) {

	fmt.Print(ProcessName())
	assert.True(t, len(ProcessName()) > 0)
}

func TestPid(t *testing.T) {
	fmt.Print(Pid())
	assert.True(t, Pid() > 0)
}
