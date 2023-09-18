package main

import (
	"czw_lol_query_tools/get_port_token"
	"fmt"
)

func main() {
	account_name := "你们别抢我补位"
	fmt.Println("I am", account_name)
	czw_port_token := get_port_token.Return_port_token()
	fmt.Println(czw_port_token)

}
