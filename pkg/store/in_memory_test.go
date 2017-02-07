package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testKey   = "key"
	testValue = "value"
)

func TestInMemoryStore_Exists(t *testing.T) {
	instance := NewInMemoryStore()

	instance.Set(testKey, testValue, 0)

	val, err := instance.Exists("foo")
	assert.Nil(t, err)
	assert.False(t, val)

	val, err = instance.Exists(testKey)
	assert.Nil(t, err)
	assert.True(t, val)
}

func TestInMemoryStore_Get(t *testing.T) {
	instance := NewInMemoryStore()

	instance.Set(testKey, testValue, 0)

	val, err := instance.Get(testKey)
	assert.Nil(t, err)
	assert.Equal(t, testValue, val)

	val, err = instance.Get("foo")
	assert.Nil(t, err)
	assert.Empty(t, val)
}
