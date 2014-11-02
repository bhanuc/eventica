package user

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"fmt"
	"github.com/dalu/mail"
	"github.com/dchest/uniuri"
	"io"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"
	"encoding/base64"
)

type (
	Users []User

	User struct {
		Id             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email          string        `bson:"email" json:"email"`
		Password       string        `bson:"password" json:"password"`
		ResetToken     string        `bson:"resettoken,omitempty" json:"resettoken"`
		FBToken        string        `bson:"token,omitempty" json:"token"`
		FBId           string        `bson:"id,omitempty" json:"id"`
		ResetSent      time.Time     `bson:"resetsent,omitempty" json:"resetsent"`
		Created        time.Time     `bson:"created" json:"created"`
		LoginHistory   []LoginEntry  `bson:"loginHistory" json:"loginhistory"`
		LastSuccessful time.Time     `bson:"lastlogin" json:"lastlogin"`
		LastFailed     time.Time     `bson:"lastfailed" json:"lastfailed"`
		FailedAttempts int8          `bson:"failedattempts" json:"failedattempts"`
		Active         bool          `bson:"active" json:"active"`
		InactiveSince  time.Time     `bson:"inactivesince" json:"inactivesince"`
		UserProfile    *Profile      `bson:"userprofile,omitempty" json:"userprofile"`
		College        string        `bson:"college" json:"college"`
		UserType       string
		ProfileStatus  bool
		ActiveStatus   bool
		ActiveCode     string
	}

	LoginEntry struct {
		Timestamp time.Time `bson:"timestamp" json:"timestamp"`
		UserAgent string    `bson:"useragent" json:"useragent"`
		Ip        string    `bson:"ip" json:"ip"`
	}

	Profile struct {
		Id              bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name            string        `bson:"name" json:"name"`
		Number          string        `bson:"number" json:"number"`
		College         string        `bson:"college" json:"college"`
		Email           string        `bson:"email" json:"email"`
		AlternateNumber string        `bson:",omitempty" json:"alternatenumber"`
		Ambassador 		string        `bson:",omitempty" json:"ambassador"`
		Sex				string        `bson:"sex" json:"sex"`
		Branch			string        `bson:"branch" json:"branch"`
		ArrivalPNR		string        `bson:"apnr" json:"apnr"`
		ArrivalDate     string        `bson:"adate" json:"adate"`
		DeparturePNR	string        `bson:"dpnr" json:"dpnr"`
		DepartureDate	string        `bson:"ddate" json:"ddate"`
	}
	UserRepository struct {
		Collection *mgo.Collection
	}
)

func (r UserRepository) Create(user *User) (err error) {
	if user.Id.Hex() == "" {
		user.Id = bson.NewObjectId()
	}
	if user.Created.IsZero() {
		user.Created = time.Now()
	}
	_, err = r.Collection.UpsertId(user.Id, user)
	return
}

func (r UserRepository) FindOneByEmail(email string) (result *User, err error) {
	result = new(User)
	err = r.Collection.Find(bson.M{"email": email}).One(result)
	return
}
func (r UserRepository) FindOneByCollege(college string) (c int) {
	c, _ = r.Collection.Find(bson.M{"college": college}).Count()
	return
}

func (r UserRepository) FindOneByResetToken(token string) (result *User, err error) {
	result = new(User)
	err = r.Collection.Find(bson.M{"resettoken": token}).One(result)
	return
}

func (r UserRepository) FindOneByIdHex(id string) (result *User, err error) {
	result = new(User)
	err = r.Collection.FindId(bson.ObjectIdHex(id)).One(result)
	return
}

func (r UserRepository) All() (users Users, err error) {
	err = r.Collection.Find(bson.M{}).All(&users)
	return
}

func (r UserRepository) CountByEmail(email string) (c int) {
	c, _ = r.Collection.Find(bson.M{"email": email}).Count()
	return
}

func (r UserRepository) Update(user *User) (err error) {
	err = r.Collection.UpdateId(user.Id, user)
	return
}

