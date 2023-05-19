package model

type Player struct {
	ID               string                 `json:"id" bson:"_id"`
	Name             string                 `json:"name"`
	PlatformInfos    []*PlatformInfo        `json:"platformInfos"`
	ExtraInfo        map[string]interface{} `json:"extraInfo,omitempty"`
	ProfessionalCert *string                `json:"professionalCert,omitempty"`
	ContestId        *string                `json:"contestId,omitempty"`
}
