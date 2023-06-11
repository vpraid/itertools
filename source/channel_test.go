package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChannel_Empty(t *testing.T) {
	t.Parallel()

	c := make(chan int)
	go func() {
		ci := Channel(c)
		assert.False(t, ci.Next())
	}()
	close(c)
}

func TestChannel_NonEmpty(t *testing.T) {
	t.Parallel()

	c := make(chan int)
	go func() {
		ci := Channel(c)
		assert.True(t, ci.Next())
		assert.Equal(t, 1, ci.Value())
		assert.True(t, ci.Next())
		assert.Equal(t, 2, ci.Value())
		assert.True(t, ci.Next())
		assert.Equal(t, 3, ci.Value())
		assert.False(t, ci.Next())
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
}

func TestChannel_Collect(t *testing.T) {
	t.Parallel()

	c := make(chan int)
	go func() {
		ci := Channel(c)
		assert.Equal(t, []int{1, 2, 3}, ci.Collect())
	}()
	c <- 1
	c <- 2
	c <- 3
	close(c)
}
