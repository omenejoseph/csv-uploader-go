
package main

import (
	"testing"
)

func TestInArray(t *testing.T) {
	ok, _ := in_array("Mr", titles)
	if !ok {
		t.Errorf("Expected 'Mr' to be found in titles")
	}

	ok, _ = in_array("RandomTitle", titles)
	if ok {
		t.Errorf("Expected 'RandomTitle' not to be found in titles")
	}
}

func TestHandleCreateName(t *testing.T) {
	createdPerson := HandleCreateName("John Doe", "")

	if createdPerson.TITLE != "" || createdPerson.FIRST_NAME != "John" || createdPerson.LAST_NAME != "Doe" {
		t.Errorf("Expected person to have TITLE: '', FIRST_NAME: 'John', LAST_NAME: 'Doe'")
	}
}

func TestHandleSpitNames(t *testing.T) {
	response := HandleSpitNames("Mr. John Smith and Mrs. Jane Smith", "and")

	if len(response) != 2 {
		t.Errorf("Expected 2 people structs to return")
	}

	if response[0].TITLE != "Mrs." || response[0].FIRST_NAME != "Jane" || response[0].LAST_NAME != "Smith" {
		t.Errorf("Expected person 1 to have TITLE: 'Mrs.', FIRST_NAME: 'Jane', LAST_NAME: 'Smith'")
	}

	if response[1].TITLE != "Mr." || response[1].FIRST_NAME != "John" || response[1].LAST_NAME != "Smith" {
		t.Errorf("Expected Person 2 to have TITLE: 'Mr.', FIRST_NAME: 'John', LAST_NAME: 'Smith'")
	}
}