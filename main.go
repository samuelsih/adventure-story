package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/samuelsih/adventure-story/pkg"
)


type(
	//Story make map of adventure
	//Create this so we're not doing repetitive map[string]Adventure
	Story map[string]pkg.Adventure

	//handler holds story type
	handler struct {
		story Story
	}
) 

//port holds the port for http server
const (
	port = ":8080"
)

var defaultTemplate = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Adventure Story</title>
	</head>

	<body>
		<h1>{{ .Title }}</h1>
		
		{{ range .Story }}
			<p>{{ . }}</p>
		{{ end }}

		<ul>
			{{ range .Options }}
			{{ . }}
				<li><a href="/{{ .Arc }}">{{ .Text }}</a></li>
			{{ end }}
		{{ end }}
		</ul>
	</body>
	</html>
`

//NewHandler return http.Handler interface with the method ServeHTTP(ResponseWriter, *Request)
//and passing data type Story
func NewHandler(story Story) http.Handler {
	return handler{story}
}

//ServeHTTP implements http.Handler interface to render our template
func (h handler) ServeHTTP (rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("").Parse(defaultTemplate))

	err := tmpl.Execute(rw, h.story["intro"])

	if err != nil {
		panic(err)
	}
}

func main() {
	story := pkg.GetStory("story/story.json")

	handler := NewHandler(story)

	log.Println("Listening on port " + port)
	http.ListenAndServe(port, handler)
}
