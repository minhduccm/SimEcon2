package storage

import (
	"errors"

	"github.com/ninjadotorg/SimEcon002/common"
	"github.com/ninjadotorg/SimEcon002/macro_economy/abstraction"
)

type Storage struct {
	Agents map[string]abstraction.Agent
	Assets map[string]map[uint]abstraction.Asset // agentID -> assetID -> asset
	Asks   map[uint][]abstraction.OrderBook
	Bids   map[uint][]abstraction.OrderBook
}

var storage *Storage

func GetStorageInstance() *Storage {
	if storage != nil {
		return storage
	}
	storage = &Storage{
		Agents: map[string]abstraction.Agent{},
		Assets: map[string]map[uint]abstraction.Asset{},
	}
	return storage
}

func (st *Storage) InsertAgent(
	agentID string,
	agentType uint,
) abstraction.Agent {
	agent := Agent{
		AgentID: agentID,
		Type:    agentType,
	}
	var newAgent abstraction.Agent = nil
	if agentType == common.PERSON {
		newAgent = &Person{
			agent,
		}
	} else if agentType == common.NECESSITY_FIRM {
		newAgent = &NecessityFirm{
			agent,
		}
	}
	st.Agents[agentID] = newAgent
	return newAgent
}

func (st *Storage) GetAgentByID(
	agentID string,
) (abstraction.Agent, error) {

	agent, ok := st.Agents[agentID]
	if !ok {
		return nil, errors.New("Could not find the agent")
	}
	return agent, nil
}

func (st *Storage) UpdateAssets(
	agentID string,
	assets []abstraction.Asset,
) {
	if _, ok := st.Assets[agentID]; !ok {
		st.Assets[agentID] = map[uint]abstraction.Asset{}
	}
	for _, asset := range assets {
		st.Assets[agentID][asset.GetType()] = asset
	}
}

func (st *Storage) UpdateAsset(
	agentID string,
	asset abstraction.Asset,
) {
	if _, ok := st.Assets[agentID]; !ok {
		st.Assets[agentID] = map[uint]abstraction.Asset{}
	}
	st.Assets[agentID][asset.GetType()] = asset
}

func (st *Storage) GetAgentAssets(
	agentID string,
) (map[uint]abstraction.Asset, error) {
	assets, ok := st.Assets[agentID]
	if !ok {
		return nil, errors.New("Could not find out assets with the agent id")
	}
	return assets, nil
}
