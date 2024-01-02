package models

import (
	_ "github.com/go-sql-driver/mysql"
	"mastengkorak.com/gowebapidb_people/controllers"

	"fmt"
)

func ListPeoplesHandler() []Peoples {
	Db := controllers.DB
	results, err := Db.Query("SELECT * FROM peoples")
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil

	}

	peoples1 := []Peoples{}
	for results.Next() {
		var ppl Peoples
		err = results.Scan(&ppl.ID, &ppl.First_name, &ppl.Last_name, &ppl.Birth_year, &ppl.Email)

		if err != nil {
			panic(err.Error())
		}
		peoples1 = append(peoples1, ppl)
	}
	return peoples1

}

func CreatePeoplesHandler(peoples Peoples) {
	Db := controllers.DB
	insert, err := Db.Query(
		"INSERT INTO peoples (id,first_name,last_name, birth_year, email) VALUES (?,?,?,?,?)",
		peoples.ID, peoples.First_name, peoples.Last_name, peoples.Birth_year, peoples.Email)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer insert.Close()
}

func UpdatePeoplesHandler(peoples Peoples) {
	Db := controllers.DB

	update, err := Db.Query(
		"UPDATE peoples SET first_name = ?, last_name = ?, birth_year = ?, email = ? WHERE id = ?",
		peoples.First_name, peoples.Last_name, peoples.Birth_year, peoples.Email, peoples.ID)

	if err != nil {
		fmt.Println("Err", err.Error())
	}
	defer update.Close()
}

func DeletePeoplesHandler(id string) {
	Db := controllers.DB

	delete, err := Db.Query("DELETE FROM peoples WHERE id=?", id)
	if err != nil {
		fmt.Println("Err", err.Error())
	} else {
		fmt.Printf("People with id : %d deleted\n", id)
	}
	defer delete.Close()
}

func GetPeoplesbyID(id string) *Peoples {
	Db := controllers.DB

	ppl := &Peoples{}
	results, err := Db.Query("SELECT * FROM peoples where id=?", id)
	if err != nil {
		fmt.Println("Err", err.Error())
		return nil
	}

	if results.Next() {
		err = results.Scan(&ppl.ID, &ppl.First_name, &ppl.Last_name, &ppl.Birth_year, &ppl.Email)
		if err != nil {
			return nil
		}
	} else {
		return nil
	}
	return ppl
}
