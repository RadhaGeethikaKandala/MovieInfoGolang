package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/dto"
)

func GetMovieList(name string) dto.OmdbResponse {
	var omdbResponse dto.OmdbResponse

	url := os.Getenv("baseUrl") + name
	fmt.Println(url)

	response, err := http.Get(url)
	logError(err, "Unable to fetch data from omdbapi")

	responseData, err := io.ReadAll(response.Body)
	logError(err, "Unable to read from response")

	err = json.Unmarshal(responseData, &omdbResponse)
	logError(err, "Unable to parse responseData to omdbResponse")
	return omdbResponse
}

func logError(err error, message string) {
	if err != nil {
		log.Fatal(message, err.Error())
	}
}
