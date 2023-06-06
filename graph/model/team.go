package model

type Team struct {
	ID          string         `json:"id,omitempty" bson:"_id,omitempty"`
	Players     []*Player      `json:"players"`
	Name        string         `json:"name"`
	LeaderIndex *int           `json:"leaderIndex,omitempty"`
	ContestId   string         `json:"contestId"`
	ExtraInfo   map[string]any `json:"extraInfo,omitempty"`
	Status      TeamStatus     `json:"status"`
	PlayerOrder []int          `json:"playerOrder"`
}
