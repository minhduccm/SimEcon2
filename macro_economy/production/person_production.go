package production

import (
	"github.com/ninjadotorg/SimEcon002/common"
	"github.com/ninjadotorg/SimEcon002/macro_economy/abstraction"
	"github.com/ninjadotorg/SimEcon002/macro_economy/dto"
)

type PersonProduction struct {
}

func (pp *PersonProduction) Produce(
	st abstraction.Storage,
	agentID string,
	assetsReq map[uint]*dto.Asset,
) (map[uint]abstraction.Asset, error) {
	agentAssets, _ := st.GetAgentAssets(agentID)
	nAsset := agentAssets[common.NECESSITY]
	mhAsset := agentAssets[common.MAN_HOUR]
	nAssetReq := assetsReq[common.NECESSITY]

	curNAsset := computeDecayNecessity(nAsset)
	curMHAsset := computeDecayManHours(mhAsset)

	spendingAmt := nAssetReq.Quantity
	if spendingAmt > curNAsset.GetQuantity() {
		spendingAmt = curNAsset.GetQuantity()
	}
	convertedMHAmt := convertLinearly(spendingAmt, 2.5)
	curMHAsset.SetQuantity(convertedMHAmt + curMHAsset.GetQuantity())
	curMHAsset.SetProducedTime()

	curNAsset.SetQuantity(curNAsset.GetQuantity() - spendingAmt)
	curNAsset.SetProducedTime()

	st.UpdateAsset(agentID, curNAsset)
	st.UpdateAsset(agentID, curMHAsset)

	res, _ := st.GetAgentAssets(agentID)
	return res, nil
}

func (pp *PersonProduction) GetActualAsset(
	asset abstraction.Asset,
) abstraction.Asset {
	return computeDecayManHours(asset)
}
