package growing_zone_history_service

import (
	"context"
	"encoding/json"
	"farm-service/domain/entity"
	"farm-service/domain/usecase/growing_zone_history"

	growingZoneHistoryP "github.com/anhvanhoa/sf-proto/gen/growing_zone_history/v1"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *GrowingZoneHistoryService) CreateHistory(ctx context.Context, req *growingZoneHistoryP.CreateHistoryRequest) (*growingZoneHistoryP.CreateHistoryResponse, error) {
	historyReq, err := s.createEntityHistoryReq(req)
	if err != nil {
		return nil, err
	}
	history, err := s.growingZoneHistoryUsecase.CreateHistory.Execute(ctx, historyReq)
	if err != nil {
		return nil, err
	}
	return &growingZoneHistoryP.CreateHistoryResponse{
		Success: true,
		Message: "History record created successfully",
		History: s.createProtoHistory(history),
	}, nil
}

func (s *GrowingZoneHistoryService) createEntityHistoryReq(req *growingZoneHistoryP.CreateHistoryRequest) (*growing_zone_history.CreateHistoryRequest, error) {
	// Convert protobuf Struct to map[string]interface{}
	var oldValue map[string]interface{}
	if req.OldValue != nil {
		oldValueBytes, err := req.OldValue.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(oldValueBytes, &oldValue); err != nil {
			return nil, err
		}
	}

	var newValue map[string]interface{}
	if req.NewValue != nil {
		newValueBytes, err := req.NewValue.MarshalJSON()
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(newValueBytes, &newValue); err != nil {
			return nil, err
		}
	}

	history := &growing_zone_history.CreateHistoryRequest{
		ZoneID:      req.ZoneId,
		Action:      req.Action.String(),
		OldValue:    oldValue,
		NewValue:    newValue,
		PerformedBy: req.PerformedBy,
		Notes:       req.Notes,
	}

	return history, nil
}

func (s *GrowingZoneHistoryService) createProtoHistory(history *entity.GrowingZoneHistory) *growingZoneHistoryP.ZoneHistory {
	// Convert action string to enum
	var action growingZoneHistoryP.HistoryAction
	switch history.Action {
	case "HISTORY_ACTION_CHANGE_SOIL":
		action = growingZoneHistoryP.HistoryAction_HISTORY_ACTION_CHANGE_SOIL
	case "HISTORY_ACTION_CHANGE_IRRIGATION":
		action = growingZoneHistoryP.HistoryAction_HISTORY_ACTION_CHANGE_IRRIGATION
	case "HISTORY_ACTION_MAINTENANCE":
		action = growingZoneHistoryP.HistoryAction_HISTORY_ACTION_MAINTENANCE
	case "HISTORY_ACTION_RESIZE":
		action = growingZoneHistoryP.HistoryAction_HISTORY_ACTION_RESIZE
	case "HISTORY_ACTION_RENAME":
		action = growingZoneHistoryP.HistoryAction_HISTORY_ACTION_RENAME
	default:
		action = growingZoneHistoryP.HistoryAction_HISTORY_ACTION_UNSPECIFIED
	}

	// Convert map[string]interface{} to protobuf Struct
	var oldValue *structpb.Struct
	if history.OldValue != nil {
		oldValueBytes, err := json.Marshal(history.OldValue)
		if err == nil {
			oldValue = &structpb.Struct{}
			oldValue.UnmarshalJSON(oldValueBytes)
		}
	}

	var newValue *structpb.Struct
	if history.NewValue != nil {
		newValueBytes, err := json.Marshal(history.NewValue)
		if err == nil {
			newValue = &structpb.Struct{}
			newValue.UnmarshalJSON(newValueBytes)
		}
	}

	response := &growingZoneHistoryP.ZoneHistory{
		Id:          history.ID,
		ZoneId:      history.ZoneID,
		Action:      action,
		OldValue:    oldValue,
		NewValue:    newValue,
		ActionDate:  timestamppb.New(history.ActionDate),
		PerformedBy: history.PerformedBy,
		Notes:       history.Notes,
	}

	return response
}
