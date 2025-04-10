package models

import "time"

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ProfilePic string `json:"profile_pic,omitempty"`
	Password string `json:"password"`
}

type Category struct {
	Blogs   *Blog	`json:"blog"`
	Links	string	`json:"links"`
	Photos	string	`json:"photos"`
	Files	string 	`json:"files"`
}

type Blog struct {
	UserId string	 `json:"user_id"`
	Title string	 `json:"title"`
	Desc string 	`json:"desc"`
	Link string 	`json:"link"`
    UploadedAt time.Time `json:"uploaded_at"`

}

type Github struct {
	UserId		string 	`json:"user_id"`
	ProjectName string	`json:"project_name"`
	Desc		string	`json:"desc"`
	Link		string	`json:"link"`
    UploadedAt time.Time `json:"uploaded_at"`

}


type Designs struct {
	UserId		string	`json:"user_id"`
	Title		string	`json:"title"`
	Image_URL	string	`json:"image_url"`
	File_URL	string	`json:"file_url"`
}


type CollectionList struct {
	Name	string
	Category *Category

}



type Links struct {
	Name	string		`json:"name"`
	Links	string		`json:"link"`
}
type