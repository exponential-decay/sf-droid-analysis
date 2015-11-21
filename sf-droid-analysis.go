/* sf-droid-analysis.go */
package main

import (
   "log"
   "os"
   "strconv"
   "github.com/codegangsta/cli"
)

func main() {

   app := cli.NewApp()

   app.Name = "sf-droid-analysis"
   app.Usage = "Automated analysis of DROID and Siegfried reports for Digital Preservation."
   app.HelpName = "sf-droid-analysis"
   app.Version = "0.1.0"
   app.Author = "exponential-decay"

   var report string
   var database string
   var export string
   var rogues, heroes bool

   app.Flags = []cli.Flag {
      cli.StringFlag{
         Name:          "analysis, report, csv, csva",
         Usage:         "Droid or Siegfried report to analyse.",
         Destination:   &report,
      }, 
      cli.StringFlag{
         Name:          "db, database",
         Usage:         "Database to output results from.",
         Destination:   &database,
      },     
      cli.BoolFlag{
         Name:          "rogues, rogue",
         Usage:         "Output Rogues Gallery",
         Destination:   &rogues,
      },
      cli.BoolFlag{
         Name:          "heroes, hero",
         Usage:         "Output Heroes Gallery",
         Destination:   &heroes,
      },
      cli.StringFlag{
         Name:          "export",
         Usage:         "Export database as CSV.",
         Destination:   &export,
      },
   }

   app.Action = func(c *cli.Context) {
      //cli Args() applies heuristics, os.Args is a purer check
      if len(os.Args) < 2 {
         cli.ShowAppHelp(c)
      } else {
         //we can do things with string flagss
         if c.IsSet("analysis") {
            if err := handleAnalysisCommands(ANALYSIS, report); err != nil {
	               log.Fatal(err)
               }
         } else if c.IsSet("db") {
            //handle rogue output
            if err := handleAnalysisCommands(DATABASE, database); err != nil {
	               log.Fatal(err)
               }
         } else if c.IsSet("export") {
            //handle rogue output
            if err := handleAnalysisCommands(EXPORT, export); err != nil {
	               log.Fatal(err)
               }
         } else if c.IsSet("rogues") {
            //handle rogue output
            if err := handleAnalysisCommands(ROGUE, strconv.FormatBool(rogues)); err != nil {
		            log.Fatal(err)
	            }
         } else {
            cli.ShowAppHelp(c)
         }
      }
   }

   app.Run(os.Args)
}
