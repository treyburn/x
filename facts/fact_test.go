package facts

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Add(x,y)
func Test_Fact_QueryOnly(t *testing.T) {
	want := 3.3

	qm := testQueryMgr{
		store: map[UUID]float64{
			UUID("1"): 1.1,
			UUID("2"): 2.2,
		},
	}

	fm := testFactMgr{
		store: map[UUID]*Fact{},
	}

	f := Fact{
		Operation: Add,
		X:         "1",
		Y:         "2",
		FactMgmt:  &fm,
		QueryMgmt: &qm,
	}

	got, err := f.Resolve()
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_Fact_NestedFact(t *testing.T) {
	want := 2.9

	qm := testQueryMgr{
		store: map[UUID]float64{
			UUID("1"): 1.1,
			UUID("X"): 5.1,
			UUID("Y"): 3.3,
		},
	}

	fm := testFactMgr{}

	f := Fact{
		Operation: Add,
		X:         "1",
		Y:         "2",
		FactMgmt:  &fm,
		QueryMgmt: &qm,
	}

	f2 := Fact{
		ID:        "2",
		Operation: 1,
		X:         "X",
		Y:         "Y",
		FactMgmt:  &fm,
		QueryMgmt: &qm,
	}

	fm.store = map[UUID]*Fact{
		"2": &f2,
	}

	got, err := f.Resolve()
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_Fact_Err(t *testing.T) {
	qm := testQueryMgr{store: map[UUID]float64{}}
	fm := testFactMgr{store: map[UUID]*Fact{}}

	f := Fact{
		Operation: 0,
		X:         "1",
		Y:         "2",
		FactMgmt:  &fm,
		QueryMgmt: &qm,
	}

	_, err := f.Resolve()
	assert.Error(t, err)
}

type testQueryMgr struct {
	store map[UUID]float64
}

func (q *testQueryMgr) LookupQuery(uuid UUID) (float64, error) {
	res, ok := q.store[uuid]
	if !ok {
		return 0, errors.New("no query with this uuid")
	}

	return res, nil
}

type testFactMgr struct {
	store map[UUID]*Fact
}

func (f *testFactMgr) LookupFact(uuid UUID) (*Fact, error) {
	res, ok := f.store[uuid]
	if !ok {
		return nil, errors.New("no fact with this uuid")
	}

	return res, nil
}
