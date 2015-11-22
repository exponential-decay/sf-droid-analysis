/* sf-output-handler.go */

// Handles SF default output and SF JSON output 

package main

import (
   "os"
   "log"
   "bufio"
   "strings"
)

func handleSFdefaultoutput(sffile string) error {

   file, err := os.Open(sffile) // For read access.
   if err != nil {
	   log.Fatal(err)
   }

   scanner := bufio.NewScanner(file)
   for scanner.Scan() {
      line := strings.SplitN(scanner.Text(), ":", 2)
      if line[0] != "---" {
         key, value := strings.Trim(line[0], " -"), strings.Trim(line[1], " ")
         log.Println(key, value)
      }
   }

   if err := scanner.Err(); err != nil {
      log.Fatal(err)
   }

   return nil
}
