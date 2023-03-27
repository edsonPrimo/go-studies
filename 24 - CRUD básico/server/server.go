package server

import (
	"crud/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if (err) != nil {
		w.Write([]byte("Internal error"))
		return
	}

	var user user
	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Json is not in a correct pattern"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Database connection failed"))
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Erro ao criar o statement"))
	}
	defer statement.Close()

	inserted, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("Error on statement exec"))
	}

	insertedId, erro := inserted.LastInsertId()
	if erro != nil {
		w.Write([]byte("Internal error"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User created succesfully. Id: %d ", insertedId)))

}

func FindUsers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Database connection failed"))
	}
	defer db.Close()

	query, err := db.Query("select * from usuarios")
	if err != nil {
		w.Write([]byte("Internal error"))
	}
	defer query.Close()

	var users []user
	for query.Next() {
		var user user

		if err := query.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Internal error"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("Internal error"))
		return
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 2)
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}

	query, err := db.Query("select * from usuarios where id = ?", id)
	if err != nil {
		w.Write([]byte("User doesn't exists"))
		return
	}

	var user user
	if query.Next() {
		if err := query.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Internal error"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(user); erro != nil {
		w.Write([]byte("Internal error"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 2)
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}

	var user user
	if err = json.Unmarshal(body, &user); err != nil {
		w.Write([]byte("Json is not in a correct pattern"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("Internal error"))
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, id); err != nil {
		w.Write([]byte("Erro on user update"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 2)
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}

	db, err := db.Connect()
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}
	defer db.Close()

	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("Internal error"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		w.Write([]byte("Error while deleting user"))
		return
	}

	w.WriteHeader(http.StatusOK)
}
