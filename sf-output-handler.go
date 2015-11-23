/* sf-output-handler.go */

// Handles SF default output and SF JSON output 

package main

import (
   "os"
   "log"
   "bufio"
   "strings"
)

var sfheader map[string]string = map[string]string {
   "siegfried" : "",
   "scandate"  : "",
   "signature" : "",
   "created"   : "",

   //identifiers, SF default style (1.4.1) (no additional signatures) 
   "name"      : "",
   "details"   : "",
}
 
var sfbody map[string]string = map[string]string {
   "filename"  : "",
   "filesize"  : "",
   "modified"  : "",
   "errors"    : "",

   //identifiers, sf default style (1.4.1) (no additional signatures)
   "id"        : "",
   "puid"      : "",
   "format"    : "",
   "version"   : "",
   "mime"      : "",
   "basis"     : "",
   "warning"   : "",
}

var mapslice []map[string]string

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func handleSFdefaultoutput(sffile string) error {

   file, err := os.Open(sffile) // For read access.
   if err != nil {
	   log.Fatal(err)
   }

   var headercount   int = 0    //we don't need to keep checking against header 
   var bodycount     int = 0

   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
      line := strings.SplitN(scanner.Text(), ":", 2)
      if line[0] != "---" {
         key, value := strings.Trim(line[0], " -"), strings.Trim(line[1], " ")
         if headercount < len(sfheader) {
            //check key in map syntax: http://stackoverflow.com/a/2050629      
            if _, ok := sfheader[key]; ok {
               headercount++
               sfheader[key] = strings.Trim(value, "'")                         
            }
         }
         if bodycount < len(sfbody) {
            if _, ok := sfbody[key]; ok {
               bodycount++
               sfbody[key] = strings.Trim(value, "'")                         
            }
         } else {
            mapslice = append(mapslice, sfbody)
            bodycount = 0
         }      
      }
   }

   for i := range mapslice {
      for k, v := range mapslice[i] {
         log.Println(k, v, "\n")
      }
   }

   if err := scanner.Err(); err != nil {
      log.Fatal(err)
   }

   dostuff()

   return nil
}
