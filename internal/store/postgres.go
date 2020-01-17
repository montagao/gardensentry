package store

import (
	"database/sql"
	"fmt"
	"log"

	"gardensentry.v1/gen/models"
)

const (
	// use companion process as proxy
	INSTANCE_CONNECTION_NAME = "35.196.103.90"
	DATABASE_NAME            = "gs"
	DATABASE_USER            = "postgres"
	PASSWORD                 = "fydp"
)

type EventStore struct {
	db *sql.DB
}

func New() (*EventStore, error) {
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
		INSTANCE_CONNECTION_NAME,
		DATABASE_NAME,
		DATABASE_USER,
		PASSWORD)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	eventStore := &EventStore{
		db: db,
	}
	log.Printf("Initialized postgres DB: %s", DATABASE_NAME)
	return eventStore, nil
}

func (s *EventStore) GetByID(id int64) (*models.Event, error) {
	rows, err := s.db.Query("select * from events where id = $1;", id)
	if rows == nil {
		log.Fatal("no rows found")
		return nil, nil
	}
	rows.Next()
	var (
		nullID      int
		description string
		occured     string
		eventType   string
		vidURL      string
	)
	err = rows.Scan(&nullID, &description, &occured, &eventType, &vidURL)
	if err != nil {
		return nil, err
	}
	e := &models.Event{
		Description: &description,
		ID:          id,
		Timestamp:   &occured,
		Type:        &eventType,
		VidURL:      vidURL,
	}
	return e, nil
}

func (s *EventStore) GetAll(limit int) ([]*models.Event, error) {
	log.Printf("Getting all events, limit: %d\n", limit)
	result := []*models.Event{}
	rows, err := s.db.Query("select * from events LIMIT $1 ", limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			id          int64
			description string
			occured     string
			eventType   string
			vidURL      string
		)
		err := rows.Scan(&id, &description, &occured, &eventType, &vidURL)
		if err != nil {
			return nil, err
		}
		e := &models.Event{
			Description: &description,
			ID:          id,
			Timestamp:   &occured,
			Type:        &eventType,
			VidURL:      vidURL,
		}
		result = append(result, e)
	}

	return result, nil
}

func (s *EventStore) Put(event *models.Event) error {
	stmt, err := s.db.Prepare("INSERT INTO events(description, occured, type, vidurl) VALUES( $1, $2, $3, $4 );")
	// Prepared statements take up server resources and should be closed after use.
	defer stmt.Close()

	if err != nil {
		return err
	}
	_, err = stmt.Exec(event.Description, event.Timestamp, event.Type, event.VidURL)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *EventStore) Update(event *models.Event, id int) error {
	// TODO: not a prioity
	return nil
}

func (s *EventStore) Delete(id int64) error {
	stmt, err := s.db.Prepare("DELETE from events where id = $1;")
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
