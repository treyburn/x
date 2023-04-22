package facts

type Fact struct {
	ID        UUID
	Operation Op
	X         Metric // this uuid might map to a query or a child fact - hence the need to resolve
	Y         Metric // this uuid might map to a query or a child fact - hence the need to resolve
	FactMgmt  FactMgr
	QueryMgmt QueryMgr
}

func (f *Fact) Resolve() (float64, error) {
	proc, err := f.Operation.Resolve()
	if err != nil {
		return -1, err
	}
	x, err := f.X.Resolve(f.QueryMgmt, f.FactMgmt)
	if err != nil {
		return -1, err
	}
	y, err := f.Y.Resolve(f.QueryMgmt, f.FactMgmt)
	if err != nil {
		return -1, err
	}
	return proc(x, y), nil
}
