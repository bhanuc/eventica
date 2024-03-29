package user

import (
	"encoding/json"
	"fmt"
	"github.com/eventica/utility"
	"github.com/gorilla/sessions"
	fb "github.com/huandu/facebook"
	"labix.org/v2/mgo/bson"
	"net/http"
	"strings"
	"sync"
)

type FlashMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

//var T = team.T
var wg sync.WaitGroup

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	turn := true
	data := make(map[string]interface{})
	e := make(map[string]FlashMessage)
	tu := struct {
		Name            string `json:"name"`
		Password        string `json:"password"`
		Email           string `json:"email"`
		Password2       string `json:"Password2"`
		Number          string `json:"number"`
		AlternateNumber string `json:"alternatenumber"`
	}{}
	utility.ReadJson(r, &tu)

	//tu.Name = strings.Trim(tu.Name, " ")

	tu.Email = strings.Trim(tu.Email, " ")
	//tu.Email2 = strings.Trim(tu.Email2, " ")

	if tu.Password == "" {
		e["Error"] = FlashMessage{"danger", "Please fill in password"}
	}
	if tu.Email == "" {
		e["Error"] = FlashMessage{"danger", "Please enter a valid email-address"}
	}
	if tu.Password != tu.Password2 {
		e["Error"] = FlashMessage{"danger", "The two Password don't match"}
	}
	//userrepo.Collection.Find(bson.M{"email": tu.Email}).One(&foundmail)
	c, _ := R.Collection.Find(bson.M{"email": tu.Email}).Count()
	//	u := R.FindOneByCollege(tu.College)

	if c > 0 && turn {
		e["Error"] = FlashMessage{"danger", "This Email ID is already registered. Please try using a different Email ID"}
	}
	//if u > 0 {
	//	e["Error"] = FlashMessage{"danger", "Your College is already registered. Please Contact ""}
	//}
	if len(e) == 0 {
		user := new(User)
		turn = false
		wg.Add(1)
		go user.Add(tu.Name, tu.Password, tu.Email, tu.Number, tu.AlternateNumber, &wg)
		wg.Wait()
		data["success"] = true
		// e["success"] = FlashMessage{"success", "Your registration is pending. Please check your inbox to activate your account"}
		e["success"] = FlashMessage{"success", "Your registration is pending. Please check your inbox to activate your account"}
		data["flashes"] = e
		data["user"] = tu
	} else {
		data["flashes"] = e
		data["user"] = tu
	}
	utility.WriteJson(w, data)
}

