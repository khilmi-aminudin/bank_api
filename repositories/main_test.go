package repositories

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/khilmi-aminudin/bank_api/db"
	"github.com/khilmi-aminudin/bank_api/utils"
)

var testDB *sql.DB
var testStore Store

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../")
	if err != nil {
		log.Fatal("cannt load config : ", err)
	}

	testDB = db.Connect(*config)
	testStore = NewStore(testDB)

	os.Exit(m.Run())
}
