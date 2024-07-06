package service

import (
	"errors"
	"fungicheck/module/entities"
	"fungicheck/module/feature/auth/domain"
	"fungicheck/module/feature/auth/mocks"
	utils "fungicheck/utils/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTest(t *testing.T) (*mocks.AuthRepositoryInterface, domain.AuthServiceInterface, *utils.HashInterface, *utils.JWTInterface) {
	repo := mocks.NewAuthRepositoryInterface(t)
	hash := utils.NewHashInterface(t)
	jwt := utils.NewJWTInterface(t)
	service := NewAuthService(repo, hash, jwt)
	return repo, service, hash, jwt
}

func TestLogin(t *testing.T) {
	email := "test@example.com"
	password := "password123"

	t.Run("Success Case - Valid Credentials", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedUser := &entities.UserModels{
			ID:       1,
			Email:    email,
			Password: "hashedPassword",
			Role:     "customer",
		}
		expectedToken := "mockedAccessToken"

		repo.On("GetUsersByEmail", email).Return(expectedUser, nil)
		hash.On("ComparePassword", expectedUser.Password, password).Return(true, nil)
		jwt.On("GenerateJWT", expectedUser.ID, expectedUser.Email, expectedUser.Role).Return(expectedToken, nil)

		user, accessToken, err := service.Login(email, password)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
		assert.Equal(t, expectedToken, accessToken)

		repo.AssertExpectations(t)
		hash.AssertExpectations(t)
		jwt.AssertExpectations(t)
	})

	t.Run("Error Case - User Not Found", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedErr := errors.New("user not found")
		repo.On("GetUsersByEmail", email).Return(nil, expectedErr)

		user, accessToken, err := service.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "user not found")
		assert.Equal(t, "", accessToken)

		repo.AssertExpectations(t)
		hash.AssertNotCalled(t, "ComparePassword")
		jwt.AssertNotCalled(t, "GenerateJWT")
	})

	t.Run("Error Case - Invalid Credentials", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedUser := &entities.UserModels{
			ID:       1,
			Email:    email,
			Password: "hashedPassword",
		}
		repo.On("GetUsersByEmail", email).Return(expectedUser, nil)

		hash.On("ComparePassword", expectedUser.Password, password).Return(false, nil)

		user, accessToken, err := service.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "wrong credential")
		assert.Equal(t, "", accessToken)

		repo.AssertExpectations(t)
		hash.AssertExpectations(t)
		jwt.AssertNotCalled(t, "GenerateJWT")
	})

	t.Run("Error Case - JWT Generation Failure", func(t *testing.T) {
		repo, service, hash, jwt := setupTest(t)
		expectedUser := &entities.UserModels{
			ID:       1,
			Email:    email,
			Password: "hashedPassword",
		}
		repo.On("GetUsersByEmail", email).Return(expectedUser, nil)

		hash.On("ComparePassword", expectedUser.Password, password).Return(true, nil)

		jwt.On("GenerateJWT", expectedUser.ID, expectedUser.Email, expectedUser.Role).Return("", errors.New("jwt generation failed"))

		user, accessToken, err := service.Login(email, password)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, "jwt generation failed")
		assert.Equal(t, "", accessToken)

		repo.AssertExpectations(t)
		hash.AssertExpectations(t)
		jwt.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	email := "test@example.com"
	password := "password123"
	name := "John Doe"
	role := "user"

	t.Run("Success Case - New User Registration", func(t *testing.T) {
		repo, service, hash, _ := setupTest(t)
		req := &domain.RegisterRequest{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     role,
		}
		hashedPassword := "hashedPassword"
		expectedUser := &entities.UserModels{
			Email:    email,
			Password: hashedPassword,
			Name:     name,
			Role:     role,
		}

		repo.On("GetUsersByEmail", email).Return(nil, nil)
		hash.On("GenerateHash", password).Return(hashedPassword, nil)
		repo.On("CreateUser", expectedUser).Return(expectedUser, nil)

		user, err := service.Register(req)

		assert.Nil(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)

		repo.AssertExpectations(t)
		hash.AssertExpectations(t)
	})

	t.Run("Error Case - Email Already Exists", func(t *testing.T) {
		repo, service, _, _ := setupTest(t)
		req := &domain.RegisterRequest{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     role,
		}
		existingUser := &entities.UserModels{
			Email: email,
		}
		expectedErr := errors.New("email already exists")

		repo.On("GetUsersByEmail", email).Return(existingUser, nil)

		user, err := service.Register(req)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, expectedErr.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Error Case - Hashing Password Failure", func(t *testing.T) {
		repo, service, hash, _ := setupTest(t)
		req := &domain.RegisterRequest{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     role,
		}
		expectedErr := errors.New("hashing password failed")

		repo.On("GetUsersByEmail", email).Return(nil, nil)
		hash.On("GenerateHash", password).Return("", expectedErr)

		user, err := service.Register(req)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, expectedErr.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Error Case - User Creation Failure", func(t *testing.T) {
		repo, service, hash, _ := setupTest(t)
		req := &domain.RegisterRequest{
			Email:    email,
			Password: password,
			Name:     name,
			Role:     role,
		}
		hashedPassword := "hashedPassword"
		expectedUser := &entities.UserModels{
			Email:    email,
			Password: hashedPassword,
			Name:     name,
			Role:     role,
		}
		expectedErr := errors.New("user creation failed")

		repo.On("GetUsersByEmail", email).Return(nil, nil)
		hash.On("GenerateHash", password).Return(hashedPassword, nil)
		repo.On("CreateUser", expectedUser).Return(nil, expectedErr)

		user, err := service.Register(req)

		assert.Error(t, err)
		assert.Nil(t, user)
		assert.EqualError(t, err, expectedErr.Error())

		repo.AssertExpectations(t)
	})
}
