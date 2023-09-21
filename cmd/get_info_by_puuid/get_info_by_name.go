package main

import (
	"czw_lol_query_tools/Only_step"
	"czw_lol_query_tools/get_port_token"
	"fmt"
	"strconv"
)

func main() {
	account_name := "LDL拐子"
	begin_index := "0"
	end_index := "2"
	fmt.Println("I am", account_name)
	czw_port_token := get_port_token.Return_port_token() + "/"

	query_command := fmt.Sprintf("lol-summoner/v1/summoners?name=%s", account_name)
	url := czw_port_token + query_command

	czw_client := Only_step.New_client()
	data_struct := Only_step.Body_to_struct_return_CurrSummoner(czw_client, url)

	fmt.Printf("%-20s%-20s\n", "puuid", data_struct.Puuid)
	fmt.Printf("%-20s%-20d\n", "AccountID", data_struct.AccountId)

	//{{server}}/lol-match-history/v1/products/lol/e07e8c6a-2f1c-5030-b96f-24665a6b6a99/matches?begIndex=0&endIndex=2
	query_match_command := czw_port_token + "lol-match-history/v1/products/lol/" + data_struct.Puuid + "/matches?begIndex=" + begin_index + "&endIndex=" + end_index
	fmt.Println(query_match_command)

	game_struct := Only_step.Body_to_struct_return_GameInfo_list(czw_client, query_match_command)
	fmt.Println("length of it", len(game_struct.Games.Games))
	fmt.Println(game_struct.Games.Games[0].GameId)
	fmt.Println(game_struct.Games.Games[0].GameType)
	fmt.Println(game_struct.Games.Games[1].GameDuration) //秒数
	fmt.Println(game_struct.Games.Games[2].Participants[0].Stats.Kills)
	fmt.Println(game_struct.Games.Games[2].Participants[0].Stats.Deaths)
	fmt.Println(game_struct.Games.Games[2].Participants[0].Stats.Assists)

	game_id := game_struct.Games.Games[0].GameId
	//这里获取了game_id
	game_id_querry_command := czw_port_token + "lol-match-history/v1/games/" + strconv.Itoa(int(game_id))
	game_info := Only_step.Body_to_struct_return_GameInfo(czw_client, game_id_querry_command)
	fmt.Println(len(game_info.ParticipantIdentities))

	fmt.Println(len(game_info.ParticipantIdentities))

	//ParticipantIdentities[i]就是第i个人
	fmt.Println(game_info.ParticipantIdentities[0].Player.SummonerName)
}
