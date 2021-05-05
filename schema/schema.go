package schema

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Team string `json:"team"`
}
