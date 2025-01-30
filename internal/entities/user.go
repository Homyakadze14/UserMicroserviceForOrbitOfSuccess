package entities

import "fmt"

type UserInfo struct {
	UserID     int
	Firstname  string
	Middlename string
	Lastname   string
	Gender     string
	Phone      string
	IconURL    string
}

func (u UserInfo) String() string {
	return fmt.Sprintf("UserID: %v; Firstname: %v; Middlename: %v; Lastname: %v, Gender: %v; Phone: %v; IconURL: %v",
		u.UserID, u.Firstname, u.Middlename, u.Lastname, u.Gender, u.Phone, u.IconURL)
}
