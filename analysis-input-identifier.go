/* analysis-input-identifier.go */
package main

const (  // iota is reset to 0
   INPUT_UNKNOWN  = iota // == 0
   INPUT_DROID    = iota // == 1
   INPUT_SFNORM   = iota // == 2
   INPUT_SFJSON   = iota // == 3
   INPUT_SFDROID  = iota // == 4
   INPUT_DATABASE = iota // == 5
)

func identifyinput(val string) int {

   return INPUT_DROID
}
