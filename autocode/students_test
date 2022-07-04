package coverage

import (
	"os"
	"testing"
	"time"
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
	},
		{
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
