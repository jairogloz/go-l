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
