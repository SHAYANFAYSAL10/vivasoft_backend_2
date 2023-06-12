package models

import (
	"VIVASOFT2/src/entities"
	"database/sql"
	"sort"
	"strconv"
	"time"
)

type UserModel struct {
	Db *sql.DB
}

type Activity struct {
	First_Name      string
	CountrY         string
	Profile_Picture string
	Total_Point     int
}

func (userModel UserModel) FindAll() (user []entities.User, err error) {
	activities_map := make(map[int]int)
	user_points := make(map[int]int)

	rows, err := userModel.Db.Query("select * from activities")

	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			var ac_id int
			var points int
			err2 := rows.Scan(&ac_id, &points)
			if err2 != nil {
				return nil, err2
			} else {
				activities_map[ac_id] = int(points)
			}
		}
	}

	rows2, err3 := userModel.Db.Query("select * from users")

	if err3 != nil {
		return nil, err3
	} else {

		for rows2.Next() {
			var id int
			var first_name string
			var last_name string
			var country string
			var profile_picture string
			err4 := rows2.Scan(&id, &first_name, &last_name, &country, &profile_picture)
			if err4 != nil {
				return nil, err4
			} else {
				user_points[id] = 0
			}
		}
	}

	rows3, err5 := userModel.Db.Query("select * from activity_logs")

	if err5 != nil {
		return nil, err5
	} else {
		for rows3.Next() {
			var id int
			var user_id int
			var activity_id int
			var logged_at string
			err6 := rows3.Scan(&id, &user_id, &activity_id, &logged_at)
			if err6 != nil {
				return nil, err6
			} else {
				today := time.Now()
				year := today.Year()
				month := int(today.Month()) - 1
				day := today.Day()
				year_string := strconv.Itoa(year)
				month_string := strconv.Itoa(month)
				day_string := strconv.Itoa(day)
				hour_string := "00:00:00"
				var a_month_ago_time string
				if month < 10 && day < 10 {
					a_month_ago_time = year_string + "-0" + month_string + "-0" + day_string + " " + hour_string
				} else if month < 10 {
					a_month_ago_time = year_string + "-0" + month_string + "-" + day_string + " " + hour_string
				} else if day < 10 {
					a_month_ago_time = year_string + "-" + month_string + "-0" + day_string + " " + hour_string
				} else {
					a_month_ago_time = year_string + "-" + month_string + "-" + day_string + " " + hour_string
				}
				if a_month_ago_time < logged_at {
					user_points[user_id] = user_points[user_id] + activities_map[activity_id]
				}
			}
		}
	}

	var activities []Activity

	rows4, err7 := userModel.Db.Query("select * from users")

	if err7 != nil {
		return nil, err7
	} else {

		for rows4.Next() {
			var id int
			var first_name string
			var last_name string
			var country string
			var profile_picture string
			err8 := rows4.Scan(&id, &first_name, &last_name, &country, &profile_picture)
			if err8 != nil {
				return nil, err8
			} else {
				var temp int = user_points[id]
				activity := Activity{
					First_Name:      first_name,
					CountrY:         country,
					Profile_Picture: profile_picture,
					Total_Point:     temp,
				}
				activities = append(activities, activity)
			}
		}
	}

	sort.SliceStable(activities, func(i, j int) bool {
		return activities[i].Total_Point > activities[j].Total_Point
	})

	var users []entities.User
	var rank = 0

	for _, i := range activities {
		rank++
		user := entities.User{
			First_name:      i.First_Name,
			Country:         i.CountrY,
			Profile_picture: i.Profile_Picture,
			Total_point:     i.Total_Point,
			Rank:            rank,
		}
		users = append(users, user)
	}

	return users, nil
}

func (userModel UserModel) Create(insert_user *entities.Insert_user) (err error) {
	result, err := userModel.Db.Exec("insert into users(id, first_name, last_name, country, profile_picture)values(?,?,?,?,?)",
		insert_user.Id, insert_user.First_name, insert_user.Last_name, insert_user.Country, insert_user.Profile_picture)
	if err != nil {
		return err
	} else {
		result.LastInsertId()
		return nil
	}
}

func (userModel UserModel) Update(insert_user *entities.Insert_user) (int64, error) {
	result, err := userModel.Db.Exec("update users set first_name = ?, last_name = ?, country = ?, profile_picture = ? where id = ?",
		insert_user.First_name, insert_user.Last_name, insert_user.Country, insert_user.Profile_picture, insert_user.Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (userModel UserModel) Delete(id int64) (int64, error) {
	result, err := userModel.Db.Exec("delete from activity_logs where user_id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (userModel UserModel) Delete_user(id int64) (int64, error) {
	result, err := userModel.Db.Exec("delete from users where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
