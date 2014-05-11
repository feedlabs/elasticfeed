package models

import (
	"errors"
	"strconv"
	"time"
)

func init() {
	Users = make(map[string]*User)
	Users["1"] = &User{"1", "foo"}
	Users["2"] = &User{"2", "bar"}
	Users["3"] = &User{"3", "foobar"}
}

func AddUser(user User) (UserId string) {
	user.UserId = strconv.FormatInt(time.Now().UnixNano(), 10)
	Users[user.UserId] = &user
	return user.UserId
}

func GetUser(UserId string) (user *User, err error) {
	if v, ok := Users[UserId]; ok {
		return v, nil
	}
	return nil, errors.New("UserId Not Exist")
}

func GetUserList() map[string]*User {
	return Users
}

func UpdateUser(UserId string, Name string) (err error) {
	if v, ok := Users[UserId]; ok {
		v.Name = Name
		return nil
	}
	return errors.New("UserId Not Exist")
}

func DeleteUser(UserId string) {
	delete(Users, UserId)
}

func (user *User) GetName() string {
	return user.Name
}
