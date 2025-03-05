package models

import (
	"time"

	"example.com/rest-api/db"
	// "github.com/pelletier/go-toml/query"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Description string `binding:"required"`
	Location string `binding:"required"`
	DateTime time.Time `binding:"required"`
	UserID int64
}


func (e *Event) Save() error {
	// late : add it to database
	query:= `
	INSERT INTO  events(name, description, location,dateTime,user_id)
	VALUES (?,?,?,?,?)
	`
	stm, err := db.DB.Prepare(query)
	if err != nil {
    	return err
    }
	result, err := stm.Exec(e.Name,e.Description,e.Location,e.DateTime,e.UserID)
	defer stm.Close()

	if err != nil {
       return err
    }
	id,err := result.LastInsertId()
	e.ID = id
	return err
}


func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	var events []Event
	if err != nil {
       return nil,err
    }
		defer rows.Close()
	
		for rows.Next(){
			var e Event
            err := rows.Scan(&e.ID,&e.Name,&e.Description,&e.Location,&e.DateTime,&e.UserID)
            if err != nil {
                return nil,err
            }
            events = append(events,e)
		}
	return events,nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query,id)

	var event Event
	err := row.Scan(&event.ID,&event.Name,&event.Description,&event.Location,&event.DateTime,&event.UserID)
	if err != nil {
        return nil,err
    }
	return &event, nil
}


func (event Event) Update() error{
	query := `
    UPDATE events 
	SET name=?, description=?, location=?, dateTime=?
	WHERE id=?
    `
    stm, err := db.DB.Prepare(query)
    if err != nil {
        	return err
    }
    defer stm.Close()

    _,err = stm.Exec(event.Name,event.Description,event.Location,event.DateTime,event.ID)
    if err != nil {
		return err
   
    }
	return nil
}

func (event Event) Delete() error{
	query := `
    DELETE FROM events WHERE id=?
    `
    stm, err := db.DB.Prepare(query)
    if err != nil {
        	return err
    }
    defer stm.Close()

    _,err = stm.Exec(event.ID)
    if err != nil {
		return err
   
    }
	return nil
}

func (e Event) Register(user_id int64) error{

	query := `
	INSERT INTO registration(user_id, event_id) VALUES (?,?)
	`
	stm, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stm.Close()
	    _,err = stm.Exec(user_id,e.ID)
    if err != nil {
		return err
    }
	return nil

}

func (e Event) CancelRegistration(user_id int64) error{

	query := `
	DELETE FROM registration  WHERE user_id = ? AND event_id = ? 
	`
	stm, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stm.Close()
	    _,err = stm.Exec(user_id,e.ID)
    if err != nil {
		return err
    }
	return nil

}