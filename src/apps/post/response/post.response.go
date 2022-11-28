package response

type PostResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
	Image   string `json:"image"`
	Name    string `json:"name"`
	UserResponse
}

type UserResponse struct {
	UserId int    `json:"user_id"`
	Phone  string `json:"phone"`
}

func NewPostResponse() PostResponse {
	return PostResponse{}
}
