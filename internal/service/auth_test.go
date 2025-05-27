package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/internal/service"
	"github.com/t-lunch/t-lunch-backend/internal/service/mocks"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type MockAuthService struct {
	mockUserRepo *mocks.MockUserRepo
	mockAuthRepo *mocks.MockAuthRepo
}

func TestRegister(t *testing.T) {
	tests := []struct {
		name        string
		input       *models.User
		mockAuthSrv func(m *MockAuthService, user *models.User)
		expected    *models.User
		err         error
	}{
		{
			name: "success",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {
				m.mockUserRepo.EXPECT().
					IsUserWithEmailExist(gomock.Any(), user.Email).
					Return(false, nil)
				m.mockUserRepo.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(nil)
			},
			expected: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: gomock.Any().String(),
			},
			err: nil,
		},

		{
			name: "repo error: IsUserWithEmailExist",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {
				m.mockUserRepo.EXPECT().
					IsUserWithEmailExist(gomock.Any(), user.Email).
					Return(false, gorm.ErrInvalidField)
			},
			expected: nil,
			err:      errors.NewErrRepository("userRepo", "IsUserWithEmailExist", gorm.ErrInvalidField),
		},

		{
			name: "user with email already exists",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {
				m.mockUserRepo.EXPECT().
					IsUserWithEmailExist(gomock.Any(), user.Email).
					Return(true, nil)
			},
			expected: nil,
			err:      errors.NewErrUserWithEmailAlreadyExists("email-test"),
		},

		{
			name: "repo error: CreateUser",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {
				m.mockUserRepo.EXPECT().
					IsUserWithEmailExist(gomock.Any(), user.Email).
					Return(false, nil)
				m.mockUserRepo.EXPECT().
					CreateUser(gomock.Any(), gomock.Any()).
					Return(gorm.ErrInvalidDB)
			},
			expected: nil,
			err:      errors.NewErrRepository("userRepo", "CreateUser", gorm.ErrInvalidDB),
		},

		{
			name: "invalid password",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {},
			expected:    nil,
			err:         errors.ErrInvalidRequest,
		},

		{
			name: "invalid email",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {},
			expected:    nil,
			err:         errors.ErrInvalidRequest,
		},

		{
			name: "invalid emoji",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {},
			expected:    nil,
			err:         errors.ErrInvalidRequest,
		},

		{
			name: "invalid office",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {},
			expected:    nil,
			err:         errors.ErrInvalidRequest,
		},

		{
			name: "invalid tg",
			input: &models.User{
				Name:           "name-test",
				Surname:        "surname-test",
				Tg:             "",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {},
			expected:    nil,
			err:         errors.ErrInvalidRequest,
		},

		{
			name: "invalid surname",
			input: &models.User{
				Name:           "name-test",
				Surname:        "",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {},
			expected:    nil,
			err:         errors.ErrInvalidRequest,
		},

		{
			name: "invalid name",
			input: &models.User{
				Name:           "",
				Surname:        "surname-test",
				Tg:             "tg-test",
				Office:         "office-test",
				Emoji:          "emoji-test",
				Email:          "email-test",
				HashedPassword: "password-test",
			},
			mockAuthSrv: func(m *MockAuthService, user *models.User) {},
			expected:    nil,
			err:         errors.ErrInvalidRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepo(ctrl)
			mockAuthSrv := MockAuthService{mockUserRepo, nil}

			test.mockAuthSrv(&mockAuthSrv, test.input)

			authService := service.NewAuthService(mockUserRepo, nil)

			user, err := authService.Register(context.Background(), test.input)
			if err == nil {
				user.HashedPassword = gomock.Any().String()
			}

			assert.Equal(t, test.expected, user)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestLogin(t *testing.T) {
	type request struct {
		email    string
		password string
	}
	type response struct {
		accessToken  string
		refreshToken string
		userID       int64
	}

	tests := []struct {
		name         string
		input        request
		mockAuthRepo func(m *MockAuthService, req *request)
		expected     response
		err          error
	}{
		{
			name: "success",
			input: request{
				email:    "email-test",
				password: "password-test",
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.password), bcrypt.DefaultCost)
				m.mockUserRepo.EXPECT().
					GetUserPasswordByEmail(gomock.Any(), req.email).
					Return(int64(1), string(hashedPassword), nil)

				m.mockAuthRepo.EXPECT().
					GenerateToken(gomock.Any(), int64(1), models.Access).
					Return("accessToken-test", nil)

				m.mockAuthRepo.EXPECT().
					GenerateToken(gomock.Any(), int64(1), models.Refresh).
					Return("refreshToken-test", nil)
			},
			expected: response{
				accessToken:  "accessToken-test",
				refreshToken: "refreshToken-test",
				userID:       1,
			},
			err: nil,
		},

		{
			name: "invalid email request",
			input: request{
				email:    "",
				password: "password-test",
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {},
			expected:     response{},
			err:          errors.ErrInvalidRequest,
		},

		{
			name: "invalid password request",
			input: request{
				email:    "email-test",
				password: "",
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {},
			expected:     response{},
			err:          errors.ErrInvalidRequest,
		},

		{
			name: "repo error: GetUserPasswordByEmail",
			input: request{
				email:    "email-test",
				password: "password-test",
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				m.mockUserRepo.EXPECT().
					GetUserPasswordByEmail(gomock.Any(), req.email).
					Return(int64(0), "", gorm.ErrRecordNotFound)
			},
			expected: response{},
			err:      errors.NewErrRepository("userRepo", "GetUserPasswordByEmail", gorm.ErrRecordNotFound),
		},

		{
			name: "incorrect password",
			input: request{
				email:    "email-test",
				password: "password-fail",
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password-test"), bcrypt.DefaultCost)
				m.mockUserRepo.EXPECT().
					GetUserPasswordByEmail(gomock.Any(), req.email).
					Return(int64(1), string(hashedPassword), nil)
			},
			expected: response{},
			err:      errors.NewErrRepository("bcrypt", "CompareHashAndPassword", errors.ErrInvalidPassword),
		},

		{
			name: "repo error: GenerateToken access",
			input: request{
				email:    "email-test",
				password: "password-test",
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.password), bcrypt.DefaultCost)
				m.mockUserRepo.EXPECT().
					GetUserPasswordByEmail(gomock.Any(), req.email).
					Return(int64(1), string(hashedPassword), nil)

				m.mockAuthRepo.EXPECT().
					GenerateToken(gomock.Any(), int64(1), gomock.Any()).
					Return("", errors.ErrUnknownTokenType)
			},
			expected: response{},
			err:      errors.NewErrRepository("authRepo", "GenerateToken access", errors.ErrUnknownTokenType),
		},

		{
			name: "repo error: GenerateToken refresh",
			input: request{
				email:    "email-test",
				password: "password-test",
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.password), bcrypt.DefaultCost)
				m.mockUserRepo.EXPECT().
					GetUserPasswordByEmail(gomock.Any(), req.email).
					Return(int64(1), string(hashedPassword), nil)

				m.mockAuthRepo.EXPECT().
					GenerateToken(gomock.Any(), int64(1), models.Access).
					Return("accessToken-test", nil)

				m.mockAuthRepo.EXPECT().
					GenerateToken(gomock.Any(), int64(1), gomock.Any()).
					Return("", errors.ErrUnknownTokenType)
			},
			expected: response{},
			err:      errors.NewErrRepository("authRepo", "GenerateToken refresh", errors.ErrUnknownTokenType),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepo(ctrl)
			mockAuthRepo := mocks.NewMockAuthRepo(ctrl)
			mockAuthSrv := MockAuthService{mockUserRepo, mockAuthRepo}

			test.mockAuthRepo(&mockAuthSrv, &test.input)

			authService := service.NewAuthService(mockUserRepo, mockAuthRepo)

			accessToken, refreshToken, userID, err := authService.Login(context.Background(), test.input.email, test.input.password)

			assert.Equal(t, test.expected.accessToken, accessToken)
			assert.Equal(t, test.expected.refreshToken, refreshToken)
			assert.Equal(t, test.expected.userID, userID)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestRefresh(t *testing.T) {
	type request struct {
		refreshToken string
		userId       int64
	}

	tests := []struct {
		name         string
		input        request
		mockAuthRepo func(m *MockAuthService, req *request)
		expected     string
		err          error
	}{
		{
			name: "success",
			input: request{
				refreshToken: "refreshToken-test",
				userId:       1,
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				payload := jwt.MapClaims{"id": float64(1), "exp": float64(time.Now().Add(time.Hour).Unix())}
				m.mockAuthRepo.EXPECT().
					GetToken(gomock.Any(), req.refreshToken).
					Return(payload, nil)
				m.mockAuthRepo.EXPECT().
					GenerateToken(gomock.Any(), req.userId, models.Access).
					Return("new_access_token", nil)
			},
			expected: "new_access_token",
			err:      nil,
		},

		{
			name: "invalid request",
			input: request{
				refreshToken: "",
				userId:       1,
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {},
			expected:     "",
			err:          errors.ErrInvalidRequest,
		},

		{
			name: "repo error: GetToken",
			input: request{
				refreshToken: "refreshToken-test",
				userId:       1,
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				m.mockAuthRepo.EXPECT().
					GetToken(gomock.Any(), req.refreshToken).
					Return(jwt.MapClaims{}, errors.ErrInvalidToken)
			},
			expected: "",
			err:      errors.NewErrRepository("authRepo", "GetToken", errors.ErrInvalidToken),
		},

		{
			name: "token expired",
			input: request{
				refreshToken: "refreshToken-test",
				userId:       1,
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				payload := jwt.MapClaims{"id": float64(1), "exp": float64(time.Now().Add(-time.Hour).Unix())}
				m.mockAuthRepo.EXPECT().
					GetToken(gomock.Any(), req.refreshToken).
					Return(payload, nil)
			},
			expected: "",
			err:      errors.NewErrRepository("authRepo", "ValidateToken", errors.ErrTokenExpired),
		},

		{
			name: "owner id mismatch",
			input: request{
				refreshToken: "refreshToken-test",
				userId:       1,
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				payload := jwt.MapClaims{"id": float64(2), "exp": float64(time.Now().Add(time.Hour).Unix())}
				m.mockAuthRepo.EXPECT().
					GetToken(gomock.Any(), req.refreshToken).
					Return(payload, nil)
			},
			expected: "",
			err:      errors.NewErrRepository("authRepo", "GetToken", errors.NewErrUserAndOwnerAreDifferent(1, 2)),
		},

		{
			name: "repo error: GenerateToken access",
			input: request{
				refreshToken: "refreshToken-test",
				userId:       1,
			},
			mockAuthRepo: func(m *MockAuthService, req *request) {
				payload := jwt.MapClaims{"id": float64(1), "exp": float64(time.Now().Add(time.Hour).Unix())}
				m.mockAuthRepo.EXPECT().
					GetToken(gomock.Any(), req.refreshToken).
					Return(payload, nil)
				m.mockAuthRepo.EXPECT().
					GenerateToken(gomock.Any(), req.userId, gomock.Any()).
					Return("", errors.ErrUnknownTokenType)
			},
			expected: "",
			err:      errors.NewErrRepository("authRepo", "GenerateToken access", errors.ErrUnknownTokenType),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockAuthRepo := mocks.NewMockAuthRepo(ctrl)
			mockAuthSrv := MockAuthService{nil, mockAuthRepo}

			test.mockAuthRepo(&mockAuthSrv, &test.input)

			authService := service.NewAuthService(nil, mockAuthRepo)

			accessToken, err := authService.Refresh(context.Background(), test.input.refreshToken, test.input.userId)

			assert.Equal(t, test.expected, accessToken)
			assert.Equal(t, test.err, err)
		})
	}
}
