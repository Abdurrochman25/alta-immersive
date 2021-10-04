package book

type PostBookRequest struct {
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}

type EditBookRequest struct {
	Title  string `json:"title" form:"title"`
	Author string `json:"author" form:"author"`
}
