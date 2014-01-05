package fungo

/*
  Same as IntFilter, but updates the value passed in by reference.

  Example:
    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    fungo.IntFilterR(&nums, isOdd)

    => nums will be []int{1, 3}
*/
func IntFilterR(ints *[]int, f func(int) bool) {
  *ints = IntFilter(*ints, f)
}

/*
  Same as IntMap, but updates the value passed in by reference.

  Example:
    nums := []int{1, 2, 3}
    double := func(a int) int {
      return a * 2
    }

    fungo.IntMapR(&nums, double)

    => nums will be []int{1, 3}
*/
func IntMapR(ints *[]int, f func(int) int) {
  *ints = IntMap(*ints, f)
}

/*
  Same as IntReject, but updates the value passed in by reference.

  Example:
    nums := []int{1, 2, 3}
    isOdd := func(a int) bool {
      if a%2 != 0 {
        return true
      }
      return false
    }

    IntRejectR(&nums, isOdd)

    => nums will be []int{2}
*/
func IntRejectR(ints *[]int, f func(int) bool) {
  *ints = IntReject(*ints, f)
}

/*
  Same as IntReject, but updates the value passed in by reference.

  Example:
    nums := []int{1, 2, 3}
    banned_nums := []int{2}

    IntWithoutR(&nums, &banned_nums)

    => nums will be []int{1, 3}
*/
func IntWithoutR(ints *[]int, without_ints *[]int) {
  *ints = IntWithout(*ints, *without_ints)
}
