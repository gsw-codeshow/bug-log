package main

import (
	"fmt"
	"path"
	"runtime"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func GetCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func InitSql() {
	config, _ := config.NewConfig("ini", GetCurrentPath()+"/conf/app.conf")
	user := config.String("mysqluser")
	passwd := config.String("mysqlpass")
	host := config.String("mysqlurls")
	port, err := config.Int("mysqlport")
	dbname := config.String("mysqldb")
	if err != nil {
		port = 3306
	}
	if config.String("runmode") == "dev" {
		orm.Debug = true
	}
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.DefaultTimeLoc = time.UTC
	mysqlConnect := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local", user, passwd, host, port, dbname)
	_ = orm.RegisterDataBase("default", "mysql", mysqlConnect)
	orm.SetMaxIdleConns("default", 0)
}

func init() {
}

type Advisor struct {
	Id    int64  `orm:"column(id);"`
	Name  string `orm:"column(name)"`
	Phone string `orm:"column(phone)"`
}

func (*Advisor) TableName() string {
	return "cms_advisors"
}

func GetAdvisor() (advisor []Advisor) {
	orm.NewOrm().Raw("SELECT id, name, phone FROM cms_advisors").QueryRows(&advisor)
	return
}

func AdvisorList(db orm.Ormer) (advisor []Advisor) {
	db.QueryTable("cms_advisors").All(&advisor)
	return
}

func TestGetAdvisor() {
	orm.Debug = true

	db, mock, err := sqlmock.New()
	if err != nil {
		return
	}

	robot := Advisor{
		1,
		"robot",
		"111222",
	}

	rows := sqlmock.NewRows([]string{"id", "name", "phone"}).AddRow(robot.Id, robot.Name, robot.Phone)
	// mock.ExpectPrepare("SELECT TIMEDIFF")
	mock.ExpectPrepare("SELECT id, name, phone FROM cms_advisors")
	mock.ExpectQuery("SELECT id, name, phone FROM cms_advisors").WillReturnRows(rows)
	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		panic(err)
	}
	orm.AddAliasWthDB("default", "mysql", db)
	orm.RegisterModel(new(Advisor))

	defer db.Close()

	// testDBRows, testErr := db.Query("SELECT id, name, phone FROM cms_advisors")
	// if testErr != nil {
	// 	fmt.Println(testErr)
	// 	fmt.Println(" --- --- ")
	// }
	// for testDBRows.Next() {
	// 	p := &Advisor{}
	// 	if err := testDBRows.Scan(&p.Id, &p.Name, &p.Phone); err != nil {
	// 		return
	// 	}
	// 	fmt.Println(p)
	// }
	fmt.Println(GetAdvisor())
	return
}

func main() {
	TestGetAdvisor()
	// fmt.Println(AdvisorList(orm.NewOrm()))
	return
}
