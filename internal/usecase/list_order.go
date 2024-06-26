package usecase

import "github.com/mbrocco/goexpert/desafio-clean-arq/internal/entity"

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]entity.Order, error) {
	orders, err := c.OrderRepository.List()

	if err != nil {
		return nil, err
	}

	return orders, nil
}
