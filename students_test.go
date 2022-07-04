package coverage

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func Test_Len(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		p        People
		expected int
	}{
		"several people": {
			p: People{Person{
				firstName: "Ivan",
				lastName:  "Ivanov",
				birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
			}, Person{
				firstName: "Petr",
				lastName:  "Petrov",
				birthDay:  time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC),
			}},
			expected: 2,
		},
		"empty": {
			p:        People{},
			expected: 0,
		},
	}

	for name, tc := range tests {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tc.p.Len()
			if got != tc.expected {
				t.Errorf("[%s] expected: %d, got %d", name, tc.expected, got)
			}
		})
	}
}

func Test_Less(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		p        People
		i        int
		j        int
		expected bool
	}{
		"less by bday": {
			p: People{
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Person{
					firstName: "Petr",
					lastName:  "Petrov",
					birthDay:  time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC),
				}},
			i:        0,
			j:        1,
			expected: true,
		},
		"greater by bday": {
			p: People{
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Person{
					firstName: "Petr",
					lastName:  "Petrov",
					birthDay:  time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC),
				}},
			i:        0,
			j:        1,
			expected: false,
		},
		"less by first name": {
			p: People{
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Person{
					firstName: "Petr",
					lastName:  "Petrov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				}},
			i:        0,
			j:        1,
			expected: true,
		},
		"greater by first name": {
			p: People{
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Person{
					firstName: "Andrew",
					lastName:  "Petrov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				}},
			i:        0,
			j:        1,
			expected: false,
		},
		"less by last name": {
			p: People{
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Person{
					firstName: "Ivan",
					lastName:  "Petrov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				}},
			i:        0,
			j:        1,
			expected: true,
		},
		"greater by last name": {
			p: People{
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Person{
					firstName: "Ivan",
					lastName:  "Averin",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				}},
			i:        0,
			j:        1,
			expected: false,
		},
		"equal": {
			p: People{
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				Person{
					firstName: "Ivan",
					lastName:  "Ivanov",
					birthDay:  time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC),
				}},
			i:        0,
			j:        1,
			expected: false,
		},
	}

	for name, tc := range tests {
		tc := tc
		name := name
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tc.p.Less(tc.i, tc.j)
			if got != tc.expected {
				t.Errorf("[%s] expected: %t, got %t", name, tc.expected, got)
			}
		})
	}
}

func Test_Swap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		p        People
		i        int
		j        int
		expected People
	}{{
		name: "swapped",
		p: People{
			Person{
				firstName: "Ivan",
				lastName:  "Ivanov",
				birthDay:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Person{
				firstName: "Petr",
				lastName:  "Petrov",
				birthDay:  time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC),
			}},
		i: 0,
		j: 1,
		expected: People{
			Person{
				firstName: "Petr",
				lastName:  "Petrov",
				birthDay:  time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC),
			},
			Person{
				firstName: "Ivan",
				lastName:  "Ivanov",
				birthDay:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			}},
	}, {
		name: "same index",
		p: People{
			Person{
				firstName: "Ivan",
				lastName:  "Ivanov",
				birthDay:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Person{
				firstName: "Petr",
				lastName:  "Petrov",
				birthDay:  time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC),
			}},
		i: 0,
		j: 0,
		expected: People{
			Person{
				firstName: "Ivan",
				lastName:  "Ivanov",
				birthDay:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			Person{
				firstName: "Petr",
				lastName:  "Petrov",
				birthDay:  time.Date(2000, 10, 11, 0, 0, 0, 0, time.UTC),
			}},
	}}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.p.Swap(tc.i, tc.j)
			if tc.p.Len() != tc.expected.Len() {
				t.Errorf("[%s] expected len: %v, got len %v", tc.name, tc.expected.Len(), tc.p.Len())
			}
			for i := range tc.p {
				if tc.p[i] != tc.expected[i] {
					t.Errorf("[%s] expected: %v, got %v", tc.name, tc.expected, tc.p)
				}
			}
		})
	}
}

//////////////////////////////////////////////////////////////////////////////////////////

