package datalayer

import (
	"database/sql"
 _ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id uint
	Name string
	Family string
}

type SqlHandler struct {
	db *sql.DB
}

func CreateConnection(conneString string)(*SqlHandler,error){
	db , err := sql.Open("mysql",conneString)
	if err != nil {
		return nil,err
	}
	return &SqlHandler{
		db:db,
	},nil
}

func (handler *SqlHandler) GetAllUsers() ([]User , error)  {
	rows , err := handler.db.Query("SELECT * FROM users")
	if err != nil {
		return nil , err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Family,
			)
		if err != nil {
			return users , err
		}
		users = append(users,user)
	}
	if err = rows.Err() ; err != nil{
		return users , err
	}
return users , nil
}

func (handler *SqlHandler) GetUserByName(name string) (User , error)  {
	row ,_:= handler.db.Query("SELECT * FROM users WHERE name = ?",name)
	var user User
	err:= row.Scan(
		&user.Id,
		&user.Name,
		&user.Family,
		)
	if err != nil {
		return user , err
	}
	return user , nil
}

func(handler *SqlHandler) AddUser(user User) error {
	_,err := handler.db.Exec("INSERT INTO users (name,family) VALUES(?,?)",user.Name,user.Family)
	if err != nil {
		return err
	}
	return nil
}
func(handler *SqlHandler) UpdateUser(user User) error {
	_,err := handler.db.Exec("UPDATE users SET name = ? , family = ? WHERE id = ?",user.Name,user.Family,user.Id)
	if err != nil {
		return err
	}
	return nil
}
func(handler *SqlHandler) DeleteUser(user User) error{
	_,err := handler.db.Exec("DELETE FROM users WHERE id = ?",user.Id)
	if err != nil {
		return err
	}
	return nil
}