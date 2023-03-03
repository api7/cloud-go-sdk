package cloud

import "net/http"

type mapHeader struct {
	Name  string
	Value string
}

func appendHeader(fs ...*mapHeader) http.Header {
	h := http.Header{}
	for _, f := range fs {
		if f == nil {
			continue
		}
		h.Add(f.Name, f.Value)
	}
	return h
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
