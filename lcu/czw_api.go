package lcu

type LolMatchHistoryMatchHistoryList struct {
	AccountId  int64    `json:"accountId"`
	Games      GameList `json:"games"`
	PlatformId string   `json:"platformId"`
}
