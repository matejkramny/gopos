package gopos

import (
	"bufio"
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