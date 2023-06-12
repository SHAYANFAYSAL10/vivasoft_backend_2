package entities

import "fmt"

type User struct {
	First_name      string
	Country         string
	Profile_picture string
	Total_point		int
	Rank			int
}

func (user User) ToString() string {
	return fmt.Sprintf("first_name: %s\ncountry: %s\nprofile_picture: %s\ntotal_point: %d\nrank: %d", user.First_name, user.Country, user.Profile_picture, user.Total_point, user.Rank)
}
