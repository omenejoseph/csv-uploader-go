package main

import "testing"

type splitParameters struct {
	NAME     string
	EXPECTED person
}

type noSplitParameters struct {
	NAME      string
	EXPECTED  []person
	CONJUCTOR string
}

var noSplitTestParameters = []splitParameters{
	{NAME: "Mrs Faye Hughes-Eastwood", EXPECTED: person{TITLE: "Mrs", FIRST_NAME: "Faye", LAST_NAME: "Hughes-Eastwood"}},
	{NAME: "Prof Alex Brogan", EXPECTED: person{TITLE: "Prof", FIRST_NAME: "Alex", LAST_NAME: "Brogan"}},
	{NAME: "Dr P Gunn", EXPECTED: person{TITLE: "Dr", FIRST_NAME: "P", LAST_NAME: "Gunn"}},
}

var splitNamesTestParameters = []noSplitParameters{
	{
		NAME: "Mr and Mrs Smith",
		EXPECTED: []person{
			{TITLE: "Mr", FIRST_NAME: "", LAST_NAME: "Smith"},
			{TITLE: "Mrs", FIRST_NAME: "", LAST_NAME: "Smith"},
		},
		CONJUCTOR: "and",
	},
	{
		NAME: "Dr & Mrs Joe Bloggs",
		EXPECTED: []person{
			{TITLE: "Dr", FIRST_NAME: "", LAST_NAME: "Bloggs"},
			{TITLE: "Mrs", FIRST_NAME: "Joe", LAST_NAME: "Bloggs"},
		},
		CONJUCTOR: "&",
	},
	{
		NAME: "Mr Tom Staff and Mr John Doe",
		EXPECTED: []person{
			{TITLE: "Mr", FIRST_NAME: "Tom", LAST_NAME: "Staff"},
			{TITLE: "Mr", FIRST_NAME: "John", LAST_NAME: "Doe"},
		},
		CONJUCTOR: "and",
	},
}

func TestCreateNamesProperlyCreatesWithTitleFirstAndLastName(t *testing.T) {

	for _, test := range noSplitTestParameters {
		if output := HandleCreateName(test.NAME, ""); output != test.EXPECTED {
			t.Errorf("Output %q not equal to expected %q", output, test.EXPECTED)

		}
	}
}

func TestSplitNamesCorrectlySplitsNames(t *testing.T) {
	for _, test := range splitNamesTestParameters {
		output := HandleSpitNames(test.NAME, test.CONJUCTOR)

		if output[0] != test.EXPECTED[1] {
			t.Errorf("Output %q not equal to expected %q", output[0], test.EXPECTED[1])
		}

		if output[1] != test.EXPECTED[0] {
			t.Errorf("Output %q not equal to expected %q", output[1], test.EXPECTED[0])
		}
	}
}
