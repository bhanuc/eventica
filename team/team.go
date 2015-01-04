package team

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"labix.org/v2/mgo"
)

var (
	Router       *mux.Router     = mux.NewRouter()
	T            *TeamRepository = new(TeamRepository)
	mgoSession   *mgo.Session
	Config       *Conf
	Devmode      bool = false
	DB           *mgo.Database
	sessionStore *sessions.CookieStore
)

type Conf struct {
	Host         string // Host or rather domain name example.com
	MailFrom     string // Mail From address, can be "Name <web@example.com>" w/o quotes
	DBName       string // Mongo Database Name
	DBCollection string // Name of the User collection, typically "user"
}

func Route_setter(ms *mgo.Session, ss *sessions.CookieStore, host, mailfrom, dbname, dbcollection string) {
	mgoSession = ms.Clone()
	sessionStore = ss
	Config = &Conf{
		Host:         host,
		MailFrom:     mailfrom,
		DBName:       dbname,
		DBCollection: dbcollection,
	}
	DB = mgoSession.DB(Config.DBName)
	T.Collection = DB.C(Config.DBCollection)

	Router.HandleFunc("/team/create", CreateHandler).Methods("POST")          //Create a new Team
	Router.HandleFunc("/team/view", ViewHandler).Methods("GET")               //View a new Team
	Router.HandleFunc("/team/register", ApprovalHandler).Methods("GET")       // Make the team ready for approval
	Router.HandleFunc("/team/payment", paymentHandler).Methods("GET")         // Once approved, this handles the payment logic
	Router.HandleFunc("/team/all", AllTeamHandler).Methods("GET")             //show all the teams
	Router.HandleFunc("/team/single", SingleTeamHandler).Methods("POST")      //show all the teams
	Router.HandleFunc("/team/adminevent", AdminSingleEvent).Methods("POST")   //show all the teams
	Router.HandleFunc("/team/event", ManagerSingleEvent).Methods("POST")      //show all the teams
	Router.HandleFunc("/team/update", TeamEditMembersHandler).Methods("POST") //show all the teams
	Router.HandleFunc("/team/all2", AllTeamHandler2).Methods("GET")           //show all the teams
	//Router.HandleFunc("/team/updatethis", UpdateAllTeamMembers).Methods("GET")        //show all the teams
	Router.HandleFunc("/team/status/{name}", StatusHandler).Methods("GET")            //show team status
	Router.HandleFunc("/team/adminall", AllAdminHandler).Methods("GET")               //show all the teams
	Router.HandleFunc("/team/adminall/{college}", CollegeAdminHandler).Methods("GET") //show all the teams
	Router.HandleFunc("/team/adminall2/{event}", EventAdminHandler).Methods("GET")    //show all the teams
	Router.HandleFunc("/team/admincount/{event}", EventCountHandler).Methods("GET")   //show all the teams
	Router.HandleFunc("/team/admintotalcount", TotalCountHandler).Methods("GET")      //show all the teams
	Router.HandleFunc("/team/adminsapprove", TeamAdminHandler).Methods("GET")         //show all the teams
	Router.HandleFunc("/team/managersapprove", TeamManagerHandler).Methods("GET")     //show all the teams
	Router.HandleFunc("/team/adminsdissapprove", DisTeamAdminHandler).Methods("GET")  //show all the teams
	Router.HandleFunc("/team/adminscomment", CommentAdminHandler).Methods("GET")      //show all the teams
	Router.HandleFunc("/team/managerscomment", CommentManagerHandler).Methods("GET")  //show all the teams

}

//}
