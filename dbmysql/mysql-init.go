package dbmysql

import (
	"github.com/phongcao0114/book-service/db"
)

func Init() db.MasterDB {
	return db.MasterDB{
		BookDBImpl: initBookDBImpl(),
	}
}
