package dto

//RegisterDTO is used when client post from /register url
type LoginDTO struct {
	Username	string `json:"username"  binding:"required"`
	Password 	string `json:"password" binding:"required"`
}