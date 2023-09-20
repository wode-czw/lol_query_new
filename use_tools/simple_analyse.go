package main

import (
	"czw_lol_query_tools/lcu"
	"fmt"
	"time"
)

//你们别抢我补位的accountID
//2602095920588480

//4017882674  站台为你买几橘
//4135692180 豹女ID

//4009959910 大笨蛋你在里面吗

// Summoner_name := "牛顿有辆小车220-240"
// rank_title := "铂金111"
func main() {

	//这个版本的任务是认为这个账号一生一共打了超过30把排位。。。。的前提条件

	port_token := lcu.Get_port_token()

	var accountID int64
	accountID = 2602095920588480
	Summoner_name := "你们别抢我补位"
	rank_title := "钻1"
	query_name := "recent_30_rank"
	file_path := "../data/rank30_json/"
	query_command_time := time.Now().Format("2006-01-02 15:04:05 Mon Jan")[11:13] + "_" + time.Now().Format("2006-01-02 15:04:05 Mon Jan")[14:16] //hour_min
	//################################################################################################################################################################
	//################################################################################################################################################################
	//################################################################################################################################################################
	//################################################################################################################################################################
	file_name := file_path + Summoner_name + query_name + "_" + query_command_time + ".json"
	file_d_t_m_list := "../data/伤害转化率-召唤师名字-段位.json"

	var rank_30_info = make([]lcu.GameInfo, 0, 30)
	var win_num int
	var if_30 bool
	//begin_number = 0
	//end_number = 20

	rank_30_info, win_num, if_30 = Get_rank30(rank_30_info, port_token, accountID, 0, 20)

	czw_champion_map := Get_champion_map("../data/champion_files/simple_champion_list.json")

	//返回的有总转化率
	all_damage_turn := show_what_you_get(rank_30_info, czw_champion_map, port_token, accountID)

	//要写的东西都准备好了，要传进来的变量有		file_name, win_num, if_30 rank_30_info
	write_rank_file(file_name, win_num, if_30, rank_30_info)
	//#-----------------------------------------------------------------------------------------
	//统计的文件和这个人的段位   file_stat, rank_title, Summoner_name all_damage_turn
	write_statistics(file_d_t_m_list, rank_title, Summoner_name, all_damage_turn)

	fmt.Println("end")
}
