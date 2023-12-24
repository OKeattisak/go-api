package main

import "github.com/google/uuid"

func main() {
	uid := uuid.New().String()
	println(uid)
}
