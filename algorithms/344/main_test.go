package main

import (
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "hello", args: args{s: []byte("hello")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			// todo:: 断言
		})
	}
}

func Benchmark_reverseString(b *testing.B) {
	s := []byte("the sky is blue")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseString(s)
	}
}

func Benchmark_reverseString2(b *testing.B) {
	s := []byte("the sky is blue")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverseString2(s)
	}
}
