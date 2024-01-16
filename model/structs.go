package model

type KeyValuePair struct {
	Key   string
	Value string
}

type User struct {
	Email string `json:"email"`
}

type InputProps struct {
	InputType string
	Name      string
}
