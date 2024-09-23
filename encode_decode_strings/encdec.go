package encodedecodestrings

import (
	"strconv"
	"strings"
)

func encodeBasic(input []string) string {
	return strings.Join(input, "|")
}

func decode(input string) []string {
	size := 0
	for i := 0; i < len(input); i++ {
		if input[i] == '|' {
			i++
		}
	}

	words := make([]string, 0, size)
	i := 0
	for i < len(input) {
		ts := i
		te := i
		for input[te] != '|' {
			te++
		}

		l := 0
		d := 1
		for j := te - 1; j >= ts; j-- {
			l += int((input[j])-48) * d
			d *= 10
		}

		start := te + 1
		end := start + l

		words = append(words, input[start:end])
		i = end
	}

	return words
}

func encode(input []string) string {
	var sb strings.Builder

	for _, str := range input {
		sb.WriteString(strconv.Itoa(len(str)))
		sb.WriteRune('|')
		sb.WriteString(str)
	}

	return sb.String()
}

func encode2(input []string) string {
	var sb strings.Builder

	for _, str := range input {
		sb.WriteString(str + "|")
	}

	return sb.String()
}

func encodeSimple(input []string) string {
	var sb string

	for _, str := range input {
		sb += (str + "|")
	}

	return sb
}

func encodeRaw(input []string) string {
	sb := make([]rune, 0, len(input)*7)

	for _, str := range input {
		sb = append(sb, []rune(str)...)
		sb = append(sb, '|')
	}

	return string(sb)
}

func encodeRaw2(input []string) string {
	sb := make([]rune, 0, len(input)*7)

	for _, str := range input {
		for _, char := range []rune(str) {
			sb = append(sb, char)
		}
		sb = append(sb, '|')
	}

	return string(sb)
}

func encodeRaw3(input []string) string {
	var sb strings.Builder

	for _, str := range input {
		for _, r := range []rune(str) {
			sb.WriteRune(r)
		}
		sb.WriteRune('|')
	}

	return sb.String()
}
