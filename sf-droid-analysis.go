/* greet.go */
package main

import (
   "fmt"
   "os"
   "github.com/codegangsta/cli"
)

func main() {
   app := cli.NewApp()
   app.Name = "sf-droid-analysis"
   app.Usage = "Automated analysis of DROID and Siegfried reports for Digital Preservation."
   app.HelpName = "sf-droid-analysis"
   app.Version = "0.1.0"
   app.Author = "exponential-decay"
   app.Action = func(c *cli.Context) {
      if len(c.Args()) < 2 {
         cli.ShowAppHelp(c)
      } else {
         fmt.Println(c.Args())
      }
   }

   app.Run(os.Args)
}
