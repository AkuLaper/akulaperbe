package config

import (
	"os"

	"github.com/AkuLaper/akulaperbe/helper/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var mongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "dblaper",
}

var Mongoconn, ErrorMongoconn = atdb.MongoConnect(mongoinfo)
