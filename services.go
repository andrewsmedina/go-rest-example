package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/ziutek/mymysql/godrv"
)

type Hello struct{}

type Service struct {
	Id   int
	Name string
}

var db *sql.DB

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {

	services := []Service{}

	rows, err := db.Query("SELECT id, name FROM services")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		services = append(services, Service{id, name})
	}

	b, _ := json.Marshal(services)
	fmt.Fprint(w, bytes.NewBuffer(b).String())

}

func main() {
	db, _ = sql.Open("mymysql", "project/andrews/123")
	defer db.Close()
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
