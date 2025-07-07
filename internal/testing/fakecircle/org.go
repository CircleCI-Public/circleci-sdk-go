package fakecircle

import (
	"fmt"
	"maps"

	"github.com/google/uuid"
)

func (s *Service) orgBySlug(slug string) *org {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, o := range s.orgs {
		if fmt.Sprintf("%s/%s", o.typ, o.name) == slug {
			return &org{
				typ:      o.typ,
				name:     o.name,
				contexts: maps.Clone(o.contexts),
			}
		}
	}
	return nil
}

type org struct {
	typ      string
	name     string
	contexts map[uuid.UUID]*context
}

type NewOrg struct {
	Type string
	Name string
}

type Org struct {
	ID   uuid.UUID
	Type string
	Name string
}

func (s *Service) AddOrg(newOrg NewOrg) (Org, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, o := range s.orgs {
		if o.name == newOrg.Name {
			return Org{}, errDuplicate
		}
	}

	id := uuid.New()
	o := &org{
		typ:      newOrg.Type,
		name:     newOrg.Name,
		contexts: make(map[uuid.UUID]*context),
	}
	s.orgs[id] = o
	return Org{
		ID:   id,
		Type: o.typ,
		Name: o.name,
	}, nil
}
