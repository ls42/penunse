package main

import (
	"log"
	"net/http"

	"github.com/alecthomas/template"
	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("penunse.bolt", 0600, nil)
	if err != nil {
		log.Fatal("cant' connect to database")
	}
	defer db.Close()

	mux := http.NewServeMux()
	t := template.Must(template.ParseGlob("templates/*"))
}

/*
func main() {

	// Render the primary application
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// TODO: Print all entries from the bolt database here
		hosts := systemdb.GetHosts(db)
		if len(hosts) == 0 {
			fmt.Println("Database empty")
		}
		err = t.ExecuteTemplate(w, "main", hosts)
		checkErr(err)
	})

	// Handle upload request of agents (JSON data)
	mux.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-Token") != conf.AuthToken {
			http.Error(w, "Upload denied", 403)
			return
		}
		r.ParseForm()
		apiContent, err := ioutil.ReadAll(r.Body)
		checkErr(err)
		var host systemdb.Host
		err = json.Unmarshal(apiContent, &host)
		if err != nil {
			http.Error(w, "Invalid JSON format", 400)
			return
		}
		host.Save(db)
	})

	// Edit or view Host details
	mux.HandleFunc("/edit/", func(w http.ResponseWriter, r *http.Request) {
		hostname := r.URL.Path[len("/edit/"):]
		host := systemdb.GetHost(hostname, db)
		err = t.ExecuteTemplate(w, "edit", host)
		checkErr(err)
	})

	// Save changes from /edit
	mux.HandleFunc("/save/", func(w http.ResponseWriter, r *http.Request) {
		hostname := r.URL.Path[len("/save/"):]
		host := systemdb.GetHost(hostname, db)
		host.Login = r.FormValue("Login")
		host.Contact = r.FormValue("Contact")
		host.Save(db)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	// Remove Host from database
	mux.HandleFunc("/delete/", func(w http.ResponseWriter, r *http.Request) {
		hostname := r.URL.Path[len("/delete/"):]
		systemdb.DeleteHost(hostname, db)
		http.Redirect(w, r, "/", http.StatusFound)
	})

	// Serve static files
	mux.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))),
	)

	fmt.Printf("Listening on port %s\n", conf.Port)
	http.ListenAndServe(conf.Port, mux)
}
*/
