package tools

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"czw_lol_query_tools/Only_step"
	"czw_lol_query_tools/get_port_token"
	"czw_lol_query_tools/lcu"
)

func Use_getAccountID_by_summonerName() {

	port_token := lcu.Get_port_token()

	//根据召唤师名字返回召唤师信息，但是只能返回同一个大区的内容
	inquire_name := "血狼灬死神OL"
	query_command := fmt.Sprintf("lol-summoner/v1/summoners?name=%s", inquire_name)
	query_name := inquire_name
	query_command_time := time.Now().Format("2006-01-02 15:04:05 Mon Jan")[11:13] + "_" + time.Now().Format("2006-01-02 15:04:05 Mon Jan")[14:16] //hour_min

	file_name := "../data/召唤师名字json/" + query_name + "_" + query_command_time + ".json"

	url := port_token + query_command

	czw_client := Only_step.New_client()
	data_struct := Only_step.Body_to_struct_return_CurrSummoner(czw_client, url)

	write_file_by_CurrSummoner(file_name, data_struct)

}

// 这个函数也执行了get_port_token.Return_port_token。根据召唤师的用户名获得召唤师的用户信息
func Return_getAccountID_by_summonerName(account_name string) *lcu.CurrSummoner {

	port_token := get_port_token.Return_port_token() + "/"

	//根据召唤师名字返回召唤师信息，但是只能返回同一个大区的内容
	inquire_name := account_name
	query_command := fmt.Sprintf("lol-summoner/v1/summoners?name=%s", inquire_name)

	url := port_token + query_command

	czw_client := Only_step.New_client()
	data_struct := Only_step.Body_to_struct_return_CurrSummoner(czw_client, url)

	return data_struct

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
