package dto

type MovieRequest struct {
	Search string `json:"search" param:"search"`
	Title  string `json:"title" `
	ID     string `json:"id" param:"id"`
}

type CreateMovieRequest struct {
	ID         int    `json:"id" `
	MovieID    string `json:"movie_id"`
	Title      string `json:"title"`
	Year       string `json:"year"`
	Runtime    string `json:"runtime"`
	Released   string `json:"released"`
	Genre      string `json:"genre"`
	Director   string `json:"director"`
	Writer     string `json:"writer"`
	Actors     string `json:"actors"`
	Plot       string `json:"plot"`
	Language   string `json:"language"`
	Country    string `json:"country"`
	Awards     string `json:"awards"`
	Poster     string `json:"poster"`
	ImdbRating string `json:"imdbRating"`
	UserID     int64  `json:"userId"`
}
