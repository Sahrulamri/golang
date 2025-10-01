package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)



func TestHelloWorld(t *testing.T) {
	result:= HelloWorld("Golang")
	if result != "Hello World Golang" {
		t.Error("Result must be 'Hello World Golang'")
	}
}


// cara running
// go test

func TestHelloKhanedy(t *testing.T) {
	result:= HelloWorld("Khanedy")
	if result != "Hello World Khanedy" {
		t.Error("Result must be 'Hello World Khanedy'")
	}
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Golang")
	assert.Equal(t, "Hello World Golang", result, "Result must be 'Hello World Golang'")
	fmt.Println("TestHelloWorldAssert running")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Golang")
	require.Equal(t, "Hello World Golang", result, "Result must be 'Hello World Golang'")
	fmt.Println("TestHelloWorldAssert running")
}

func TestMain(m *testing.M) {
	fmt.Println("TestMain running")
	m.Run()
	fmt.Println("TestMain finished")
}

func TestSubTest(t *testing.T) {
	t.Run("Khanedy", func(t *testing.T) {
		result := HelloWorld("Khanedy")
		require.Equal(t, "Hello World Khanedy", result, "Result must be 'Hello World Khanedy'")
	})
	t.Run("Golang", func(t *testing.T) {
		result := HelloWorld("Golang")
		require.Equal(t, "Hello World Golang", result, "Result must be 'Hello World Golang'")
	})
}

func TestHelloWorldtTable(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Khanedy", "Khanedy", "Hello World Khanedy"},
		{"Golang", "Golang", "Hello World Golang"},
		{"Java", "Java", "Hello World Java"},
		{"Python", "Python", "Hello World Python"},
		{"Budi", "Budi", "Hello World Budi"},
		{"Joko", "Joko", "Hello World Joko"},
		{"Sandhika", "Sandhika", "Hello World Sandhika"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.input)
			require.Equal(t, test.expected, result, "Result must be 'Hello World Golang'")
		})
	}
	}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Golang")
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Khanedy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Khanedy")
		}
	})
	b.Run("Golang", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Golang")
		}
	})
}

func BenchmarkHelloWorldTable(b *testing.B) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Khanedy", "Khanedy", "Hello World Khanedy"},
		{"Golang", "Golang", "Hello World Golang"},
		{"Java", "Java", "Hello World Java"},
		{"Python", "Python", "Hello World Python"},
		{"Budi", "Budi", "Hello World Budi"},
		{"Joko", "Joko", "Hello World Joko"},
		{"Sandhika", "Sandhika", "Hello World Sandhika"},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.input)
			}
		})
	}
}