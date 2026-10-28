package domain

import (
	"errors"
	"fmt"
	"testing"
)

func TestManageError(t *testing.T) {
	type args struct {
		err error
		msg string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "Test ManageError with ErrDuplicateKey error",
			args: args{
				err: ErrDuplicateKey,
			},
			wantErr: AppError{
				Code: ErrCodeDuplicateKey,
				Msg:  "Duplicate key",
			},
		},
		{
			name: "Test ManageError with wrap ErrDuplicateKey error",
			args: args{
				err: fmt.Errorf("wrap: %w", ErrDuplicateKey),
			},
			wantErr: AppError{
				Code: ErrCodeDuplicateKey,
				Msg:  "Duplicate key",
			},
		},
		{
			name: "Test ManageError with ErrIncorrectID error",
			args: args{
				err: ErrIncorrectID,
			},
			wantErr: AppError{
				Code: ErrCodeInvalidParams,
				Msg:  "Incorrect id",
			},
		},
		{
			name: "Test ManageError with wrap ErrIncorrectID error",
			args: args{
				err: fmt.Errorf("wrap: %w", ErrIncorrectID),
			},
			wantErr: AppError{
				Code: ErrCodeInvalidParams,
				Msg:  "Incorrect id",
			},
		},
		{
			name: "Test ManageError with ErrNotFound error",
			args: args{
				err: ErrNotFound,
			},
			wantErr: AppError{
				Code: ErrCodeNotFound,
				Msg:  "Not found",
			},
		},
		{
			name: "Test ManageError with wrap ErrNotFound error",
			args: args{
				err: fmt.Errorf("wrap: %w", ErrNotFound),
			},
			wantErr: AppError{
				Code: ErrCodeNotFound,
				Msg:  "Not found",
			},
		},
		{
			name: "Test ManageError with nil error",
			args: args{
				err: nil,
			},
			wantErr: nil,
		},
		{
			name: "Test ManageError with ErrTimeout error",
			args: args{
				err: ErrTimeout,
			},
			wantErr: AppError{
				Code: ErrCodeTimeout,
				Msg:  "Timeout",
			},
		},
		{
			name: "Test ManageError with custom message",
			args: args{
				err: ErrNotFound,
				msg: "resource not found",
			},
			wantErr: AppError{
				Code: ErrCodeNotFound,
				Msg:  "Not found: resource not found",
			},
		},
		{
			name: "Test ManageError with unknown error",
			args: args{
				err: errors.New("unknown error"),
			},
			wantErr: AppError{
				Code: ErrCodeInternalServerError,
				Msg:  "Server Error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ManageError(tt.args.err, tt.args.msg); !errors.Is(err, tt.wantErr) {
				t.Errorf("ManageError() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAppError(t *testing.T) {
	err := NewAppError(ErrCodeNotFound, "item not found")
	if err.Code != ErrCodeNotFound {
		t.Errorf("expected code %s, got %s", ErrCodeNotFound, err.Code)
	}
	if err.Msg != "item not found" {
		t.Errorf("expected msg %s, got %s", "item not found", err.Msg)
	}
	expectedStr := "not_found: item not found"
	if err.Error() != expectedStr {
		t.Errorf("expected error string %s, got %s", expectedStr, err.Error())
	}
}
