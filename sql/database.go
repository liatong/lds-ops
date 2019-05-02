package sql 
import(
	"log"
	
	//导入database/sql
	"database/sql"
	_ "github.com/go-sql-driver/mysql"


)
var Pool *sql.DB

func init() {
	var err error
	Pool, err = sql.Open("mysql","root:Password@tcp(127.0.0.1:3306)/lds_ops")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = Pool.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
