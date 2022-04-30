package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func CheckDB() *sql.DB {
	_, err := os.Stat("./internal/adapters/database/database-flights.db")
	if os.IsNotExist(err) {
		createFile()
	}
	var d Database
	d.open("./internal/adapters/database/database-flights.db")
	//d.createTable()
	d.DropTables()
	if err = d.CreateTable(); err != nil {
		log.Println(err)
	}
	d.Insert()
	return d.db
}

func createFile() {
	file, err := os.Create("./internal/adapters/database/database-flights.db")
	if err != nil {
		log.Fatalf("file doesn't create %v", err)
	}
	defer file.Close()
}

func (d *Database) open(file string) {
	var err error
	d.db, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("this error is in dbase/open() %v", err)
	}
}

func (d *Database) CreateTable() error {
	dir, err := ioutil.ReadDir("./internal/adapters/database/schema/up/")
	if err != nil {
		log.Println(err)
		return err
	}
	for _, v := range dir {
		body, err := ioutil.ReadFile("./internal/adapters/database/schema/up/" + v.Name())
		if err != nil {
			log.Println(err)
		}
		_, err = d.db.Exec(string(body))
		//log.Println(string(body))
		if err != nil {
			log.Println("---->", v.Name(), err)
		}
	}
	return nil
}

func (d *Database) DropTables() {
	dir, err := ioutil.ReadDir("./internal/adapters/database/schema/drop/")
	if err != nil {
		log.Println(err)
	}

	for _, v := range dir {
		body, err := ioutil.ReadFile("./internal/adapters/database/schema/drop/" + v.Name())
		if err != nil {
			log.Println(err)
		}

		_, err = d.db.Exec(string(body))
		if err != nil {
			log.Println(err)
		}
	}
}

func (d *Database) Insert() {
	_, err := d.db.Exec(`INSERT INTO bus_flights (departure_city, arrival_city,distance, travel_time)
		VALUES('Almaty', 'Astana', '1200 km', '11:30'),
			('Astana', 'Almaty', '1200 km', '11:30'),
			('Semey', 'Astana', '800 km', '8:00'),
			('Semey', 'Almaty', '1000 km', '10:00'),
			('Kokshetau', 'Karaganda', '1500 km', '14:00'),
			('Aktobe', 'Aktau', '600 km', '6:00');`)
	if err != nil {
		log.Printf("%v\n", err)
	}

	_, err = d.db.Exec(`INSERT INTO intermediate_bus_station (id, station,number_station, arrival_time, departure_time)
		VALUES(1, 'Balkhash', 1, '10-00', '10-30'),
		       (1, 'Karaganda', 2, '15-30', '16-30'),
		       (2, 'Karaganda', 1, '14-00', '15-00'),
		       (2, 'Balkhash', 2, '18-30', '19-00'),
		       (3, 'Pavlodar', 1, '14-00', '14-30'),
		       (4, 'Ayaguz', 1, '15-30', '16-30'),
		       (5, 'Astana', 1, '14-00', '15-00'),
		       (5, 'Balkhash', 2, '18-30', '19-00'),
		       (5, 'Balkhash', 3, '20-30', '21-00');`)
	if err != nil {
		log.Printf("%v\n", err)
	}

	_, err = d.db.Exec(`INSERT INTO time_flights (id, start_flight)
		VALUES(1, '2022-04-29'),
		       (2, '2022-04-30'),
		       (3, '2022-04-28'),
		       (4, '2022-04-30'),
		       (5, '2022-04-30'),
		       (6, '2022-04-30');`)
	if err != nil {
		log.Printf("-----%v\n", err)
	}
}
