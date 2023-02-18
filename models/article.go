package models

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID           uint
	Title        string `json:"title"`
	Author       string `json:"author"`
	PublishTime  string `json:"publish_time"`
	Category     string `json:"category"`
	PreviewImage string `json:"preview_image"`
	Content      string `json:"content"`
}

func GetArticles(numbers int, maps interface{}) (articles []Article) {
	db.Where(maps).Limit(numbers).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	return
}

func GetArticlesTotal(maps interface{}) (count int64) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func ExistArticleByTitle(title string) bool {
	var article Article
	//fmt.Println(title)
	db.Select("id").Where("title = ?", title).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}

func AddArticle(title, author, publishTime, category, preview, content string) bool {
	db.Create(&Article{
		Title:        title,
		Author:       author,
		PublishTime:  publishTime,
		Category:     category,
		PreviewImage: preview,
		Content:      content,
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(&Article{})
	return true
}

func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}
