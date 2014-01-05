# Functional goodness for Go.

## Disclaimer

This library is work in progress.
For now, you can use it on slices of integers and strings.
There is no warranty that the API won't change (it shouldn't, though)

## Usage

### API Documentation

You can find the complete API documentation here:
http://godoc.org/github.com/yagooar/fungo

These are some of the functions Fungo provides:

```go
func IntContain(ints []int, num int) bool
func IntEvery(ints []int, f func(int) bool) bool
func IntFilter(ints []int, f func(int) bool) (ret_ints []int)
func IntFilterR(ints *[]int, f func(int) bool)
func IntFind(ints []int, f func(int) bool) (elem int, ok bool)
func IntMap(ints []int, f func(int) int) (ret_ints []int)
func IntMapR(ints *[]int, f func(int) int)
func IntReduce(ints []int, f func(int, int) int, memo int) int
func IntReject(ints []int, f func(int) bool) (ret_ints []int)
func IntRejectR(ints *[]int, f func(int) bool)
func IntSome(ints []int, f func(int) bool) bool
func IntWithout(ints []int, without_ints []int) (ret_ints []int)
func IntWithoutR(ints *[]int, without_ints *[]int)
func StringContain(strings []string, s string) bool
func StringEvery(strings []string, f func(string) bool) bool
func StringFilter(strings []string, f func(string) bool) (ret_strings []string)
func StringFilterR(strings *[]string, f func(string) bool)
func StringFind(strings []string, f func(string) bool) (s string, ok bool)
func StringMap(strings []string, f func(string) string) (ret_strings []string)
func StringMapR(strings *[]string, f func(string) string)
func StringReject(strings []string, f func(string) bool) (ret_strings []string)
func StringRejectR(strings *[]string, f func(string) bool)
func StringSome(strings []string, f func(string) bool) bool
func StringWithout(strings []string, without_s []string) (res []string)
func StringWithoutR(strings *[]string, without_s *[]string)
```

### Example

```go
package testing_fungo

import (
  "fmt"
  "github.com/yagooar/fungo"
)

func main() {
  nums := []int{1, 2, 3}
  isOdd := func(a int) bool {
    if a%2 != 0 {
      return true
    }
    return false
  }

  odd_nums := fungo.IntFilter(nums, isOdd)
  fmt.Printf("Odd numbers: %s\n", odd_nums)
}
```

In this example, odd_nums will be `[]int{1, 3}`

## Installation

```bash
go get github.com/yagooar/fungo
```


## FAQ

### What is this?

Fungo is not intended to be a full-fledged functional extension for Go,
but rather a collection of tools for common Go types, like integers or strings.

### Why isn't it generic?

Go provides a set of tools for reflection, but it is known that using them
can make the code slower than using explicit typing.
You can read more on this topic here:
http://blog.burntsushi.net/type-parametric-functions-golang

### What is Fungo good for?

I found writing a lot of code that used typical functional operations,
like find, filter or map. Usually, those functions where on integers or
strings. This package allows to avoid repeating the same patterns over and
over again.

### Why isn't there function X for type Y?

Well, because no one has written it yet. Feature & pull requests are welcome.

### Why did you write it?

There are plenty of reasons:
* I wanted to get rid off commonly repeated patterns to avoid reinventing
the wheel every time and make my codebase cleaner and more explicit
* I didn't want to spend time testing those patterns every time I needed
them in an application.
* I wanted to share these patterns with people, so they can use them freely
* I wanted to get feedback on the current algorithms to further improve them.
* For fun. Go is fun, so we shall have fun.

