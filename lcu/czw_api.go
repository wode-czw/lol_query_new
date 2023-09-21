package lcu

//旧的命令无法查到那些结构体了，我只能用新的命令了，该命令返回的application/json，这个作者并没有写出来，我就只能自己写了
type LolMatchHistoryMatchHistoryList struct {
	AccountId  int64    `json:"accountId"`
	Games      GameList `json:"games"`
	PlatformId string   `json:"platformId"`
}
