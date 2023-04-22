package facts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOp_Resolve(t *testing.T) {
	type testCase struct {
		name string
		op   Op
		want OpFunc
	}

	var tests = []testCase{
		{"add", Add, add},
		{"sub", Subtract, sub},
		{"mult", Multiply, mlt},
		{"div", Divide, div},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			got, err := test.op.Resolve()
			assert.NoError(t, err)
			// go disallows function ptr comparison - so we will compare the outputs for equality given equal inputs
			assert.Equal(t, test.want(2.0, 2.0), got(2.0, 2.0))
		})
	}
}

func TestOp_Resolve_Err(t *testing.T) {
	type testCase struct {
		name string
		op   Op
	}

	var tests = []testCase{
		{"invalid - below", Op(-1)},
		{"invalid - above", Op(4)},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			_, err := test.op.Resolve()
			assert.Error(t, err)
		})
	}
}

func Test_toFixed(t *testing.T) {
	type testCase struct {
		name      string
		x         float64
		precision int
		want      float64
	}

	var tests = []testCase{
		{"exact", 1.23, 2, 1.23},
		{"round down", 1.23456, 2, 1.23},
		{"round up", 1.23856, 2, 1.24},
		{"add precision", 1, 2, 1.00},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			got := toFixed(test.x, test.precision)
			assert.Equal(t, test.want, got)
		})
	}
}
