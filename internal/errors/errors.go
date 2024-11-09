package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrCanNotBeEmpty(argumentName string) error {
	return status.Errorf(codes.Internal, fmt.Sprintf("%s cannot be empty", argumentName))
}

func ErrPointerIsNil(argumentName string) error {
	return status.Errorf(codes.Internal, fmt.Sprintf("%s is nil", argumentName))
}

func ErrFailedToBuildQuery(argumentName error) error {
	return status.Errorf(codes.Internal, fmt.Sprintf("Failed to build query: %v", argumentName))
}

func ErrFailedToInsertQuery(argumentName error) error {
	return status.Errorf(codes.Internal, fmt.Sprintf("Failed to insert query: %v", argumentName))
}

func ErrFailedToDeleteQuery(argumentName error) error {
	return status.Errorf(codes.Internal, fmt.Sprintf("Failed to delete query: %v", argumentName))
}

func ErrObjectNotFount(objectName string, objectId int64) error {
	return status.Errorf(codes.Internal, fmt.Sprintf("%s with id %d not found", objectName, objectId))
}
