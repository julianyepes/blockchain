package views

import "html/template"

func NewView(funcMap template.FuncMap, files ...string) *View {
	files = append(files,
		"views/layouts/bulma.gohtml",
		"views/layouts/footer.gohtml",
	)

	var t *template.Template
	var err error
	if funcMap != nil {
		t, err = template.New("").Funcs(funcMap).ParseFiles(files...)
	} else {
		t, err = template.ParseFiles(files...)
	}
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

type View struct {
	Template *template.Template
}
