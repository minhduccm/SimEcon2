package abstraction

import "github.com/ninjadotorg/SimEcon002/macro_economy/dto"

type Market interface {
	Buy(string, *dto.OrderItem, Storage, AccountManager) (float64, error)
	Sell(string, *dto.OrderItem, Storage, AccountManager) (float64, error)
}
