package werewolf

type GameState struct {
	ChatId int64 `json:"ChatId"`

	Running bool `json:"running"`
}

func NewGameState(chatId int64) *GameState {
	return &GameState{
		ChatId:  chatId,
		Running: false,
	}
}
