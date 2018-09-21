package storage

import (
	"time"

	"github.com/ninjadotorg/SimEcon002/common"
	"github.com/ninjadotorg/SimEcon002/macro_economy/abstraction"
)

type Person struct {
	Agent
}

func (p *Person) InitAgentAssets(
	st abstraction.Storage,
) {
	// necessity asset
	nAsset := &Asset{
		AgentID:      p.AgentID,
		Type:         common.NECESSITY,
		Quantity:     common.PERSON_NECESSITY,
		ProducedTime: time.Now().Unix(),
	}

	mhAsset := &Asset{
		AgentID:      p.AgentID,
		Type:         common.MAN_HOUR,
		Quantity:     common.PERSON_MAN_HOURS,
		ProducedTime: time.Now().Unix(),
	}
	st.UpdateAssets(p.AgentID, []abstraction.Asset{nAsset, mhAsset})
}

func (p *Person) GetType() uint {
	return p.Type
}
