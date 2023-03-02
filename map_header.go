package cloud

type mapHeader struct {
	Name  string
	Value string
}

func appendHeader(fs ...*mapHeader) map[string]string {
	header := make(map[string]string)
	for _, f := range fs {
		if f == nil {
			continue
		}
		header[f.Name] = f.Value
	}
	return header
}

func mapClusterIdFromStore(s StoreInterface) (ret *mapHeader) {
	if s == nil || s.GetGlobalClusterID() <= 0 {
		return
	}

	ret = &mapHeader{
		Name:  ClusterHeaderName,
		Value: s.GetGlobalClusterID().String(),
	}

	return
}

func mapClusterIdFromOpts(opts ResourceCommonOpts) (ret *mapHeader) {
	if opts == nil || opts.GetCluster() == nil || opts.GetCluster().ID <= 0 {
		return
	}

	ret = &mapHeader{
		Name:  ClusterHeaderName,
		Value: opts.GetCluster().ID.String(),
	}

	return
}

func mapClusterId(clusterID ID) (ret *mapHeader) {
	if clusterID <= 0 {
		return
	}

	ret = &mapHeader{
		Name:  ClusterHeaderName,
		Value: clusterID.String(),
	}
	return
}
