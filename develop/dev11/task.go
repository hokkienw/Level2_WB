package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Event struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

var events []*Event

func (e *Event) toJSON() []byte {
	data, err := json.Marshal(e)
	if err != nil {
		log.Println("Error marshalling event:", err)
		return []byte(`{"error": "internal server error"}`)
	}
	return data
}

func parseEventParams(r *http.Request) (*Event, error) {
	var event Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		return nil, fmt.Errorf("invalid request body: %v", err)
	}
	return &event, nil
}

func createEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventParams(r)
	if err != nil {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}

	if err := createEvent(event); err != nil {
	    http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
	    return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "event created"}`))
}

func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	event, err := parseEventParams(r)
	if err != nil {
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}
	if err := updateEvent(event); err != nil {
	    http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
	    return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "event updated"}`))
}

func deleteEventHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "event deleted"}`))
}

func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "events for day"}`))
}

func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "events for week"}`))
}

func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "events for month"}`))
}

func main() {
	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	port := ":8080"
	log.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}


func createEvent(event *Event) error {
	if event.UserID == 0 {
		return fmt.Errorf("missing user ID")
	}
	if event.Title == "" {
		return fmt.Errorf("missing title")
	}
	if event.StartTime.IsZero() {
		return fmt.Errorf("missing start time")
	}
	if event.EndTime.IsZero() {
		return fmt.Errorf("missing end time")
	}
	if event.StartTime.After(event.EndTime) {
		return fmt.Errorf("start time cannot be after end time")
	}
	event.ID = len(events) + 1
	events = append(events, event)

	return nil
}

func updateEvent(event *Event) error {
	var found bool
	for i, e := range events {
		if e.ID == event.ID {
			found = true
			events[i] = event
			break
		}
	}
	if !found {
		return fmt.Errorf("event not found")
	}

	return nil
}
