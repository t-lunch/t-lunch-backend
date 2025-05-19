package transport

import (
	"context"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
	"go.uber.org/zap"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type LunchService interface {
	CreateLunch(ctx context.Context, userID int64, place string, lunchTime time.Time, description string) (*models.Lunch, error)
	GetLunches(ctx context.Context, userID int64, offset, limit int) ([]*models.Lunch, int64, error)
	GetLunchByID(ctx context.Context, lunchID int64) (*models.Lunch, error)
	JoinLunch(ctx context.Context, lunchID, userID int64) (*models.Lunch, error)
	LeaveLunch(ctx context.Context, lunchID, userID int64) (*models.Lunch, error)
}

type UserService interface {
	GetUsersByIDs(ctx context.Context, userIDs []int64) ([]*models.UserResponse, error)
}

type LunchTransport struct {
	tlunch.UnimplementedTlunchServer
	lunchService LunchService
	userService  UserService
	zapLogger    *zap.Logger
}

func NewLunchTransport(lunchService LunchService, userService UserService, zapLogger *zap.Logger) *LunchTransport {
	return &LunchTransport{
		lunchService: lunchService,
		userService:  userService,
		zapLogger:    zapLogger,
	}
}

func (t *LunchTransport) CreateLunch(ctx context.Context, request *tlunch.CreateLunchRequest) (*tlunch.LunchResponse, error) {
	t.zapLogger.Info("CreateLunch request",
		zap.Int64("user_id", request.GetUserId()),
		zap.String("place", request.GetPlace()),
		zap.Time("time", request.GetTime().AsTime()),
	)

	response, err := t.lunchService.CreateLunch(
		ctx,
		request.GetUserId(),
		request.GetPlace(),
		request.GetTime().AsTime(),
		request.GetDescription(),
	)
	if err != nil {
		t.zapLogger.Error("CreateLunch failed", zap.Error(err))
		return nil, err
	}

	rsafe := pointer.Get(response)
	t.zapLogger.Info("CreateLunch success", zap.Int64("lunch_id", rsafe.ID))

	var description *string = nil
	if rsafe.Description != "" {
		description = &rsafe.Description
	}
	return &tlunch.LunchResponse{
		Lunch: &tlunch.Lunch{
			Id:                   rsafe.ID,
			Name:                 rsafe.Creator.Name,
			Surname:              rsafe.Creator.Surname,
			Place:                rsafe.Place,
			Time:                 timestamppb.New(rsafe.Time),
			NumberOfParticipants: rsafe.NumberOfParticipants,
			Description:          description,
			UsersId:              rsafe.Participants,
		},
	}, nil
}

func (t *LunchTransport) GetLunches(ctx context.Context, request *tlunch.LunchRequest) (*tlunch.GetLunchesResponse, error) {
	t.zapLogger.Info("GetLunches request",
		zap.Int64("user_id", request.GetUserId()),
		zap.Int32("offset", request.GetOffset()),
		zap.Int32("limit", request.GetLimit()),
	)

	response, lunchID, err := t.lunchService.GetLunches(
		ctx,
		request.GetUserId(),
		int(request.GetOffset()),
		int(request.GetLimit()),
	)
	if err != nil {
		t.zapLogger.Error("GetLunches failed", zap.Error(err))
		return nil, err
	}

	lunchesResponse := &tlunch.GetLunchesResponse{
		Lunches: make([]*tlunch.Lunch, len(response)),
		LunchId: nil,
	}

	if lunchID > 0 {
		lunchesResponse.LunchId = &lunchID
	}

	for i, lunch := range response {
		rsafe := pointer.Get(lunch)

		var description *string = nil
		if rsafe.Description != "" {
			description = &rsafe.Description
		}
		lunchesResponse.Lunches[i] = &tlunch.Lunch{
			Id:                   rsafe.ID,
			Name:                 rsafe.Creator.Name,
			Surname:              rsafe.Creator.Surname,
			Place:                rsafe.Place,
			Time:                 timestamppb.New(rsafe.Time),
			NumberOfParticipants: rsafe.NumberOfParticipants,
			Description:          description,
			UsersId:              rsafe.Participants,
		}
	}

	t.zapLogger.Info("GetLunches success", zap.Int("lunches_count", len(response)))

	return lunchesResponse, nil
}

func (t *LunchTransport) GetDetailLunch(ctx context.Context, request *tlunch.DetailLunchRequest) (*tlunch.DetailLunchResponse, error) {
	t.zapLogger.Info("GetDetailLunch request", zap.Int64("lunch_id", request.GetLunchId()), zap.Int64s("user_id", request.GetUsersId()))

	lunch, err := t.lunchService.GetLunchByID(ctx, request.GetLunchId())
	if err != nil {
		t.zapLogger.Error("GetDetailLunch failed", zap.Error(err))
		return nil, err
	}

	users, err := t.userService.GetUsersByIDs(ctx, request.GetUsersId())
	if err != nil {
		t.zapLogger.Error("GetDetailLunch failed", zap.Error(err))
		return nil, err
	}

	rsafe := pointer.Get(lunch)
	var description *string = nil
	if rsafe.Description != "" {
		description = &rsafe.Description
	}

	response := &tlunch.DetailLunchResponse{
		Lunch: &tlunch.Lunch{
			Id:                   rsafe.ID,
			Name:                 rsafe.Creator.Name,
			Surname:              rsafe.Creator.Surname,
			Place:                rsafe.Place,
			Time:                 timestamppb.New(rsafe.Time),
			NumberOfParticipants: rsafe.NumberOfParticipants,
			Description:          description,
			UsersId:              rsafe.Participants,
		},
		Users: make([]*tlunch.User, len(users)),
	}

	for i, user := range users {
		rsafe := pointer.Get(user)
		response.Users[i] = &tlunch.User{
			UserId:  rsafe.ID,
			Name:    rsafe.Name,
			Surname: rsafe.Surname,
			Tg:      rsafe.Tg,
			Office:  rsafe.Office,
			Emoji:   rsafe.Emoji,
		}
	}

	t.zapLogger.Info("GetDetailLunch success", zap.Int64("lunch_id", request.GetLunchId()), zap.Int("users_count", len(users)))

	return response, nil
}

func (t *LunchTransport) JoinLunch(ctx context.Context, request *tlunch.ActionLunchRequest) (*tlunch.LunchResponse, error) {
	t.zapLogger.Info("JoinLunch request", zap.Int64("user_id", request.GetUserId()), zap.Int64("lunch_id", request.GetLunchId()))

	response, err := t.lunchService.JoinLunch(ctx, request.GetLunchId(), request.GetUserId())
	if err != nil {
		t.zapLogger.Error("JoinLunch failed", zap.Error(err))
		return nil, err
	}

	rsafe := pointer.Get(response)
	t.zapLogger.Info("JoinLunch success", zap.Int64("user_id", request.GetUserId()), zap.Int64("lunch_id", request.GetLunchId()))

	var description *string = nil
	if rsafe.Description != "" {
		description = &rsafe.Description
	}
	return &tlunch.LunchResponse{
		Lunch: &tlunch.Lunch{
			Id:                   rsafe.ID,
			Name:                 rsafe.Creator.Name,
			Surname:              rsafe.Creator.Surname,
			Place:                rsafe.Place,
			Time:                 timestamppb.New(rsafe.Time),
			NumberOfParticipants: rsafe.NumberOfParticipants,
			Description:          description,
			UsersId:              rsafe.Participants,
		},
	}, nil
}

func (t *LunchTransport) LeaveLunch(ctx context.Context, request *tlunch.ActionLunchRequest) (*tlunch.LunchResponse, error) {
	t.zapLogger.Info("LeaveLunch request", zap.Int64("user_id", request.GetUserId()), zap.Int64("lunch_id", request.GetLunchId()))

	response, err := t.lunchService.LeaveLunch(ctx, request.GetLunchId(), request.GetUserId())
	if err != nil {
		t.zapLogger.Error("LeaveLunch failed", zap.Error(err))
		return nil, err
	}

	rsafe := pointer.Get(response)
	t.zapLogger.Info("LeaveLunch success", zap.Int64("user_id", request.GetUserId()), zap.Int64("lunch_id", request.GetLunchId()))

	var description *string = nil
	if rsafe.Description != "" {
		description = &rsafe.Description
	}
	return &tlunch.LunchResponse{
		Lunch: &tlunch.Lunch{
			Id:                   rsafe.ID,
			Name:                 rsafe.Creator.Name,
			Surname:              rsafe.Creator.Surname,
			Place:                rsafe.Place,
			Time:                 timestamppb.New(rsafe.Time),
			NumberOfParticipants: rsafe.NumberOfParticipants,
			Description:          description,
			UsersId:              rsafe.Participants,
		},
	}, nil
}
