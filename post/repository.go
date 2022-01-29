package post

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Post, error)
	FindByID(ID int) (Post, error)
	Save(post Post) (Post, error)
	Update(post Post) (Post, error)
	Delete(ID int) (Post, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Post, error) {
	var posts []Post

	err := r.db.Find(&posts).Error
	if err != nil {
		return posts, err
	}

	return posts, nil
}
func (r *repository) FindByID(ID int) (Post, error) {
	var post Post

	err := r.db.Where("id = ?", ID).Find(&post).Error

	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) Save(post Post) (Post, error) {
	err := r.db.Create(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) Update(post Post) (Post, error) {
	err := r.db.Save(&post).Error

	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) Delete(ID int) (Post, error) {
	var post Post
	err := r.db.Where("id = ?", ID).Delete(&post).Error

	if err != nil {
		return post, err
	}

	return post, nil
}
