package repository

import (
	"github.com/Buddy-Git/golang-git-pro/model/entity"
	"github.com/jinzhu/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Failed to close database")
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Update(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set(`gorm:auto_preload`, true).Find(&videos)
	return videos
}
