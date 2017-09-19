package rolodex

import (
       "database/sql"
       "log"
)

func InitDB(db * sql.DB) {
     _, err := db.Exec(`
       create table rolodex (
           firstname text,
	   lastname text, 
	   street_address text, 
	   city text, 
	   state text,
	   zipcode int
       );
     `)
     if err != nil {
     	log.Fatal(err)
     }
}

func Serialize(c chan AddressRecord, db * sql.DB) {
     log.Print("Serializing to database")
     for elem := range c {
         _, err := db.Exec(
	           "insert into rolodex(firstname, lastname, street_address, city, state, zipcode) values (?, ?, ?, ?, ?, ?)",
		   elem.FirstName,
		   elem.LastName,
		   elem.StreetAddress,
		   elem.City,
		   elem.State,
		   elem.Zipcode,
         )
	 if err != nil {
             log.Fatal(err)
	 }
     }
}