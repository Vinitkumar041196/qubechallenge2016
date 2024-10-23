package localstore

import (
	"distributor-manager/internal/store"
	"distributor-manager/internal/types"
	"fmt"
)

type DistributorDB struct {
	Code            string
	Permissions     *types.DistributorPermissions
	SubDistributors map[string]struct{}
}

type localDistributorStore struct {
	store MapStore[*DistributorDB]
}

func NewLocalDistributorStore() store.DistributorStorage {
	return &localDistributorStore{
		store: newMapStore[*DistributorDB](),
	}
}

func toDistributor(s *localDistributorStore, d *DistributorDB) *types.Distributor {
	dist := new(types.Distributor)
	dist.Code = d.Code
	dist.Permissions = d.Permissions
	if d.SubDistributors != nil {
		for k := range d.SubDistributors {
			subD, ok := s.store.Get(k)
			if ok {
				dist.SubDistributors = append(dist.SubDistributors, toDistributor(s, subD))
			}
		}
	}
	return dist
}

func toDistributorDB(d *types.Distributor) *DistributorDB {
	dist := new(DistributorDB)
	dist.Code = d.Code
	dist.Permissions = d.Permissions
	dist.SubDistributors = make(map[string]struct{})

	if len(d.SubDistributors) != 0 {
		for _, sD := range d.SubDistributors {
			dist.SubDistributors[sD.Code] = struct{}{}
		}
	}
	return dist
}

func (s *localDistributorStore) GetDistributorByCode(code string) (*types.Distributor, error) {
	dist, ok := s.store.Get(code)
	if !ok {
		return nil, ErrNotFound
	}
	return toDistributor(s, dist), nil
}

func (s *localDistributorStore) PutDistributorByCode(dist *types.Distributor) error {
	if dist == nil {
		return fmt.Errorf("record cannot be nil")
	}

	s.store.Set(dist.Code, toDistributorDB(dist))

	if len(dist.SubDistributors) > 0 {
		for _, subD := range dist.SubDistributors {
			s.store.Set(subD.Code, toDistributorDB(subD))
		}
	}
	return nil
}

func (s *localDistributorStore) DeleteDistributorByCode(code string) error {
	s.store.Delete(code)
	return nil
}
