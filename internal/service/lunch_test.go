package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/t-lunch/t-lunch-backend/internal/errors"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	"github.com/t-lunch/t-lunch-backend/internal/service"
	"github.com/t-lunch/t-lunch-backend/internal/service/mocks"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func TestCreateLunch(t *testing.T) {
	type request struct {
		userID      int64
		place       string
		lunchTime   time.Time
		description string
	}

	tests := []struct {
		name          string
		input         request
		mockLunchRepo func(m *mocks.MockLunchRepo, req *request)
		expected      *models.Lunch
		err           error
	}{
		{
			name: "success",
			input: request{
				userID:      1,
				place:       "place",
				lunchTime:   time.Now(),
				description: "description",
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().CreateLunch(gomock.Any(), &models.Lunch{
					CreatorID:            req.userID,
					Place:                req.place,
					Time:                 req.lunchTime,
					NumberOfParticipants: 1,
					Description:          req.description,
					Participants:         []int64{req.userID},
					LikedBy:              []int64{},
				}).Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), gomock.Any()).
					Return(&models.Lunch{ID: 1}, nil)
			},
			expected: &models.Lunch{ID: 1},
			err:      nil,
		},

		{
			name: "invalid userId",
			input: request{
				userID:      0,
				place:       "place",
				lunchTime:   time.Now(),
				description: "description",
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "invalid place",
			input: request{
				userID:      1,
				place:       "",
				lunchTime:   time.Now(),
				description: "description",
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "repo error: CreateLunch",
			input: request{
				userID:      1,
				place:       "place",
				lunchTime:   time.Now(),
				description: "description",
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().CreateLunch(gomock.Any(), &models.Lunch{
					CreatorID:            req.userID,
					Place:                req.place,
					Time:                 req.lunchTime,
					NumberOfParticipants: 1,
					Description:          req.description,
					Participants:         []int64{req.userID},
					LikedBy:              []int64{},
				}).Return(gorm.ErrInvalidData)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "CreateLunch", gorm.ErrInvalidData),
		},

		{
			name: "repo error: GetLunchByID",
			input: request{
				userID:      1,
				place:       "place",
				lunchTime:   time.Now(),
				description: "description",
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().CreateLunch(gomock.Any(), &models.Lunch{
					CreatorID:            req.userID,
					Place:                req.place,
					Time:                 req.lunchTime,
					NumberOfParticipants: 1,
					Description:          req.description,
					Participants:         []int64{req.userID},
					LikedBy:              []int64{},
				}).Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), gomock.Any()).
					Return(nil, gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "GetLunchByID", gorm.ErrRecordNotFound),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLunchRepo := mocks.NewMockLunchRepo(ctrl)
			test.mockLunchRepo(mockLunchRepo, &test.input)

			lunchService := service.NewLunchService(mockLunchRepo)

			lunch, err := lunchService.CreateLunch(
				context.Background(),
				test.input.userID,
				test.input.place,
				test.input.lunchTime,
				test.input.description,
			)
			assert.Equal(t, test.expected, lunch)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestGetLunches(t *testing.T) {
	type request struct {
		userID int64
		offset int
		limit  int
	}

	tests := []struct {
		name          string
		input         request
		mockLunchRepo func(m *mocks.MockLunchRepo, req *request)
		expected      []*models.Lunch
		expectedID    int64
		err           error
	}{
		{
			name: "success #1",
			input: request{
				userID: 1,
				offset: 0,
				limit:  3,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().GetLunchIdByUserID(gomock.Any(), req.userID).
					Return(int64(1), nil)

				m.EXPECT().GetLunches(gomock.Any(), req.userID, req.offset, req.limit).
					Return([]*models.Lunch{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					}, nil)
			},
			expected: []*models.Lunch{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			expectedID: 1,
			err:        nil,
		},

		{
			name: "success #2",
			input: request{
				userID: 1,
				offset: -3,
				limit:  3,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().GetLunchIdByUserID(gomock.Any(), req.userID).
					Return(int64(0), gorm.ErrRecordNotFound)

				m.EXPECT().GetLunches(gomock.Any(), req.userID, 0, req.limit).
					Return([]*models.Lunch{
						{ID: 1},
						{ID: 2},
						{ID: 3},
					}, nil)
			},
			expected: []*models.Lunch{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			expectedID: 0,
			err:        nil,
		},

		{
			name: "invalid userID",
			input: request{
				userID: 0,
				offset: 0,
				limit:  3,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			expectedID:    0,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "repo error: GetLunchIdByUserID",
			input: request{
				userID: 1,
				offset: 0,
				limit:  100,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().GetLunchIdByUserID(gomock.Any(), req.userID).
					Return(int64(0), gorm.ErrInvalidData)
			},
			expected:   nil,
			expectedID: 0,
			err:        errors.NewErrRepository("lunchRepo", "GetLunchIdByUserID", gorm.ErrInvalidData),
		},

		{
			name: "repo error: GetLunches",
			input: request{
				userID: 1,
				offset: 0,
				limit:  0,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().GetLunchIdByUserID(gomock.Any(), req.userID).
					Return(int64(1), nil)

				m.EXPECT().GetLunches(gomock.Any(), req.userID, req.offset, 5).
					Return(nil, gorm.ErrInvalidData)
			},
			expected:   nil,
			expectedID: 0,
			err:        errors.NewErrRepository("lunchRepo", "GetLunches", gorm.ErrInvalidData),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLunchRepo := mocks.NewMockLunchRepo(ctrl)
			test.mockLunchRepo(mockLunchRepo, &test.input)

			lunchService := service.NewLunchService(mockLunchRepo)

			lunches, lunchID, err := lunchService.GetLunches(
				context.Background(),
				test.input.userID,
				test.input.offset,
				test.input.limit,
			)
			assert.Equal(t, test.expected, lunches)
			assert.Equal(t, test.expectedID, lunchID)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestGetLunchByID(t *testing.T) {
	tests := []struct {
		name          string
		input         int64
		mockLunchRepo func(m *mocks.MockLunchRepo, id int64)
		expected      *models.Lunch
		err           error
	}{
		{
			name:  "success",
			input: 1,
			mockLunchRepo: func(m *mocks.MockLunchRepo, id int64) {
				m.EXPECT().GetLunchByID(gomock.Any(), id).
					Return(&models.Lunch{ID: 1}, nil)
			},
			expected: &models.Lunch{ID: 1},
			err:      nil,
		},

		{
			name:          "invalid id",
			input:         0,
			mockLunchRepo: func(m *mocks.MockLunchRepo, id int64) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name:  "repo error: GetLunchByID",
			input: 7,
			mockLunchRepo: func(m *mocks.MockLunchRepo, id int64) {
				m.EXPECT().GetLunchByID(gomock.Any(), id).
					Return(nil, gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "GetLunchByID", gorm.ErrRecordNotFound),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLunchRepo := mocks.NewMockLunchRepo(ctrl)
			test.mockLunchRepo(mockLunchRepo, test.input)

			lunchService := service.NewLunchService(mockLunchRepo)

			lunch, err := lunchService.GetLunchByID(context.Background(), test.input)
			assert.Equal(t, test.expected, lunch)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestJoinLunch(t *testing.T) {
	type request struct {
		lunchID int64
		userID  int64
	}

	tests := []struct {
		name          string
		input         request
		mockLunchRepo func(m *mocks.MockLunchRepo, req *request)
		expected      *models.Lunch
		err           error
	}{
		{
			name: "success",
			input: request{
				lunchID: 1,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchParticipants(gomock.Any(), models.Add, req.lunchID, req.userID).
					Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), req.lunchID).
					Return(&models.Lunch{ID: 1}, nil)
			},
			expected: &models.Lunch{ID: 1},
			err:      nil,
		},

		{
			name: "invalid lunchID",
			input: request{
				lunchID: 0,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "invalid userID",
			input: request{
				lunchID: 1,
				userID:  0,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "repo error: UpdateLunchParticipants",
			input: request{
				lunchID: 7,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchParticipants(gomock.Any(), models.Add, req.lunchID, req.userID).
					Return(gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "UpdateLunchParticipants", gorm.ErrRecordNotFound),
		},

		{
			name: "repo error: GetLunchByID",
			input: request{
				lunchID: 7,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchParticipants(gomock.Any(), models.Add, req.lunchID, req.userID).
					Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), req.lunchID).
					Return(nil, gorm.ErrInvalidDB)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "GetLunchByID", gorm.ErrInvalidDB),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLunchRepo := mocks.NewMockLunchRepo(ctrl)
			test.mockLunchRepo(mockLunchRepo, &test.input)

			lunchService := service.NewLunchService(mockLunchRepo)

			lunch, err := lunchService.JoinLunch(context.Background(), test.input.lunchID, test.input.userID)
			assert.Equal(t, test.expected, lunch)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestLeaveLunch(t *testing.T) {
	type request struct {
		lunchID int64
		userID  int64
	}

	tests := []struct {
		name          string
		input         request
		mockLunchRepo func(m *mocks.MockLunchRepo, req *request)
		expected      *models.Lunch
		err           error
	}{
		{
			name: "success",
			input: request{
				lunchID: 1,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchParticipants(gomock.Any(), models.Del, req.lunchID, req.userID).
					Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), req.lunchID).
					Return(&models.Lunch{ID: 1}, nil)
			},
			expected: &models.Lunch{ID: 1},
			err:      nil,
		},

		{
			name: "invalid lunchID",
			input: request{
				lunchID: 0,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "invalid userID",
			input: request{
				lunchID: 1,
				userID:  0,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "repo error: UpdateLunchParticipants",
			input: request{
				lunchID: 7,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchParticipants(gomock.Any(), models.Del, req.lunchID, req.userID).
					Return(gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "UpdateLunchParticipants", gorm.ErrRecordNotFound),
		},

		{
			name: "repo error: GetLunchByID",
			input: request{
				lunchID: 7,
				userID:  1,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchParticipants(gomock.Any(), models.Del, req.lunchID, req.userID).
					Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), req.lunchID).
					Return(nil, gorm.ErrInvalidDB)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "GetLunchByID", gorm.ErrInvalidDB),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLunchRepo := mocks.NewMockLunchRepo(ctrl)
			test.mockLunchRepo(mockLunchRepo, &test.input)

			lunchService := service.NewLunchService(mockLunchRepo)

			lunch, err := lunchService.LeaveLunch(context.Background(), test.input.lunchID, test.input.userID)
			assert.Equal(t, test.expected, lunch)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestGetLunchHistory(t *testing.T) {
	type request struct {
		userID int64
		offset int
		limit  int
	}

	tests := []struct {
		name          string
		input         request
		mockLunchRepo func(m *mocks.MockLunchRepo, req *request)
		expected      []*models.LunchFeedback
		err           error
	}{
		{
			name: "success",
			input: request{
				userID: 1,
				offset: -5,
				limit:  100,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().GetUsersLunches(gomock.Any(), req.userID, 0, 10).
					Return([]*models.Lunch{
						{ID: 1, LikedBy: pq.Int64Array{1}},
						{ID: 2, LikedBy: pq.Int64Array{2}},
						{ID: 3, LikedBy: pq.Int64Array{3}},
					}, nil)
			},
			expected: []*models.LunchFeedback{
				{
					Lunch:   &models.Lunch{ID: 1, LikedBy: pq.Int64Array{1}},
					IsLiked: true,
				},
				{
					Lunch:   &models.Lunch{ID: 2, LikedBy: pq.Int64Array{2}},
					IsLiked: false,
				},
				{
					Lunch:   &models.Lunch{ID: 3, LikedBy: pq.Int64Array{3}},
					IsLiked: false,
				},
			},
			err: nil,
		},

		{
			name: "invalid userID",
			input: request{
				userID: 0,
				offset: 0,
				limit:  3,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "repo error: GetUsersLunches",
			input: request{
				userID: 1,
				offset: 0,
				limit:  0,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().GetUsersLunches(gomock.Any(), req.userID, req.offset, 5).
					Return(nil, gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "GetUsersLunches", gorm.ErrRecordNotFound),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLunchRepo := mocks.NewMockLunchRepo(ctrl)
			test.mockLunchRepo(mockLunchRepo, &test.input)

			lunchService := service.NewLunchService(mockLunchRepo)

			lunches, err := lunchService.GetLunchHistory(context.Background(), test.input.userID, test.input.offset, test.input.limit)
			assert.Equal(t, test.expected, lunches)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestRateLunch(t *testing.T) {
	type request struct {
		userID  int64
		lunchID int64
		isLiked bool
	}

	tests := []struct {
		name          string
		input         request
		mockLunchRepo func(m *mocks.MockLunchRepo, req *request)
		expected      *models.LunchFeedback
		err           error
	}{
		{
			name: "success",
			input: request{
				userID:  1,
				lunchID: 1,
				isLiked: true,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchLikedBy(gomock.Any(), models.Add, req.lunchID, req.userID).
					Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), req.lunchID).
					Return(&models.Lunch{ID: 1, LikedBy: pq.Int64Array{1}}, nil)
			},
			expected: &models.LunchFeedback{
				Lunch:   &models.Lunch{ID: 1, LikedBy: pq.Int64Array{1}},
				IsLiked: true,
			},
			err: nil,
		},

		{
			name: "invalid userID",
			input: request{
				userID:  0,
				lunchID: 1,
				isLiked: false,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "invalid lunchID",
			input: request{
				userID:  1,
				lunchID: 0,
				isLiked: true,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {},
			expected:      nil,
			err:           errors.ErrInvalidRequest,
		},

		{
			name: "repo error: UpdateLunchLikedBy",
			input: request{
				userID:  1,
				lunchID: 7,
				isLiked: false,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchLikedBy(gomock.Any(), models.Del, req.lunchID, req.userID).
					Return(gorm.ErrRecordNotFound)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "UpdateLunchLikedBy", gorm.ErrRecordNotFound),
		},

		{
			name: "repo error: GetLunchByID",
			input: request{
				userID:  1,
				lunchID: 7,
				isLiked: true,
			},
			mockLunchRepo: func(m *mocks.MockLunchRepo, req *request) {
				m.EXPECT().UpdateLunchLikedBy(gomock.Any(), models.Add, req.lunchID, req.userID).
					Return(nil)

				m.EXPECT().GetLunchByID(gomock.Any(), req.lunchID).
					Return(nil, gorm.ErrInvalidDB)
			},
			expected: nil,
			err:      errors.NewErrRepository("lunchRepo", "GetLunchByID", gorm.ErrInvalidDB),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLunchRepo := mocks.NewMockLunchRepo(ctrl)
			test.mockLunchRepo(mockLunchRepo, &test.input)

			lunchService := service.NewLunchService(mockLunchRepo)

			lunch, err := lunchService.RateLunch(context.Background(), test.input.userID, test.input.lunchID, test.input.isLiked)
			assert.Equal(t, test.expected, lunch)
			assert.Equal(t, test.err, err)
		})
	}
}

func TestValidTime(t *testing.T) {
	tests := []struct {
		name      string
		now       time.Time
		lunchTime time.Time
		expected  bool
	}{
		{
			name:      "Время обеда внутри допустимого интервала",
			now:       time.Now(),
			lunchTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 12, 0, 0, 0, time.Now().Location()),
			expected:  true,
		},
		{
			name:      "Время обеда раньше начала допустимого интервала",
			now:       time.Now(),
			lunchTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 10, 0, 0, 0, time.Now().Location()),
			expected:  false,
		},
		{
			name:      "Время обеда позже конца допустимого интервала",
			now:       time.Now(),
			lunchTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 15, 0, 0, 0, time.Now().Location()),
			expected:  false,
		},
		{
			name:      "Время обеда равно началу допустимого интервала",
			now:       time.Now(),
			lunchTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 11, 0, 0, 0, time.Now().Location()),
			expected:  true,
		},
		{
			name:      "Время обеда равно концу допустимого интервала",
			now:       time.Now(),
			lunchTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 14, 0, 0, 0, time.Now().Location()),
			expected:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := service.ValidTime(context.Background(), test.now, test.lunchTime)
			assert.Equal(t, test.expected, result)
		})
	}
}
