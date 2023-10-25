package services

import "github.com/stretchr/testify/mock"

type PromotionServiceMock struct {
	mock.Mock
}

func NewPromotionServiceMock() *PromotionServiceMock {
	return &PromotionServiceMock{}
}

func (s *PromotionServiceMock) CalculateDiscount(amount int) (int, error) {
	args := s.Called(amount)

	return args.Int(0), args.Error(1)
}
