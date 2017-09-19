package rolodex

import (
       "log"
       "os"
       "io"
       "fmt"
       "encoding/csv"
       "strconv"
)

var zipcodes map[int]string = make(map[int]string)

func InitZipCodes(fname string) {
     f, err := os.Open(fname)
     if err != nil {
     	log.Fatal(err)
     }
     reader := csv.NewReader(f)
     for {
     	 record, err := reader.Read()
	 if err == io.EOF {
	    break
	 }
	 zipcode, err := strconv.Atoi(record[0])
	 if err != nil {
	    log.Fatal(err)
	 }
	 zipcodes[zipcode] = record[1]
     }
}

type AddressRecord struct {
     FirstName, LastName, StreetAddress, City, State string
     Zipcode int
}

func ReadAddressCSV(fname string, c chan AddressRecord) {
     var n_records_total, n_records_duplicate, n_records_invalid int
     seen := make(map[AddressRecord]bool)

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
	 n_records_total += 1
	 if err != nil {
	    log.Print(err)
	    n_records_invalid += 1
	    continue
	 }
         zipcode, err := strconv.Atoi(record[5])
         if err != nil {
             n_records_invalid += 1
             log.Print(err)
	     continue
         }

	 if(zipcodes[zipcode] != record[4]) {
             log.Print(fmt.Sprintf("Zipcode %d is not valid for state %s", zipcode, record[4]))
             n_records_invalid += 1
	     continue
	 }

         address_record := AddressRecord{record[0], record[1], record[2], record[3], record[4], zipcode}

	 if seen[address_record] {
             n_records_duplicate += 1
             continue
	 }
	 seen[address_record] = true
         c <- address_record
     }
     close(c)
     log.Print(fmt.Sprintf("File %s - %d records total; %d invalid; %d duplicate", fname, n_records_total, n_records_invalid, n_records_duplicate))
}
