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

func (h Hello) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request) {

	services := []Service{}

    db, err := sql.Open("mymysql", "project/andrews/123")
	if err != nil {
		panic(err)
	}

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

	err = db.Close()
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(services)
	fmt.Fprint(w, bytes.NewBuffer(b).String())

}

func main() {
	var h Hello
	http.ListenAndServe("localhost:4000", h)
}
