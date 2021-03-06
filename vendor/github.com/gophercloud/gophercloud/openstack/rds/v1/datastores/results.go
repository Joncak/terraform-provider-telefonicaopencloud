package datastores

import "github.com/gophercloud/gophercloud"

type DataStore struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	DataStore string `json:"datastore"`
	Image     string `json:"image"`
	Packages  string `json:"packages"`
	Active    int    `json:"active"`
}

type ListResult struct {
	gophercloud.Result
}

func (lr ListResult) Extract() ([]DataStore, error) {
	var a struct {
		DataStores []DataStore `json:"datastores"`
	}
	err := lr.Result.ExtractInto(&a)
	return a.DataStores, err
}
