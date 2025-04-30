package entities

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Area     Area   `json:"area"`
	Role     string `json:"role"`
}
