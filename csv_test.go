package rolodex

import (
       "testing"
       "os"
)

func TestReadFile(t *testing.T) {
     _, err := os.Open("testdata/enemies.csv")
     if err != nil {
     	t.Errorf("Could not open test data")
     }
}

func TestRecord(t *testing.T) {

     InitZipCodes("testdata/zipcodes.csv")

     c := make(chan AddressRecord)
     go ReadAddressCSV("testdata/enemies.csv", c)

     first_record := <- c
     if first_record.FirstName != "Greg" {
     	t.Errorf("Expected FirstName Greg; got %s", first_record.FirstName)
     }
     if first_record.StreetAddress != "404 N. Found St." {
     	t.Errorf("Expected StreetAddress 404 N. Found St.; got %s", first_record.StreetAddress)
     }
     if first_record.Zipcode != 97221 {
     	t.Errorf("Expected zip code 97221; got %d", first_record.Zipcode)
     }

     second_valid_record := <- c
     if second_valid_record.FirstName != "Swiper" {
     	t.Errorf("Expected FirstName Swiper; got %s", second_valid_record.FirstName)
     }
}