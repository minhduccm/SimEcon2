package economy

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ninjadotorg/SimEcon002/common"
	"github.com/ninjadotorg/SimEcon002/macro_economy/dto"
)

// POST /types/{AGENT_TYPE}/agents
func Join(w http.ResponseWriter, r *http.Request, econ *Economy) {
	newAgentID := common.UUID()
	am := econ.AccountManager
	st := econ.Storage
	agentType, _ := strconv.Atoi(mux.Vars(r)["AGENT_TYPE"])

	// open wallet account
	am.OpenWalletAccount(newAgentID, 0.0)

	// insert new agent
	agent := st.InsertAgent(newAgentID, uint(agentType))
	agent.InitAgentAssets(st)

	jsInBytes, _ := json.Marshal(agent)
	w.Write(jsInBytes)
}

// POST /agents/{AGENT_ID}/produce
func Produce(w http.ResponseWriter, r *http.Request, econ *Economy) {
	st := econ.Storage
	prod := econ.Production

	var assetsReq map[uint]*dto.Asset
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&assetsReq)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	// TODO: validate assets requested by agent type

	agentID := mux.Vars(r)["AGENT_ID"]
	agent, _ := st.GetAgentByID(agentID)

	agentProd, err := prod.GetProductionByAgentType(agent.GetType())
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	updatedAssets, err := agentProd.Produce(st, agentID, assetsReq)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsInBytes, _ := json.Marshal(updatedAssets)
	w.Write(jsInBytes)
}

// GET /agents/{AGENT_ID}/assets
func GetAgentAssets(w http.ResponseWriter, r *http.Request, econ *Economy) {
	agentID := mux.Vars(r)["AGENT_ID"]
	st := econ.Storage
	assets, err := st.GetAgentAssets(agentID)
	res := map[string]interface{}{}
	if err != nil {
		res["error"] = err.Error()
	} else {
		res["result"] = assets
	}
	jsInBytes, _ := json.Marshal(res)
	w.Write(jsInBytes)
}

// POST /agents/{AGENT_ID}/buy
func Buy(w http.ResponseWriter, r *http.Request, econ *Economy) {

}

// POST /agents/{AGENT_ID}/sell
func Sell(w http.ResponseWriter, r *http.Request, econ *Economy) {

}
