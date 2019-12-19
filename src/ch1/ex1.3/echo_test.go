package main

import (
	"bytes"
	"strconv"
	"testing"
)

var inputSize = 10000

func BenchmarkEchoSlow(b *testing.B) {
	var buf bytes.Buffer
	var s = make([]string, 0, inputSize)
	for i := 0; i < inputSize; i++ {
		s = append(s, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		echoSlow(&buf, s)
	}
}

func BenchmarkEchoFast(b *testing.B) {
	var buf bytes.Buffer
	var s = make([]string, 0, inputSize)
	for i := 0; i < inputSize; i++ {
		s = append(s, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		echoFast(&buf, s)
	}
}

func TestEchoSlow(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{""}, "\n"},
		{[]string{"abcde", "123", "-test"}, "abcde 123 -test\n"},
		{[]string{" ", " "}, "   \n"}, // never happens through command line
	}
	for _, test := range tests {
		var buf bytes.Buffer
		if echoSlow(&buf, test.input); buf.String() != test.want {
			t.Errorf("echoSlow(%q) = %v", test.input, buf.String())
		}
	}
}

func TestEchoFast(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
	}{
		{[]string{""}, "\n"},
		{[]string{"abcde", "123", "-test"}, "abcde 123 -test\n"},
		{[]string{" ", " "}, "   \n"}, // never happens through command line
	}
	for _, test := range tests {
		var buf bytes.Buffer
		if echoFast(&buf, test.input); buf.String() != test.want {
			t.Errorf("echoSlow(%q) = %v", test.input, buf.String())
		}
	}
}
