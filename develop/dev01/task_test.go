package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetExactTime(t *testing.T) {
	time, err := GetTime("0.beevik-ntp.pool.ntp.org")

	assert.NoError(t, err)
	assert.False(t, time.IsZero())
}
