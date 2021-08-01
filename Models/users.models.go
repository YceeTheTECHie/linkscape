package Models

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
	Password  int64  `json:"password"`
}
