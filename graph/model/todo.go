package model

type Todo struct {
	ID            int    `json:"id"`
	Text          string `json:"text"`
	Done          bool   `json:"done"`
	UserID int `json:"user_id"`
}

func (Todo) IsNode()         {}
func (this Todo) GetID() int { return this.ID }
