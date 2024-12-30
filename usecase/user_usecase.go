package usecase

import (
	"errors"
	"onlineshopbackend/domain"
	"time"
)

type UserUseCase struct {
	UserRepository domain.UserRepository
	contextTimeout time.Duration
	TokenGen       domain.TokenGenerator
	PasswordSvc    domain.PasswordService
}

func NewUserUseCase(userRepository domain.UserRepository, timeout time.Duration, tokenGen domain.TokenGenerator, passwordSvc domain.PasswordService) domain.UserUseCase {
	return &UserUseCase{
		UserRepository: userRepository,
		contextTimeout: timeout,
		TokenGen:       tokenGen,
		PasswordSvc:    passwordSvc,
	}
}

func (uc *UserUseCase) CreateAccount(user domain.User) (domain.User, error) {

	check, _ := uc.UserRepository.GetAllUserByEmial(user.Email)

	if check.Email != "" {
		return domain.User{}, errors.New("email already exist")
	}

	hashedPassword, err := uc.PasswordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}
	user.Password = hashedPassword

	res, err := uc.UserRepository.CreateAccount(user)

	if err != nil {
		return domain.User{}, err
	}

	return res, nil
}

func (uc *UserUseCase) Login(newUser domain.User) (string, error) {

	user, err := uc.UserRepository.GetAllUserByEmial(newUser.Email)
	if err != nil {
		return "", err
	}

	match := uc.PasswordSvc.CheckPasswordHash(newUser.Password, user.Password)

	if !match {
		return "", errors.New("invalid email or password")
	}

	token, err := uc.TokenGen.GenerateToken(user)

	if err != nil {
		return "", err
	}

	user.Password = ""

	return token, nil
}

func (uc *UserUseCase) GetByID(id string) (domain.User, error) {
	user, err := uc.UserRepository.GetByID(id)

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateProfile(id string, user domain.User) (domain.User, error) {

	Password, err := uc.PasswordSvc.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, errors.New("failed to hash password")
	}
	user.Password = Password
	_, err = uc.UserRepository.UpdateProfile(id, user)

	if err != nil {

		return domain.User{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetAllUser() ([]domain.User, error) {
	users, err := uc.UserRepository.GetAllUser()

	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

// func (uc *UserUseCase) GetUserByID(id string) (domain.User, error) {

// 	user, err := uc.UserRepository.GetByID(id)

// 	if err != nil {
// 		return domain.User{}, err

// 	}

// 	return user, nil
// }
