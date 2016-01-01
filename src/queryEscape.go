package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	token := "CkAn_wSQtzyM2fB5-wXfH9t0enQye5Tz1qy_VoQ0vHFMzdGqgzAMxvF3-a5z0faoZ_RlRrBhCtVKGxlDfPdlG4PlJoF_4Hdg3-eECBDeR3ADoZYsiI7AN9u-7wiL6FRSQzywcWv3Uu3Zd70Pl-E_uJMwVmGVq86L_BYLkrMJ_s99x7Bx4nWV_JETK-tje5nnMwAA__8="
	bytes, _ := json.Marshal(map[string]string{"ticket": token})

	fmt.Println(url.QueryEscape(string(bytes)))
}
