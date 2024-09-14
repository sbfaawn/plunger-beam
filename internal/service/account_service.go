package service

import (
	"errors"
	"plunger-beam/internal/models"
	"plunger-beam/internal/repository"
	util "plunger-beam/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AccountService interface {
	SaveAccount(ctx *gin.Context, account *models.Account) error
	ChangePassword(ctx *gin.Context, account *models.Account) error
	AccountVerification(ctx *gin.Context, account *models.Account) error
	Login(ctx *gin.Context, account *models.Account) error
}

type accountServiceImpl struct {
	accountRepository repository.AccountRepository
	encryptor         util.PasswordEncryptor
}

func NewAccountService(accountRepository repository.AccountRepository, encryptor util.PasswordEncryptor) AccountService {
	return &accountServiceImpl{
		accountRepository: accountRepository,
		encryptor:         encryptor,
	}
}

func (s *accountServiceImpl) SaveAccount(ctx *gin.Context, account *models.Account) error {
	_, err := s.accountRepository.GetAccountByUsername(ctx, account.Username)

	if err == nil {
		return errors.New("Account with username " + account.Username + " is already exist")
	}

	passEncrypt, err := s.encryptor.Encrypt(account.Password)

	if err != nil {
		return errors.New("Error is occured when encrypt password")
	}

	account.Password = passEncrypt
	err = s.accountRepository.SaveAccount(ctx, account)

	if err != nil {
		return errors.New("Error when trying create an account")
	}

	return nil
}

func (s *accountServiceImpl) ChangePassword(ctx *gin.Context, account *models.Account) error {
	err := s.accountRepository.UpdatePasswordByUsername(ctx, account.Username, account.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *accountServiceImpl) AccountVerification(ctx *gin.Context, account *models.Account) error {
	err := s.accountRepository.UpdateVerifiedByEmail(ctx, account.Email)
	if err != nil {
		return err
	}

	return nil
}

func (s *accountServiceImpl) Login(ctx *gin.Context, account *models.Account) error {
	result, err := s.accountRepository.GetAccountByUsername(ctx, account.Username)

	if err != nil {
		return errors.New("Account with username " + account.Username + " is not exist")
	}

	isMatchPass := s.encryptor.IsHashedPasswordMatch(result.Password, account.Password)

	if !isMatchPass {
		return errors.New("password is not match")
	}

	return nil
}
