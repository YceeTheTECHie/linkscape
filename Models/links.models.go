package Models

type Link struct {
    ID       	int64  `json:"id"`
    Title     	string `json:"title"`
    Link  		string `json:"link"`
	CategoryId   int64  `json:"categoryId"`
	UserId		 int64	`json:"userId"`
}