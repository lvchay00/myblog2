package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int    `orm:"pk;auto"`
	Username string `orm:"unique"`
	Password string
	Token    string    `orm:"unique"`
	Email    string    `orm:"null"`
	InTime   time.Time `orm:"auto_now_add;type(datetime)"`
}

func FindUserById(id int) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Id", id).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByToken(token string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Token", token).One(&user)
	return err != orm.ErrNoRows, user
}

func Login(username string, password string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).Filter("Password", password).One(&user)
	return err != orm.ErrNoRows, user
}

func FindUserByUserName(username string) (bool, User) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(user).Filter("Username", username).One(&user)
	return err != orm.ErrNoRows, user
}

func SaveUser(user *User) int64 {
	o := orm.NewOrm()
	id, _ := o.Insert(user)
	return id
}

func UpdateUser(user *User) {
	o := orm.NewOrm()
	o.Update(user)
}

func DeleteUser(user *User) {
	o := orm.NewOrm()
	o.Delete(user)
}

func DeleteUserRolesByUserId(user_id int) {
	o := orm.NewOrm()
	o.Raw("delete from user_roles where user_id = ?", user_id).Exec()
}

func SaveUserRole(user_id int, role_id int) {
	o := orm.NewOrm()
	o.Raw("insert into user_roles (user_id, role_id) values (?, ?)", user_id, role_id).Exec()
}

func FindUserRolesByUserId(user_id int) []orm.Params {
	o := orm.NewOrm()
	var res []orm.Params
	o.Raw("select id, user_id, role_id from user_roles where user_id = ?", user_id).Values(&res, "id", "user_id", "role_id")
	return res
}
