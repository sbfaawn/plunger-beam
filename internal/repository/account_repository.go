package repository

import (
	"errors"
	"fmt"
	"plunger-beam/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccountRepository interface {
	SaveAccount(ctx *gin.Context, account *models.Account) error
	GetAccountByUsername(ctx *gin.Context, username string) (models.Account, error)
	GetAccountByEmail(ctx *gin.Context, email string) (models.Account, error)
	UpdatePasswordByUsername(ctx *gin.Context, username string, newPassword string) error
	UpdateVerifiedByEmail(ctx *gin.Context, email string) error
}

type accountRepositoryImpl struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) AccountRepository {
	return &accountRepositoryImpl{
		DB: DB,
	}
}

func (storage *accountRepositoryImpl) SaveAccount(ctx *gin.Context, account *models.Account) error {
	var err error
	db := storage.DB

	tx := db.Begin()
	err = tx.Create(&account).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (storage *accountRepositoryImpl) GetAccountByUsername(ctx *gin.Context, username string) (models.Account, error) {
	db := storage.DB
	var err error
	var account models.Account

	fmt.Println("username : " + username)
	tx := db.Begin()
	err = tx.First(&account, "username = ?", username).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	return account, nil
}

func (storage *accountRepositoryImpl) UpdatePasswordByUsername(ctx *gin.Context, username string, newPassword string) error {
	db := storage.DB
	var err error

	tx := db.Begin()
	update := tx.Model(&models.Account{}).Where("username = ? AND deleted_at IS null", username).Updates(map[string]any{
		"password": newPassword,
	})

	err = update.Error
	if err != nil {
		tx.Rollback()
		return err
	}

	result := update.Commit().WithContext(ctx)

	if result.Error != nil {
		tx.Rollback()
		return err
	}

	if update.RowsAffected == result.RowsAffected {
		tx.Rollback()
		return errors.New("record with username " + username + " is not found")
	}

	return nil
}

func (storage *accountRepositoryImpl) GetAccountByEmail(ctx *gin.Context, email string) (models.Account, error) {
	db := storage.DB
	var err error
	var account models.Account

	tx := db.Begin()
	err = tx.First(&account, "email = ?", email).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	err = tx.Commit().WithContext(ctx).Error

	if err != nil {
		tx.Rollback()
		return account, err
	}

	return account, nil
}

func (storage *accountRepositoryImpl) UpdateVerifiedByEmail(ctx *gin.Context, email string) error {
	db := storage.DB
	var err error

	tx := db.Begin()
	update := tx.Model(&models.Account{}).Where("email = ? AND deleted_at IS null", email).Updates(map[string]any{
		"verified": true,
	})

	err = update.Error
	if err != nil {
		tx.Rollback()
		return err
	}

	result := update.Commit().WithContext(ctx)

	if result.Error != nil {
		tx.Rollback()
		return err
	}

	if update.RowsAffected == result.RowsAffected {
		tx.Rollback()
		return errors.New("record with email " + email + " is not found")
	}

	return nil
}
