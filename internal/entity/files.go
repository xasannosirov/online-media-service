// Package entity defines main entities for business logic (services), database mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// File -.
type File struct {
	FileKey  string `json:"file_key"      example:"0b4a4e6e-e502-454f-a8a0-fb7d445e5efa"`
	Filename string `json:"file_name"     example:"template.png"`
	FileURL  string `json:"file_url"      example:"https://example.com/images/template.png"`
}
