package db

import (
	"reflect"
	"sunlight_automated_upgrade_service/server/configuration"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

//CreateObject creates a new Object
func CreateObject(r interface{}) error {
	return Error.ErrCreate(db.Create(r).Error, reflect.TypeOf(r).String())
}

//UpdateObject updates an Object
func UpdateObject(r interface{}) error {
	return Error.ErrUpdate(db.Save(&r).Error, reflect.TypeOf(r).String())
}

//DeleteObject deletes an Object
func DeleteObject(r interface{}) error {
	return Error.ErrDelete(db.Delete(r).Error, reflect.TypeOf(r).String())
}

//InitDB initilizes the db object and the tables
func InitDB() {

	// Open a new connection to our sqlite database.
	d, err := gorm.Open("sqlite3", configuration.Conf.DBFname)
	if err != nil {
		panic("Failed to open the SQLite database." + err.Error())
	}

	d.AutoMigrate(&User{})

	db = d
}

//GetDB returns the db object
func GetDB() *gorm.DB {
	return db
}
