package gointerview_test


import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

var ids []string
var input string = `["9100000001,1,dzf,10001,1200","9100000002,2,inv,10002,200","9100000001,12,dzf,10001,3000","9100000001,123,dzf,10003,2100"]`

func TestTopic004(t *testing.T) {
	e := json.Unmarshal([]byte(input), &ids)

	fmt.Println(ids, e)

	dzfMap := map[string][]string{}
	dzfTotal := map[string]int{}
	invMap := map[string][]string{}
	invTotal := map[string]int{}
	speMap := map[string][]string{}
	speTotal := map[string]int{}
	for _, v := range ids {
		r := strings.Split(v, ",")
		switch r[2] {
		case "dzf":
			dzfMap[r[3]] = append(dzfMap[r[3]], r[0]+","+r[1])
			price, _ := strconv.Atoi(r[4])
			dzfTotal[r[3]] = dzfTotal[r[3]] + price
		case "inv":
			invMap[r[3]] = append(invMap[r[3]], r[0]+","+r[1])
			price, _ := strconv.Atoi(r[4])
			invTotal[r[3]] = invTotal[r[3]] + price
		case "spe":
			speMap[r[3]] = append(speMap[r[3]], r[0]+","+r[1])
			price, _ := strconv.Atoi(r[4])
			speTotal[r[3]] = speTotal[r[3]] + price
		default:
		}
	}
	// for k, v := range invMap {
	// 	fmt.Printf("%v %v\n", k, v)
	// }
	for k, v := range dzfMap {
		vv, _ := json.Marshal(v)
		fmt.Printf("%v %v\n", k, string(vv))
	}
	for k, v := range invMap {
		vv, _ := json.Marshal(v)
		fmt.Printf("%v %v\n", k, string(vv))
	}
	for k, v := range speMap {
		vv, _ := json.Marshal(v)
		fmt.Printf("%v %v\n", k, string(vv))
	}
	// fmt.Printf("%v \n", dzfMap)
	fmt.Println(dzfTotal, invTotal, speTotal)
}
