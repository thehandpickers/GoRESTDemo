package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Created     string `json:"created"`
	Language    string `json:"language"`
	ShortDesc   string `json:"short_description"`
	LastUpdated string `json:"last_updated"`
	Likes       int    `json:"likes"`
	Shares      int    `json:"shares"`
}

type Projects struct {
	ProjectsList []Project `json:"projects"`
}

func handler(rw http.ResponseWriter, req *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3307)/projectdb")

	rows, err := db.Query("SELECT * FROM projects")

	if err != nil {
		fmt.Println("Error")
	}

	Response := Projects{}
	for rows.Next() {
		project := Project{}
		rows.Scan(&project.Id, &project.Name, &project.LastUpdated, &project.Language,
			&project.ShortDesc, &project.Likes, &project.Shares, &project.Created)
		Response.ProjectsList = append(Response.ProjectsList, project)
	}

	output, err := json.Marshal(Response)

	if err != nil {
		fmt.Println("Error")
	}
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(rw, string(output))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4040", nil)
}
