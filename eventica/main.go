package main

import (
	"fmt"
	"github.com/eventica/team"
	"github.com/eventica/user"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"labix.org/v2/mgo"
	"log"
	"net/http"
	"text/template"
)

type FlashMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

var (
	router          *mux.Router = mux.NewRouter()
	mgoSession      *mgo.Session
	sessionStore    *sessions.CookieStore
	sessionAuthKey  []byte = make([]byte, 64)
	sessionCryptKey []byte = make([]byte, 32)
)

var R = user.R

func main() {
	var err error
	mgoSession, err = mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	sessionAuthKey = securecookie.GenerateRandomKey(64)
	sessionCryptKey = securecookie.GenerateRandomKey(32)
	sessionStore = sessions.NewCookieStore(sessionAuthKey, sessionCryptKey)
	sessionStore.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 0,
	}
	user.New(mgoSession, sessionStore, "techkriti.org", "noreply@techkriti.org", "local", "users")
	team.Route_setter(mgoSession, sessionStore, "techkriti.org", "noreply@techkriti.org", "local", "teams")
	// Thanks to Kamil Kisiel
	router.PathPrefix("/user").Handler(user.Router)
	router.PathPrefix("/team").Handler(team.Router)
	router.PathPrefix("/api").Handler(http.StripPrefix("/api", router))
	router.HandleFunc("/profilea/{name}", profile)
	router.HandleFunc("/login", login).Methods("GET")
	router.HandleFunc("/profile", profile_view).Methods("GET")
	router.HandleFunc("/register_event", register_event).Methods("GET")
	router.HandleFunc("/showteam", show_team).Methods("GET")
	router.HandleFunc("/print", print_team).Methods("GET")
	router.HandleFunc("/signup", signup).Methods("GET")
	router.HandleFunc("/app", app).Methods("GET")
	router.HandleFunc("/fb", fblogin).Methods("GET")
	router.HandleFunc("/", home)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", router)
	log.Println("Listening...")
	http.ListenAndServe(":2368", nil)
}

func print_team(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	_, ok := session.Values["user"].(string)
	if ok && session.Values["usertype"] == "admin" {
		t, err := template.New("print_team.html").ParseFiles("templates/print_team.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func profile(w http.ResponseWriter, r *http.Request) { //Example to get variables from request
	params := mux.Vars(r)
	name := params["name"]
	fmt.Println(r.URL.Query())
	w.Write([]byte("Hello " + name))
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("login.html").ParseFiles("templates/login.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
func fblogin(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("facebook.html").ParseFiles("templates/facebook.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
func profile_view(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	_, ok := session.Values["user"].(string)
	if ok {
		t, err := template.New("CompleteProfile.html").ParseFiles("templates/CompleteProfile.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func register_event(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	id, ok := session.Values["user"].(string)
	if ok {
		u, _ := R.FindOneByIdHex(id)
		profilestatus := u.ProfileStatus
		data := struct {
			Status bool
		}{profilestatus}
		t, err := template.New("registerforevent.html").ParseFiles("templates/registerforevent.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, data)
		if err != nil {
			log.Println(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func show_team(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	_, ok := session.Values["user"].(string)
	if ok {
		t, err := template.New("showteam.html").ParseFiles("templates/showteam.html")
		if err != nil {
			log.Println(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("register.html").ParseFiles("templates/register.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
func app(w http.ResponseWriter, r *http.Request) {
	session, err := sessionStore.Get(r, "p")
	if err != nil {
		log.Println(err)
	}
	id, ok := session.Values["user"].(string)
	if ok {
		log.Println(session.Values["usertype"])
		if session.Values["usertype"] == "user" {
			u, _ := R.FindOneByIdHex(id)
			profilestatus := u.ProfileStatus
			data := struct {
				Status bool
			}{profilestatus}
			fmt.Println(profilestatus)
			t, err := template.New("dashboard.html").ParseFiles("templates/dashboard.html")
			if err != nil {
				log.Println(err)
			}
			err = t.Execute(w, data)
			if err != nil {
				log.Println(err)
			}
		} else {
			t, err := template.New("admindashboard.html").ParseFiles("templates/admindashboard.html")
			if err != nil {
				log.Println(err)
			}
			err = t.Execute(w, nil)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("home.html").ParseFiles("templates/home.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
