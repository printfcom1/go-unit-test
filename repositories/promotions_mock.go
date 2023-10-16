package repositories

import "github.com/stretchr/testify/mock"

type promptionRepositoryMock struct {
	mock.Mock
}

func NewPromotionRipositoryMock() *promptionRepositoryMock {
	return &promptionRepositoryMock{}
}

func (p *promptionRepositoryMock) GetPromotion() (Promotion, error) {
	args := p.Called()

	return args.Get(0).(Promotion), args.Error(1)
}
