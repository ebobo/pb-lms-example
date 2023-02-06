package service

import (
	"context"

	// load postgres driver
	_ "github.com/lib/pq"
	"google.golang.org/protobuf/types/known/emptypb"

	"atlassian.carcgl.com/bitbucket/ls/lms/db"
	"atlassian.carcgl.com/bitbucket/ls/lms/pkg/api/proto/v1"
	"atlassian.carcgl.com/bitbucket/ls/lms/pkg/model"
	"atlassian.carcgl.com/bitbucket/ls/lms/pkg/utility"
)

//ListOperators implements the gRPC ListOperators method
func (lms *LMSService) ListOperators(ctx context.Context, _ *emptypb.Empty) (*proto.OperatorList, error) {
	list := &proto.OperatorList{}
	results, err := db.ListOperators(lms.pgDatabase)
	if err != nil {
		return list, err
	}
	for _, o := range results {
		list.Operators = append(list.Operators, &proto.Operator{Id: o.ID, Name: o.Name, Valid: o.Valid, Created: o.Created.Format("2006-01-02T15:04:05Z07:00")})
	}
	return list, nil
}

//AddOperator implements the gRPC AddOperator method
func (lms *LMSService) CreateOperator(ctx context.Context, req *proto.CreateOperatorRequest) (*proto.CreateOperatorResponse, error) {
	id, err := utility.GenerateID(12)
	if err != nil {
		return nil, err
	}
	o := &model.Operator{
		ID:      id,
		Name:    req.Name,
		Valid:   req.Valid,
		Created: utility.GetTimeNow(),
	}
	err = db.AddOperator(lms.pgDatabase, o)
	if err != nil {
		return nil, err
	}
	return &proto.CreateOperatorResponse{OperatorId: id}, nil
}

//GetOperator implements the gRPC GetOperator method
func (lms *LMSService) GetOperator(ctx context.Context, req *proto.GetOperatorRequest) (*proto.Operator, error) {
	result, err := db.GetOperator(lms.pgDatabase, req.OperatorId)
	if err != nil {
		return nil, err
	}
	return &proto.Operator{Id: result.ID, Name: result.Name, Valid: result.Valid, Created: result.Created.Format("2006-01-02T15:04:05Z07:00")}, nil
}

//UpdateOperator implements the gRPC UpdateOperator method
func (lms *LMSService) UpdateOperator(ctx context.Context, req *proto.UpdateOperatorRequest) (*proto.UpdateOperatorResponse, error) {
	o := &model.Operator{
		ID:    req.OperatorId,
		Name:  req.Name,
		Valid: req.Valid,
	}
	err := db.UpdateOperator(lms.pgDatabase, o)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateOperatorResponse{OperatorId: req.OperatorId}, nil
}

//DeleteOperator implements the gRPC DeleteOperator method
func (lms *LMSService) DeleteOperator(ctx context.Context, req *proto.DeleteOperatorRequest) (*proto.DeleteOperatorResponse, error) {
	err := db.DeleteOperator(lms.pgDatabase, req.OperatorId)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteOperatorResponse{OperatorId: req.OperatorId}, nil
}
