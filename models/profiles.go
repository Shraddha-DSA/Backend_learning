package models

type UserURI struct {
	ID int `uri:"id" binding:"required"`
}
type ProfileQuery struct {
	Verbose bool   `form:"verbose"`
	Lang    string `form:"lang"`
}
