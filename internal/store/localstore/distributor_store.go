package localstore

import (
	"distributor-manager/internal/store"
	"distributor-manager/internal/types"
	"fmt"
)

type DistributorDB struct {
	Code        string
	Permissions *types.DistributorPermissions
	ParentCode  string
}

type localDistributorStore struct {
	store MapStore[*DistributorDB]
}

func NewLocalDistributorStore() store.DistributorStorage {
	return &localDistributorStore{
		store: newMapStore[*DistributorDB](),
	}
}

func toDistributor(d *DistributorDB) *types.Distributor {
	dist := new(types.Distributor)
	dist.Code = d.Code
	dist.Permissions = d.Permissions
	dist.ParentCode = d.ParentCode
	return dist
}

func toDistributorDB(d *types.Distributor) *DistributorDB {
	dist := new(DistributorDB)
	dist.Code = d.Code
	dist.Permissions = d.Permissions
	dist.ParentCode = d.ParentCode
	return dist
}

func (s *localDistributorStore) GetDistributorByCode(code string) (*types.Distributor, error) {
	dist, ok := s.store.Get(code)
	if !ok {
		return nil, ErrNotFound
	}
	return toDistributor(dist), nil
}

func (s *localDistributorStore) PutDistributorByCode(dist *types.Distributor) error {
	if dist == nil {
		return fmt.Errorf("record cannot be nil")
	}

	s.store.Set(dist.Code, toDistributorDB(dist))
	return nil
}

func (s *localDistributorStore) DeleteDistributorByCode(code string) error {
	s.store.Delete(code)
	return nil
}
