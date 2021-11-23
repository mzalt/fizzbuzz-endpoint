package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gol4ng/logger"

	"github.com/fizzbuzz-endpoint/internal/model"
)

// MyFizzBuzz endpoint
func MyFizzBuzz(log logger.LoggerInterface) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Type", "application/json; charset=UTF-8")

		query := request.URL.Query()
		int1, _ := strconv.Atoi(query.Get("int1"))
		int2, _ := strconv.Atoi(query.Get("int2"))
		limit, _ := strconv.Atoi(query.Get("limit"))
		input := &model.FizzbuzzInpute{
			Number1: int1,
			Number2: int2,
			Limit:   limit,
			Str1:    query.Get("str1"),
			Str2:    query.Get("str2"),
		}

		result, err := myFizzBuzz(input)
		if err != nil {
			log.Error("invalid input : %err%", logger.Error("err", err))
			response.WriteHeader(http.StatusBadRequest)
			return
		}

		payload, err := json.Marshal(result)
		if err != nil {
			log.Error("faild to marshal result : %err%", logger.Error("err", err))
		}
		response.WriteHeader(http.StatusOK)
		response.Write(payload)
	}
}

func myFizzBuzz(req *model.FizzbuzzInpute) (*model.FizzbuzzResponse, error) {
	if !req.CheckInput() {
		return nil, fmt.Errorf("bad values for number1 or/and number2")
	}

	output := "1"
	for i := 2; i <= req.Limit; i++ {
		result := ""
		if i%req.Number1 == 0 {
			result += req.Str1
		}
		if i%req.Number2 == 0 {
			result += req.Str2
		}
		if result != "" {
			output += ", " + result
			continue
		}
		output += ", " + strconv.Itoa(i)
	}

	return &model.FizzbuzzResponse{Output: output}, nil
}
