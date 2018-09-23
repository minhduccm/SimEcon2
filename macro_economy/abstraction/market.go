package abstraction

import "github.com/ninjadotorg/SimEcon002/macro_economy/dto"

type Market interface {
	Buy(string, *dto.OrderItem, Storage, AccountManager) error
	Sell()
}