func MakeManagerHandler(w http.ResponseWriter, r *http.Request) {
	turn := true
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	e := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	if ok && session.Values["usertype"] == "admin" {
		tu := struct {
			Name            string `json:"name"`
			Password        string `json:"password"`
			Email           string `json:"email"`
			Password2       string `json:"Password2"`
			Number          string `json:"number"`
			AlternateNumber string `json:"alternatenumber"`
			Event           string `json:"event"`
		}{}
		utility.ReadJson(r, &tu)

		//tu.Name = strings.Trim(tu.Name, " ")

		tu.Email = strings.Trim(tu.Email, " ")
		//tu.Email2 = strings.Trim(tu.Email2, " ")

		if tu.Password == "" {
			e["Error"] = FlashMessage{"danger", "Please fill in password"}
		}
		if tu.Email == "" {
			e["Error"] = FlashMessage{"danger", "Please enter a valid email-address"}
		}
		if tu.Password != tu.Password2 {
			e["Error"] = FlashMessage{"danger", "The two Password don't match"}
		}
		//userrepo.Collection.Find(bson.M{"email": tu.Email}).One(&foundmail)
		c, _ := R.Collection.Find(bson.M{"email": tu.Email}).Count()
		//	u := R.FindOneByCollege(tu.College)

		if c > 0 && turn {
			e["Error"] = FlashMessage{"danger", "This Email ID is already registered. Please try using a different Email ID"}
		}
		//if u > 0 {
		//	e["Error"] = FlashMessage{"danger", "Your College is already registered. Please Contact ""}
		//}
		if len(e) == 0 {
			user := new(User)
			turn = false
			go user.AddManager(tu.Name, tu.Password, tu.Email, tu.Number, tu.AlternateNumber, tu.Event)
			data["success"] = true
			e["success"] = FlashMessage{"success", "Your registration is pending. Please check your inbox to activate your account"}
			data["flashes"] = e
			data["user"] = tu
		} else {
			data["flashes"] = e
			data["user"] = tu
		}
		utility.WriteJson(w, data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func FbHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	e := make(map[string]FlashMessage)
	tu := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Id    string `json:"id"`
		Token string `json:"token"`
	}{}
	utility.ReadJson(r, &tu)
	res, err := fb.Get("/me", fb.Params{
		"access_token": tu.Token,
	})

	if err != nil {
		e["Error"] = FlashMessage{"danger", "Facebook Login was not successful. Please use the traditional method"}
	} else {
		tr := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Id    string `json:"id"`
			Token string `json:"token"`
		}{}
		//buf := new(bytes.Buffer)
		b, err := json.Marshal(res)
		if err != nil {

			e["Error"] = FlashMessage{"danger", "Response could not be parsed. Please login again."}
		}
		nerr := json.Unmarshal(b, &tr)
		if nerr != nil {

			e["Error"] = FlashMessage{"danger", "Response could not be parsed. Please login again."}
		}
		tu.Email = strings.Trim(tu.Email, " ")
		tr.Email = strings.Trim(tr.Email, " ")

		if tu.Email == "" {
			e["Error"] = FlashMessage{"danger", "Please enter a valid email-address"}
		}
		if tu.Email != tr.Email {
			e["Error"] = FlashMessage{"danger", "Facebook Session Expired. Please try again."}
		}
		//userrepo.Collection.Find(bson.M{"email": tu.Email}).One(&foundmail)
		c, _ := R.Collection.Find(bson.M{"email": tr.Email}).Count()
		uzer, err := R.FindOneByEmail(tr.Email)
		if c > 0 {
			// login into the account
			uzer.CheckTechID()
			session, _ := sessionStore.Get(r, "p")
			session.Values["user"] = uzer.Id.Hex()
			session.Values["usertype"] = uzer.UserType
			session.Values["profilestatus"] = uzer.ProfileStatus
			session.Save(r, w)
			e["success"] = FlashMessage{"success", "Welcome back, your login was successfull"}
			data["flashes"] = e
			data["user"] = uzer
		}
		if len(e) == 0 {
			user := new(User)
			user.FbAdd(tu.Name, tu.Email, tu.Id, tu.Token)
			session, _ := sessionStore.Get(r, "p")
			session.Values["user"] = user.Id.Hex()
			session.Values["usertype"] = user.UserType
			session.Values["profilestatus"] = user.ProfileStatus
			session.Save(r, w)
			e["success"] = FlashMessage{"success", "Your registration was successful. Your registered Email ID is " + tu.Email + "."}
			data["flashes"] = e
			data["user"] = tu
		} else {
			data["flashes"] = e
			data["user"] = tu
		}
		data["flashes"] = e
	}
	data["flashes"] = e
	utility.WriteJson(w, data)
}

/**
func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	// get email + password
	valid := false
	unactive := false
	data := make(map[string]interface{})
	tc := struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		Rememberme bool   `json:"string"`
	}{}
	tu := struct {
		Name string `json:"name"`
	}{}
	flashes := make(map[string]FlashMessage)

	utility.ReadJson(r, &tc)
	tc.Email = strings.Trim(tc.Email, " ")
	user, err := R.FindOneByEmail(tc.Email)

	// if user not found
	if err != nil {
		valid = false
	} else {
		fmt.Println(user.ActiveStatus)
//		if user.ActiveStatus { //check email activation
			// check if login allowed
			if user.LoginAllowed() {
				if valid = user.VerifyCredentials(tc.Email, tc.Password); valid == false {
					user.FailLogin()
				}
			} else {
				// login not allowed
				flashes["Error"] = FlashMessage{"warning", "You have failed 3 login attempts in the last 15 Minutes. Please wait 15 Minutes from now on and try again."}
			}
//		} else {
//			unactive = true
//			flashes["Error"] = FlashMessage{"warning", "Your Account is not yet Activated. Please Check your inbox or spam for the activation link. "}
//		}

	}
	data["valid"] = valid
	if valid {
		tu.Name = user.UserProfile.Name
		data["user"] = tu
		data["redirect"] = "/user/profile"
		user.Login(r.UserAgent(), r.RemoteAddr)
		session, _ := sessionStore.Get(r, "p")
		session.Values["user"] = user.Id.Hex()
		session.Values["usertype"] = user.UserType
		session.Values["profilestatus"] = user.ProfileStatus
		if tc.Rememberme {
			session.Options = &sessions.Options{
				Path:   "/",
				MaxAge: 86400 * 30 * 12,
			}
		}
		session.Save(r, w)
	} else if unactive {
		data["flashes"] = flashes
	} else {
		flashes["Error"] = FlashMessage{"danger", "Login not successful. Either a user with this email address doesn't exist or the email and password combination is wrong"}
		data["flashes"] = flashes
	}

	utility.WriteJson(w, data)
}**/

