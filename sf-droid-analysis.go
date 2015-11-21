/* greet.go */
package main

import (
   //"fmt"
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

   var rogues bool
   var report string
   var database string

   app.Flags = []cli.Flag {
      cli.StringFlag{
         Name:          "report, csv, csva",
         Usage:         "Droid or Siegfried report to analyse.",
         Destination:   &report,
      }, 
      cli.StringFlag{
         Name:          "db",
         Usage:         "Database to output results from.",
         Destination:   &database,
      },     
      cli.BoolFlag{
         Name:          "rogues, rogue",
         Usage:         "Output Rogues Gallery",
         Destination:   &rogues,
      },
      cli.StringFlag{
         Name:          "export",
         Usage:         "Export database as CSV.",
         Destination:   &database,
      },
   }

   app.Action = func(c *cli.Context) {
      //cli Args() applies heuristics, os.Args is a purer check
      if len(os.Args) < 2 {
         cli.ShowAppHelp(c)
      } else {
         if len(c.Args()) > 0 {
            //we can do things with string flagss
         }
         if rogues == true {
            //handle rogue output
         }
      }
   }

   app.Run(os.Args)
}
