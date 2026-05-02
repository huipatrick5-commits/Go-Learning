package repository

import (
    "gorm.io/gorm"
)

type User struct {
    ID   int    `gorm:"primaryKey"`
    Name string `gorm:"column:name"`
    Age  int    `gorm:"column:age"`
}

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id int) (User, error) {
    var user User
    result := r.db.First(&user, id)
    if result.Error != nil {
        return User{}, result.Error
    }
    return user, nil
}

func (r *UserRepository) AddUser(user User) error {
    result := r.db.Create(&user)
    return result.Error
}

func (r *UserRepository) DelByID(id int) error {
    result := r.db.Delete(&User{}, id)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (r *UserRepository) ModUser(user User) error {
    result := r.db.Save(&user)
    return result.Error
}