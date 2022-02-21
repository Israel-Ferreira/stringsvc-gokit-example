package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Israel-Ferreira/stringsvc/endpoints"
	"github.com/Israel-Ferreira/stringsvc/models"
	"github.com/Israel-Ferreira/stringsvc/services"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := services.StrService{}

	uppercaseHandler := httptransport.NewServer(endpoints.MakeUppercaseEndPoint(svc), decodeUppercase, encodeUppercase)

	http.Handle("/uppercase", uppercaseHandler)

	log.Fatalln(http.ListenAndServe(":9030", nil))

}

func encodeUppercase(_ context.Context, rw http.ResponseWriter, response interface{}) error {

	rw.Header().Add("Content-Type", "application/json")

	return json.NewEncoder(rw).Encode(response)
}

func decodeUppercase(_ context.Context, r *http.Request) (interface{}, error) {
	var request models.UppercaseRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}
