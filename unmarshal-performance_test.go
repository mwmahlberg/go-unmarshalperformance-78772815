package gounmarshalperformance78772815

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type unknownJSON map[string]interface{}

type testStruct struct {
	Name     string  `json:"name"`
	Language string  `json:"language"`
	ID       string  `json:"id"`
	Bio      string  `json:"bio"`
	Version  float64 `json:"version"`
}

func benchmark(path string, arraysize int, b *testing.B) {
	d, err := os.ReadFile(path)
	require.NoError(b, err)
	foo := make([]unknownJSON, arraysize)
	err = json.Unmarshal(d, &foo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err := json.Marshal(foo)
		require.NoError(b, err)
		require.NotEmpty(b, r)
	}
}

func benchmarkStruct(path string, arraysize int, b *testing.B) {
	d, err := os.ReadFile(path)
	require.NoError(b, err)
	foo := make([]testStruct, arraysize)
	err = json.Unmarshal(d, &foo)
	require.NoError(b, err)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r, err := json.Marshal(foo)
		require.NoError(b, err)
		require.NotEmpty(b, r)
	}
}

func benchmarkEncoder(path string, arraysize int, b *testing.B) {
	d, err := os.ReadFile(path)
	require.NoError(b, err)
	foo := make([]unknownJSON, arraysize)
	err = json.Unmarshal(d, &foo)
	require.NoError(b, err)
	buf := bytes.NewBuffer(make([]byte, len(d)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		enc := json.NewEncoder(buf)
		err := enc.Encode(foo)
		require.NoError(b, err)
		buf.Reset()
	}
}

func Benchmark64kMapStringInterface(b *testing.B) {
	// Unmarshal the JSON data to determine the size of the array
	// Array sizes were determined by running the following command:
	// $ jq length testdata/78772815-{64k,128k,256k}.json
	benchmark("testdata/78772815-64k.json", 197, b)
}

func Benchmark128kMapStringInterface(b *testing.B) {
	benchmark("testdata/78772815-128k.json", 788, b)
}

func Benchmark256kMapStringInterface(b *testing.B) {
	benchmark("testdata/78772815-256k.json", 792, b)
}

func Benchmark64kStruct(b *testing.B) {
	benchmarkStruct("testdata/78772815-64k.json", 197, b)
}

func Benchmark128kStruct(b *testing.B) {
	benchmarkStruct("testdata/78772815-128k.json", 788, b)
}

func Benchmark256kStruct(b *testing.B) {
	benchmarkStruct("testdata/78772815-256k.json", 792, b)
}

func Benchmark64kEncoder(b *testing.B) {
	benchmarkEncoder("testdata/78772815-64k.json", 197, b)
}

func Benchmark128kEncoder(b *testing.B) {
	benchmarkEncoder("testdata/78772815-128k.json", 788, b)
}

func Benchmark256kEncoder(b *testing.B) {
	benchmarkEncoder("testdata/78772815-256k.json", 792, b)
}
