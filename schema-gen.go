package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"os/exec"
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

	GenerateMigrator(mnames)
	GenerateImpoter(mnames)
}

func GenerateImpoter(mnames []string) {
	t, err := template.ParseFiles("templates/importer.go.tmpl")
	if err != nil {
		panic(err)
	}

	var o *os.File
	if _, err := os.Stat(filepath.Join("models", "importer", "main.go")); !os.IsNotExist(err) {
		err := os.Remove(filepath.Join("models", "importer", "main.go"))
		if err != nil {
			panic(err)
		}
	}
	o, err = os.Create(filepath.Join("models", "importer", "main.go"))
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

	err = t.Execute(o, res)
	if err != nil {
		panic(err)
	}
}

func GenerateMigrator(mnames []string) {
	t, err := template.ParseFiles("templates/migrator.go.tmpl")
	if err != nil {
		panic(err)
	}

	var o *os.File
	if _, err := os.Stat(filepath.Join("models", "exec", "main.go")); !os.IsNotExist(err) {
		err := os.Remove(filepath.Join("models", "exec", "main.go"))
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

	err = t.Execute(o, res)
	if err != nil {
		panic(err)
	}
}

func GenerateSchemaForCSV(filename string) string {
	t, err := template.ParseFiles("templates/model.go.tmpl")
	if err != nil {
		panic(err)
	}

	path := filepath.Join("data/csv", filename)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	title, err := r.Read()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("cp", path, filepath.Join("models", filename))

	fmt.Println(cmd.String())

	err = cmd.Run()
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

	res["Filename"] = filename
	res["TableName"] = s

	res["Fields"] = []map[string]interface{}{}
	for i := range title {
		n := strings.ReplaceAll(title[i], "_", " ")
		n = c.String(n)
		n = strings.ReplaceAll(n, "-", "")
		n = strings.ReplaceAll(n, "_", "")
		n = strings.ReplaceAll(n, " ", "")
		v := map[string]interface{}{
			"FieldNameProper": n,
			"FieldNameSnake":  title[i],
			"FieldExample":    first[i],
		}
		ty := ParseType(first[i])
		switch ty {
		case "int":
			v["Statement"] = fmt.Sprintf("utils.ToInt(records[i][%d])", i)
		case "float64":
			v["Statement"] = fmt.Sprintf("utils.ToFloat64(records[i][%d])", i)
		case "bool":
			v["Statement"] = fmt.Sprintf("utils.ToBool(records[i][%d])", i)
		case "string":
			v["Statement"] = fmt.Sprintf("records[i][%d]", i)
		}

		v["FieldType"] = ty

		res["Fields"] = append(res["Fields"].([]map[string]interface{}), v)
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
	// _, err := strconv.Atoi(s)
	// if err == nil {
	// 	return "int"
	// }

	// if able to be parsed as float, return float
	_, err := strconv.ParseFloat(s, 64)
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
