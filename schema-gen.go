package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateSchema() {
	// see all csv files
	c, err := os.ReadDir("data/csv")
	if err != nil {
		panic(err)
	}

	fmt.Println(c)

	mnames := []string{}
	// for each csv file
	for _, entry := range c {
		if strings.HasSuffix(entry.Name(), ".csv") {
			// generate schema
			mnames = append(mnames, GenerateSchemaForCSV(entry.Name()))
		}
	}

	t, err := template.ParseFiles("templates/migrator.go.tmpl")
	if err != nil {
		panic(err)
	}

	var o *os.File
	if _, err := os.Stat(filepath.Join("models", "main.go")); !os.IsNotExist(err) {
		err := os.Remove(filepath.Join("models", "main.go"))
		if err != nil {
			panic(err)
		}

	}
	o, err = os.Create(filepath.Join("models", "exec", "main.go"))
	if err != nil {
		panic(err)
	}

	defer o.Close()

	res := map[string]interface{}{
		"Models": []map[string]interface{}{},
	}

	for i := range mnames {
		res["Models"] = append(res["Models"].([]map[string]interface{}), map[string]interface{}{
			"Name": mnames[i],
		})
	}

	t.Execute(o, res)
}

func GenerateSchemaForCSV(filename string) string {
	t, err := template.ParseFiles("templates/model.go.tmpl")
	if err != nil {
		panic(err)
	}

	f, err := os.Open(filepath.Join("data/csv", filename))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	title, err := r.Read()
	if err != nil {
		panic(err)
	}

	first, err := r.Read()
	if err != nil {
		panic(err)
	}

	res := map[string]interface{}{}

	c := cases.Title(language.English)

	s := strings.ReplaceAll(filename, ".csv", "")
	s = c.String(s)
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ReplaceAll(s, "_", "")

	if len(s) <= 3 {
		s = strings.ToUpper(s)
	}

	mname := s

	res["ModelName"] = mname

	s = strings.ReplaceAll(filename, ".csv", "")
	s = strings.ReplaceAll(s, "-", "_")
	if strings.HasSuffix(s, "y") {
		s = s[:len(s)-1] + "ies"
	} else if strings.HasSuffix(s, "S") {
		s += "es"
	} else if strings.HasSuffix(s, "s") {
	} else {
		s += "s"
	}

	res["TableName"] = s

	res["Fields"] = []map[string]interface{}{}
	for i := range title {
		n := strings.ReplaceAll(title[i], "_", " ")
		n = c.String(n)
		n = strings.ReplaceAll(n, "-", "")
		n = strings.ReplaceAll(n, "_", "")
		n = strings.ReplaceAll(n, " ", "")

		res["Fields"] = append(res["Fields"].([]map[string]interface{}), map[string]interface{}{
			"FieldNameProper": n,
			"FieldNameSnake":  title[i],
			"FIeldExample":    first[i],
			"FieldType":       ParseType(first[i]),
		})
	}

	fname := strings.ReplaceAll(filename, ".csv", "")

	var o *os.File
	if _, err := os.Stat(filepath.Join("models", fname+".go")); !os.IsNotExist(err) {
		err := os.Remove(filepath.Join("models", fname+".go"))
		if err != nil {
			panic(err)
		}

	}
	o, err = os.Create(filepath.Join("models", fname+".go"))
	if err != nil {
		panic(err)
	}

	defer o.Close()

	err = t.Execute(o, res)
	if err != nil {
		panic(err)
	}

	return mname
}

// Parse Type
func ParseType(s string) string {
	// if able to be parsed as int, return int
	_, err := strconv.Atoi(s)
	if err == nil {
		return "int"
	}

	// if able to be parsed as float, return float
	_, err = strconv.ParseFloat(s, 64)
	if err == nil {
		return "float64"
	}

	// if able to be parsed as bool, return bool
	_, err = strconv.ParseBool(s)
	if err == nil {
		return "bool"
	}

	return "string"
}
