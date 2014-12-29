package team

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type (
	Teams []Team

	Team struct {
		Id             bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy      string        `bson:"createdby" json:"createdby"`
		Name           string        `bson:"name" json:"name"`
		Active         bool          `bson:"active" json:"active"`
		InactiveSince  time.Time     `bson:"inactivesince" json:"inactivesince"`
		TeamProfile    *Profile      `bson:"Teamprofile,omitempty" json:"Teamprofile"`
		Members        string        `bson:"members" json:"members"`
		College        string        `bson:"college" json:"college"`
		RequestMod     bool          `bson:"requestmod" json:"requestmod"`
		PaymentAllowed bool          `bson:"payment" json:"payment"`
		PaymentStatus  bool          `bson:",omitempty" json:"paymentstatus"`
		Approved       string        `bson:",omitempty" json:"approved"`
		Event          string        `bson:",omitempty" json:"event"`
		Comments       string        `bson:",omitempty" json:"comments"`
	}

	Profile struct {
		Id      bson.ObjectId `bson:"_id,omitempty" json:"id"`
		City    string        `bson:",omitempty" json:"city"`
		Country string        `bson:",omitempty" json:"country"`
	}
	TeamRepository struct {
		Collection *mgo.Collection
	}
)

func (r TeamRepository) FindOneByName(name string) (result *Team, err error) {
	result = new(Team)
	err = r.Collection.Find(bson.M{"name": name}).One(result)
	return
}

func (r TeamRepository) FindAllByCreatedBy(createdby string) (teams Teams, err error) {
	err = r.Collection.Find(bson.M{"createdby": createdby}).All(&teams)
	return
}
func (r TeamRepository) FindAllByEvent(event string) (teams Teams, err error) {
	err = r.Collection.Find(bson.M{"event": event}).All(&teams)
	return
}
func (r TeamRepository) FindAllByCollege(college string) (teams Teams, err error) {
	err = r.Collection.Find(bson.M{"college": college}).All(&teams)
	return
}
func (r TeamRepository) FindAllByRequestMod() (teams Teams, err error) {
	err = r.Collection.Find(bson.M{"requestmod": true}).All(&teams)
	return
}

func (r TeamRepository) FindOneByResetToken(token string) (result *Team, err error) {
	result = new(Team)
	err = r.Collection.Find(bson.M{"resettoken": token}).One(result)
	return
}

func (r TeamRepository) FindOneByIdHex(id string) (result *Team, err error) {
	result = new(Team)
	err = r.Collection.FindId(bson.ObjectIdHex(id)).One(result)
	return
}

func (r TeamRepository) All() (teams Teams, err error) {
	err = r.Collection.Find(bson.M{}).All(&teams)
	return
}

func (r TeamRepository) CountByEmail(email string) (c int) {
	c, _ = r.Collection.Find(bson.M{"email": email}).Count()
	return
}
func (r TeamRepository) CountByCollege(college string) (c int) {
	c, _ = r.Collection.Find(bson.M{"college": college}).Count()
	return
}
func (r TeamRepository) CountByCollegenEvent(college string, event string) (c int) {
	c, _ = r.Collection.Find(bson.M{"college": college, "event": event}).Count()
	return
}

func (r TeamRepository) Update(team *Team) (err error) {
	err = r.Collection.UpdateId(team.Id, team)
	return
}

func (r TeamRepository) Delete(id string) (err error) {
	bid := bson.ObjectIdHex(id)
	err = r.Collection.RemoveId(bid)
	return
}

func (r TeamRepository) Create(team *Team) (err error) {
	if team.Id.Hex() == "" {
		team.Id = bson.NewObjectId()
	}
	_, err = r.Collection.UpsertId(team.Id, team)
	return
}
func (u *Team) Update() {
	if err := T.Update(u); err != nil {
		panic(err)
	}
}

func (u *Team) Add(name string, members string, event string, gender string, createdby string, college string) {
	p := new(Profile)
	p.Id = bson.NewObjectId()
	//p.Name = name
	//p.Surname = surname
	u.Name = name
	u.Members = members
	u.Event = event
	u.CreatedBy = createdby
	u.Approved = "Not Requested"
	u.College = college
	if err := T.Create(u); err != nil {
		panic(err)
	}
}

func (t *Team) Paymentallowed() bool {
	if t.PaymentAllowed {
		return true
	} else {
		return false
	}
}

func (t *Team) RequestModeration() bool {
	if t.RequestMod {
		return true
	} else {
		t.RequestMod = true
		t.Approved = "Pending"
		t.Update()
		return true
	}
}
