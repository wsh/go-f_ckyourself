// f_ckyourself is a simple sanitization library for Go.
package f_ckyourself

import "strings"

var bowdlerizations map[string]string

func Sanitize(dirty string) (clean string){
  clean = dirty
  bowdlerizations = make(map[string]string)
  bowdlerizations["fuck"] = "f___"
  for expletive, substitute := range bowdlerizations {
    clean = strings.Replace(clean, expletive, substitute, -1)
  }
  return
}
