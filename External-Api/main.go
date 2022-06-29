package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/PratikVarute/Bulk-Data-Upload-CSV/model"
)

func main() {
	var userData model.User

	response, err := http.Get("https://random-data-api.com/api/users/random_user")
	err = json.NewDecoder(response.Body).Decode(&userData)
	fmt.Println("data in structure \n", userData)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
