package werewolf

type GameState struct {
	ChatId int64 `json:"ChatId"`

	Running       bool  `json:"Running"`
	CurrentPlayer int64 `json:"CurrentPlayer"`
}

func NewGameState(chatId int64) *GameState {
	return &GameState{
		ChatId:        chatId,
		Running:       false,
		CurrentPlayer: 0,
	}
}

func SetCurrentPlayerState(chatId int64, currentPlayer int64) *GameState {
	return &GameState{
		ChatId:        chatId,
		Running:       false,
		CurrentPlayer: currentPlayer,
	}
}

type PlayerState struct {
	ChatId int64 `json:"ChatId"`
	UserId int64 `json:"UserId"`
}

func NewPlayerState(chatId int64, userId int64) *PlayerState {
	return &PlayerState{
		ChatId: chatId,
		UserId: userId,
	}
}

type RoomState struct {
	ChatId int64   `json:"ChatId"`
	UserId []int64 `json:"UserId"`
}

func NewRoomState(chatId int64, userId []int64) *RoomState {
	return &RoomState{
		ChatId: chatId,
		UserId: userId,
	}
}
