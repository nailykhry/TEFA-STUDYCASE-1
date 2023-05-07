package database

import (
	"fmt"     //O(1)
	"log"     //O(1)
	"os"      //O(1)
	"strconv" //O(1)

	"gopkg.in/mgo.v2" //O(1)
)

type Connection interface {
	Close()            //O(1)
	DB() *mgo.Database //O(1)
}

type conn struct {
	session *mgo.Session //O(1)
}

func NewConnection() Connection {
	var c conn                     //O(1)
	var err error                  //O(1)
	url := getURL()                //O(1)
	c.session, err = mgo.Dial(url) //O(1)
	if err != nil {                //O(1)
		log.Panicln(err.Error()) //O(1)
	}
	return &c //O(1)
}

func (c *conn) Close() {
	c.session.Close() //O(1)
}

func (c *conn) DB() *mgo.Database {
	return c.session.DB(os.Getenv("DATABASE_NAME")) //O(1)
}

func getURL() string {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT")) //O(1)
	if err != nil {                                       //O(1)
		log.Println("error on load db port from env:", err.Error()) //O(1)
		port = 27017                                                //O(1)
	}
	return fmt.Sprintf("mongodb://%s:%d/%s", //O(1)
		os.Getenv("DATABASE_HOST"),
		port,
		os.Getenv("DATABASE_NAME"))
}
