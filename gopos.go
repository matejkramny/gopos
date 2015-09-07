package gopos

import (
	"bufio"
	"bytes"
	"strings"
	"text/template"
)

func (pos *ESCPOS) PrintTemplate(tmplData string) {
	funcMap := template.FuncMap{
		"lf": LineFeed,
		"feed": Feed,
		"reverseFeed": ReverseFeed,
		"cut": Cut,
		"underline": Underline,
		"emphesize": Emphesize,
		"doubleStrike": DoubleStrike,
		"font": Font,
		"justify": Justify,
		"generatePulse": GeneratePulse,
		"reverseBlackWhite": ReversePrint,
	}

	tmpl := template.New("receipt")
	tmpl.Delims("[[", "]]")
	tmpl.Funcs(funcMap)

	tmpl = template.Must(tmpl.Parse(tmplData))

	buffer := bufio.NewWriter(pos.Connection)

	if err := tmpl.Execute(buffer, map[string]interface{}{}); err != nil {
		panic(err)
	}

	buffer.Flush()
}

func RenderTemplate(tmplData string) *bytes.Buffer {
	funcMap := template.FuncMap{
		"lf": LineFeed,
		"feed": Feed,
		"reverseFeed": ReverseFeed,
		"cut": Cut,
		"underline": Underline,
		"emphesize": Emphesize,
		"doubleStrike": DoubleStrike,
		"font": Font,
		"justify": Justify,
		"generatePulse": GeneratePulse,
		"reverseBlackWhite": ReversePrint,
		"at": func () string {
			return "\xa3"
		},
		"spaces": func (one string, two string) string {
			return strings.Repeat(" ", 48 - (len(one) + len(two) + 1))
		},
	}

	tmpl := template.New("receipt")
	tmpl.Delims("[[", "]]")
	tmpl.Funcs(funcMap)

	tmpl = template.Must(tmpl.Parse(tmplData))

	buffer := new(bytes.Buffer)
	if err := tmpl.Execute(buffer, map[string]interface{}{}); err != nil {
		panic(err)
	}

	return buffer
}