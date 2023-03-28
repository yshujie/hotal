package dingtalk

type ActionCardType struct {
	Msgtype    string     `json:"msgtype"`
	At         at         `json:"at"`
	ActionCard actionCard `json:"action_card"`
}

type at struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool
}

type actionCard struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleUrl      string `json:"singleUrl"`
}
