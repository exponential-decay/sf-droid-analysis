/* analysis-input-identifier.go */
package main

import (
   "encoding/hex"
   "strings"
   "log"
   "os"
)

const (  // iota is reset to 0
   INPUT_UNKNOWN  = iota // == 0
   INPUT_DROID    = iota // == 1
   INPUT_SFNORM   = iota // == 2
   INPUT_SFJSON   = iota // == 3
   INPUT_SFDROID  = iota // == 4
   INPUT_DATABASE = iota // == 5
)

//filetype signatures
var inputtypes map[string]string = map[string]string{
   "SFNORMAL"  : "2D2D2D0A736965676672696564",                    //---\r\nsiegfried
   "SFJSON"    : "7B2273696567667269656422",                      //{"siegfried"
   "SFDROID"   : "49442C504152454E545F49442C555249",              //ID,PARENT_ID,URI
   "DROID"     : "224944222C22504152454E545F4944222C2255524922",  //"ID","PARENT_ID","URI"
   "SQLITE"    : "53514C69746520666F726D61742033",                //SQLite format 3
}

func getBytes(val string) []byte {

   file, err := os.Open(val) // For read access.
   if err != nil {
	   log.Fatal(err)
   }

   data := make([]byte, 22)
   count, err := file.Read(data)
   if err != nil {
	   log.Fatal(err)
   }

   log.Printf("INFO: Read %d bytes from input report: %x\n", count, data[:count])
   return data
}

func compareBytes(data []byte) int {
   inputFound := ""

   for k := range inputtypes {
      needlelen := len(inputtypes[k])
      haystack := hex.EncodeToString(data)[0:needlelen]
      if strings.ToLower(inputtypes[k]) == haystack {
         inputFound = k
      }
   }

   inputType := INPUT_UNKNOWN
   switch inputFound {
      case "SFNORMAL":
         inputType = INPUT_SFNORM
      case "SFJSON": 
         inputType = INPUT_SFJSON
      case "SFDROID":
         inputType = INPUT_SFDROID
      case "DROID": 
         inputType = INPUT_DROID
      case "SQLITE":
         inputType = INPUT_DATABASE
   }

   return inputType 
}

func identifyinput(val string) int {
   data := getBytes(val)     
   return compareBytes(data)
}
