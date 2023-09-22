package main

import (
	"czw_lol_query_tools/Only_step"
	"czw_lol_query_tools/get_port_token"
	"czw_lol_query_tools/lcu"
	"czw_lol_query_tools/tools"
	"fmt"
)

func main() {
	account_name := "你们别抢我补位"
	fmt.Println("I am", account_name)
	czw_port_token := get_port_token.Return_port_token() + "/"
	fmt.Println(czw_port_token, "这是你用来在apipost设置{{server}}的值")
	czw_client := Only_step.New_client()

	ID_info := tools.Return_getAccountID_by_summonerName(account_name)
	PrintCurrSummonerInfo(ID_info)

	get_ranked_command := czw_port_token + "lol-ranked/v1/ranked-stats/" + ID_info.Puuid
	rank_data := Only_step.Get_ranked_data_by_puuid(get_ranked_command, czw_client)
	Print_rank_data(rank_data)
}

func PrintCurrSummonerInfo(summoner *lcu.CurrSummoner) {

	fmt.Printf("%-30s %-35d\n", "AccountId:", summoner.AccountId)
	fmt.Printf("%-30s %-35s\n", "DisplayName:", summoner.DisplayName)
	fmt.Printf("%-30s %-35s\n", "Puuid:", summoner.Puuid)
	fmt.Printf("%-30s %-35d\n", "SummonerId:", summoner.SummonerId)

}

func Print_rank_data(rank *lcu.LolRankedRankedStats) {

	fmt.Printf("%-30s %-15s %-10s\n", "现在的单双段位:", rank.HighestCurrentSeasonReachedTierSR, rank.HighestPreviousSeasonAchievedDivision)
	fmt.Printf("%-30s %-15s %-10s\n", "过去的单双段位:", rank.HighestPreviousSeasonAchievedTier, rank.HighestPreviousSeasonEndDivision)
	fmt.Printf("%-30s %-15s %-10s\n", "现在的灵活段位:", rank.QueueMap["RANKED_FLEX_SR"].Tier, rank.QueueMap["RANKED_FLEX_SR"].Division)
	fmt.Printf("%-30s %-15s %-10s\n", "过去的灵活段位:", rank.QueueMap["RANKED_FLEX_SR"].HighestTier, rank.QueueMap["RANKED_FLEX_SR"].HighestDivision)

}
