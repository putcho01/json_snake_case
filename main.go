package main

//go:generate go run cmd/json-snake-case/main.go -out ./gen/json_gen.go -type=User
type User struct {
	UserID string
	Name   string
}

func main() {
	// Do nothing
}
