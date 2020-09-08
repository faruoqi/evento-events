package db

import (
	"github.com/faruoqi/evento/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	DBNAME = "evento"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDbLayer struct {
	session *mgo.Session
}

func (m *MongoDbLayer) FindEventByID(ID string) (model.Event, error) {

	session := m.getFreshSession()
	defer session.Close()
	event := model.Event{}
	err := session.DB(DBNAME).C(EVENTS).Find(bson.M{"id": ID}).One(&event)

	return event, err

}

func (m *MongoDbLayer) FindAllEvents() ([]model.Event, error) {

	s := m.getFreshSession()
	defer s.Close()
	var events []model.Event
	err := s.DB(DBNAME).C(EVENTS).Find(nil).All(&events)
	return events, err
}

func (m *MongoDbLayer) FindEventByName(eventName string) (model.Event, error) {

	s := m.getFreshSession()
	defer s.Close()
	event := model.Event{}
	err := s.DB(DBNAME).C(EVENTS).Find(bson.M{"name": eventName}).One(&event)
	return event, err

}

func (m *MongoDbLayer) AddEvent(event model.Event) error {

	session := m.getFreshSession()
	defer session.Close()
	return session.DB(DBNAME).C(EVENTS).Insert(event)
}

func (m *MongoDbLayer) getFreshSession() *mgo.Session {
	return m.session.Copy()
}

func NewMongoDbLayer(conn string) (*MongoDbLayer, error) {
	s, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}
	return &MongoDbLayer{session: s}, nil
}
