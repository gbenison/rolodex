package rolodex

import (
       "log"
       "os"
       "io"
       "encoding/csv"
       "strconv"
)

type AddressRecord struct {
     FirstName, LastName, StreetAddress, City, State string
     Zipcode int
}

func ReadAddressCSV(fname string, c chan AddressRecord) {
     f, err := os.Open(fname)
     if err != nil {
     	log.Fatal(err)
     }
     reader := csv.NewReader(f)
     reader.Comment = '#'
     for {
         record, err := reader.Read()
	 if err == io.EOF {
	    break
	 }
	 if err != nil {
	    log.Print(err)
	    continue
	 }
         zipcode, err := strconv.Atoi(record[5])
         if err != nil {
             log.Print(err)
	     continue
         }
         address_record := AddressRecord{record[0], record[1], record[2], record[3], record[4], zipcode}
         c <- address_record
     }
     close(c)
}
