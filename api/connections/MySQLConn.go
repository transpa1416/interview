package conexion

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	MySQL_SERVER   = "localhost"
	MySQL_USER     = "root"
	MySQL_PASS     = ""
	MySQL_PORT     = "3306"
	MySQL_DATABASE = "interview"
)

var DB *gorm.DB
var err error

//conexión a MySQL
func GetMySQLConn() (*gorm.DB, error) {
	dbDriver := "mysql"
	dbName := MySQL_DATABASE
	dbUser := MySQL_USER
	dbPassword := MySQL_PASS
	DB, err = gorm.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName+"?charset=utf8&parseTime=True")
	if err != nil {
		return nil, err
	}
	return DB, nil
}
