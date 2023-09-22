package main

import (
	"czw_lol_query_tools/Only_step"
	"czw_lol_query_tools/lcu"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
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
	account_name := "躲在她后面"
	begin_index := "0"
	end_index := "10"
	//------------------------------------------------------------------------------------
	fmt.Println("I am", account_name)
	czw_port_token := "https://riot:VnvYlvlf4RaJc176Qj0iNg@127.0.0.1:2502" + "/"
	query_command := fmt.Sprintf("lol-summoner/v1/summoners?name=%s", account_name)
	url := czw_port_token + query_command
	czw_client := Only_step.New_client()

	//------------------------------------------------------------------------------------
	//目的是拿到这个人的puuid
	data_struct := Only_step.Body_to_struct_return_CurrSummoner(czw_client, url)
	show_Curr_data(data_struct)

	//------------------------------------------------------------------------------------
	query_match_command := czw_port_token + "lol-match-history/v1/products/lol/" + data_struct.Puuid + "/matches?begIndex=" + begin_index + "&endIndex=" + end_index
	game_struct := Only_step.Body_to_struct_return_GameInfo_list(czw_client, query_match_command)
	//show_GameInfo_list_data(game_struct)

	//------------------------------------------------------------------------------------
	//根据game id获取数据
	length := len(game_struct.Games.Games)
	for i := 0; i < length; i++ {
		game_id_querry_command := czw_port_token + "lol-match-history/v1/games/" + strconv.Itoa(int(game_struct.Games.Games[i].GameId))
		game_info := Only_step.Body_to_struct_return_GameInfo(czw_client, game_id_querry_command)
		//ParticipantIdentities[i]就是第i个人
		show_game_info(game_info, account_name)
	}

	//根据召唤师名字返回召唤师信息，但是只能返回同一个大区的内容

	query_command_time := time.Now().Format("2006-01-02 15:04:05 Mon Jan")[11:13] + "_" + time.Now().Format("2006-01-02 15:04:05 Mon Jan")[14:16] //hour_min
	file_name := "data/召唤师名字json/" + account_name + "_" + query_command_time + ".json"
	write_file_by_CurrSummoner(file_name, data_struct)

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

// 显示这把游戏里的10个队友
func show_game_info(Game_info_data *lcu.GameInfo, name string) {
	for i := 0; i < 10; i++ {

		if Game_info_data.ParticipantIdentities[i].Player.SummonerName == name {
			the_id := Game_info_data.ParticipantIdentities[i].ParticipantId
			for j := 0; j < 10; j++ {
				if the_id == Game_info_data.Participants[j].ParticipantId {
					k := Game_info_data.Participants[j].Stats.Kills
					d := Game_info_data.Participants[j].Stats.Deaths
					a := Game_info_data.Participants[j].Stats.Assists
					fmt.Println("K-D-A:", k, "-", d, "-", a)
					fmt.Println("--------------------------------------")
				}
			}

		}

	}

	//fmt.Println(Game_info_data.ParticipantIdentities[1].Player.SummonerName)
	//fmt.Println("------------------------------------------------------------------------------------------")
}

func write_file_by_CurrSummoner(filename string, data *lcu.CurrSummoner) {
	fileobj, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("os.OpenFile failed ,err: ", err)
		return

	}
	defer fileobj.Close()

	write_string, _ := json.MarshalIndent(data, "", "    ")

	fileobj.Write([]byte(string(write_string)))

}
