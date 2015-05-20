package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"
	"strings"
)

var (
	flagBalancers string
)

func init() {
	flag.StringVar(&flagBalancers, "balancers", "", "processes that need a load balancer frontend")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "architect: create a cloudformation stack for a convox application\n\nUsage:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample:\n  architect -balancers web:3000,admin:4000 > formation.json\n")
	}
}

type Balancer struct {
	Name string
	Port string
}

func main() {
	flag.Parse()

	params := map[string]interface{}{
		"App":       nil,
		"Balancers": parseBalancers(flagBalancers),
	}

	data, err := buildTemplate("formation", "app", params)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error building json template: %s\n", err)
		os.Exit(1)
	}

	pretty, err := prettyJson(data)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error building json template: %s\n", err)
		printLines(data)
		displaySyntaxError(data, err)
		os.Exit(1)
	}

	fmt.Println(pretty)
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

func displaySyntaxError(data string, err error) {
	syntax, ok := err.(*json.SyntaxError)

	if !ok {
		fmt.Println(err)
		return
	}

	start, end := strings.LastIndex(data[:syntax.Offset], "\n")+1, len(data)

	if idx := strings.Index(data[start:], "\n"); idx >= 0 {
		end = start + idx
	}

	line, pos := strings.Count(data[:start], "\n"), int(syntax.Offset)-start-1

	fmt.Printf("Error in line %d: %s \n", line, err)
	fmt.Printf("%s\n%s^\n", data[start:end], strings.Repeat(" ", pos))
}

func parseBalancers(s string) []Balancer {
	list := parseList(s)

	b := make([]Balancer, len(list))

	for i, l := range list {
		parts := strings.SplitN(l, ":", 2)

		if len(parts) != 2 {
			fmt.Fprintf(os.Stderr, "balancers must be name:port pairs\n")
			os.Exit(1)
		}

		b[i] = Balancer{Name: parts[0], Port: parts[1]}
	}

	return b
}

func parseList(list string) []string {
	if list == "" {
		return []string{}
	}

	parts := strings.Split(list, ",")

	parsed := make([]string, len(parts))

	for i, p := range parts {
		parsed[i] = strings.TrimSpace(p)
	}

	return parsed
}

func prettyJson(raw string) (string, error) {
	var parsed map[string]interface{}

	if err := json.Unmarshal([]byte(raw), &parsed); err != nil {
		return "", err
	}

	bp, err := json.MarshalIndent(parsed, "", "  ")

	if err != nil {
		return "", err
	}

	clean := strings.Replace(string(bp), "\n\n", "\n", -1)

	return clean, nil
}

func printLines(data string) {
	lines := strings.Split(data, "\n")

	for i, line := range lines {
		fmt.Printf("%d: %s\n", i, line)
	}
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
			us := strings.ToUpper(name[0:1]) + name[1:]

			for {
				i := strings.Index(us, "-")

				if i == -1 {
					break
				}

				s := us[0:i]

				if len(us) > i+1 {
					s += strings.ToUpper(us[i+1 : i+2])
				}

				if len(us) > i+2 {
					s += us[i+2:]
				}

				us = s
			}

			return us
		},
	}
}
