package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	m "modul3/models"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT id, room_name FROM rooms"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.Room_name); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	roomID := r.URL.Query().Get("id")

	query := "SELECT r.id, r.room_name, a.id, a.username FROM participants p JOIN rooms r ON p.id_room = r.id JOIN accounts a ON p.id_account = a.id WHERE r.id = ?"

	rows, err := db.Query(query, roomID)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w)
		return
	}

	var detailParticipant m.DetailParticipants
	var detailParticipants []m.DetailParticipants
	for rows.Next() {
		if err := rows.Scan(&detailParticipant.Room.ID, &detailParticipant.Room.Room_name, &detailParticipant.Account.ID, &detailParticipant.Account.Username); err != nil {
			log.Println(err)
			sendErrorResponse(w)
			return
		} else {
			detailParticipants = append(detailParticipants, detailParticipant)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response m.DetailParticipantsResponse
	response.Status = 200
	response.Data = detailParticipants
	response.Message = "Success"
	json.NewEncoder(w).Encode(response)
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	// // Parse request body
	var participant m.Participants
	err := json.NewDecoder(r.Body).Decode(&participant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if room exists
	var room m.Rooms
	err = db.QueryRow("SELECT id, id_game FROM Rooms WHERE id = ?", participant.ID_room).Scan(&room.ID, &room.ID_game)
	if err != nil {
		http.Error(w, "Room not found", http.StatusNotFound)
		return
	}

	// Check if game exists
	var game m.Games
	err = db.QueryRow("SELECT max_player FROM Games WHERE id = ?", room.ID_game).Scan(&game.Max_player)
	if err != nil {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	// Check number of participants in the room
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Participants WHERE id_room = ?", participant.ID_room).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if room is full
	if count >= game.Max_player {
		response := m.Response{Message: "Failed to join room: Room is full"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Insert participant into the room
	_, err = db.Exec("INSERT INTO Participants (id_room, id_account) VALUES (?, ?)", participant.ID_room, participant.ID_account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := m.Response{Message: "Successfully joined room"}
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter) {
	var response m.ParticipantResponse
	response.Status = 200
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter) {
	var response m.ParticipantResponse
	response.Status = 400
	response.Message = "Failed"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
