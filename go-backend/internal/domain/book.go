package domain

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	ISBN   string `json:"isbn"`
}

func (b Book) Validate() error {
	if b.Title == "" {
		return ErrValidation
	}
	if b.Year < 1000 || b.Year > 2026 {
		return ErrValidation
	}
	l := len(b.ISBN)
	if (l != 10 && l != 13) || !isDigits(b.ISBN) {
		return ErrValidation
	}
	return nil
}

func isDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
