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
/*
//这里如果使用meddler的话，那么我们对与新增这种动作就不必要再重新自定义一个函数来处理了。仅需要直接使用meddler即可。

func (u *User) AddUser(){

}
*/
