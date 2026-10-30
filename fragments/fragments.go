package fragments

import (
	"fmt"
	"html/template"
	"os"
)


func Init() {
	page := template.Must(template.ParseGlob("./fragments/*.html"))
	fmt.Println(page.DefinedTemplates())

	page.ExecuteTemplate(os.Stdout, "test.html", map[string]string { })

}
