package localstore

import (
	"distributor-manager/internal/store"
	"distributor-manager/internal/types"
	"fmt"
)

type localDistributorStore struct {
	store MapStore[*types.Distributor]
}

func NewLocalDistributorStore() store.DistributorStorage {
	return &localDistributorStore{
		store: newMapStore[*types.Distributor](),
	}
}

func (s *localDistributorStore) GetDistributorByCode(code string) (*types.Distributor, error) {
	dist, ok := s.store.Get(code)
	if !ok {
		return nil, ErrNotFound
	}
	return dist, nil
}

func (s *localDistributorStore) PutDistributorByCode(dist *types.Distributor) error {
	if dist == nil {
		return fmt.Errorf("record cannot be nil")
	}

	s.store.Set(dist.Code, dist)

	if len(dist.SubDistributors) > 0 {
		for _, subD := range dist.SubDistributors {
			s.store.Set(subD.Code, subD)
		}
	}
	return nil
}

func (s *localDistributorStore) DeleteDistributorByCode(code string) error {
	s.store.Delete(code)
	return nil
}
