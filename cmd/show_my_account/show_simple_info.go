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
	fmt.Println(czw_port_token, "这是你用来在apipost设置{{server}}的值")

	ID_info := tools.Return_getAccountID_by_summonerName(account_name)
	PrintCurrSummonerInfo(ID_info)
}

func PrintCurrSummonerInfo(summoner *lcu.CurrSummoner) {
	fmt.Printf("%-15s %-35d\n", "AccountId:", summoner.AccountId)
	fmt.Printf("%-15s %-35s\n", "DisplayName:", summoner.DisplayName)
	fmt.Printf("%-15s %-35s\n", "Puuid:", summoner.Puuid)
	fmt.Printf("%-15s %-35d\n", "SummonerId:", summoner.SummonerId)
	fmt.Printf("%-15s %-35d\n", "SummonerLevel:", summoner.SummonerLevel)
	fmt.Printf("%-15s %-35d\n", "CurrentPoints:", summoner.RerollPoints.CurrentPoints)
}
