package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"wedding-service/interfaces"
)

type DB struct {
	*sql.DB
}

func ConnectToDB() (*DB, error) {
	db, err := sql.Open("postgres", "user=admin dbname=postgres password=123456 sslmode=disable")
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Printf("failed to ping database: %v", err)
		return nil, err
	}
	log.Println("connected to database")
	return &DB{db}, nil
}

func (db *DB) InsertWedding(wedding interfaces.Wedding) error {
	// Insert persons
	coupleID, err := db.InsertCouple(wedding.Couple)
	if err != nil {
		return err
	}
	// Insert wedding
	_, err = db.Exec("INSERT INTO weddings (couple_id, date, location, guest_count) VALUES ($1, $2, $3, $4)",
		coupleID, wedding.Date, wedding.Location, wedding.GuestCount)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) InsertCouple(couple interfaces.Couple) (int64, error) {
	// [todo] make concurrent
	// [todo] int64?
	person1ID, err := db.InsertPerson(couple.Person1)
	if err != nil {
		return 0, err
	}
	person2ID, err := db.InsertPerson(couple.Person2)
	if err != nil {
		return 0, err
	}

	var coupleID int64
	err = db.QueryRow(`
        INSERT INTO couples (person_1_id, person_2_id)
        VALUES ($1, $2)
        RETURNING id
    `, person1ID, person2ID).Scan(&coupleID)
	if err != nil {
		return 0, err
	}

	return coupleID, nil
}

func (db *DB) InsertPerson(person interfaces.Person) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO persons (first_name, last_name, gender) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		log.Printf("failed to add person: %v, %v", person, err)
		return 0, err
	}
	defer stmt.Close()

	var id int64
	err = stmt.QueryRow(person.FirstName, person.LastName, person.Gender).Scan(&id)
	if err != nil {
		log.Printf("failed to find id of added person: %v, %v", person, err)
		return 0, err
	}

	return id, nil
}
