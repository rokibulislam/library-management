package bookloan

import (
	"context"
	"time"
)

type service struct {
	repository Repository
}

// NewService creates a service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Create(ctx context.Context, bl *BookLoan) error {
	bl.Status = BookLoanInitiated

	if err := s.repository.AddBookLoan(ctx, bl); err != nil {
		return err
	}

	return nil
}

func (s *service) Get(ctx context.Context, ID string) (*BookLoan, error) {
	bl, err := s.repository.GetBookLoan(ctx, ID)
	if err != nil {
		return nil, err
	}
	if bl == nil {
		return nil, NewBookLoanNotFound("BookLoan not found", []string{})
	}

	return bl, nil
}

func (s *service) List(ctx context.Context, skip *int64, limit *int64) ([]BookLoan, error) {
	bls, err := s.repository.ListBookLoan(ctx, skip, limit)
	if err != nil {
		return nil, err
	}

	return bls, nil
}

func (s *service) ListByUser(ctx context.Context, userID string, skip *int64, limit *int64) ([]BookLoan, error) {
	bls, err := s.repository.ListBookLoanByUserID(ctx, userID, skip, limit)
	if err != nil {
		return nil, err
	}

	return bls, nil
}

func (s *service) Accept(ctx context.Context, ID string) error {
	bl, err := s.repository.GetBookLoan(ctx, ID)
	if err != nil {
		return err
	}
	if bl == nil {
		return NewBookLoanNotFound("BookLoan not found", []string{})
	}

	now := time.Now()

	bl.Status = BookLoanAccepted
	bl.AcceptedAt = &now

	err = s.repository.UpdateBookLoan(ctx, ID, bl)

	return err
}

func (s *service) Reject(ctx context.Context, ID string, rejectionCause string) error {
	bl, err := s.repository.GetBookLoan(ctx, ID)
	if err != nil {
		return err
	}
	if bl == nil {
		return NewBookLoanNotFound("BookLoan not found", []string{})
	}

	now := time.Now()

	bl.Status = BookLoanRejected
	bl.RejectionCause = rejectionCause
	bl.RejectedAt = &now

	err = s.repository.UpdateBookLoan(ctx, ID, bl)

	return err
}
