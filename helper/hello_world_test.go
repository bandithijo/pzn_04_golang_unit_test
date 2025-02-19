package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func BenchmarkHelloWorldRizqi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rizqi")
	}
}

func BenchmarkHelloWorldNur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Nur")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Rizqi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Rizqi")
		}
	})

	b.Run("Nur", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Nur")
		}
	})
}

func BenchmarkHelloWorldtable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name: "HelloWorld(Rizqi)",
			request: "Rizqi",
		},
		{
			name: "HelloWorld(Nur)",
			request: "Nur",
		},
		{
			name: "HelloWorld(Assyaufi)",
			request: "Assyaufi",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Rizqi", func(t *testing.T) {
		result := HelloWorld("Rizqi")
		require.Equal(t, "Hello, Rizqi", result)
	})

	t.Run("Nur", func(t *testing.T) {
		result := HelloWorld("Nur")
		require.Equal(t, "Hello, Nur", result)
	})
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name: "HellowWorld(Rizqi)",
			request: "Rizqi",
			expected: "Hello, Rizqi",
		},
		{
			name: "HelloWorld(Nur)",
			request: "Nur",
			expected: "Hello, Nur",
		},
		{
			name: "HelloWorld(Assyaufi)",
			request: "Assyaufi",
			expected: "Hello, Assyaufi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run test on Linux")
	}

	result := HelloWorld("Rizqi")
	require.Equal(t, "Hello, Rizqi", result)
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rizqi")
	assert.Equal(t, "Hello, Rizqi", result, "Result must be Hello, Rizqi")
	fmt.Println("TestHelloWorldAssert Done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Rizqi")
	require.Equal(t, "Hello, Rizqi", result, "Result must be Hello, Rizqi")
	fmt.Println("TestHelloWorldRequire Done")
}

func TestHelloWorldRizqi(t *testing.T) {
	result := HelloWorld("Rizqi")

	if result != "Hello, Rizqi" {
		// failed
		// panic("Result is not Hello, Rizqi")
		// t.Fail()
		t.Error("Result must be Hello, Rizqi")
	}

	fmt.Println("TestHelloWorldRizqi Done")
}

func TestHelloWorldNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello, Nur" {
		// failed
		// panic("Result is not Hello, Nur")
		// t.FailNow()
		t.Fatal("Result must be Hello, Nur")
	}

	fmt.Println("TestHelloWorldNur Done")
}

// === RUN   TestHelloWorldRizqi
//     hello_world_test.go:15: Result must be Hello, Rizqi
// TestHelloWorldRizqi Done
// --- FAIL: TestHelloWorldRizqi (0.00s)
// === RUN   TestHelloWorldNur
//     hello_world_test.go:28: Result must be Hello, Nur
// --- FAIL: TestHelloWorldNur (0.00s)
// FAIL
// exit status 1
// FAIL    pzn_04_golang_unit_test/helper  0.004s
