package database

import (
	"database/sql"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

const pathDbFile = "./internal/adapters/database/database-flights.db"

func InitDb() (*sql.DB, error) {
	_, err := os.Stat(pathDbFile)
	if os.IsNotExist(err) {
		if err = createFile(); err != nil {
			return nil, err
		}
	}
	var d Database
	if err = d.openDbFile(pathDbFile); err != nil {
		return nil, err
	}
	// d.createTable()
	d.DropTables()
	if err = d.CreateTable(); err != nil {
		return nil, err
	}
	d.InsertDummyData()

	return d.db, nil
}

func createFile() error {
	file, err := os.Create(pathDbFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func (d *Database) openDbFile(file string) error {
	var err error
	d.db, err = sql.Open("sqlite3", file)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) CreateTable() error {
	dir, err := ioutil.ReadDir("./internal/adapters/database/schema/up/")
	if err != nil {
		return err
	}
	for _, v := range dir {
		body, err := ioutil.ReadFile("./internal/adapters/database/schema/up/" + v.Name())
		if err != nil {
			return err
		}
		_, err = d.db.Exec(string(body))
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) DropTables() error {
	dir, err := ioutil.ReadDir("./internal/adapters/database/schema/drop/")
	if err != nil {
		return err
	}

	for _, v := range dir {
		body, err := ioutil.ReadFile("./internal/adapters/database/schema/drop/" + v.Name())
		if err != nil {
			return err
		}
		_, err = d.db.Exec(string(body))
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) InsertDummyData() error {
	_, err := d.db.Exec(`INSERT INTO bus_flights (departure_city, arrival_city,distance, travel_time)
		VALUES('Almaty', 'Astana', '1200 km', '11:30'),
			('Astana', 'Almaty', '1200 km', '11:30'),
			('Semey', 'Astana', '800 km', '8:00'),
			('Semey', 'Almaty', '1000 km', '10:00'),
			('Kokshetau', 'Karaganda', '1500 km', '14:00'),
			('Aktobe', 'Aktau', '600 km', '6:00');`)
	if err != nil {
		return err
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
		return err
	}

	_, err = d.db.Exec(`INSERT INTO time_flights (id, start_flight)
		VALUES(1, '2022-04-29'),
		       (2, '2022-04-30'),
		       (3, '2022-04-28'),
		       (4, '2022-04-30'),
		       (5, '2022-04-30'),
		       (6, '2022-04-30');`)
	if err != nil {
		return err
	}
	return nil
}
