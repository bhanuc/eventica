package user

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/dchest/uniuri"
	"io"
	"io/ioutil"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

type (
	Users []User

	EmailConfig struct {
		Username string
		Password string
		Host     string
		Port     int
	}

	User struct {
		Id             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Email          string        `bson:"email" json:"email"`
		Password       string        `bson:"password" json:"password"`
		ResetToken     string        `bson:"resettoken,omitempty" json:"resettoken"`
		FBToken        string        `bson:"token,omitempty" json:"token"`
		FBId           string        `bson:"FBId,omitempty" json:"FBId"`
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
		EventName      string
		Teams          string
		Tech_id        string `bson:"Tech_id" json:"Tech_id"`
		ProfileStatus  bool
		ActiveStatus   bool
		ActiveCode     string
		PaymentStatus  string
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
		Ambassador      string        `bson:",omitempty" json:"ambassador"`
		Sex             string        `bson:"sex" json:"sex"`
		Branch          string        `bson:"branch" json:"branch"`
		BookingID       string        `bson:"bookingid" json:"bookingid"`
		Year            string        `bson:"year" json:"year"`
		Tech_id         string
	}
	UserRepository struct {
		Collection *mgo.Collection
	}
)

func setup_tekid() (num int) {
	fmt.Println("started read")
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	length := len(b)
	fmt.Println(b[0], b[1], b[2], length)
	num = 0

	for i := 0; i < length; i++ {
		num = num*10 + int((float64(b[i] - '0')))
	}
	num = num + 1
	fmt.Println("in read", num)
	for j := 0; j < 1; {
		if R.FindCountByTechid(strconv.Itoa(num)) == 0 {
			j = 2
		} else {
			num = num + 2
		}
	}
	err = ioutil.WriteFile("input.txt", []byte(strconv.Itoa(num)), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("finished read")
	return num
}

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

func (r UserRepository) FindOneByTechID(id string) (result *User, err error) {
	result = new(User)
	fmt.Println(id)
	err = r.Collection.Find(bson.M{"Tech_id": id}).One(result)
	return
}
func (r UserRepository) FindCountByTechid(id string) (c int) {
	c, _ = r.Collection.Find(bson.M{"Tech_id": id}).Count()
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
	return
}

func (u *User) CheckTechID() {
	if u.Tech_id == "" && u.UserProfile.Tech_id == "" {
		tid := strconv.Itoa(setup_tekid())
		u.Tech_id = tid
		u.UserProfile.Tech_id = tid
		u.Update()
	} else if u.Tech_id != "" && u.UserProfile.Tech_id == "" {
		u.UserProfile.Tech_id = u.Tech_id
		u.Update()
	}
	return
}

func (u *User) CheckProfileStatus() bool {
	return u.ProfileStatus
}

func (u *User) Add(name, password, email, number, alternatenumber string) {
	tid := strconv.Itoa(setup_tekid())
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
	u.UserType = "user"
	u.ProfileStatus = false
	u.Password = strings.Trim(string(b[:]), "\x00")
	u.UserProfile = p

	u.Tech_id = tid
	p.Tech_id = tid
	u.ActiveStatus = false
	//Making a random string for checking email
	size := 32 // change the length of the generated random string here

	rb := make([]byte, size)
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

	// smtpServer := "smtp.163.com"
	// auth := smtp.PlainAuth(
	// 	"",
	// 	"fledna@163.com",
	// 	"password*******",
	// 	smtpServer,
	// )

	body := "Hi ,\n\n"
	body += "welcome to " + Config.Host + ".\nYour account has been created.To Activate your account, please visit http://portal.techkriti.org/user/activate?ui=" + slice + "&us=" + u.ActiveCode + " . Copy and paste the link in the browser to activate.\nYou login credentials are \nEmail-Address.\n"
	body += email + "\n"
	body += "Password:\n"
	body += password + "\n"
	body += "Regards,\n\n"
	body += Config.Host + " team"

	auth := smtp.PlainAuth(
		"",
		"webadmin@techkriti.org",
		"rememberyourpassword",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err2 := smtp.SendMail(
		"smtp.gmail.com:25",
		auth,
		"noreply@techkriti.org",
		[]string{email},
		[]byte(body),
	)
	if err2 != nil {
		log.Fatal(err)
	}
	// m := mail.NewMail(Config.MailFrom, []string{email}, "Welcome to "+Config.Host, body)
	// if err := m.Send(); err != nil {
	// 	fmt.Printf("The error is %s", err)
	// }
}

func (u *User) ResendActEmail(email string) {

	//rs := u.ActiveCode
	uid := u.Id.String()
	slice := uid[13:37]

	body := "Hi ,\n\n"
	body += "welcome to " + Config.Host + ".\nYour account is already created.To Activate your account, please visit http://portal.techkriti.org/user/activate?ui=" + slice + "&us=" + u.ActiveCode + " . Copy and paste the link in the browser to activate.\nYou login credentials are \nEmail-Address.\n"
	body += "Regards,\n\n"
	body += Config.Host + " team"

	// m := mail.NewMail(Config.MailFrom, []string{email}, "Welcome to "+Config.Host, body)
	// if err := m.Send(); err != nil {
	// 	fmt.Printf("The error is %s", err) techkriti@sharklasers.com
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// msg := gomail.NewMessage()
	// msg.SetHeader("From", "noreply@techkriti.org")
	// msg.SetHeader("To", mail)
	// msg.SetHeader("Subject", "Hello!")
	// msg.SetBody("text/html", body)

	// // Send the email to Bob, Cora and Dan
	// mailer := gomail.NewMailer("smtp.gmail.com", "webadmin@techkriti.org", "rememberyourpassword", 587)
	// if err := mailer.Send(msg); err != nil {
	// 	fmt.Println(err)
	// }

	// e := email.NewEmail()
	// e.From = "Techkriti <noreply@techkriti.org>"
	// e.To = []string{email}
	// e.Subject = "Activation Email"
	// e.Text = []byte("Text Body is, of course, supported!")
	// e.HTML = []byte(body)
	// e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "webadmin@techkriti.org", "rememberyourpassword", "smtp.gmail.com"))

	auth := smtp.PlainAuth(
		"",
		"webadmin@techkriti.org",
		"rememberyourpassword",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:25",
		auth,
		"noreply@techkriti.org",
		[]string{email},
		[]byte(body),
	)
	if err != nil {
		log.Fatal(err)
	}

}

func (u *User) AddManager(name, password, email, number, alternatenumber, event string) {
	tid := strconv.Itoa(setup_tekid())
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
	u.UserType = "user"
	u.ProfileStatus = false
	u.Password = strings.Trim(string(b[:]), "\x00")
	u.UserProfile = p

	u.Tech_id = tid
	p.Tech_id = tid
	u.ActiveStatus = true
	u.EventName = event

	u.ActiveCode = "manager"

	if err := R.Create(u); err != nil {
		panic(err)
	}

}

func (u *User) FbAdd(name, email, id, token string) {
	tid := strconv.Itoa(setup_tekid())
	p := new(Profile)
	p.Id = bson.NewObjectId()
	//p.Name = name
	//p.Surname = surname
	u.Email = email
	p.Email = email
	p.Name = name
	u.FBId = id
	u.FBToken = token
	u.Tech_id = tid
	p.Tech_id = tid
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

	// m := mail.NewMail(Config.MailFrom, []string{u.Email}, "Password Reset", body)
	// if err := m.Send(); err != nil {
	// 	panic(err)
	// } else {
	// 	u.ResetSent = time.Now()
	// 	u.Update()
	// }
	auth := smtp.PlainAuth(
		"",
		"webadmin@techkriti.org",
		"rememberyourpassword",
		"smtp.gmail.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.gmail.com:25",
		auth,
		"noreply@techkriti.org",
		[]string{u.Email},
		[]byte(body),
	)
	if err != nil {
		log.Fatal(err)
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
		// m := mail.NewMail(Config.MailFrom, []string{u.Email}, "Your new password for "+Config.Host, body)
		// if err := m.Send(); err != nil {
		// 	panic(err)
		// 	return false
		// } else {
		// 	u.Update()
		// 	return true
		// }

		auth := smtp.PlainAuth(
			"",
			"webadmin@techkriti.org",
			"rememberyourpassword",
			"smtp.gmail.com",
		)
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		err := smtp.SendMail(
			"smtp.gmail.com:25",
			auth,
			"noreply@techkriti.org",
			[]string{u.Email},
			[]byte(body),
		)
		if err != nil {
			log.Fatal(err)
			return false
		} else {
			u.Update()
			return true
		}

	}
}
