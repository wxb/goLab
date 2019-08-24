package version2

import (
	"strconv"
)

func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		payload[i] = i
	}

	req := Request{"demo_transaction", payload}
	v, err := req.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return string(v)
}

func processRequest(reqs []string) []string {
	reps := []string{}

	for _, req := range reqs {
		obj := Request{}
		obj.UnmarshalJSON([]byte(req))

		ret := ""
		for _, e := range obj.PayLoad {
			ret += strconv.Itoa(e) + ","
		}

		rep := Response{obj.TransactionID, ret}
		repJSON, err := rep.MarshalJSON()
		if err != nil {
			panic(err)
		}

		reps = append(reps, string(repJSON))
	}

	return reps
}
