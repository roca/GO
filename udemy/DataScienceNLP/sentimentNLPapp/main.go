package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)

type SentimentDetails struct {
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Neutral  float64 `json:"neutral"`
	Compound float64 `json:"compound"`
}

func main() {

	engine := html.New("./views", ".html")
	// Reload
	engine.Reload(true)

	// instance of app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Route
	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello App")
		initMessage := "Sentiment Analysis App on the Go"
		return c.Render("index", fiber.Map{
			"initMessage": initMessage,
		})
	})

	app.Post("/", func(c *fiber.Ctx) error {
		// return c.SendString("Hello App")
		initMessage := "Sentiment Analysis App on the Go"
		message := c.FormValue("message")
		sentimentResults := sentimentize(message)
		return c.Render("index", fiber.Map{
			"initMessage":      initMessage,
			"originalMsg":      message,
			"sentimentDetails": sentimentResults,
		})
	})

	// API Route
	// localhost:3000/api/?text="this is your sentiment"

	app.Get("/api/:text?", func(c *fiber.Ctx) error {
		message := c.Query("text")
		sentimentResults := analyzeSentiment(message)
		return c.JSON(fiber.Map{
			"message":   message,
			"sentiment": sentimentResults,
		})

	})

	// Swagger
	app.Get("/swagger/*", swagger.Handler)

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	app.Get("/api/:text?", func(c *fiber.Ctx) error {
		message := c.Query("text")
		sentimentResults := analyzeSentiment(message)
		return c.JSON(fiber.Map{
			"message":   message,
			"sentiment": sentimentResults,
		})

	})

	// Listen Route
	app.Listen(":3000")
}

// Functions for processing
func sentimentize(docx string) SentimentDetails {
	// Parse
	parsedtext := sentitext.Parse(docx, lexicon.DefaultLexicon)
	// Process
	results := sentitext.PolarityScore(parsedtext)
	sentimentScores := SentimentDetails{Positive: results.Positive, Negative: results.Negative, Neutral: results.Neutral, Compound: results.Compound}
	return sentimentScores
}

func analyzeSentiment(docx string) float64 {
	// Parse
	parsedtext := sentitext.Parse(docx, lexicon.DefaultLexicon)
	// Process
	results := sentitext.PolarityScore(parsedtext)
	return results.Compound
}
