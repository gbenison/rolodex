package rolodex

import (
       "testing"
       "database/sql"
       _ "github.com/mattn/go-sqlite3"
)

func mockAddressSource(c chan AddressRecord) {
     c <- AddressRecord{"TestFirst1", "TestSecond1", "TestAddress1", "TestCity1", "TestState1", 12345}
     c <- AddressRecord{"TestFirst2", "TestSecond2", "TestAddress2", "TestCity2", "TestState2", 12345}
     c <- AddressRecord{"TestFirst3", "TestSecond3", "TestAddress3", "TestCity3", "TestState3", 12345}
     close(c)
}

func TestSerialize(t *testing.T) {
     c := make(chan AddressRecord)
     go mockAddressSource(c)

     db, err := sql.Open("sqlite3", ":memory:")
     if err != nil {
     	t.Errorf("Booboo opening database connection.")
     }
     db.SetMaxOpenConns(1)

     InitDB(db)
     Serialize(c, db)

     var n int
     err = db.QueryRow("select count(1) from rolodex").Scan(&n)
     if err != nil {
     	t.Errorf(err.Error())
     }
     if (n != 3) {
        t.Errorf("Expected 3 records; got %d", n)
     }
     var firstname string
     err = db.QueryRow("select firstname from rolodex where lastname='TestSecond2'").Scan(&firstname)
     if err != nil {
     	t.Errorf(err.Error())
     }
     if firstname != "TestFirst2" {
     	t.Errorf("Expected firstname=TestFirst2; got %s", firstname)
     }
}
