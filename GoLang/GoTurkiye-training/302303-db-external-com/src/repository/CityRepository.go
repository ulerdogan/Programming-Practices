package repository

import (
	"database/sql"
	"db-ext-com/entity"
	"fmt"
)

type CityRepo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *CityRepo {
	return &CityRepo{
		db: db,
	}
}

func (repo CityRepo) Insert(city entity.City) {
	stmt, err := repo.db.Prepare("insert into cities(name, code) values($1, $2)")
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := stmt.Exec(city.Name, city.Code)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(r.RowsAffected())
	}

}

func (repo CityRepo) List() []entity.City {
	var cityList []entity.City
	rows, err := repo.db.Query("select id, name, code from cities")
	if err != nil {
		fmt.Println(err)
		return cityList
	} else {
		for rows.Next() {
			var city entity.City
			err := rows.Scan(&city.Id, &city.Name, &city.Code)
			if err != nil {
				fmt.Println(err)
			} else {
				cityList = append(cityList, city)
			}
		}
		rows.Close()
		return cityList
	}
}

func (repo CityRepo) GetById(id int) *entity.City {
	var city entity.City
	fSql := fmt.Sprintf("select id, name, code from cities where id = %v", id) // vulnerability
	err := repo.db.QueryRow(fSql).Scan(&city.Id, &city.Name, &city.Code)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &city
}
