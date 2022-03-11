package dao



import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"time"
)

var DB *sql.DB
func init(){
	//	执行main函数之前执行的init方法
	dataSourceName:=fmt.Sprintf("root:libing0804.@tcp(localhost:3306)/goblog?charset=utf8&loc=%s&parseTime=true",url.QueryEscape("Asia/Shanghai"))
	db,err:=sql.Open("mysql",dataSourceName)
	if err!=nil{
		log.Println(err)
		panic(err)
	}
	//	最大空闲连接数不配置是两个
	db.SetConnMaxIdleTime(5)
	//	最大连接数  不配置是无上限
	db.SetMaxOpenConns(100)
	//	连接最大存活时间
	db.SetConnMaxLifetime(time.Minute*3)
	//空闲连接最大存活时间

	db.SetConnMaxIdleTime(time.Minute*1)
	err = db.Ping()
	if err != nil {
		log.Println("数据库无法连接")
		_ = db.Close()
		panic(err)
	}
	DB = db
}

