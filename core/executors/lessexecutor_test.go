package executors

import (
	"testing"
	"time"

	"casbinginxorm/core/timex"

	"github.com/stretchr/testify/assert"
)

func TestLessExecutor_DoOrDiscard(t *testing.T) {
	executor := NewLessExecutor(time.Minute)
	assert.True(t, executor.DoOrDiscard(func() {}))
	assert.False(t, executor.DoOrDiscard(func() {}))
	executor.lastTime.Set(timex.Now() - time.Minute - time.Second*30)
	assert.True(t, executor.DoOrDiscard(func() {}))
	assert.False(t, executor.DoOrDiscard(func() {}))
}

func BenchmarkLessExecutor(b *testing.B) {
	exec := NewLessExecutor(time.Millisecond)
	for i := 0; i < b.N; i++ {
		exec.DoOrDiscard(func() {
		})
	}
}
