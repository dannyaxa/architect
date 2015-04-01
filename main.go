package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"strings"
)

var (
	flagScale    StringSet
	flagServices StringSet
)

func init() {
	flagScale = StringSet{"web=1"}

	flag.Var(&flagScale, "scale", "process concurrency (can specify multiple times)")
	flag.Var(&flagServices, "service", "service env in KEY=value format (can specify multiple times)")
}

func main() {
	flag.Parse()

	fmt.Printf("flagServices %+v\n", flagServices)
	fmt.Printf("flag.Args() %+v\n", flag.Args())

	data, err := buildTemplate("formation", "app", nil)
	fmt.Printf("data %+v\n", data)
	fmt.Printf("err %+v\n", err)
}

func buildTemplate(name, section string, data interface{}) (string, error) {
	tmpl, err := template.New(section).Funcs(templateHelpers()).ParseFiles(fmt.Sprintf("template/%s.tmpl", name))

	if err != nil {
		return "", err
	}

	var formation bytes.Buffer

	err = tmpl.Execute(&formation, data)

	if err != nil {
		return "", err
	}

	return formation.String(), nil
}

func templateHelpers() template.FuncMap {
	return template.FuncMap{
		"array": func(ss []string) template.HTML {
			as := make([]string, len(ss))
			for i, s := range ss {
				as[i] = fmt.Sprintf("%q", s)
			}
			return template.HTML(strings.Join(as, ", "))
		},
		"join": func(s []string, t string) string {
			return strings.Join(s, t)
		},
		"ports": func(nn []int) template.HTML {
			as := make([]string, len(nn))
			for i, n := range nn {
				as[i] = fmt.Sprintf("%d", n)
			}
			return template.HTML(strings.Join(as, ","))
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"upper": func(name string) string {
			return strings.ToUpper(name[0:1]) + name[1:]
		},
	}
}

type StringSet []string

func (ss *StringSet) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

func (ss *StringSet) String() string {
	return fmt.Sprintf("[%s]", strings.Join(*ss, ","))
}
