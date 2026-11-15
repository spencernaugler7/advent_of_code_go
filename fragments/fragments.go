package fragments

import (
	"fmt"
	"html/template"
	"io"
	"log"
)

func Register(registrar map[string]func(io.Writer)) map[string]func(io.Writer) {
	registrar["fragments"] = Run
	return registrar
}

func Run(writer io.Writer) {
	page := template.Must(template.ParseGlob("./fragments/*.html"))
	fmt.Fprintln(writer, page.DefinedTemplates())

	page.ExecuteTemplate(writer, "test.html", map[string]string{})

	// template test. test out defining blocks in different templates.
	const (
		master = `
		<body>
			{{block "body" .}}<h1>Default</h1>{{end}}
		</body>`
		overlay  = `{{define "body"}}<h1>{{ . }}</h1>{{end}}`
		overlay2 = `{{define "body"}}<h1>Hello from overlay2</h1>{{end}}`
	)
	var testVar string = "Spencer"
	masterTmpl, err := template.New("master").Parse(master)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	if err := masterTmpl.Execute(writer, testVar); err != nil {
		log.Fatal(err)
		panic(err)
	}
	if err := overlayTmpl.Execute(writer, testVar); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
