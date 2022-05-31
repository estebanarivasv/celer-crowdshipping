package entities

type PackageInDTO struct {
	Name        string  `json:"name"`
	ImageURL    string  `json:"imageUrl"`
	Description string  `json:"description"`
	Dimensions  string  `json:"dimensions"`
	Value       float64 `json:"value"`
}

type PackageOutDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	ImageURL    string  `json:"imageUrl"`
	Description string  `json:"description"`
	Dimensions  string  `json:"dimensions"`
	Value       float64 `json:"value"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
	// DeletedAt time.Time `json:"deletedAt"`
}
