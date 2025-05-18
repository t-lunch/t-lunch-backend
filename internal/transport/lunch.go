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
}

type LunchTransport struct {
	tlunch.UnimplementedTlunchServer
	lunchService LunchService
	zapLogger    *zap.Logger
}

func NewLunchTransport(lunchService LunchService, zapLogger *zap.Logger) *LunchTransport {
	return &LunchTransport{
		lunchService: lunchService,
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
