package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
    "scraper/internal/scraper"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Usage: scraper <input.csv> <output.json>")
	}

	csvFileName := os.Args[1]
	jsonFileName := os.Args[2]

	csvFile, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

    var profiles []scraper.Profile

	for i, record := range records[1:] {
		if len(record) < 3 {
			log.Printf("Invalid record at line %d: %v", i+2, record)
			continue
		}

		profileURL := record[2]
		userName := record[0]
		userEmail := record[1]

        var assignments = []string{
            "Get Started with Pub/Sub",
            "Develop GenAI Apps with Gemini and Streamlit",
            "Prompt Design in Vertex AI",
            "Analyze Images with the Cloud Vision API",
            "Networking Fundamentals on Google Cloud",
            "Monitoring in Google Cloud",
            "Cloud Speech API: 3 Ways",
            "App Engine: 3 Ways",
            "Cloud Functions: 3 Ways",
            "Get Started with Google Workspace Tools",
            "Get Started with Dataplex",
            "The Basics of Google Cloud Compute",
            "Get Started with Looker",
            "Get Started with API Gateway",
            "Get Started with Cloud Storage",
            "Level 3: Google Cloud Adventures",
        }

		profile, err := scraper.ScrapeProfile(profileURL, assignments)
		if err != nil {
			log.Printf("Error scraping profile %s: %v", profileURL, err)
			continue
		} 
		profile.Name = strings.TrimSpace(userName)
		profile.Email = strings.TrimSpace(userEmail)

        if err == nil {
			log.Printf("Successfully scraped profile: %s", profile.Name)
		}

        profiles = append(profiles, profile)
	}

    if err := scraper.WriteToJSONFile(profiles, jsonFileName); err != nil {
        log.Fatal(err)
    }

	log.Printf("Response written to %s", jsonFileName)
}
