/* analysis-command-handler.go */

package main

const (  // iota is reset to 0
         ANALYSIS  = iota   // == 0
         DATABASE  = iota   // == 1
         ROGUE     = iota   // == 2
         EXPORT    = iota   // == 3
)

func handleAnalysisCommands(cmd int, val string) error {
   //handle command routing here...
   return nil
}
