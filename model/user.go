package model 
import (
	"github.com/liatong/lds-ops/sql"
)

type User struct{
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (u *User) QueryUser(name string)(users []User,err error ){
	//rows, err := sql.Pool.Query("select `id`, `name` from `test_name` where `name` = ?", name)
	
	users = make([]User, 0)
	rows, err := sql.Pool.Query("select `id`, `name` from `test_name` where `name` = ?", name)
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
 		return
	}
	return  
	
}