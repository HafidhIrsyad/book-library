package category

import "time"

type (
	CategoryInput struct {
		Name        string     `json:"name"`
		Description string     `json:"description"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at"`
	}

	CategoryResponse struct {
		ID          int64      `json:"id"`
		Name        string     `json:"name"`
		Description string     `json:"description"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   *time.Time `json:"updated_at"`
	}

	CategoryResponseJoin struct {
		ID          int64  `json:"category_id"`
		Name        string `json:"category_name"`
		Description string `json:"category_description"`
	}

	CategorySearch struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
)
