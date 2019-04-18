package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

// Session is a Cassandra session
var Session *gocql.Session

func init() {

	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "127.0.0.1"
	}

	portString := os.Getenv("DB_PORT")

	var port int
	if portString != "" {
		port, _ = strconv.Atoi(portString)
	}
	var err error

	cluster := gocql.NewCluster(host)
	cluster.ProtoVersion = 4
	cluster.Keyspace = "go_do"

	if port != 0 {
		cluster.Port = port
	}

	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true

	fmt.Println("Connecting now")
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Printf("error creating cluster: %v", err)
		return
	}

	fmt.Println("cassandra init done")

	keySpaceMeta, _ := Session.KeyspaceMetadata("go_do")

	if _, exists := keySpaceMeta.Tables["task"]; exists != true {
		Session.Query("CREATE TABLE task (id UUID, title text, description text, userId text, createdDate timestamp, completedDate timestamp, dailyDo Boolean, PRIMARY KEY(id))").Exec()
		Session.Query("create index UserIdIndex on task(userId)").Exec()
		Session.Query("create index DailyDoIndex on task(dailyDo)").Exec()
		Session.Query("create index CompletedIndex on task(completedDate)").Exec()
	}
}
