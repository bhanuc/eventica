package user

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"labix.org/v2/mgo"
)

var (
	Config       *Conf
	Devmode      bool        = false
	Router       *mux.Router = mux.NewRouter()
	DB           *mgo.Database
	R            *UserRepository = new(UserRepository)
	mgoSession   *mgo.Session
	sessionStore *sessions.CookieStore
)

type Conf struct {
	Host         string // Host or rather domain name example.com
	MailFrom     string // Mail From address, can be "Name <web@example.com>" w/o quotes
	DBName       string // Mongo Database Name
	DBCollection string // Name of the User collection, typically "user"
}

func New(ms *mgo.Session, ss *sessions.CookieStore, host, mailfrom, dbname, dbcollection string) {
	mgoSession = ms.Clone()
	sessionStore = ss
	Config = &Conf{
		Host:         host,
		MailFrom:     mailfrom,
		DBName:       dbname,
		DBCollection: dbcollection,
	}
	DB = mgoSession.DB(Config.DBName)
	R.Collection = DB.C(Config.DBCollection)

	Router.HandleFunc("/user/register", RegisterHandler).Methods("POST")
	Router.HandleFunc("/user/makemanager", MakeManagerHandler).Methods("POST")
	Router.HandleFunc("/user/register", RegisterHandler).Methods("POST")
	Router.HandleFunc("/user/authenticate", AuthenticateHandler).Methods("POST")
	Router.HandleFunc("/user/logout", LogoutHandler).Methods("GET")
	Router.HandleFunc("/user/profile", ProfileHandler).Methods("GET")
	Router.HandleFunc("/user/vprofile", AdminViewProfile).Methods("GET")
	Router.HandleFunc("/user/mprofile", ManagerViewProfile).Methods("GET")
	Router.HandleFunc("/user/tek-profile", ManagerViewTekProfile).Methods("GET")
	Router.HandleFunc("/user/profile", UpdateProfileHandler).Methods("POST")
	Router.HandleFunc("/user/resetrequest", ResetRequestHandler).Methods("POST")
	Router.HandleFunc("/user/resetpassword", ResetPasswordHandler).Methods("GET")
	Router.HandleFunc("/user/status", StatusHandler).Methods("GET")
	Router.HandleFunc("/user/fb", FbHandler).Methods("POST")
	Router.HandleFunc("/user/activate", ActHandler).Methods("GET")
	Router.HandleFunc("/user/payactivate", PayHandler).Methods("GET")
}
