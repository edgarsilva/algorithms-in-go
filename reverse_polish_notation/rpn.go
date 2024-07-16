package reverse_polish_notation

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Assert() bool {
	exp := "(4*3)+2(3)-5"
	fmt.Println("Initial expresion", exp)

	exp = handleImplicitMult(exp)
	fmt.Println("W/o implicit mult", exp)

	rpn := ExpToRPN(exp)
	fmt.Println("RPN notation", strings.Join(rpn, " "))

	result := EvalMathExp(exp)
	fmt.Println("Exp evaluated result", result)

	return true
}

func EvalMathExp(exp string) string {
	exp = handleImplicitMult(exp)
	rpn := ExpToRPN(exp)
	result := EvalRPN(rpn)

	return result
}

func ExpToRPN(exp string) []string {
	var operators Stack[rune]
	var output Stack[string]

	runeExp := []rune(exp)
	n := len(runeExp)

	for i := 0; i < n; i++ {
		c := runeExp[i]
		if unicode.IsDigit(c) {
			num := []rune{}
			for i+1 < n && unicode.IsDigit(runeExp[i+1]) {
				i++
				num = append(num, runeExp[i])
			}
			output.Push(string(num))
		} else if c == '(' {
			operators.Push(c)
		} else if c == ')' {
			for operators.Len() > 0 && operators.Last() != '(' {
				output.Push(string(operators.Last()))
				operators.Pop()
			}
			operators.Pop()
		} else if isOp(string(c)) {
			for operators.Len() > 0 &&
				isOp(string(operators.Last())) &&
				precedenceOf(c) <= precedenceOf(operators.Last()) {
				output.Push(string(operators.Last()))
				operators.Pop()
			}
			operators.Push(c)
		}
	}

	for operators.Len() > 0 {
		output.Push(string(operators.Last()))
		operators.Pop()
	}

	return output.Elems()
}

func handleImplicitMult(exp string) string {
	var sb strings.Builder

	for i := 0; i < len(exp); i++ {
		if i > 0 && exp[i] == '(' && (unicode.IsDigit(rune(exp[i-1])) || exp[i-1] == ')') {
			sb.WriteString("*(")
		} else {
			sb.WriteByte(exp[i])
		}
	}

	return sb.String()
}

func doMath(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	}
	panic("operation not supported")
}

func EvalRPN(rpn []string) string {
	var stack []float64

	for _, token := range rpn {
		if isOp(token) {
			a := stack[len(stack)-2]
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			result := doMath(a, b, token)
			stack = append(stack, result)
		} else {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		}
	}

	return strconv.FormatFloat(stack[0], 'f', 2, 64)
}

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Elems() []T {
	return s.elements
}

func (s *Stack[T]) Push(elem T) int {
	s.elements = append(s.elements, elem)

	return len(s.elements)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var t T
		return t, false
	}

	elem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return elem, true
}

func (s *Stack[T]) Last() T {
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.elements)
}

func (s *Stack[T]) ElemAt(idx int) T {
	return s.elements[idx]
}

func (s *Stack[T]) Replace(src []T) {
	s.elements = src
}

func precedenceOf(op rune) int {
	prec := map[rune]int{
		'-': 1,
		'+': 2,
		'/': 3,
		'*': 4,
	}

	return prec[op]
}

func isOp(op string) bool {
	valid := map[string]bool{
		"-": true,
		"+": true,
		"/": true,
		"*": true,
	}

	_, ok := valid[op]
	return ok
}
