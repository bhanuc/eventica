package team

import (
	"fmt"
	"github.com/eventica/user"
	"github.com/eventica/utility"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/sessions"
	"net/http"
	//"strconv"
	"strings"
)

type FlashMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

var R = user.R

func Addteam(pid string, tek_id string) {
	u, err := R.FindOneByTechID(tek_id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(u.Teams), u, pid, tek_id)
		fmt.Println("start push")
		if u.Teams == "" {
			u.Teams = pid
		} else {
			u.Teams = u.Teams + "," + pid
		}
		fmt.Println("pushed", u.Teams)
		u.Update()
	}
}
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	id, ok := session.Values["user"].(string)
	if ok {
		e := make(map[string]FlashMessage)
		tu := struct {
			Name    string `json:"name"`
			Members string `json:"members"`
			Event   string `json:"event"`
			Gender  string `json:"gender"`
		}{}
		utility.ReadJson(r, &tu)

		if tu.Name == "" {
			e["Error"] = FlashMessage{"danger", "Please fill in the name of team"}
		}
		if tu.Members == "" {
			e["Error"] = FlashMessage{"danger", "Please add team members"}
		}
		members := strings.Split(tu.Members, ",")
		mlength := len(members)
		fmt.Println(members, mlength)
		index := 0
		index2 := 0
		for i := 0; i < mlength; i++ {
			if members[i] != "" || members != nil {
				index++
				if R.FindCountByTechid(members[i]) != 0 {
					index2++
				}
			}
		}
		/**for _, member := range members {
							//cnum, _ := strconv.Atoi(member)
		 			if R.FindCountByTechid(member) == 0 {
		 				index++
		 			}
		    	}**/
		fmt.Println(index2, index)
		if index != index2 {
			e["Error"] = FlashMessage{"danger", "Only Techid Needs to be entered. One of your members seems to be not registered"}
			data["flashes"] = e
		} else {
			u, _ := R.FindOneByIdHex(id)

			//	d := T.CountByCollegenEvent(u.College, tu.Event)
			//	fmt.Println(d)

			//if d > 0 {
			//	e["Error"] = FlashMessage{"danger", "You have already registered for this event. Repeat registration of same game is not allowed"}
			//	data["flashes"] = e
			//	data["team"] = tu
			//} else {
			_, err := T.FindOneByName(tu.Name)
			if err != nil {
				fmt.Println("inside add team")
				team := new(Team)
				team.Add(tu.Name, tu.Members, tu.Event, tu.Gender, id, u.UserProfile.College)
				fmt.Println("made team", members, team.Id)
				for i := 0; i < len(members); i++ {
					fmt.Println("in loop", i, members[i], string(team.Id))
					if members[i] != "" || members != nil {
						Addteam(team.Id.String(), members[i])
					}
				}
				fmt.Println("completed team")
				data["success"] = true
				e["success"] = FlashMessage{"success", "Your Team registration was successful."}
				data["flashes"] = e
				data["team"] = tu
			} else {
				e["Error"] = FlashMessage{"danger", "This team Name is already registered. Please choose a difference name for your team"}
				data["flashes"] = e
				data["team"] = tu
			}
		}
		utility.WriteJson(w, data)

	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func ApprovalHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	_, ok := session.Values["user"].(string)
	if ok {
		status := false
		tc := struct {
			Name      string `json:"name"`
			CreatedBy string `json:"createdby"`
		}{r.FormValue("name"), r.FormValue("createdby")}

		flashes := make(map[string]FlashMessage)
		team, err := T.FindOneByName(tc.Name)

		// if team not found
		if err != nil {
			status = false
		} else {
			//check owner
			if team.CreatedBy == tc.CreatedBy {
				// request for moderation
				team.RequestModeration()
				http.Redirect(w, r, "/showteam", 302)
			} else {
				flashes["Error"] = FlashMessage{"danger", "Not your Team"}
				data["flashes"] = flashes
				utility.WriteJson(w, data)

			}

		}

		data["success"] = status

		//if !status {
		//	flashes["invalid"] = FlashMessage{"danger", "Your Team seems to already passed by moderators"}
		//	data["flashes"] = flashes
		//	}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func SingleTeamHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	_, ok := session.Values["user"].(string)
	if ok {
		status := false
		tc := struct {
			Id string `json:"id"`
		}{}
		utility.ReadJson(r, &tc)

		team, err := T.FindOneByIdHex(tc.Id)

		// if team not found
		if err != nil {
			status = false
		} else {
			data["team"] = team
			status = true
			utility.WriteJson(w, data)
		}

		data["success"] = status

		//if !status {
		//	flashes["invalid"] = FlashMessage{"danger", "Your Team seems to already passed by moderators"}
		//	data["flashes"] = flashes
		//	}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func TeamEditHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	_, ok := session.Values["user"].(string)
	if ok {
		status := false
		tc := struct {
			Id    string `json:"id"`
			Name  string `json:"name"`
			Event string `json:"event"`
		}{}
		utility.ReadJson(r, &tc)

		team, err := T.FindOneByIdHex(tc.Id)

		// if team not found
		if err != nil {
			status = false
		} else {
			team.Name = tc.Name
			team.Event = tc.Name
			team.Update()
			status = true
		}

		data["success"] = status
		utility.WriteJson(w, data)

		//if !status {
		//	flashes["invalid"] = FlashMessage{"danger", "Your Team seems to already passed by moderators"}
		//	data["flashes"] = flashes
		//	}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func paymentHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	flashes["success"] = FlashMessage{"success", "Payment portal still under creation. Please be patience!!"}
	data["flashes"] = flashes
	utility.WriteJson(w, data)
}

func AllTeamHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	id, ok := session.Values["user"].(string)
	if ok && session.Values["usertype"] == "user" {
		u, err := T.FindAllByCreatedBy(id)
		if err != nil {
			flashes["Error"] = FlashMessage{"danger", "You need to Create a Team Before you can see all the teams"}
			data["flashes"] = flashes
		} else {
			if len(u) == 0 {
				flashes["Error"] = FlashMessage{"danger", "You need to Create a Team Before you can see all the teams"}
				data["flashes"] = flashes
			} else {
				data["teams"] = u
				data["success"] = true
				flashes["AllTeams"] = FlashMessage{"success", "Your Teams have been fetched."}
				data["flashes"] = flashes
			}
		}
		utility.WriteJson(w, data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func AllTeamHandler2(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	id, ok := session.Values["user"].(string)
	//t := new(Team)
	if ok && session.Values["usertype"] == "user" {
		u, err := R.FindOneByIdHex(id)
		if err != nil {
			http.Redirect(w, r, "/login", 302)
		} else {
			if len(u.Teams) == 0 {
				flashes["Error"] = FlashMessage{"danger", "No One has Added to your Team."}
				data["flashes"] = flashes
			} else {
				var teamz = strings.Split(u.Teams, ",")
				temp_team := make([]Team, len(teamz))
				fmt.Println(teamz, len(teamz))
				for i := 0; i < len(teamz); i++ {
					fmt.Println(i, teamz[i], teamz[i][13:37])
					if teamz[i] != "" {
						fmt.Println(teamz[i])
						t, err := T.FindOneByIdHex(teamz[i][13:37])
						if err == nil {
							temp_team[i] = *t
						}
					}
				}
				/**for _, member := range u.Teams {
					u, err := T.FindOneByIdHex(member)
					if err == nil {
						temp_team[len(temp_team)] = *u
					}
				} **/
				data["teams"] = temp_team
				data["success"] = true
				flashes["AllTeams"] = FlashMessage{"success", "Your Teams have been fetched."}
				data["flashes"] = flashes
			}
		}
		utility.WriteJson(w, data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	tc := struct {
		Name string `json:"name"`
	}{r.FormValue("name")}
	if ok && session.Values["usertype"] == "admin" {
		u, err := T.FindOneByName(tc.Name)
		if err != nil {
			flashes["No Team Present"] = FlashMessage{"danger", "Wrong team name"}
			data["flashes"] = flashes
			utility.WriteJson(w, data)
		} else {
			data["success"] = u
			utility.WriteJson(w, data)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func TeamAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	tc := struct {
		Name string `json:"name"`
	}{r.FormValue("name")}
	if ok && session.Values["usertype"] == "admin" {
		u, err := T.FindOneByName(tc.Name)
		if err != nil {
			flashes["No Team Present"] = FlashMessage{"danger", "Wrong team name"}
			data["flashes"] = flashes
			utility.WriteJson(w, data)
		} else {
			u.Approved = "Approved"
			u.Update()
			http.Redirect(w, r, "/app", 302)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func DisTeamAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	tc := struct {
		Name string `json:"name"`
	}{r.FormValue("name")}
	if ok && session.Values["usertype"] == "admin" {
		u, err := T.FindOneByName(tc.Name)
		if err != nil {
			flashes["No Team Present"] = FlashMessage{"danger", "Wrong team name"}
			data["flashes"] = flashes
			utility.WriteJson(w, data)
		} else {
			u.Approved = "Declined"
			u.Update()
			http.Redirect(w, r, "/app", 302)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func CommentAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	tc := struct {
		Name     string `json:"name"`
		Comments string `json:"comments"`
	}{r.FormValue("name"), r.FormValue("comments")}
	if ok && session.Values["usertype"] != "user" {
		u, err := T.FindOneByName(tc.Name)
		if err != nil {
			flashes["No Team Present"] = FlashMessage{"danger", "Wrong team name"}
			data["flashes"] = flashes
			utility.WriteJson(w, data)
		} else {
			u.Comments = tc.Comments
			u.Update()
			http.Redirect(w, r, "/app", 302)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func AllAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	if ok && session.Values["usertype"] != "user" {
		u, err := T.FindAllByRequestMod()
		if err != nil {
			flashes["No Team Present"] = FlashMessage{"danger", "No teams are available for mod"}
			data["Error"] = flashes
		} else {
			data["teams"] = u
			data["success"] = true
			flashes["AllTeams"] = FlashMessage{"success", "The Teams have been fetched."}
			data["flashes"] = flashes
		}
		utility.WriteJson(w, data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

func UpdateAllTeam(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	if ok {
		u, err := T.All()
		if err != nil {
			flashes["No Team Present"] = FlashMessage{"danger", "No teams are available for mod"}
			data["Error"] = flashes
		} else {
			for i, c := range u {
				if c.College == "" {
					fmt.Println(i, c.CreatedBy)
					u, _ := R.FindOneByIdHex(c.CreatedBy)
					c.College = u.UserProfile.College
					c.Update()
				}
				//

			}
		}
		utility.WriteJson(w, data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	id, ok := session.Values["user"].(string)
	tu := struct {
		Name string `json:"name"`
	}{}
	params := mux.Vars(r)
	tu.Name = params["name"]
	fmt.Println(tu.Name)
	if ok {
		u, err := T.FindOneByName(tu.Name)
		fmt.Println(err)
		if err != nil {
			data["Error"] = "Your Team was not found"
		} else {
			if id == u.CreatedBy {
				data["sucess"] = u
			} else {
				data["Error"] = "Team was not created by you. So you cannot access it"
			}
		}
	}
	utility.WriteJson(w, data)
}
func CollegeAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	_, ok := session.Values["user"].(string)
	tu := struct {
		College string `json:"college"`
	}{}
	params := mux.Vars(r)
	tu.College = params["college"]
	if ok && session.Values["usertype"] != "admin" {
		u, err := T.FindAllByCollege(tu.College)
		fmt.Println(err)
		if err != nil {
			data["Error"] = "No Teams Registered yet !!"
		} else {
			data["teams"] = u
			data["success"] = true
		}
	}
	utility.WriteJson(w, data)
}

func EventAdminHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	_, ok := session.Values["user"].(string)
	tu := struct {
		Event string `json:"event"`
	}{}
	params := mux.Vars(r)
	tu.Event = params["event"]
	if ok && session.Values["usertype"] != "admin" {
		u, err := T.FindAllByEvent(tu.Event)
		fmt.Println(err)
		if err != nil {
			data["Error"] = "No Teams Registered yet !!"
		} else {
			data["teams"] = u
			data["success"] = true
		}
	}
	utility.WriteJson(w, data)
}
