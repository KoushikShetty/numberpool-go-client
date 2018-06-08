package numberpool

// Resource - Numberpool resource
type Resource struct {
	ID                     string        `json:"id"`
	Status                 string        `json:"status"`
	Name                   string        `json:"name"`
	SubAccount             string        `json:"sub_account"`
	TotalCount             string        `json:"total_number_count"`
	RequestURL             string        `json:"resource_url"`
	PhoneNumberResourceURL string        `json:"phone_numbers_resource_uri"`
	Composition            []Composition `json:"composition"`
}
