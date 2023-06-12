package entities

import "fmt"

type Insert_user struct {
	Id				int64
	First_name      string
	Last_name      	string
	Country         string
	Profile_picture string
}

func (insert_user Insert_user) ToString() string {
	return fmt.Sprintf("id: %d\nfirst_name: %s\nlast_name: %s\ncountry: %s\nprofile_picture: %s", insert_user.Id, insert_user.First_name, insert_user.Last_name, insert_user.Country, insert_user.Profile_picture)
}
