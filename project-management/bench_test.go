package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var benchText = "Dont communicate by sharing memory, share memory by communicating"

// need to install Graphviz (graphviz.org)

func BenchmarkTokenize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tokens := Tokenize(benchText)
		require.Equal(b, 10, len(tokens))
	}
}
