/* analysis-command-handler.go */
package main

import (
   //"log"
)

const (  // iota is reset to 0
   ANALYSIS = iota // == 0
   DATABASE = iota // == 1
   ROGUE    = iota // == 2
   EXPORT   = iota // == 3
   HERO     = iota // == 4
)

func handleAnalysisCommands(cmd int, val string) error {
   //handle command routing here...
   
   if cmd == ANALYSIS || cmd == DATABASE {
      identifyinput(val)
      // we can analyse the content
      // id if we're looking at a db/csv/or json
      // handle appropriately
   }

   if cmd == ROGUE {
      // output rogue list
   }

   if cmd == HERO {
      // output hero list
   }

   if cmd == EXPORT {
      // we can export the database
   }
   return nil
}
