package main

import (
	"czw_lol_query_tools/Only_step"
	"czw_lol_query_tools/get_port_token"
	"czw_lol_query_tools/lcu"
	"fmt"
	"strconv"
)

/*
就算这个人隐藏战绩也没用
输入是召唤师名字，找到CurrSummoner吗，从中提取到puuid
根据puuid，找到GameInfo_list
再找game ID


最终目的是这个人最近10把的kda，英雄名字，伤害占比，经济占比
并计算出一个比wegame更加合理的分数

*/

func main() {
	account_name := "只剩冬季"
	begin_index := "0"
	end_index := "10"
	//------------------------------------------------------------------------------------
	fmt.Println("I am", account_name)
	czw_port_token := get_port_token.Return_port_token() + "/"
	query_command := fmt.Sprintf("lol-summoner/v1/summoners?name=%s", account_name)
	url := czw_port_token + query_command
	czw_client := Only_step.New_client()

	//------------------------------------------------------------------------------------
	data_struct := Only_step.Body_to_struct_return_CurrSummoner(czw_client, url)
	show_Curr_data(data_struct)

	//------------------------------------------------------------------------------------
	query_match_command := czw_port_token + "lol-match-history/v1/products/lol/" + data_struct.Puuid + "/matches?begIndex=" + begin_index + "&endIndex=" + end_index
	game_struct := Only_step.Body_to_struct_return_GameInfo_list(czw_client, query_match_command)
	show_GameInfo_list_data(game_struct)

	//------------------------------------------------------------------------------------
	game_id_querry_command := czw_port_token + "lol-match-history/v1/games/" + strconv.Itoa(int(game_struct.Games.Games[0].GameId))
	game_info := Only_step.Body_to_struct_return_GameInfo(czw_client, game_id_querry_command)
	//ParticipantIdentities[i]就是第i个人
	show_game_info(game_info)

}

func show_Curr_data(CurrSummoner_data *lcu.CurrSummoner) {
	fmt.Printf("%-20s%-20s\n", "puuid", CurrSummoner_data.Puuid)
	fmt.Printf("%-20s%-20d\n", "AccountID", CurrSummoner_data.AccountId)
	fmt.Println("------------------------------------------------------------------------------------------")
}

func show_GameInfo_list_data(GameInfo_list_data *lcu.LolMatchHistoryMatchHistoryList) {
	fmt.Println("查了多少把比赛\t\t", len(GameInfo_list_data.Games.Games))
	fmt.Println("Game ID\t\t", GameInfo_list_data.Games.Games[0].GameId)
	fmt.Println("Game Type\t\t", GameInfo_list_data.Games.Games[0].GameType)
	fmt.Println("Game Mode\t\t", GameInfo_list_data.Games.Games[0].GameMode)
	k := GameInfo_list_data.Games.Games[0].Participants[0].Stats.Kills
	d := GameInfo_list_data.Games.Games[0].Participants[0].Stats.Deaths
	a := GameInfo_list_data.Games.Games[0].Participants[0].Stats.Assists
	fmt.Println("K-D-A:", k, "-", d, "-", a)
	fmt.Println("------------------------------------------------------------------------------------------")
}

func show_game_info(Game_info_data *lcu.GameInfo) {
	for i := 0; i < 10; i++ {
		fmt.Println(Game_info_data.ParticipantIdentities[i].Player.SummonerName)
	}

	fmt.Println(Game_info_data.ParticipantIdentities[1].Player.SummonerName)
	fmt.Println("------------------------------------------------------------------------------------------")
}
