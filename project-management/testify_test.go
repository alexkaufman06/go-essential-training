package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSimple2(t *testing.T) {
	val, err := Sqrt(2)
	require.NoError(t, err)
	require.InDelta(t, 1.414214, val, 0.001)
}
