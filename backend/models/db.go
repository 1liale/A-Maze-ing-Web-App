package models

import (
	"github.com/gofrs/uuid"
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
	Data      MazeData     `gorm:"type:jsonb"`
	Solution  MazeSolution `gorm:"type:jsonb"`
	SolveTime int
}
