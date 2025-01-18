# Cartesian Product

Giving a string `"abc"` and  `n=3` generate all possible combinations of each character in a string with length `n`; e.g.

```go
var chars string = "abc"
var result []string = CartesianProduct(chars, 2)
// Cartesian Product ->
// aa
// ab
// ac
// ba
// bb
// bc
// ca
// cb
// cc
// Product length 9

var result []string = CartesianProduct(chars, 3)
// Cartesian Product ->
// aaa
// aab
// aac
// aba
// abb
// abc
// aca
// acb
// acc
// baa
// bab
// bac
// bba
// bbb
// bbc
// bca
// bcb
// bcc
// caa
// cab
// cac
// cba
// cbb
// cbc
// cca
// ccb
// ccc
// Product length 27
```
