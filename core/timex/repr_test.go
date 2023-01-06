package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReprOfDuration(t *testing.T) {

	fmt.Println(ReprOfDuration(time.Second))

	fmt.Println(ReprOfDuration(
		time.Second + time.Millisecond*111 + time.Microsecond*555))

	assert.Equal(t, "1000.0ms", ReprOfDuration(time.Second))
	assert.Equal(t, "1111.6ms", ReprOfDuration(
		time.Second+time.Millisecond*111+time.Microsecond*555))
}
