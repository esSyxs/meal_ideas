package server

import (
	logger "System/Log"
	"System/food"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const idName = "id"

func Start(port string) {
	r := mux.NewRouter()

	r.HandleFunc(fmt.Sprintf("/recipes/{%s}", idName), recpiesGet).Methods(http.MethodGet)

	logger.Default.Println("starting server")
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	logger.Default.Println("stopping server")
}

func recpiesGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recepiesID := vars[idName]
	logger.Default.Printf("got recepie request with id: %s\n", recepiesID)

	i, err := strconv.Atoi(recepiesID)
	if err != nil {
		logger.Default.Printf("failed to parse recepie id: %s, err: %s\n", recepiesID, err.Error())

		writeJsonRespone(w, http.StatusBadRequest, err.Error())
		return
	}

	rec, err := food.GetRecepie(uint(i))
	if err != nil {
		logger.Default.Printf("failed to find recepie id: %s, err: %s\n", recepiesID, err.Error())

		writeJsonRespone(w, http.StatusBadRequest, err.Error())
		return
	}

	writeJsonRespone(w, http.StatusOK, rec)
}

func writeJsonRespone(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		//TODO Should log instead of panic
		panic("There was an error encoding the initialized struct")
	}
}
