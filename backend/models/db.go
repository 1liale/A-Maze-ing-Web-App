package models

import (
	"github.com/google/uuid"
)

// has-many(1:N) relationship between User and MazeRecord

type User struct {
	ID               uuid.UUID `gorm:"PRIMARY_KEY"`
	Name             string    `gorm:"unique"`
	FastestSolveTime int
	Records          []MazeRecord
}

type MazeRecord struct {
	ID        uuid.UUID `gorm:"PRIMARY_KEY"`
	UserId    uuid.UUID
	Data      MazeData     `gorm:"serializer:json"`
	Solution  MazeSolution `gorm:"serializer:json"`
	SolveTime int
}
