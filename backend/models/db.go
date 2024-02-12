package models

import (
	"github.com/google/uuid"
)

// has-many(1:N) relationship between User and MazeRecord

type User struct {
	ID               uuid.UUID `gorm:"primaryKey; not null"`
	Name             string    `gorm:"primaryKey; not null"`
	FastestSolveTime int
	// Records          []MazeRecord `gorm:"foreignKey:UserId,UserName;references:ID,Name"`
}

type MazeRecord struct {
	ID        uuid.UUID `gorm:"primaryKey; not null"`
	UserId    uuid.UUID
	UserName  string
	User      User         `gorm:"ForeignKey:UserId,UserName;References:ID,Name"`
	Data      MazeData     `gorm:"serializer:json"`
	Solution  MazeSolution `gorm:"serializer:json"`
	SolveTime int
}
