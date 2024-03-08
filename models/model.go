package models

type Accounts struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Games struct {
	ID         int    `json:"id"`
	Name       string `json:"games_name"`
	Max_player int    `json:"max_player"`
}

type Rooms struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	ID_game   int    `json:"id_game"`
}

type Participants struct {
	ID         int `json:"id"`
	ID_room    int `json:"id_room"`
	ID_account int `json:"id_accounts"`
}

type DetailParticipants struct {
	ID      int      `json:"id_participant"`
	Room    Rooms    `json:"room"`
	Account Accounts `json:"accounts"`
}

type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

type DetailParticipantsResponse struct {
	Status  int                  `json:"status"`
	Message string               `json:"message"`
	Data    []DetailParticipants `json:"data"`
}

type ParticipantResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    Participants `json:"data"`
}

type ParticipantsResponse struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    []Participants `json:"data"`
}

type Response struct {
	Message string `json:"message"`
}
