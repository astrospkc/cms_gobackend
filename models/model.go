package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Name 		string `bson:"name" json:"name"`
	Email 		string `bson:"email" json:"email"`
	ProfilePic  string `bson:"profile_pic,omitempty" json:"profile_pic"`
	Password 	string `bson:"password" json:"password"`
	Role 		string	`bson:"role" json:"role"`
	
}

type Project struct{
	Id           primitive.ObjectID `bson:"id,omitempty" json:"id"`
	UserId		 primitive.ObjectID 	`bson:"user_id" json:"user_id"`
	Title		 string	`bson:"title" json:"title"`
	Description	 string	`bson:"description,omitempty" json:"description"`
	Tags		 string	`bson:"tags,omitempty" json:"tags"`
	Thumbnail 	 string	`bson:"thumbnail,omitempty" json:"thumbnail"`
	GithubLink	 string	`bson:"githublink,omitempty" json:"githublink"`
	LiveDemoLink string	`bson:"livedemolink,omitempty" json:"liveddemolink"`
	CreatedAt	time.Time	`bson:"time" json:"time"`
}

type Category struct {
	Blogs   Blog	`bson:"blog" json:"blog"`
	Links	string	`bson:"links" json:"links"`
	Photos	string	`bson:"photos" json:"photos"`
	Files	string 	`bson:"files" json:"files"`
}

type Blog struct{
	
	UserId		string 	`bson:"user_id" json:"user_id"`
	Title		string	`bson:"title" json:"title"`
	Content		string	`bson:"content" json:"content"`
	Tags		string	`bson:"tags,omitempty" json:"tags"`
	CoverImage 	string	`bson:"coverImage,omitempty" json:"coverImage"`
	Published 	time.Time	`bson:"published" json:"published "`
	CreatedAt	time.Time	`bson:"time" jsonjson:"time"`
}

type Media struct{
	
	UserId		string 	`bson:"user_id" json:"user_id"`
	FileUrl		string	`bson:"fileurl" json:"fileurl"`
	Type		string	`bson:"content" json:"content"` 
	FileName	string	`bson:"filename" json:"filename"`
	CreatedAt	time.Time	`bson:"time" json:"time"`
}
// type (image, video, audio, doc, pdf, etc.)
type Link struct{
	
	UserId		string 	`bson:"user_id" json:"user_id"`
	Title		string	`bson:"title" json:"title"`
	Url			string	`bson:"url" json:"url"` 
	Description	string	`bson:"description" json:"description"`
	Category	string	`bson:"category" json:"category"`
}
// category (e.g., Social, Project, Resume)

type Resume struct {
	UserId 		string	`bson:"user_id" json:"user_id"`
	FileUrl		string	`bson:"fileurl" json:"fileurl"`
	UploadData	time.Time	`bson:"uploadData" json:"uploadData"`
}

type SubscriptionPlan struct {
	Name		string		`bson:"name" json:"name"`
	Price 		string		`bson:"price" json:"price"`
	Features 	[]string	`bson:"features" json:"features"`
	Duration 	string		`bson:"duration" json:"duration"`
}

type UserSubscription struct {
	Userid		string		`bson:"user_id" json:"user_id"`
	PlanId 		string		`bson:"plan_id" json:"plan_id"`
	StartDate	time.Time	`bson:"startDate" json:"startDate"`
	EndDate		time.Time	`bson:"endDate" json:"endDate"`
	Status		string		`bson:"status" json:"status"`
}

// status- active or expired
type APIkey struct{
	Userid		string	`bson:"user_id" json:"user_id"`
	Key 		string	`bson:"key" json:"key"`
	UsageLimit	string	`bson:"usagelimit" json:"usagelimit"`
	CreatedAt	time.Time	`bson:"createdat" json:"createdat"`
	Revoked		bool	`bson:"revoked" json:"revoked"`
	
}