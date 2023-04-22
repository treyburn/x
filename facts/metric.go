package facts

import (
	"errors"
)

type Metric UUID

type QueryMgr interface {
	LookupQuery(UUID) (float64, error)
}

type FactMgr interface {
	LookupFact(UUID) (*Fact, error)
}

func (m Metric) Resolve(qm QueryMgr, fm FactMgr) (float64, error) {
	res, err := qm.LookupQuery(m.self())
	if err == nil {
		return res, nil
	}

	fact, err := fm.LookupFact(m.self())
	if err == nil {
		return fact.Resolve()
	}

	return 0, errors.New("could not resolve metric")
}

func (m Metric) self() UUID {
	return UUID(m)
}
