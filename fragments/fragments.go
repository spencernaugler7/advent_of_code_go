package fragments

import (
	"html/template"
	"os"
)

var page *template.Template

func Init() {
	page = template.New("main")

	page = template.Must(page.Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<script src="https://unpkg.com/htmx.org@1.8.0"></script>
			<link rel="stylesheet" href="https://unpkg.com/missing.css@1.1.1"/>
			<title>Template Fragment Example</title>
		</head>
		<body>
			<h1>Template Fragment Example</h1>
			
			<p>This page demonstrates how to create and serve 
			<a href="https://htmx.org/essays/template-fragments/">template fragments</a> 
			using the <a href="https://pkg.go.dev/text/template">built-in template package</a> in Go.</p>
			
			<p>This is accomplished by using the "block" action in the template, which lets you
			define and execute a sub-template in a single step.</p>
			<!-- Here's the fragment.  We can target it by executing the "buttonOnly" template. -->
			{{block "buttonOnly" .}}
				<button hx-get="/?counter={{.next}}&template=buttonOnly" hx-swap="outerHTML">
					This Button Has Been Clicked {{.counter}} Times
				</button>
			{{end}}
		</body>
		</html>`))

	data := map[string]string {
		"next": "blah",
		"counter": "3",
	}

	page.ExecuteTemplate(os.Stdout, "buttonOnly", data)

	// page.Execute(os.Stdout, data)
}
