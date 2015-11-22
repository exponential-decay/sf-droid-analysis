/* analysis-command-handler.go */
package main

import (
   "os"
   "log"
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
      inputtype := identifyinput(val)
      switch inputtype { 
         case INPUT_DROID:    
         case INPUT_SFDEFAULT: 
            log.Println("handling normal SF input.") 
            handleSFdefaultoutput(val) 
         case INPUT_SFJSON:   
         case INPUT_SFDROID:  
         case INPUT_DATABASE: 
         default: 
            log.Println("INFO: Unknown input type.")
            os.Exit(0)
      }
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
