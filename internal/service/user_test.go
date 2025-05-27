package service_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/internal/service"
	"github.com/t-lunch/t-lunch-backend/internal/service/mocks"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestGetUsersByIDs(t *testing.T) {
	tests := []struct {
		name         string
		input        []int64
		mockUserRepo func(m *mocks.MockUserRepo, ids []int64)
		expected     []*models.UserResponse
		err          error
	}{
		{
			name:  "success",
			input: []int64{1, 2, 3},
			mockUserRepo: func(m *mocks.MockUserRepo, ids []int64) {
				m.EXPECT().GetUsersByIDs(gomock.Any(), ids).
					Return([]*models.UserResponse{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					}, nil)
			},
			expected: []*models.UserResponse{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			err: nil,
		},

		{
			name:         "empty input",
			input:        []int64{},
			mockUserRepo: func(m *mocks.MockUserRepo, ids []int64) {},
			expected:     nil,
			err:          errors.ErrInvalidRequest,
		},

		{
			name:         "invalid id",
			input:        []int64{231, 0, 77},
			mockUserRepo: func(m *mocks.MockUserRepo, ids []int64) {},
			expected:     nil,
			err:          errors.ErrInvalidRequest,
		},

		{
			name:  "repo error: GetUsersByIDs",
			input: []int64{4, 5, 6},
			mockUserRepo: func(m *mocks.MockUserRepo, ids []int64) {
				m.EXPECT().GetUsersByIDs(gomock.Any(), ids).
					Return(nil, gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("userRepo", "GetUsersByIDs", gorm.ErrRecordNotFound),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepo(ctrl)
			test.mockUserRepo(mockUserRepo, test.input)

			userService := service.NewUserService(mockUserRepo)

			users, err := userService.GetUsersByIDs(context.Background(), test.input)
			assert.Equal(t, test.expected, users)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	tests := []struct {
		name         string
		input        int64
		mockUserRepo func(m *mocks.MockUserRepo, id int64)
		expected     *models.UserResponse
		err          error
	}{
		{
			name:  "success",
			input: 1,
			mockUserRepo: func(m *mocks.MockUserRepo, id int64) {
				m.EXPECT().GetUserByID(gomock.Any(), id).
					Return(&models.UserResponse{ID: 1}, nil)
			},
			expected: &models.UserResponse{ID: 1},
			err:      nil,
		},

		{
			name:         "invalid id",
			input:        0,
			mockUserRepo: func(m *mocks.MockUserRepo, id int64) {},
			expected:     nil,
			err:          errors.ErrInvalidRequest,
		},

		{
			name:  "repo error: GetUserByID",
			input: 7,
			mockUserRepo: func(m *mocks.MockUserRepo, id int64) {
				m.EXPECT().GetUserByID(gomock.Any(), id).
					Return(nil, gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("userRepo", "GetUserByID", gorm.ErrRecordNotFound),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepo(ctrl)
			test.mockUserRepo(mockUserRepo, test.input)

			userService := service.NewUserService(mockUserRepo)

			users, err := userService.GetUserByID(context.Background(), test.input)
			assert.Equal(t, test.expected, users)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name         string
		input        *models.UserResponse
		mockUserRepo func(m *mocks.MockUserRepo, user *models.UserResponse)
		expected     *models.UserResponse
		err          error
	}{
		{
			name: "success",
			input: &models.UserResponse{
				ID:      1,
				Name:    "new-name",
				Surname: "new-surname",
				Tg:      "new-tg",
				Office:  "new-office",
				Emoji:   "new-emoji",
			},
			mockUserRepo: func(m *mocks.MockUserRepo, user *models.UserResponse) {
				m.EXPECT().UpdateUserByID(gomock.Any(), user.ID, map[string]interface{}{
					"name":    user.Name,
					"surname": user.Surname,
					"tg":      user.Tg,
					"office":  user.Office,
					"emoji":   user.Emoji,
				}).Return(nil)

				m.EXPECT().GetUserByID(gomock.Any(), user.ID).Return(user, nil)
			},
			expected: &models.UserResponse{
				ID:      1,
				Name:    "new-name",
				Surname: "new-surname",
				Tg:      "new-tg",
				Office:  "new-office",
				Emoji:   "new-emoji",
			},
			err: nil,
		},

		{
			name:         "invalid id",
			input:        &models.UserResponse{ID: 0},
			mockUserRepo: func(m *mocks.MockUserRepo, user *models.UserResponse) {},
			expected:     nil,
			err:          errors.ErrInvalidRequest,
		},

		{
			name: "repo error: UpdateUserByID",
			input: &models.UserResponse{
				ID:      7,
				Name:    "new-name",
				Surname: "new-surname",
				Tg:      "new-tg",
				Office:  "new-office",
				Emoji:   "new-emoji",
			},
			mockUserRepo: func(m *mocks.MockUserRepo, user *models.UserResponse) {
				m.EXPECT().UpdateUserByID(gomock.Any(), user.ID, map[string]interface{}{
					"name":    user.Name,
					"surname": user.Surname,
					"tg":      user.Tg,
					"office":  user.Office,
					"emoji":   user.Emoji,
				}).Return(gorm.ErrInvalidData)
			},
			expected: nil,
			err:      errors.NewErrRepository("userRepo", "UpdateUserByID", gorm.ErrInvalidData),
		},

		{
			name: "repo error: GetUserByID",
			input: &models.UserResponse{
				ID:      7,
				Name:    "new-name",
				Surname: "new-surname",
				Tg:      "new-tg",
				Office:  "new-office",
				Emoji:   "new-emoji",
			},
			mockUserRepo: func(m *mocks.MockUserRepo, user *models.UserResponse) {
				m.EXPECT().UpdateUserByID(gomock.Any(), user.ID, map[string]interface{}{
					"name":    user.Name,
					"surname": user.Surname,
					"tg":      user.Tg,
					"office":  user.Office,
					"emoji":   user.Emoji,
				}).Return(nil)

				m.EXPECT().GetUserByID(gomock.Any(), user.ID).Return(user, gorm.ErrInvalidData)
			},
			expected: nil,
			err:      errors.NewErrRepository("userRepo", "GetUserByID", gorm.ErrInvalidData),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepo(ctrl)
			test.mockUserRepo(mockUserRepo, test.input)

			userService := service.NewUserService(mockUserRepo)

			users, err := userService.UpdateUser(context.Background(), test.input)
			assert.Equal(t, test.expected, users)
			assert.Equal(t, test.err, err)
		})
	}
}
