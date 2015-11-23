/* build-analysis-db.go */

package main

import (
   "log"
   "github.com/mattn/go-sqlite3"
   "database/sql"
)

func dostuff() error {

   //http://golang-basic.blogspot.co.nz/2014/06/golang-database-step-by-step-guide-on.html
   var DB_DRIVER string
   sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{})

   _, err := sql.Open(DB_DRIVER, "mysqlite_3") 
   if err != nil {
      log.Println("ERROR: Failed to create the handle.")
   }

   return nil
}

