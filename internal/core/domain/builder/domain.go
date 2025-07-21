package domain_builder

import "github.com/alfariiizi/vandor/internal/infrastructure/db"

type Domain[EntType any, DomainType any] struct {
	wrap   func(ent EntType, client *db.Client) DomainType
	client *db.Client
}

func NewDomain[EntType any, DomainType any](
	wrapFunc func(ent EntType, client *db.Client) DomainType,
	client *db.Client,
) Domain[EntType, DomainType] {
	return Domain[EntType, DomainType]{
		wrap:   wrapFunc,
		client: client,
	}
}

func (d Domain[EntType, DomainType]) Convert(ent EntType) DomainType {
	return d.wrap(ent, d.client)
}

func (d Domain[EntType, DomainType]) One(ent EntType, err error) (DomainType, error) {
	if err != nil {
		var zero DomainType
		return zero, err
	}
	return d.wrap(ent, d.client), nil
}

func (d Domain[EntType, DomainType]) Many(ents []EntType, err error) ([]DomainType, error) {
	if err != nil {
		return nil, err
	}
	out := make([]DomainType, len(ents))
	for i, ent := range ents {
		out[i] = d.wrap(ent, d.client)
	}
	return out, nil
}
