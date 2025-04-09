package models

import "time"

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ProfilePic string `json:"profile_pic,omitempty"`
	Password string `json:"password"`
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

type Sketches struct {
    UserId		string	`json:"user_id"`
    Name      string    `json:"name"`
    URL       string    `json:"url"`
    UploadedAt time.Time `json:"uploaded_at"`
}

type Designs struct {
	UserId		string	`json:"user_id"`
	Title		string	`json:"title"`
	Image_URL	string	`json:"image_url"`
	File_URL	string	`json:"file_url"`
}
type Tea struct {
	Type     string
	Category string
	Toppings []string
	Price    float32
}