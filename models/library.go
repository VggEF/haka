package models

type Book struct {
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
