package entities

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Area     int    `json:"area"`
	Role     int    `json:"role"`
}
