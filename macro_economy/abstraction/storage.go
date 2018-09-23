package abstraction

type Storage interface {
	InsertAgent(string, uint) Agent
	UpdateAssets(string, []Asset)
	UpdateAsset(string, Asset)
	GetAgentAssets(string) (map[uint]Asset, error)
	GetAgentByID(string) (Agent, error)
	GetSortedBidsByAssetType(uint, bool) OrderItems
	RemoveBidsByAgentIDs([]string, uint) error
	AppendAsk(uint, string, float64, float64)
}
