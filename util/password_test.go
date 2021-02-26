package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPassword(t *testing.T) {
	password1 := GetPassword("123456")
	password2 := GetPassword("correct")
	password3 := GetPassword("wrong")

	fmt.Println("\"123456\" encryption is", password1)
	fmt.Println("\"correct\" encryption is", password2)
	fmt.Println("\"wrong\" encryption is", password3)

	assert.Equal(t, CheckPassword("123456", password1), true)
	assert.Equal(t, CheckPassword("correct", password2), true)
	assert.Equal(t, CheckPassword("wrong", password3), true)
	assert.Equal(t, CheckPassword("123456", password3), false)
}
