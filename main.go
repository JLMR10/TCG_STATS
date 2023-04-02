package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type game struct {
	Date         string `json:"date"`
	FirstDeck    string `json:"firstDeck"`
	SecondDeck   string `json:"secondDeck"`
	WinnerDeck   string `json:"winnerDeck"`
	FirstPlayer  string `json:"firstPlayer"`
	SecondPlayer string `json:"secondPlayer"`
	WinnerPlayer string `json:"winnerPlayer"`
}

var games = []game{
	{Date: time.Now().Format("2006-01-02"), FirstDeck: "SDS Escanor", SecondDeck: "MTI 8 Stby", WinnerDeck: "SDS Escanor", FirstPlayer: "Jose", SecondPlayer: "Will", WinnerPlayer: "Jose"},
	{Date: time.Now().Format("2006-01-02"), FirstDeck: "SDS Escanor", SecondDeck: "DAL 8 Choice", WinnerDeck: "SDS Escanor", FirstPlayer: "Poe", SecondPlayer: "Jose", WinnerPlayer: "Jose"},
}

func getGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, games)
}

func gamesByDeck(c *gin.Context) {
	deck := c.Param("deck")
	games, err := getGamesByDeck(deck)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Games not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, games)
}

func getGamesByDeck(deck string) ([]*game, error) {
	var res []*game

	for i, b := range games {
		if b.FirstDeck == deck || b.SecondDeck == deck {
			res = append(res, &games[i])
		}
	}
	if len(res) == 0 {
		return nil, errors.New("Games not found")
	}
	return res, nil
}

func createGame(c *gin.Context) {
	var newGame game

	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	games = append(games, newGame)
	c.IndentedJSON(http.StatusCreated, newGame)
}

func main() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/games/:deck", gamesByDeck)
	router.POST("/games", createGame)
	router.Run("localhost:8080")
}
