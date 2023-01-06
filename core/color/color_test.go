package color

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithColor(t *testing.T) {
	output := WithColor("Hello", BgRed)
	fmt.Print(output)
	assert.Equal(t, "Hello", output)
}

func TestWithColorPadding(t *testing.T) {
	output := WithColorPadding("Hello", BgRed)
	fmt.Print(output)
	assert.Equal(t, " Hello ", output)
}
