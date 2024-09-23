package encodedecodestrings

import (
	"fmt"
	"slices"
	"testing"
)

func TestEncode(t *testing.T) {
	input := []string{"hello", "world", ":", "encode"}
	got := encode(input)
	want := "5|hello5|world1|:6|encode"

	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func TestDecode(t *testing.T) {
	input := []string{"hello", "world", ":", "encode"}
	encstr := encode(input)
	got := decode(encstr)
	want := input

	fmt.Println("got", got)
	fmt.Println("want", want)

	if !slices.Equal(got, want) {
		t.Errorf("got %#v want %#v", got, want)
	}
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"hello", "world", ":", "encode", "some", "random", "strins"}
		encode(input)
	}
}

func BenchmarkEncodeBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"hello", "world", ":", "encode", "some", "random", "strins"}
		encodeBasic(input)
	}
}

func BenchmarkEncode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"hello", "world", ":", "encode", "some", "random", "strins"}
		encode2(input)
	}
}

func BenchmarkSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"hello", "world", ":", "encode", "some", "random", "strins"}
		encodeSimple(input)
	}
}

func BenchmarkRaw(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"hello", "world", ":", "encode", "some", "random", "strins"}
		encodeRaw(input)
	}
}

func BenchmarkRaw2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"hello", "world", ":", "encode", "some", "random", "strins"}
		encodeRaw2(input)
	}
}

func BenchmarkRaw3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"hello", "world", ":", "encode", "some", "random", "strins"}
		encodeRaw3(input)
	}
}
