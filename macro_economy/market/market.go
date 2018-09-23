package market

import (
	"github.com/ninjadotorg/SimEcon002/common"
	"github.com/ninjadotorg/SimEcon002/macro_economy/abstraction"
	"github.com/ninjadotorg/SimEcon002/macro_economy/dto"
)

type Market struct {
}

var market *Market

func GetMarketInstance() *Market {
	if market != nil {
		return market
	}
	market = &Market{}
	return market
}

func (m *Market) Buy(
	agentID string,
	orderItemReq *dto.OrderItem,
	st abstraction.Storage,
	am abstraction.AccountManager,
) error {
	sortedBidsByAssetType := st.GetSortedBidsByAssetType(orderItemReq.AssetType, false)

	removingBidAgentIDs := []string{}
	for _, bid := range sortedBidsByAssetType {
		if bid.GetPricePerUnit() > orderItemReq.PricePerUnit {
			continue
		}
		if bid.GetQuantity() >= orderItemReq.Quantity {
			am.Pay(
				agentID,
				bid.GetAgentID(),
				bid.GetPricePerUnit()*orderItemReq.Quantity,
				common.PRIIC,
			)
			bid.SetQuantity(bid.GetQuantity() - orderItemReq.Quantity)
			if bid.GetQuantity() == 0 {
				removingBidAgentIDs = append(removingBidAgentIDs, bid.GetAgentID())
			}
			break
		}
		am.Pay(
			agentID,
			bid.GetAgentID(),
			bid.GetPricePerUnit()*bid.GetQuantity(),
			common.PRIIC,
		)
		orderItemReq.Quantity -= bid.GetQuantity()
		bid.SetQuantity(0)
		removingBidAgentIDs = append(removingBidAgentIDs, bid.GetAgentID())
	}
	// re-update bid list: remove bid with qty = 0 and append new ask if remaning qty > 0
	if len(removingBidAgentIDs) > 0 {
		err := st.RemoveBidsByAgentIDs(removingBidAgentIDs, orderItemReq.AssetType)
		if err != nil {
			return err
		}
	}

	if orderItemReq.Quantity > 0 {
		st.AppendAsk(
			orderItemReq.AssetType,
			orderItemReq.AgentID,
			orderItemReq.Quantity,
			orderItemReq.PricePerUnit,
		)
	}

	return nil
}

func (m *Market) Sell() {

}