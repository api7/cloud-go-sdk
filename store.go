package cloud

// StoreInterface for store some global state
type StoreInterface interface {
	SetGlobalClusterID(id ID)
	GetGlobalClusterID() ID
}

type store struct {
	clusterID ID
}

func (g *store) SetGlobalClusterID(id ID) {
	g.clusterID = id
}
func (g *store) GetGlobalClusterID() ID {
	return g.clusterID
}

func newStore() StoreInterface {
	return &store{}
}
