package main

import (
	"fmt"
	"time"
	"log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Name 			string			`json: "username"`
	Password 		string			`json: "password"`
	Age 			int				`json: "age"`
	LastUpdated     time.Time
}


// GetProfiles - Returns all the profile in the Profiles Collection
func GetProfiles() []Profile {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return nil
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ProfileService").C("Profiles")
	var profiles []Profile
	err = c.Find(bson.M{}).All(&profiles)

	return profiles
}


// ShowProfile - Returns the profile in the Profiles Collection with name equal to the id parameter (id == name)
func ShowProfile(id string) Profile {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return Profile{}
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ProfileService").C("Profiles")
	profile := Profile{}
	err = c.Find(bson.M{"name": id}).One(&profile)

	return profile
}


// DeleteProfile - Deletes the profile in the Profiles Collection with name equal to the id parameter (id == name)
func DeleteProfile(id string) bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ProfileService").C("Profiles")
	err = c.RemoveId(id)

	return true
}


// CreateOrUpdateProfile - Creates or Updates (Upsert) the profile in the Profiles Collection with id parameter
func (p *Profile) CreateOrUpdateProfile() bool {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Println("Could not connect to mongo: ", err.Error())
		return false
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("ProfileService").C("Profiles")
	_ , err = c.UpsertId( p.Name, p )
	if err != nil {
		log.Println("Error creating Profile: ", err.Error())
		return false
	}
	return true
}

func main(){
	profile := Profile{Name: "Muhamad Enrinal", Password : "qwerty123", Age : 20, LastUpdated: time.Now()}
	err := profile.CreateOrUpdateProfile()
	if err != true {
		fmt.Println("Error")
	}
}