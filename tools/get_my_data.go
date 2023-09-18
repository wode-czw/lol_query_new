package tools

type GameData struct {
	AccountId  int    `json:"accountId"`
	PlatformId string `json:"platformId"`
	Games      Games  `json:"games"`
}

type Games struct {
	GameBeginDate  string `json:"gameBeginDate"`
	GameCount      int    `json:"gameCount"`
	GameEndDate    string `json:"gameEndDate"`
	GameIndexBegin int    `json:"gameIndexBegin"`
	GameIndexEnd   int    `json:"gameIndexEnd"`
	Games          []Game `json:"games"`
}

type Game struct {
	GameCreation          int64                 `json:"gameCreation"`
	GameDuration          int                   `json:"gameDuration"`
	GameId                int64                 `json:"gameId"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities"`
	Participants          []Participant         `json:"participants"`
	Teams                 []Team                `json:"teams"`
	// ... other fields
}

type ParticipantIdentity struct {
	ParticipantId int    `json:"participantId"`
	Player        Player `json:"player"`
}

type Player struct {
	AccountId         int    `json:"accountId"`
	CurrentAccountId  int    `json:"currentAccountId"`
	CurrentPlatformId string `json:"currentPlatformId"`
	MatchHistoryUri   string `json:"matchHistoryUri"`
	PlatformId        string `json:"platformId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerId        int    `json:"summonerId"`
	SummonerName      string `json:"summonerName"`
}

type Participant struct {
	// ... other fields
	Stats    ParticipantStats `json:"stats"`
	TeamId   int              `json:"teamId"`
	Timeline Timeline         `json:"timeline"`
}

type ParticipantStats struct {
	Assists          int  `json:"assists"`
	ChampLevel       int  `json:"champLevel"`
	Deaths           int  `json:"deaths"`
	FirstBloodAssist bool `json:"firstBloodAssist"`
	FirstBloodKill   bool `json:"firstBloodKill"`
	// ... other fields
}

type Timeline struct {
	CreepsPerMinDeltas map[string]float64 `json:"creepsPerMinDeltas"`
	// ... other fields
}

type Team struct {
	BaronKills  int    `json:"baronKills"`
	DragonKills int    `json:"dragonKills"`
	Win         string `json:"win"`
	// ... other fields
}

func main() {

}
