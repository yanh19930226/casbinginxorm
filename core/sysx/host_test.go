package sysx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostname(t *testing.T) {

	fmt.Print(Hostname())

	assert.True(t, len(Hostname()) > 0)

}
