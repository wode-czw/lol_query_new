package lcu

//旧的命令无法查到那些结构体了，我只能用新的命令了，该命令返回的application/json，这个作者并没有写出来，我就只能自己写了
type LolMatchHistoryMatchHistoryList struct {
	AccountId  int64    `json:"accountId"`
	Games      GameList `json:"games"`
	PlatformId string   `json:"platformId"`
}

type LolRankedRankedStats struct {
	EarnedRegaliaRewardIds                []interface{} `json:"earnedRegaliaRewardIds"`
	HighestCurrentSeasonReachedTierSR     string        `json:"highestCurrentSeasonReachedTierSR"`
	HighestPreviousSeasonAchievedDivision string        `json:"highestPreviousSeasonAchievedDivision"`
	HighestPreviousSeasonAchievedTier     string        `json:"highestPreviousSeasonAchievedTier"`
	HighestPreviousSeasonEndDivision      string        `json:"highestPreviousSeasonEndDivision"`
	HighestPreviousSeasonEndTier          string        `json:"highestPreviousSeasonEndTier"`
	HighestRankedEntry                    struct {
		Division                       string      `json:"division"`
		HighestDivision                string      `json:"highestDivision"`
		HighestTier                    string      `json:"highestTier"`
		IsProvisional                  bool        `json:"isProvisional"`
		LeaguePoints                   int         `json:"leaguePoints"`
		Losses                         int         `json:"losses"`
		MiniSeriesProgress             string      `json:"miniSeriesProgress"`
		PreviousSeasonAchievedDivision string      `json:"previousSeasonAchievedDivision"`
		PreviousSeasonAchievedTier     string      `json:"previousSeasonAchievedTier"`
		PreviousSeasonEndDivision      string      `json:"previousSeasonEndDivision"`
		PreviousSeasonEndTier          string      `json:"previousSeasonEndTier"`
		ProvisionalGameThreshold       int         `json:"provisionalGameThreshold"`
		ProvisionalGamesRemaining      int         `json:"provisionalGamesRemaining"`
		QueueType                      string      `json:"queueType"`
		RatedRating                    int         `json:"ratedRating"`
		RatedTier                      string      `json:"ratedTier"`
		Tier                           string      `json:"tier"`
		Warnings                       interface{} `json:"warnings"`
		Wins                           int         `json:"wins"`
	} `json:"highestRankedEntry"`
	HighestRankedEntrySR struct {
		Division                       string      `json:"division"`
		HighestDivision                string      `json:"highestDivision"`
		HighestTier                    string      `json:"highestTier"`
		IsProvisional                  bool        `json:"isProvisional"`
		LeaguePoints                   int         `json:"leaguePoints"`
		Losses                         int         `json:"losses"`
		MiniSeriesProgress             string      `json:"miniSeriesProgress"`
		PreviousSeasonAchievedDivision string      `json:"previousSeasonAchievedDivision"`
		PreviousSeasonAchievedTier     string      `json:"previousSeasonAchievedTier"`
		PreviousSeasonEndDivision      string      `json:"previousSeasonEndDivision"`
		PreviousSeasonEndTier          string      `json:"previousSeasonEndTier"`
		ProvisionalGameThreshold       int         `json:"provisionalGameThreshold"`
		ProvisionalGamesRemaining      int         `json:"provisionalGamesRemaining"`
		QueueType                      string      `json:"queueType"`
		RatedRating                    int         `json:"ratedRating"`
		RatedTier                      string      `json:"ratedTier"`
		Tier                           string      `json:"tier"`
		Warnings                       interface{} `json:"warnings"`
		Wins                           int         `json:"wins"`
	} `json:"highestRankedEntrySR"`
	QueueMap           map[string]LolRankedQueueEntry `json:"queueMap"`
	Queues             []LolRankedQueueEntry          `json:"queues"`
	RankedRegaliaLevel int                            `json:"rankedRegaliaLevel"`
	Seasons            map[string]LolRankedSeason     `json:"seasons"`
	SplitsProgress     map[string]int                 `json:"splitsProgress"`
}

type LolRankedQueueEntry struct {
	Division                       string      `json:"division"`
	HighestDivision                string      `json:"highestDivision"`
	HighestTier                    string      `json:"highestTier"`
	IsProvisional                  bool        `json:"isProvisional"`
	LeaguePoints                   int         `json:"leaguePoints"`
	Losses                         int         `json:"losses"`
	MiniSeriesProgress             string      `json:"miniSeriesProgress"`
	PreviousSeasonAchievedDivision string      `json:"previousSeasonAchievedDivision"`
	PreviousSeasonAchievedTier     string      `json:"previousSeasonAchievedTier"`
	PreviousSeasonEndDivision      string      `json:"previousSeasonEndDivision"`
	PreviousSeasonEndTier          string      `json:"previousSeasonEndTier"`
	ProvisionalGameThreshold       int         `json:"provisionalGameThreshold"`
	ProvisionalGamesRemaining      int         `json:"provisionalGamesRemaining"`
	QueueType                      string      `json:"queueType"`
	RatedRating                    int         `json:"ratedRating"`
	RatedTier                      string      `json:"ratedTier"`
	Tier                           string      `json:"tier"`
	Warnings                       interface{} `json:"warnings"`
	Wins                           int         `json:"wins"`
}

type LolRankedSeason struct {
	CurrentSeasonEnd int64 `json:"currentSeasonEnd"`
	CurrentSeasonId  int   `json:"currentSeasonId"`
	NextSeasonStart  int   `json:"nextSeasonStart"`
}
