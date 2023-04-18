package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type person struct {
	TITLE      string
	FIRST_NAME string
	LAST_NAME  string
}

var titles = []string{"Mr", "Mrs", "Dr", "Ms", "Prof"}

var people []person

func main() {
	fd, err := os.Open("data.csv")

	if err != nil {
		fmt.Println(err)
	}
	defer fd.Close()

	fileReader := csv.NewReader(fd)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}

	for key, val := range records {
		if key != 0 {
			if !strings.Contains(val[0], "and") && !strings.Contains(val[0], "&") {
				HandleCreateName(val[0], "")
			} else {
				if strings.Contains(val[0], "and") {
					HandleSpitNames(val[0], "and")
				}
				if strings.Contains(val[0], "&") {
					HandleSpitNames(val[0], "&")
				}
			}
		}

	}

	fmt.Println("people", people)
}

func HandleSpitNames(n string, conjuctor string) [] person {
	couple := strings.Split(n, conjuctor)
	namedPerson1 := HandleCreateName(couple[1], "")
	namedPerson2 := HandleCreateName(couple[0], namedPerson1.LAST_NAME)

	return [] person {
		namedPerson1,
		namedPerson2,
	}
}

func HandleCreateName(n string, ln string) person {
	isTitle, _ := in_array(strings.TrimSpace(n), titles)

	var title string
	var firstName string
	var lastName string

	if isTitle {
		title = n
		firstName = ""
		lastName = ln
	} else {
		parts := strings.Split(strings.TrimSpace(n), " ")

		switch len(parts) {
		case 1:
			lastName = parts[0]
		case 2:
			isTitle, _ := in_array(strings.TrimSpace(parts[0]), titles)
			if isTitle {
				title = parts[0]
			} else {
				firstName = parts[0]
			}	
			lastName = parts[1]
		case 3:
			title = parts[0]
			firstName = parts[1]
			lastName = parts[2]	
		}

		if ln != "" && len(parts) == 1 {
			lastName = ln
		}
	}

	namedPerson := person{TITLE: strings.TrimSpace(title), FIRST_NAME: strings.TrimSpace(firstName), LAST_NAME: strings.TrimSpace(lastName)}
	people = append(people, namedPerson)

	return namedPerson
}

func in_array(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}
