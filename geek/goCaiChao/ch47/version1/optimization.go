package version1

import (
	"encoding/json"
	"strconv"
)

func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		payload[i] = i
	}

	req := Request{"demo_transaction", payload}
	v, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	return string(v)
}

func processRequest(reqs []string) []string {
	reps := []string{}

	for _, req := range reqs {
		obj := Request{}
		json.Unmarshal([]byte(req), &obj)

		ret := ""
		for _, e := range obj.PayLoad {
			ret += strconv.Itoa(e) + ","
		}

		rep := Response{obj.TransactionID, ret}
		repJSON, err := json.Marshal(rep)
		if err != nil {
			panic(err)
		}

		reps = append(reps, string(repJSON))
	}

	return reps
}