func AuthenticateHandler(w http.ResponseWriter, r *http.Request) {
	// get email + password
	valid := false
	unactive := false
	data := make(map[string]interface{})
	tc := struct {
		Email      string `json:"email"`
		Password   string `json:"password"`
		Rememberme bool   `json:"string"`
	}{}
	tu := struct {
		Name string `json:"name"`
	}{}
	flashes := make(map[string]FlashMessage)

	utility.ReadJson(r, &tc)
	tc.Email = strings.Trim(tc.Email, " ")
	user, err := R.FindOneByEmail(tc.Email)

	// if user not found
	if err != nil {
		valid = false
	} else {
		if user.ActiveStatus { //check email activation
			// check if login allowed
			if user.LoginAllowed() {
				if valid = user.VerifyCredentials(tc.Email, tc.Password); valid == false {
					user.FailLogin()
				}
			} else {
				// login not allowed
				flashes["Error"] = FlashMessage{"warning", "You have failed 3 login attempts in the last 15 Minutes. Please wait 15 Minutes from now on and try again."}
			}
		} else {
			unactive = true
			flashes["Error"] = FlashMessage{"warning", "Your Account is not yet Activated. Please Check your inbox or spam for the activation link.Activation email has been resent to your email id. "}
			user.ResendActEmail(user.Email)
			valid = false
		}
	}
	data["valid"] = valid
	if valid {
		user.CheckTechID()
		tu.Name = user.UserProfile.Name
		data["user"] = tu
		data["redirect"] = "/user/profile"
		user.Login(r.UserAgent(), r.RemoteAddr)
		session, _ := sessionStore.Get(r, "p")
		session.Values["user"] = user.Id.Hex()
		session.Values["usertype"] = user.UserType
		session.Values["profilestatus"] = user.ProfileStatus
		if tc.Rememberme {
			session.Options = &sessions.Options{
				Path:   "/",
				MaxAge: 86400 * 30 * 12,
			}
		}
		session.Save(r, w)
	} else if unactive {
		data["flashes"] = flashes
	} else {
		flashes["Error"] = FlashMessage{"danger", "Login not successful. Either a user with this email address doesn't exist or the email and password combination is wrong"}
		data["flashes"] = flashes
	}

	utility.WriteJson(w, data)
}

func PayHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	_, ok2 := session.Values["user"].(string)
	tc := struct {
		Id string `json:"id"`
	}{r.FormValue("id")}
	if ok2 && session.Values["usertype"] == "admin" {
		data := make(map[string]interface{})
		flashes := make(map[string]FlashMessage)
		user, err := R.FindOneByIdHex(tc.Id)
		// if user not found
		if err != nil {
		} else {
			if user.ActiveStatus {
				flashes["Error"] = FlashMessage{"warning", "The Profile payment is updated"}
			} else {
				user.UserProfile.PaymentStatus = "Completed"
				user.Update()
				flashes["success"] = FlashMessage{"success", "Your Profile has been successfully activated. You Can login using the login credentials"}
			}
		}
		data["flashes"] = flashes
		utility.WriteJson(w, data)
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
func ActHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	params := r.URL.Query()

	Uid := params["ui"][0]
	Ustring := params["us"][0]
	flashes := make(map[string]FlashMessage)

	Uid = strings.Trim(Uid, " ")
	user, err := R.FindOneByIdHex(Uid)
	// if user not found
	if err != nil {
	} else {
		if user.ActiveCode == Ustring {
			user.ActiveStatus = true
			user.Update()
			flashes["success"] = FlashMessage{"success", "The Account has been Activated.You Can login using the login credentials"}
		} else {
			flashes["Error"] = FlashMessage{"warning", "The Activation url is wrong. Please Contact Support"}
		}
	}
	data["flashes"] = flashes
	utility.WriteJson(w, data)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	session.Values["user"] = nil
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: -1,
	}
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	id, ok := session.Values["user"].(string)
	if ok {
		u, _ := R.FindOneByIdHex(id)
		data["profile"] = u.UserProfile
	} else {
		flashes["no_session"] = FlashMessage{"danger", "You are not logged in"}
		data["flashes"] = flashes
		data["redirect"] = "/login"
	}
	utility.WriteJson(w, data)
}

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	id, ok := session.Values["user"].(string)
	if ok {
		u, err := R.FindOneByIdHex(id)
		if err != nil {
			flashes["no_session"] = FlashMessage{"danger", "You are not logged in"}
			data["flashes"] = flashes
			data["redirect"] = "/login"
		} else {
			p := new(Profile)
			utility.ReadJson(r, &p)
			u.UserProfile.Number = p.Number
			u.UserProfile.College = p.College
			u.UserProfile.AlternateNumber = p.AlternateNumber
			u.UserProfile.Ambassador = p.Ambassador
			u.UserProfile.Sex = p.Sex
			u.UserProfile.Name = p.Name
			u.UserProfile.Branch = p.Branch
			u.UserProfile.Year = p.Year
			u.UserProfile.BookingID = p.BookingID
			u.Update()
			if u.CheckProfile() {
				u.ProfileStatus = true
			} else {
				u.ProfileStatus = false
			}
			u.Update()
			data["success"] = true
			flashes["profile_updated"] = FlashMessage{"success", "Your Profile has been updated"}
			data["flashes"] = flashes
		}
	}
	utility.WriteJson(w, data)
}

func ResetRequestHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	tc := struct {
		Email string `json:"email"`
	}{}
	utility.ReadJson(r, &tc)
	tc.Email = strings.Trim(tc.Email, " ")
	fmt.Println(tc.Email)
	user, err := R.FindOneByEmail(tc.Email)
	fmt.Println(user, err)
	if err != nil {
		flashes["Error"] = FlashMessage{"danger", "This user does not exist"}
	} else {
		user.CreateResetToken()
		flashes["success"] = FlashMessage{"success", "An Email has been sent to " + tc.Email + ". Please check your mailbox."}
	}
	data["flashes"] = flashes
	utility.WriteJson(w, data)
}

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	tc := struct {
		Token string `json:"token"`
	}{r.FormValue("token")}
	tc.Token = strings.Trim(tc.Token, " ")
	fmt.Println(tc.Token)
	if tc.Token != "" {
		user, err := R.FindOneByResetToken(tc.Token)
		if err != nil && user.ResetToken != tc.Token {
			flashes["user_not_found"] = FlashMessage{"danger", "Invalid token"}
		} else {
			s := user.ResetPassword()
			if s == true {
				flashes["success"] = FlashMessage{"success", "An Email with your new password has been sent to " + user.Email + ". Please check your mailbox."}
			} else {
				flashes["Error"] = FlashMessage{"danger", "The Token expired. Please request a new password reset token."}
			}
		}
		data["flashes"] = flashes
		utility.WriteJson(w, data)
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	id, ok := session.Values["user"].(string)
	if ok {
		u, err := R.FindOneByIdHex(id)
		if err != nil {
		} else {
			data["name"] = u.UserProfile.Name
		}
	}
	utility.WriteJson(w, data)
}

func AdminViewProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	_, ok := session.Values["user"].(string)
	tc := struct {
		Id string `json:"id"`
	}{r.FormValue("id")}
	if ok && session.Values["usertype"] == "admin" {
		u, err := R.FindOneByIdHex(tc.Id)
		if err != nil {
			flashes["User not Found"] = FlashMessage{"danger", "User seems to be not present in the database"}
			data["flashes"] = flashes
			utility.WriteJson(w, data)
		} else {
			data["success"] = u.UserProfile
			utility.WriteJson(w, data)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
func ManagerViewProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	flashes := make(map[string]FlashMessage)
	id, ok := session.Values["user"].(string)
	tc := struct {
		Id string `json:"id"`
	}{r.FormValue("id")}
	if ok {
		user, err1 := R.FindOneByIdHex(id)
		if err1 != nil && user.EventName != "" {
			http.Redirect(w, r, "/login", 302)
		} else {
			u, err := R.FindOneByIdHex(tc.Id)
			if err != nil {
				flashes["User not Found"] = FlashMessage{"danger", "User seems to be not present in the database"}
				data["flashes"] = flashes
				utility.WriteJson(w, data)
			} else {
				data["success"] = u.UserProfile
				utility.WriteJson(w, data)
			}
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
func ManagerViewTekProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "p")
	data := make(map[string]interface{})
	_, ok := session.Values["user"].(string)
	tc := struct {
		Id string `json:"id"`
	}{r.FormValue("id")}
	if ok {
		user, err1 := R.FindOneByTechID(tc.Id)
		if err1 != nil && user.EventName != "" {
			http.Redirect(w, r, "/login", 302)
		} else if user != nil {
			data["success"] = true
			data["user"] = user.UserProfile
			utility.WriteJson(w, data)
		} else {
			data["success"] = false
			utility.WriteJson(w, data)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
