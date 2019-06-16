package dbmanagers

import (
	"../models"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-xorm/xorm"
	"log"
)

var (
	isShowSQLCode    bool   = true
	connectionString string = "driver={SQL Server};server=localhost;port=14330;database=barfootDB;user id=sa;password=Lei123456;"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

// Create ORM object
func getOrm() *xorm.Engine {

	orm, err := xorm.NewEngine("mssql", connectionString)
	checkErr(err, "Creating ORM object was failed")

	orm.ShowSQL(isShowSQLCode)
	return orm
}

// InitDB method
// parameter isShowSql - true - show sql script in debug console/ false - will skip it
func InitDB(isShowSQL bool) {

	orm := getOrm()
	defer orm.Close()
	isShowSQLCode = isShowSQL
	err := orm.Sync2(new(models.Barfoottable1))
	checkErr(err, "Sync structures with DB was failed")
}

// Get list of customers
func GetCustomers() []models.Barfoottable1 {
	var customers []models.Barfoottable1

	orm := getOrm()
	defer orm.Close()

	err := orm.Find(&customers)
	checkErr(err, "Get Data was failed")

	return customers
}

// Update Customer
func UpdateCustomer(customer models.Barfoottable1) bool {

	orm := getOrm()
	defer orm.Close()

	affected, err := orm.Id(customer.ID).Update(&customer)
	checkErr(err, "Update was failed")

	return affected > 0
}

// Insert Customer
func InsertCustomer(customer models.Barfoottable1) bool {

	//	if (models.Barfoottable1{}) == customer {
	//		return false
	//	}

	orm := getOrm()
	defer orm.Close()

	affected, err := orm.Insert(&customer)
	checkErr(err, "Insert was failed")

	return affected > 0
}

// Delete Customer
func DeleteCustomer(customer models.Barfoottable1) bool {

	//	if (models.Barfoottable1{}) == customer {
	//		return false
	//	}

	orm := getOrm()
	defer orm.Close()

	affected, err := orm.Id(customer.ID).Delete(&customer)
	checkErr(err, "Delete was failed")

	return affected > 0
}
