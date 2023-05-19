package model

type Contest struct {
	ID            string        `json:"id" bson:"_id"`
	Name          string        `json:"name"`
	Status        ContestStatus `json:"status"`
	MahjongType   MahjongType   `json:"mahjongType"`
	MaxTeamMember int           `json:"maxTeamMember"`
	MinTeamMember int           `json:"minTeamMember"`
	IsIndividual  bool          `json:"isIndividual"`
}
