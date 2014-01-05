package fungo

/*
  Looks through each value in the list.
  Returns a slice of all the values that pass the truth test (f).

  Example:
    words := []string{"man", "woman", "child"}
    func longerThanThree(s string) bool {
      if len(s) > 3 {
        return true
      }
      return false
    }

    long_words := fungo.StringFilter(words, longerThanThree)

    => long_words will be []string{"woman", "child"}
*/
func StringFilter(strings []string, f func(string) bool) (ret_strings []string) {
  for _, value := range(strings) {
    if f(value) {
      ret_strings = append(ret_strings, value)
    }
  }

  return
}

/*
  Produces a new slice of values by mapping each value in
  the list through a transformation function (f).

  Example:
    words := []string{"man", "woman", "child"}
    func withPrefix(s string) string {
      return "good" + s
    }

    prefixed_words := fungo.StringMap(words, withPrefix)

    => prefixed_words will be []string{"goodman", "goodwoman", "goodchild"}
*/
func StringMap(strings []string, f func(string) string) (ret_strings []string) {
  for _, value := range(strings) {
    ret_strings = append(ret_strings, f(value))
  }

  return
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

    words := []string{"man", "woman", "child"}
    func longerThanThree(s string) bool {
      if len(s) > 3 {
        return true
      }
      return false
    }

    long_word := fungo.StringFind(words, longerThanThree)

    => long_word will equal "woman"
*/
func StringFind(strings []string, f func(string) bool) (s string, ok bool) {
  for _, value := range(strings) {
    if f(value) {
      return value, true
    }
  }

  return s, false
}

/*
  Returns the values in list without the elements that the
  truth test (f) passes. The opposite of filter.

  Example:

    words := []string{"man", "woman", "child"}
    func longerThanThree(s string) bool {
      if len(s) > 3 {
        return true
      }
      return false
    }

    short_words := fungo.StringReject(words, longerThanThree)

    => short_words will equal []string{"man"}
*/
func StringReject(strings []string, f func(string) bool) (ret_strings []string) {
  for _, value := range(strings) {
    if !f(value) {
      ret_strings = append(ret_strings, value)
    }
  }

  return
}

/*
  Returns true if all of the values in the list pass the truth test (f).

  Example:

    words := []string{"woman", "child"}
    func longerThanThree(s string) bool {
      if len(s) > 3 {
        return true
      }
      return false
    }

    every := fungo.StringEvery(words, longerThanThree)

    => every will be true
*/
func StringEvery(strings []string, f func(string) bool) bool {
  for _, value := range(strings) {
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

    words := []string{"man", "woman", "child"}
    func longerThanThree(s string) bool {
      if len(s) > 3 {
        return true
      }
      return false
    }

    some := fungo.StringSome(words, longerThanThree)

    => some will be true
*/
func StringSome(strings []string, f func(string) bool) bool {
  for _, value := range(strings) {
    if f(value) {
      return true
    }
  }

  return false
}

/*
  Returns true if the value is present in the list.

  Example:

    words := []string{"man", "woman", "child"}
    contains := fungo.StringContain(words, "woman")

    => contains will be true
*/
func StringContain(strings []string, s string) bool {
  for _, value := range(strings) {
    if value == s {
      return true
    }
  }

  return false
}

/*
  Returns a slice with all banned values removed.

  Example:

    words := []string{"man", "woman", "child"}
    banned_words := []string{"man", "woman"}
    allowed_words := fungo.StringWithout(words, banned_words)

    => allowed_words will equal []string{"child"}
*/
func StringWithout(strings []string, without_s []string) (res []string) {
  var found bool
  for _, value := range(strings) {
    found = false
    for _, without_value := range(without_s) {
      if value == without_value {
        found = true
        break
      }
    }

    if !found {
      res = append(res, value)
    }
  }

  return
}