func (r UserRepository) Delete(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

func (u *User) Update() {
	if err := R.Update(u); err != nil {
		panic(err)
	}
}

//THis just checks for profile completeness
func (u *User) CheckProfile() (stat bool) {
	stat = true
	if u.UserProfile.Name == "" {
		stat = false
	}
	if u.UserProfile.Number == "" {
		stat = false
	}
	if u.UserProfile.College == "" {
		stat = false
	}
	if u.UserProfile.Email == "" {
		stat = false
	}
	if u.UserProfile.AlternateNumber == "" {
		stat = false
	}
	return
}

func (u *User) CheckProfileStatus() bool {
	return u.ProfileStatus
}

func (u *User) Add(name, password, email, number, alternatenumber string) {
	//password := uniuri.New()
	//if Devmode {
	//fmt.Printf("New User: %s, %s\n", email, password)
	//}
	b := []byte(password)
	b, _ = bcrypt.GenerateFromPassword(b, 12)

	p := new(Profile)
	p.Id = bson.NewObjectId()
	//p.Name = name
	//p.Surname = surname
	u.Email = email
	p.Email = email
	p.Name = name
	p.Number = number
	p.AlternateNumber = alternatenumber

	u.UserType = "admin"
	u.ProfileStatus = false
	u.Password = strings.Trim(string(b[:]), "\x00")
	u.UserProfile = p
	u.ActiveStatus = false
	//Making a random string for checking email
	size := 32 // change the length of the generated random string here

   rb := make([]byte,size)
   _, err := rand.Read(rb)


   if err != nil {
      fmt.Println(err)
   }

   rs := base64.URLEncoding.EncodeToString(rb)

	u.ActiveCode = rs

	if err := R.Create(u); err != nil {
		panic(err)
	}
	uid := u.Id.String()
		slice := uid[13:37]
	   			fmt.Println(slice)

	body := "Hi ,\n\n"
	body += "welcome to " + Config.Host + ".\nYour account has been created.To Activate your account, please visit techkrit.org/user/activate?ui=" + slice + "&us=" + u.ActiveCode + " . Copy and paste the link in the browser to activate.\nYou login credentials are \n password:\n"
	body += password + "\n"
	body += "email address.\n"
	body += email+"\n"
	body += "Regards,\n\n"
	body += Config.Host + "team"

	m := mail.NewMail(Config.MailFrom, []string{email}, "Welcome to "+Config.Host, body)
	if err := m.Send(); err != nil {
		panic(err)
	}
}
func (u *User) FbAdd(name, email, id, token string) {
	p := new(Profile)
	p.Id = bson.NewObjectId()
	//p.Name = name
	//p.Surname = surname
	u.Email = email
	p.Email = email
	p.Name = name
	u.FBId = id
	u.FBToken = token
	u.UserType = "user"
	u.ProfileStatus = false
	u.UserProfile = p
	if err := R.Create(u); err != nil {
		panic(err)
	}
	/*body := "Hi ,\n\n"
	body += "welcome to " + Config.Host + ".\n"
	body += "Your account has been created. You may log in with the following password:\n"
	body += password + "\n"
	body += "and this email address.\n\n"
	body += "Regards,\n\n"
	body += Config.Host + "team"

	m := mail.NewMail(Config.MailFrom, []string{email}, "Welcome to "+Config.Host, body)
	if err := m.Send(); err != nil {
		panic(err)
	}*/
}

func (u *User) GenerateToken(l int) string {
	b := make([]byte, l)
	_, _ = io.ReadFull(rand.Reader, b)
	return fmt.Sprintf("%02x", b)
}

func (u *User) VerifyCredentials(email, password string) bool {
	if email != u.Email {
		return false
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return false
	}
	return true
}

func (u *User) Login(ua, ip string) {
	u.LastSuccessful = time.Now()
	u.Active = true
	l := LoginEntry{
		u.LastSuccessful,
		ua,
		strings.Split(ip, ":")[0],
	}
	if len(u.LoginHistory) > 10 {
		u.LoginHistory = u.LoginHistory[len(u.LoginHistory)-10 : len(u.LoginHistory)]
	}
	u.LoginHistory = append(u.LoginHistory, l)
	u.FailedAttempts = 0
	u.Update()
}

func (u *User) LoginAllowed() bool {
	if u.FailedAttempts >= 3 {
		if time.Since(u.LastFailed) >= time.Minute*15 {
			u.FailedAttempts = 0
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func (u *User) FailLogin() {
	u.FailedAttempts++
	if u.FailedAttempts >= 3 {
		u.LastFailed = time.Now()
	}
	u.Update()
}

func (u *User) CreateResetToken() {
	u.ResetToken = u.GenerateToken(42)

	body := "Hello " + u.UserProfile.Name + "  ,\n\n"
	body += "a password reset token for your " + Config.Host + " account has been created.\n"
	body += "Please click the following link to generate a new password\n"
	body += "http://" + Config.Host + "/#/user/resetpassword?token=" + u.ResetToken + "\n\n"
	body += "Regards,\n\n"
	body += Config.Host + " team"

	m := mail.NewMail(Config.MailFrom, []string{u.Email}, "Password Reset", body)
	if err := m.Send(); err != nil {
		panic(err)
	} else {
		u.ResetSent = time.Now()
		u.Update()
	}
}

func (u *User) ResetPassword() bool {
	password := uniuri.New()
	b := []byte(password)
	b, _ = bcrypt.GenerateFromPassword(b, 12)

	body := "Hello " + u.UserProfile.Name + ",\n\n"
	body += "Your password was reset.\n"
	body += "Please use the following password to log into your account:\n"
	body += password + "\n\n"
	body += "Regards,\n\n"
	body += Config.Host + " team"

	if time.Since(u.ResetSent) >= time.Hour*24 {
		u.ResetToken = ""
		u.ResetSent = time.Time{}
		u.Update()
		return false
	} else {
		u.Password = strings.Trim(string(b[:]), "\x00")
		if Devmode {
			fmt.Printf("User Password Changed: %s, %s\n", u.Email, password)
		}
		u.ResetToken = ""
		u.ResetSent = time.Time{}
		m := mail.NewMail(Config.MailFrom, []string{u.Email}, "Your new password for "+Config.Host, body)
		if err := m.Send(); err != nil {
			panic(err)
			return false
		} else {
			u.Update()
			return true
		}
	}
}
