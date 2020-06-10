package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber"
	helper "github.com/scraper_v2/helper"
	models "github.com/scraper_v2/models"
)

func ImdbController(fibCon *fiber.Ctx) {
	id := fibCon.Query("id")
	_, err := os.Stat(id + ".json")
	api_key := os.Getenv("omdb_key")
	if os.IsNotExist(err) {
		url := fmt.Sprintf("https://www.omdbapi.com/?i=%s&apikey=%s", id, api_key)
		res, _ := helper.GetResponse(url)
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err.Error())
		}
		var data models.Imdb
		json.Unmarshal(body, &data)
		print(data.Response)
		if data.Response == "False" {
			fibCon.Status(204)
		} else {
			file, _ := json.MarshalIndent(data, "", " ")
			_ = ioutil.WriteFile(id+".json", file, 0644)
			fibCon.Status(200).JSON(data)
		}

	} else {
		file, _ := ioutil.ReadFile(id + ".json")
		data := models.Imdb{}
		_ = json.Unmarshal([]byte(file), &data)
		fibCon.Status(200).JSON(data)

	}

}
