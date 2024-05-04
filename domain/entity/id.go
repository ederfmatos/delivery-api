package entity

import "github.com/google/uuid"

func NewId() string {
	id, _ := uuid.NewV7()
	return id.String()
}
