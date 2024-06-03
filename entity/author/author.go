package author

import (
	"time"
)

type (
	AuthorInput struct {
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	AuthorResponse struct {
		ID        int64      `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt *time.Time `json:"updated_at"`
	}

	AuthorResponseJoin struct {
		ID    int64  `json:"author_id"`
		Name  string `json:"author_name"`
		Email string `json:"author_email"`
	}

	AuthorSearch struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
