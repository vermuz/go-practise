package main

import (
	//"encoding/json"
	//"os"
	"log"
	"reflect"
	"regexp"
)

type User struct {
	FirstName  string `json:"first_name"`
	LastName   string
	Email      string `json:"email, omitempty"` // Omit if empty
	Password   string `json:"-"` // Do not marshal pass
}

// Returns String Array
func ColsForStruct(s interface{}) []string {
	// Give us all the different names of columns
	// of Json annotations
	cols := []string{};
	st := reflect.TypeOf(s)
	field_count := st.NumField()
	// log.Println(field_count)
	for i := 0; i < field_count; i++ {
		field := st.Field(i)
		// log.Println(field.Name)
		tag := field.Tag.Get("json")
		// cols = append(cols, field.Tag.Get("json"))
		if tag == "" {
			tag = field.Name
		}
		if tag != "-" {
			s := regexp.MustCompile(",").Split(tag, -1)
			//cols = append(cols, tag)
			cols = append(cols, s[0])
		}
	}
	return cols
}

func main() {
	u := User{"Mani", "Ali", "mani@example.com", "password"}
	// b, _ := json.Marshal(u)
	// os.Stdout.Write(b)
	cols := ColsForStruct(u)

	for _, column := range cols {
		log.Println(column)
	}
}

