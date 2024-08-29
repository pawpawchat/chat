package status

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MissingFields(fields []string) error {
	badRequest := new(errdetails.BadRequest)

	for idx := range fields {
		violation := &errdetails.BadRequest_FieldViolation{
			Field:       fields[idx],
			Description: "empty",
		}

		badRequest.FieldViolations = append(badRequest.FieldViolations, violation)
	}
	status, _ := status.New(codes.InvalidArgument, "missing required fields").WithDetails(badRequest)
	return status.Err()
}

func Exists(msg string) error {
	return status.New(codes.AlreadyExists, msg).Err()
}

func NotFound(msg string, keyDesc string, keyVal interface{}) error {
	searchedBy := &errdetails.BadRequest_FieldViolation{
		Field:       keyDesc,
		Description: fmt.Sprintf("%v", keyVal),
	}
	status, _ := status.New(codes.NotFound, msg).WithDetails(searchedBy)
	return status.Err()
}

func Internal(msg string) error {
	return status.New(codes.Internal, msg).Err()
}

func Unexpected(msg string) error {
	return status.New(codes.Unknown, msg).Err()
}
