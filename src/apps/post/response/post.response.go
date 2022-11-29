package response

type PostResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
	Image   string `json:"image"`
	UserId  int    `json:"user_id"`
}

func NewPostResponse() PostResponse {
	return PostResponse{}
}
