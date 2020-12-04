package main

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	return
}

func TestGetAdvisor(t *testing.T) {
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}
	orm.Debug = true

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err)
		return
	}
	orm.AddAliasWthDB("default", "mysql", db)
	orm.RegisterModel(new(Advisor))

	defer db.Close()

	robot := Advisor{
		1,
		"robot",
		"111222",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "phone"}).AddRow(robot.Id, robot.Name, robot.Phone)

	mock.ExpectQuery("SELECT id, name, phone FROM cms_advisors").WillReturnRows(rows)

	fmt.Println(GetAdvisor())
	return
}
