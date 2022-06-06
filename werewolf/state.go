package werewolf

type GameState struct {
	ChatId int `json:"ChatId"`

	Running bool `json:"running"`
}

func NewGameState(chatId int) *GameState {
	return &GameState{
		ChatId:  chatId,
		Running: false,
	}
}
