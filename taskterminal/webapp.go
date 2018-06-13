package taskterminal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	VERSION = "0.6"
	S3URL   = "https://s3.amazonaws.com/todolist-local/" + VERSION
)

type Webapp struct {
	Router *httprouter.Router
}

func NewWebapp() *Webapp {
	return &Webapp{Router: setupRoutes()}
}

func (w *Webapp) Run() {
	log.Fatal(http.ListenAndServe(":7890", w.Router))
}

func setupRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", IndexScaffold)
	router.OPTIONS("/tasks", TaskOptions)
	router.GET("/tasks", GetTasks)
	router.POST("/tasks", SaveTasks)
	router.NotFound = http.HandlerFunc(RedirectScaffold)
	return router
}

func IndexScaffold(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	template := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <link rel="stylesheet" href="https://bootswatch.com/flatly/bootstrap.min.css">
    <title>TaskTerminal</title>
    <link href="` + urlFor("main.css") + `" rel="stylesheet">
  </head>
  <body>
    <div id="app"></div>
    <script type="text/javascript" src="` + urlFor("common.js") + `"></script>
    <script type="text/javascript" src="` + urlFor("main.js") + `"></script>
  </body>
</html>
	`
	fmt.Fprintf(w, template)
}

func RedirectScaffold(w http.ResponseWriter, r *http.Request) {
	template := `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width,initial-scale=1">
    <link rel="stylesheet" href="https://bootswatch.com/flatly/bootstrap.min.css">
    <title>TaskTerminal</title>
    <link href="` + urlFor("main.css") + `" rel="stylesheet">
  </head>
  <body>
    <div id="app"></div>
    <script type="text/javascript" src="` + urlFor("common.js") + `"></script>
    <script type="text/javascript" src="` + urlFor("main.js") + `"></script>
  </body>
</html>
	`
	fmt.Fprintf(w, template)
}

func urlFor(file string) string {
	return S3URL + "/" + file
}

func RedirectToIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, S3URL+r.URL.Path, 301)
}

func GetTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	app := NewApp()
	app.Load()
	json, _ := json.Marshal(app.TaskTerminal.Data)
	fmt.Fprintf(w, string(json))
}
func TaskOptions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "")
}

func SaveTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	decoder := json.NewDecoder(r.Body)
	var tasks []*Task
	err := decoder.Decode(&tasks)
	if err != nil {
		log.Fatal("encountered an error parsing json, ", err)
	}
	app := NewApp()
	app.TaskStore.Load()
	app.TaskStore.Save(tasks)
}
