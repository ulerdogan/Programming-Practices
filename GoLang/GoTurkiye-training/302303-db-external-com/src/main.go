package main

import (
	"database/sql"
	"db-ext-com/brokers"
	"db-ext-com/cache"
	"db-ext-com/entity"
	"db-ext-com/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

var (
	db    *sql.DB
	dbErr error
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "gotr-pass"
	dbName := "godb"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password = %s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, dbErr = sql.Open("postgres", psqlInfo)
	if dbErr != nil {
		panic(dbErr)
	}

	cityRepo := repository.NewRepo(db)
	rabbitmq := brokers.NewRabbitMQ()
	redisCache := cache.NewRedis()

	http.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {

			if r.URL.Query().Has("id") {
				cityIdStr := r.URL.Query().Get("id")
				cityId, _ := strconv.Atoi(cityIdStr)
				city := cityRepo.GetById(cityId)
				if city == nil {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				cityBytes, _ := json.Marshal(city)
				w.Write(cityBytes)
				return
			}

			cityCache := redisCache.Get()
			if cityCache != nil {
				fmt.Println("in cache")
				w.Write(cityCache)
				return
			}

			cl := cityRepo.List()
			//json.NewEncoder(w).Encode(cl)

			clByte, _ := json.Marshal(cl)
			go func(data []byte) {
				fmt.Println("to cache")
				redisCache.Put(data)
			}(clByte)
			w.Write(clByte)
		} else if r.Method == http.MethodPost {
			var city entity.City
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			if err := json.Unmarshal(bodyBytes, &city); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			cityRepo.Insert(city)
			rabbitmq.Publish([]byte("city created with name: " + city.Name))
			w.WriteHeader(http.StatusCreated)
		} else {
			http.Error(w, "unsupported http method", http.StatusMethodNotAllowed)
			return
		}

	})

	func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println(err)
		}
	}()

}

// func selectWithPreparedStatement(name string) {
// 	stmt, err := db.Prepare("select id, name, code from cities where name = $1")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	} else {
// 		var city City
// 		err := stmt.QueryRow(name).Scan(&city.Id, &city.Name, &city.Code)
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println(city)
// 		}
// 	}
// }