func Test_New(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		str      string
		expected *Matrix
		err      error
	}{
		{
			name: "square matrix",
			str:  "1 2 3\n4 5 6\n7 8 9",
			expected: &Matrix{
				rows: 3,
				cols: 3,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			err: nil,
		},
		{
			name: "non-square matrix",
			str:  "1 2 3 4 \n 5 6 7 8 ",
			expected: &Matrix{
				rows: 2,
				cols: 4,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8}},
			err: nil,
		},
		{
			name: "one row",
			str:  "1 2 3 4 5 6 7 8 9",
			expected: &Matrix{
				rows: 1,
				cols: 9,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			err: nil,
		},
		{
			name: "one column",
			str:  "1\n2\n3\n4",
			expected: &Matrix{
				rows: 4,
				cols: 1,
				data: []int{1, 2, 3, 4}},
			err: nil,
		},
		{
			name:     "unequal row lengths",
			str:      "1 2 3 4\n 5 6\n7 8 9",
			expected: nil,
			err:      errors.New("Rows need to be the same length"),
		},
		{
			name:     "incorrect character",
			str:      "1 a 3\n4 5 6\n7 8 9",
			expected: nil,
			err:      strconv.ErrSyntax,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := New(tc.str)
			if err != nil && err.Error() != tc.err.Error() && !errors.Is(err, tc.err) {
				t.Errorf("[%s] unexpected error occured: %v", tc.name, err.Error())
			}
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("[%s] expected: %v, got %v", tc.name, tc.expected, got)
			}
		})
	}
}

func Test_Rows(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		m        Matrix
		expected [][]int
	}{
		{
			name: "square matrix",
			m: Matrix{
				rows: 3,
				cols: 3,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			expected: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name: "non-square matrix",
			m: Matrix{
				rows: 2,
				cols: 4,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			expected: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
		},
		{
			name: "one row",
			m: Matrix{
				rows: 1,
				cols: 4,
				data: []int{1, 2, 3, 4},
			},
			expected: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "one column",
			m: Matrix{
				rows: 4,
				cols: 1,
				data: []int{1, 2, 3, 4},
			},
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name: "empty matrix",
			m: Matrix{
				rows: 0,
				cols: 0,
				data: []int{},
			},
			expected: [][]int{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.m.Rows()
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Cols(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		m        Matrix
		expected [][]int
	}{
		{
			name: "square matrix",
			m: Matrix{
				rows: 3,
				cols: 3,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			expected: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "non-square matrix",
			m: Matrix{
				rows: 2,
				cols: 4,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			expected: [][]int{{1, 5}, {2, 6}, {3, 7}, {4, 8}},
		},
		{
			name: "one row",
			m: Matrix{
				rows: 1,
				cols: 4,
				data: []int{1, 2, 3, 4},
			},
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name: "one column",
			m: Matrix{
				rows: 4,
				cols: 1,
				data: []int{1, 2, 3, 4},
			},
			expected: [][]int{{1, 2, 3, 4}},
		},
		{
			name: "empty matrix",
			m: Matrix{
				rows: 0,
				cols: 0,
				data: []int{},
			},
			expected: [][]int{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := tc.m.Cols()
			assert.Equal(t, tc.expected, got)
		})
	}
}

func Test_Set(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		m               Matrix
		row, col, value int
		expected        Matrix
		ok              bool
	}{
		{
			name: "success",
			m: Matrix{
				rows: 3,
				cols: 3,
				data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			row:   1,
			col:   1,
			value: 0,
			expected: Matrix{
				rows: 3,
				cols: 3,
				data: []int{1, 2, 3, 4, 0, 6, 7, 8, 9},
			},
			ok: true,
		},
		{
			name: "negative row",
			m: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			row:   -2,
			col:   0,
			value: 0,
			expected: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			ok: false,
		},
		{
			name: "row out of range",
			m: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			row:   5,
			col:   0,
			value: 0,
			expected: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			ok: false,
		},
		{
			name: "negative col",
			m: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			row:   1,
			col:   -9,
			value: 0,
			expected: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			ok: false,
		},
		{
			name: "col out of range",
			m: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			row:   0,
			col:   10,
			value: 0,
			expected: Matrix{
				rows: 2,
				cols: 2,
				data: []int{1, 2, 3, 4},
			},
			ok: false,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			ok := tc.m.Set(tc.row, tc.col, tc.value)
			assert.Equal(t, tc.ok, ok)
			assert.Equal(t, tc.expected, tc.m)
		})
	}
}
