package model

type User struct {
	ID                      int64       `json:id`
	UID                     string      `json:uid`
	Password                string      `json:password`
	First_name              string      `json:first_name`
	Last_name               string      `json:last_name`
	Username                string      `json:username`
	Email                   string      `json:email`
	Avatar                  string      `json:avatar`
	Gender                  string      `json:gender`
	Phone_number            string      `json:phone_number`
	Social_insurance_number string      `json:social_insurance_number`
	Date_of_birth           string      `json:date_of_birth`
	Employment              Employment  `json:"employment"`
	Address                 Address     `json:"address"`
	Credit_card             Credit_card `json:Credit_card`
}
