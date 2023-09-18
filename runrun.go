package main

import (
	"czw_lol_query_tools/get_port_token"
	"czw_lol_query_tools/lcu"
	"czw_lol_query_tools/tools"
	"fmt"
)

func main() {
	account_name := "你们别抢我补位"
	fmt.Println("I am", account_name)
	czw_port_token := get_port_token.Return_port_token() + "/"
	fmt.Println(czw_port_token, "==========-=-============")

	ID_info := tools.Return_getAccountID_by_summonerName(account_name)
	PrintCurrSummonerInfo(ID_info)
}

// chatgpt写的一行一行展示lcu.CurrSummoner的数据
func PrintCurrSummonerInfo(summoner *lcu.CurrSummoner) {
	fmt.Printf("AccountId:\t\t\t%d\n", summoner.AccountId)
	fmt.Printf("DisplayName:\t\t\t%s\n", summoner.DisplayName)
	fmt.Printf("InternalName:\t\t\t%s\n", summoner.InternalName)
	fmt.Printf("NameChangeFlag:\t\t\t%v\n", summoner.NameChangeFlag)
	fmt.Printf("PercentCompleteForNextLevel:\t\t\t%d\n", summoner.PercentCompleteForNextLevel)
	fmt.Printf("ProfileIconId:\t\t\t%d\n", summoner.ProfileIconId)
	fmt.Printf("Puuid:\t\t\t%s\n", summoner.Puuid)
	fmt.Printf("RerollPoints.CurrentPoints:\t\t\t%d\n", summoner.RerollPoints.CurrentPoints)
	fmt.Printf("RerollPoints.MaxRolls:\t\t\t%d\n", summoner.RerollPoints.MaxRolls)
	fmt.Printf("RerollPoints.NumberOfRolls:\t\t\t%d\n", summoner.RerollPoints.NumberOfRolls)
	fmt.Printf("RerollPoints.PointsCostToRoll:\t\t\t%d\n", summoner.RerollPoints.PointsCostToRoll)
	fmt.Printf("RerollPoints.PointsToReroll:\t\t\t%d\n", summoner.RerollPoints.PointsToReroll)
	fmt.Printf("SummonerId:\t\t\t%d\n", summoner.SummonerId)
	fmt.Printf("SummonerLevel:\t\t\t%d\n", summoner.SummonerLevel)
	fmt.Printf("Unnamed:\t\t\t%v\n", summoner.Unnamed)
	fmt.Printf("XpSinceLastLevel:\t\t\t%d\n", summoner.XpSinceLastLevel)
	fmt.Printf("XpUntilNextLevel:\t\t\t%d\n", summoner.XpUntilNextLevel)
}
