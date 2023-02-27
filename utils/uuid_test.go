package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	a := assert.New(t)
	uid, err := CreateUid()
	a.NotEmpty(uid)
	a.Nil(err)
}
