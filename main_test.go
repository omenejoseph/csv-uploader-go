
package main

import (
	"reflect"
	"testing"
)

func TestHandleSpitNames(t *testing.T) {
	cases := []struct {
		name           string
		n              string
		conjuctor      string
		expectedResult []person
	}{
		{
			name:      "Test split with 'and'",
			n:         "Mr Smith and Mrs Smith",
			conjuctor: "and",
			expectedResult: []person{
				{TITLE: "Mrs", FIRST_NAME: "", LAST_NAME: "Smith"},
				{TITLE: "Mr", FIRST_NAME: "", LAST_NAME: "Smith"},
			},
		},
		{
			name:      "Test split with '&'",
			n:         "Mr Jones & Mrs Jones",
			conjuctor: "&",
			expectedResult: []person{
				{TITLE: "Mrs", FIRST_NAME: "", LAST_NAME: "Jones"},
				{TITLE: "Mr", FIRST_NAME: "", LAST_NAME: "Jones"},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := HandleSpitNames(c.n, c.conjuctor)

			if !reflect.DeepEqual(result, c.expectedResult) {
				t.Errorf("Expected %v, but got %v", c.expectedResult, result)
			}
		})
	}
}

func TestHandleCreateName(t *testing.T) {
	cases := []struct {
		name           string
		n              string
		ln             string
		expectedResult person
	}{
		{
			name:      "Test single name",
			n:         "Smith",
			ln:        "",
			expectedResult: person{TITLE:"", FIRST_NAME: "", LAST_NAME: "Smith"},
		},
		{
			name:      "Test title and single name",
			n:         "Mr Smith",
			ln:        "",
			expectedResult: person{TITLE: "Mr", FIRST_NAME: "", LAST_NAME: "Smith"},
		},
		{
			name:      "Test full name",
			n:         "Mr John Smith",
			ln:        "",
			expectedResult: person{TITLE: "Mr", FIRST_NAME: "John", LAST_NAME: "Smith"},
		},
		{
			name:      "Test last name provided",
			n:         "Mrs",
			ln:        "Doe",
			expectedResult: person{TITLE: "Mrs", FIRST_NAME: "", LAST_NAME: "Doe"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result := HandleCreateName(c.n, c.ln)

			if !reflect.DeepEqual(result, c.expectedResult) {
				t.Errorf("Expected %v, but got %v", c.expectedResult, result)
			}
		})
	}
}

func TestInArray(t *testing.T) {
	cases := []struct {
		name           string
		v              interface{}
		in             interface{}
		expectedResult bool
		expectedIndex  int
	}{
		{
			name:           "Test in array",
			v:              "b",
			in:             []string{"a", "b", "c"},
			expectedResult: true,
			expectedIndex:  1,
		},
		{
			name:           "Test not in array",
			v:              "d",
			in:             []string{"a", "b", "c"},
			expectedResult: false,
			expectedIndex:  -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			result, idx := in_array(c.v, c.in)

			if result != c.expectedResult {
				t.Errorf("Expected %v, but got %v", c.expectedResult, result)
			}

			if idx != c.expectedIndex {
				t.Errorf("Expected index %d, but got %d", c.expectedIndex, idx)
			}
		})
	}
}