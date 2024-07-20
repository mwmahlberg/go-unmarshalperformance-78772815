package gounmarshalperformance78772815

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type unknownJson map[string]interface{}

func benchmark(path string, arraysize int, b *testing.B) {
	d, err := os.ReadFile(path)
	require.NoError(b, err)
	for i := 0; i < b.N; i++ {
		foo := make([]unknownJson, arraysize)
		err := json.Unmarshal(d, &foo)
		require.NoError(b, err)
	}
}

func Benchmark64k(b *testing.B) {
	// Array sizes were determined by running the following command:
	// $ jq length testdata/78772815-{64k,128k,256k}.json
	benchmark("testdata/78772815-64k.json", 197, b)
}

func Benchmark128k(b *testing.B) {
	benchmark("testdata/78772815-128k.json", 788, b)
}

func Benchmark256k(b *testing.B) {
	benchmark("testdata/78772815-256k.json", 792, b)
}
