package models

type Student struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
