/* build-analysis-db.go */

package main

import (
   "github.com/mattn/go-sqlite3"
   "database/sql"
)

func dostuff() error {

   var DB_DRIVER string
   sql.Register(DB_DRIVER, &sqlite3.SQLiteDriver{})

   return nil
}
