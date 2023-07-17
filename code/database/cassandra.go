package database

import (
	"log"
	"github.com/gocql/gocql"
)

var session *gocql.Session

func initCassandra() error {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace= ""
	cluster.Consistency := gocql.Quorum 


	var err error

	session, err = cluster.CreateSession()

	if(err!=nil){
		return err
	}

	log.Println("Cassandra Initialized")
	return nil

}


func CloseCassandra(){
	session.Close()
	log.Println("Cassandra Session Suspended")
}