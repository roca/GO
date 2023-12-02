package grpc

import (
	"fmt"
	"grpc-go-server/data"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StatusError interface {
	Err() error
}

var StatusErrors map[error]StatusError

type InsufficientBalanceError struct {
	Acount string
	Amount float64
	Error  error
}

func NewInsufficientBalanceError(accountNumber string, amount float64) StatusError {
	return &InsufficientBalanceError{
		Acount: accountNumber,
		Amount: amount,
		Error:  data.ErrInsufficientBalance,
	}
}

func (i *InsufficientBalanceError) Err() error {
	s := status.New(codes.InvalidArgument, fmt.Sprintf("DB table BankAccounts error: %s", i.Error))
	s, _ = s.WithDetails(&errdetails.ErrorInfo{
		Reason: "INSUFFICIENT_BALANCE",
		Domain: "DB table BankAccounts",
		Metadata: map[string]string{
			"account_number": i.Acount,
			"amount":         fmt.Sprintf("Withdrawal amount: %v", i.Amount),
		},
	})
	s, _ = s.WithDetails(&errdetails.DebugInfo{})
	return s.Err()
}

type AccountNotFoundError struct {
	Account string
	Error   error
}

func NewAccountNotFoundError(accountNumber string) StatusError {
	return &AccountNotFoundError{
		Account: accountNumber,
		Error:   data.ErrAccountNotFound,
	}
}

func (i *AccountNotFoundError) Err() error {
	s := status.New(codes.FailedPrecondition, fmt.Sprintf("DB table BankAccounts error: %s", i.Error))
	s, _ = s.WithDetails(&errdetails.ErrorInfo{
		Reason: "INVALID_ACCOUNT, Account not found",
		Domain: "DB table BankAccounts",
		Metadata: map[string]string{
			"account_number": i.Account,
		},
	})
	s, _ = s.WithDetails(&errdetails.DebugInfo{})
	return s.Err()
}

type InternalTableError struct {
	TableName string
	Error     error
}

func NewInternalTableError(tableName string, err error) StatusError {
	return &InternalTableError{
		TableName: tableName,
		Error:     err,
	}
}

func (i *InternalTableError) Err() error {
	s := status.New(codes.Internal, fmt.Sprintf("DB table %s error: %s", i.TableName, i.Error))
	s, _ = s.WithDetails(&errdetails.ErrorInfo{Reason: fmt.Sprintf("INTERNAL_ERROR: %s", i.Error), Domain: fmt.Sprintf("DB table %s", i.TableName)})
	s, _ = s.WithDetails(&errdetails.DebugInfo{})
	return s.Err()
}
