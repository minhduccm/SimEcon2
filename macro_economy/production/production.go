package production

import (
	"errors"

	"github.com/ninjadotorg/SimEcon002/common"
	"github.com/ninjadotorg/SimEcon002/macro_economy/abstraction"
)

type Production struct {
	AgentTypeToAgentProduction map[uint]abstraction.AgentProduction
}

var production *Production

func GetProductionInstance() *Production {
	if production != nil {
		return production
	}
	production = &Production{
		AgentTypeToAgentProduction: map[uint]abstraction.AgentProduction{
			common.PERSON:         &PersonProduction{},
			common.NECESSITY_FIRM: &NFirmProduction{},
			common.CAPITAL_FIRM:   &CFirmProduction{},
		},
	}
	return production
}

func computeDecayNecessity(asset abstraction.Asset) abstraction.Asset {
	return asset
}

func computeDecayCapital(asset abstraction.Asset) abstraction.Asset {
	return asset
}

func computeDecayManHours(asset abstraction.Asset) abstraction.Asset {
	return asset
}

func convertLinearly(input float64, a float64) float64 {
	return a * input
}

func (prod *Production) GetProductionByAgentType(
	agentType uint,
) (abstraction.AgentProduction, error) {
	agentProd, ok := prod.AgentTypeToAgentProduction[agentType]
	if !ok {
		return nil, errors.New("Agent ID not found")
	}
	return agentProd, nil
}
