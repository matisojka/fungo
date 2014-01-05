# Functional goodness for Go.

## Usage

In a source file, import the Fungo package:

```go
package my_cool_project

import "github.com/yagooar/fungo"
```

And then, use the available functions in your code.

### Example

```go
package testing_fungo

import "github.com/yagooar/fungo"

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
but rather a collection of tools for common Go types, link integers or strings.

### Why isn't it generic?

Go provides a set of tools for reflection, but it is known that using them
can make the code slower than using explicit typing.

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

