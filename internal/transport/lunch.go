package transport

import (
	"context"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/t-lunch/t-lunch-backend/internal/models"
	tlunch "github.com/t-lunch/t-lunch-backend/pkg/api/generated"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type LunchService interface {
	CreateLunch(ctx context.Context, userID int64, place string, lunchTime time.Time, description string) (*models.Lunch, error)
	GetLunches(ctx context.Context, userID int64, offset, limit int) ([]*models.Lunch, int64, error)
}

type LunchTransport struct {
	tlunch.UnimplementedTlunchServer
	lunchService LunchService
}

func NewLunchTransport(lunchService LunchService) *LunchTransport {
	return &LunchTransport{lunchService: lunchService}
}

func (t *LunchTransport) CreateLunch(ctx context.Context, request *tlunch.CreateLunchRequest) (*tlunch.LunchResponse, error) {
	response, err := t.lunchService.CreateLunch(
		ctx,
		request.GetUserId(),
		request.GetPlace(),
		request.GetTime().AsTime(),
		request.GetDescription(),
	)
	if err != nil {
		return nil, err
	}
	rsafe := pointer.Get(response)
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
	response, lunchID, err := t.lunchService.GetLunches(
		ctx,
		request.GetUserId(),
		int(request.GetOffset()),
		int(request.GetLimit()),
	)
	if err != nil {
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
	return lunchesResponse, nil
}
