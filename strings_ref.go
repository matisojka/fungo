package fungo

/*
  Same as StringFilter, but updates the value passed in by reference.

  Example:
    words := []string{"man", "woman", "child"}
    func longerThanThree(s string) bool {
      if len(s) > 3 {
        return true
      }
      return false
    }

    fungo.StringFilterR(&words, longerThanThree)

    => words will be []string{"woman", "child"}
*/
func StringFilterR(strings *[]string, f func(string) bool) {
  *strings = StringFilter(*strings, f)
}

/*
  Same as StringFilter, but updates the value passed in by reference.

  Example:
    words := []string{"man", "woman", "child"}
    func withPrefix(s string) string {
      return "good" + s
    }

    fungo.StringMapR(&words, withPrefix)

    => words will be []string{"goodman", "goodwoman", "goodchild"}
*/
func StringMapR(strings *[]string, f func(string) string) {
  *strings = StringMap(*strings, f)
}

/*
  Same as StringReject, but updates the value passed in by reference.

  Example:
    words := []string{"man", "woman", "child"}
    func longerThanThree(s string) bool {
      if len(s) > 3 {
        return true
      }
      return false
    }

    fungo.StringRejectR(&words, longerThanThree)

    => words will equal []string{"man"}
*/
func StringRejectR(strings *[]string, f func(string) bool) {
  *strings = StringReject(*strings, f)
}

/*
  Same as StringWithout, but updates the value passed in by reference.

  Example:
    words := []string{"man", "woman", "child"}
    banned_words := []string{"man", "woman"}
    fungo.StringWithout(&words, &banned_words)

    => words will equal []string{"child"}
*/
func StringWithoutR(strings *[]string, without_s *[]string) {
  *strings = StringWithout(*strings, *without_s)
}

