package main

import (
       "log"
       "os"
       "fmt"
       "database/sql"
       _ "github.com/mattn/go-sqlite3"
       "github.com/gbenison/rolodex/lib"
)

var db_file = "rolodex.db"

func main() {

     if len(os.Args) != 3 {

         log.Print(`
Usage: rolodex <address file> <zipcode file>

Load the contents of a csv file containing addresses into an sqlite database.
The second argument is a mandatory csv file containing <zip code>,<state name> pairs
of valid zip codes.
         `)

	 os.Exit(1)
     }

     var input_file = os.Args[1]
     var zipcode_file = os.Args[2]

     db, err := sql.Open("sqlite3", db_file)
     if err != nil {
         log.Fatal(err)
     }
     log.Print(fmt.Sprintf("Importing file %s to sqlite database %s", input_file, db_file))
     lib.InitDB(db)

     lib.InitZipCodes(zipcode_file)
     c := make(chan lib.AddressRecord)
     go lib.ReadAddressCSV(input_file, c)
     lib.Serialize(c, db)
}