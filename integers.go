package fungo

/*
  Looks through each value in the list.
  Returns a slice of all the values that pass the truth test (f).

  Example:
    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    odd_nums := fungo.IntFilter(nums, isOdd)

    => odd_nums will be []int{1, 3}
*/
func IntFilter(ints []int, f func(int) bool) (ret_ints []int) {
  for _, value := range(ints) {
    if f(value) {
      ret_ints = append(ret_ints, value)
    }
  }

  return
}

/*
  Produces a new array of values by mapping each value in
  the list through a transformation function (f).

  Example:
    nums := []int{1, 2, 3}
    double := func(a int) int {
      return a * 2
    }

    double_nums := fungo.IntMap(nums, double)

    => double_nums will be []int{1, 3}
*/
func IntMap(ints []int, f func(int) int) (ret_ints []int) {
  for _, value := range(ints) {
    ret_ints = append(ret_ints, f(value))
  }

  return
}

/*
  Sometimes called inject or foldl, reduce boils down a list of values
  into a single value.
  Memo is the initial state of the reduction, and each successive step
  of it should be returned by the iterator (f).

  Example:

    nums := []int{1, 2, 3}
    memo := 5
    sum := func(a int, b int) int {
      return a + b
    }

    reduced_num := IntReduce(nums, sum, memo)

    => reduced_num will be 11
*/
func IntReduce(ints []int, f func(int, int) int, memo int) int {
  for _, value := range(ints) {
    memo = f(value, memo)
  }

  return memo
}

/*
  Looks through each value in the list, returning the first one that passes
  a truth test (f) and a boolean ok flag set to true.
  If no value matches the truth test, the ok flag will be
  returned as false.
  The function returns as soon as it finds an acceptable element,
  and doesn't traverse the entire list.
  In order to ensure having gotten a result, you have to check
  for the ok flag every time you use this function.

  Example:

    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    first_odd, ok := IntFind(nums, isOdd)

    => first_odd will be 1
    => ok will be true
*/
func IntFind(ints []int, f func(int) bool) (elem int, ok bool) {
  for _, value := range(ints) {
    if f(value) {
      return value, true
    }
  }

  return elem, false
}

/*
  Returns the values in list without the elements that the
  truth test (f) passes. The opposite of filter.

  Example:

    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    pair_nums := IntReject(nums, isOdd)

    => pair_nums will be []int{2}
*/
func IntReject(ints []int, f func(int) bool) (ret_ints []int) {
  for _, value := range(ints) {
    if !f(value) {
      ret_ints = append(ret_ints, value)
    }
  }

  return
}

/*
  Returns true if all of the values in the list pass the truth test (f).

  Example:

    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    all_odds := IntEvery(nums, isOdd)

    => all_odds will be false
*/
func IntEvery(ints []int, f func(int) bool) bool {
  for _, value := range(ints) {
    if !f(value) {
      return false
    }
  }

  return true
}

/*
  Returns true if any of the values in the list pass the truth test (f).
  The function will stop and return as soon as there is a passing value.

  Example:

    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    some_odds := IntSome(nums, isOdd)

    => some_odds will be true
*/
func IntSome(ints []int, f func(int) bool) bool {
  for _, value := range(ints) {
    if f(value) {
      return true
    }
  }

  return false
}

/*
  Returns true if the value is present in the list.

  Example:

    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    contains_odd := IntContains(nums, isOdd)

    => contains_odd will be true
*/
func IntContains(ints []int, f func(int) bool) bool {
  _, ok := IntFind(ints, f)

  return ok
}

/*
  Returns a slice with all banned values removed.

  Example:

    nums := []int{1, 2, 3}
    banned_nums := []int{2}

    allowed_nums := IntWithout(nums, banned_nums)

    => allowed_nums will be []int{1, 3}
*/
func IntWithout(ints []int, without_ints []int) (ret_ints []int) {
  var found bool
  for _, value := range(ints) {
    found = false
    for _, without_value := range(without_ints) {
      if value == without_value {
        found = true
        break
      }
    }

    if !found {
      ret_ints = append(ret_ints, value)
    }
  }

  return
}

