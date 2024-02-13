package models

// has-many(1:N) relationship between User and MazeRecord

type User struct {
	ID        string `gorm:"primaryKey; not null"`
	Name      string `gorm:"primaryKey; not null"`
	Highscore int
	Records   []MazeRecord `gorm:"ForeignKey:UserId,UserName;References:ID,Name"`
}

type MazeRecord struct {
	ID       uint `gorm:"primaryKey; not null"`
	UserId   string
	UserName string
	Data     MazeData     `gorm:"serializer:json"`
	Solution MazeSolution `gorm:"serializer:json"`
	Score    int
}
