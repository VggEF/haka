package library

type AddBookRequest struct {
	Title       string   `json:"title" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Publisher   string   `json:"publisher"`
	Year        string   `json:"year"`
	ISBN        string   `json:"isbn"`
	Description string   `json:"description"`
	CoverURL    string   `json:"cover_url"`
	FilePath    string   `json:"file_path"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
}

type UpdateBookRequest struct {
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Publisher   string   `json:"publisher"`
	Year        string   `json:"year"`
	ISBN        string   `json:"isbn"`
	Description string   `json:"description"`
	CoverURL    string   `json:"cover_url"`
	FilePath    string   `json:"file_path"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
}

type BookResponse struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	Author          string   `json:"author"`
	Publisher       string   `json:"publisher"`
	Year            string   `json:"year"`
	ISBN            string   `json:"isbn"`
	Description     string   `json:"description"`
	CoverURL        string   `json:"cover_url"`
	FilePath        string   `json:"file_path"`
	Category        string   `json:"category"`
	Tags            []string `json:"tags"`
	AvailableCopies int      `json:"available_copies"`
}

type SearchBooksQuery struct {
	Query  string `form:"q"`
	Author string `form:"author"`
	Year   string `form:"year"`
	Limit  int    `form:"limit"`
	Offset int    `form:"offset"`
}

type ExternalBook struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Publisher   string `json:"publisher"`
	Year        string `json:"year"`
	Description string `json:"description"`
	CoverURL    string `json:"cover_url"`
}
