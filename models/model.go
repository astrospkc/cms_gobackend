package models

import "time"

type User struct {
	Name string `json:"name"`
	Email string `json:"email"`
	ProfilePic string `json:"profile_pic,omitempty"`
	Password string `json:"password"`
	Role 		string	`json:"role"`
	
}

type Project struct{
	
	UserId		string 	`json:"user_id"`
	Title		string	`json:"title"`
	Description	string	`json:"description"`
	Tags		string	`json:"tags,omitempty"`
	Thumbnail 	string	`json:"thumbnail,omitempty"`
	GithubLink	string	`json:"githublink,omitempty"`
	LiveDemoLink string	`json:"liveddemolink,omitempty"`
	CreatedAt	time.Time	`json:"time"`
}

type Category struct {
	Blogs   *Blog	`json:"blog"`
	Links	string	`json:"links"`
	Photos	string	`json:"photos"`
	Files	string 	`json:"files"`
}

type Blog struct{
	
	UserId		string 	`json:"user_id"`
	Title		string	`json:"title"`
	Content		string	`json:"content"`
	Tags		string	`json:"tags,omitempty"`
	CoverImage 	string	`json:"coverImage,omitempty"`
	Published 	time.Time	`json:"published "`
	CreatedAt	time.Time	`json:"time"`
}

type Media struct{
	
	UserId		string 	`json:"user_id"`
	FileUrl		string	`json:"fileurl"`
	Type		string	`json:"content"` 
	FileName	string	`json:"filename"`
	CreatedAt	time.Time	`json:"time"`
}
// type (image, video, audio, doc, pdf, etc.)
type Link struct{
	
	UserId		string 	`json:"user_id"`
	Title		string	`json:"title"`
	Url		string	`json:"url"` 
	Description	string	`json:"description"`
	Category	string	`json:"category"`
}
// category (e.g., Social, Project, Resume)

type Resume struct {
	UserId 		string	`json:"user_id"`
	FileUrl		string	`json:"fileurl"`
	UploadData	time.Time	`json:"uploadDate"`
}

type SubscriptionPlan struct {
	Name	string		`json:"name"`
	Price 	string		`json:"price"`
	Features []string	`json:"features"`
	Duration string		`json:"duration"`
}

type UserSubscription struct {
	Userid	string	`json:"user_id"`
	PlanId 	string	`json:"plan_id"`
	StartDate	time.Time	`json:"startDate"`
	EndDate		time.Time	`json:"endData"`
	Status	string	`json:"status"`
}

// status- active or expired
type APIkey struct{
	Userid		string	`json:"user_id"`
	Key 		string	`json:"key"`
	UsageLimit	string	`json:"usagelimit"`
	CreatedAt	time.Time	`json:"createdat"`
	Revoked		bool	`json:"revoked"`
	
}