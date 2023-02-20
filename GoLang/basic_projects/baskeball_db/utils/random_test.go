package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	str := GetRandomString(10)
	require.Equal(t, len(str), 10)
}

func TestGetRandomRole(t *testing.T) {
	str := GetRandomRole()

	var catch bool
	for i := range roles {
		if str == roles[i] {
			catch = true
			break
		}
	}
	require.True(t, catch)
}