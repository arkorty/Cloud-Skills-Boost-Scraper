package main

import (
	"encoding/csv"
	"encoding/json"
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Profile struct {
	Name                     string   `json:"name"`
	Email                    string   `json:"email"`
	ProfileURL               string   `json:"profile_url"`
	BadgesCount              int      `json:"badges_count"`
	IncompleteAssignmentsCount int      `json:"incomplete_assignments_count"`
	CompletedAssignmentsCount int      `json:"completed_assignments_count"`
	Badges                   []string `json:"badges"`
	IncompleteAssignments     []string `json:"incomplete_assignments"`
	CompletedAssignments      []string `json:"completed_assignments"`
}


func fetchProfile(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func scrapeProfile(url string, assignments []string) (Profile, error) {
	res, err := fetchProfile(url)
	if err != nil {
		return Profile{}, err
	}
	defer res.Body.Close()

	
	if res.StatusCode != 200 {
		return Profile{}, fmt.Errorf("failed to fetch profile: %s", res.Status)
	}

	
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return Profile{}, err
	}

	
	profile := Profile{
		ProfileURL: url,
	}

	
	doc.Find(".profile-badges .profile-badge").Each(func(i int, s *goquery.Selection) {
		badge := strings.TrimSpace(s.Find(".ql-title-medium").Text())
		if badge != "" {
			profile.Badges = append(profile.Badges, badge)
		}
	})

	
	profile.BadgesCount = len(profile.Badges)
	profile.IncompleteAssignmentsCount = len(assignments) - profile.BadgesCount
	profile.CompletedAssignmentsCount = countCompletedAssignments(profile.Badges, assignments)

	
	profile.CompletedAssignments = getCompletedAssignments(profile.Badges, assignments)
	profile.IncompleteAssignments = getIncompleteAssignments(profile.Badges, assignments)

	return profile, nil
}


func countCompletedAssignments(badges []string, assignments []string) int {
	count := 0
	for _, badge := range badges {
		for _, assignment := range assignments {
			if badge == assignment {
				count++
			}
		}
	}
	return count
}


func getCompletedAssignments(badges []string, assignments []string) []string {
	var completed []string
	for _, badge := range badges {
		for _, assignment := range assignments {
			if badge == assignment {
				completed = append(completed, badge)
			}
		}
	}
	return completed
}


func getIncompleteAssignments(badges []string, assignments []string) []string {
	var incomplete []string
	for _, assignment := range assignments {
		found := false
		for _, badge := range badges {
			if assignment == badge {
				found = true
				break
			}
		}
		if !found {
			incomplete = append(incomplete, assignment)
		}
	}
	return incomplete
}


func appendToJSONFile(profile Profile, filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") 
	if err := encoder.Encode(profile); err != nil {
		return err
	}

	return nil
}


func readAssignments(filename string) ([]string, error) {
    var assignments []string

    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close() 

    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        
        if trimmed := strings.TrimSpace(line); trimmed != "" {
            assignments = append(assignments, trimmed)
        }
    }

    
    if err := scanner.Err(); err != nil {
        return nil, err
    }

    return assignments, nil
}


func main() {
	csvFile, err := os.Open("data/profile_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	jsonFileName := "data/scraped_data.json"
	if err := os.WriteFile(jsonFileName, []byte("["), 0644); err != nil { 
		log.Fatal(err)
	}
	
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	assignmentsFile := "data/assignments.txt"
	assignments, err := readAssignments(assignmentsFile)
	if err != nil {
		fmt.Println("Error reading assignments:", err)
		return
	}
	
	for i, record := range records[1:] {
		if len(record) < 3 {
			log.Printf("Invalid record at line %d: %v", i+2, record) 
			continue
		}

		profileURL := record[2] 
		userName := record[0]   
		userEmail := record[1]  

		profile, err := scrapeProfile(profileURL, assignments)
		if err != nil {
			log.Printf("Error scraping profile %s: %v", profileURL, err)
			continue
		}

		profile.Name = strings.TrimSpace(userName)
		profile.Email = strings.TrimSpace(userEmail)

		if err := appendToJSONFile(profile, jsonFileName); err != nil {
			log.Printf("Error saving profile to JSON: %v", err)
		} else {
			log.Printf("Successfully scraped and saved profile: %s", profile.Name)
		}
	}

	
	if err := appendClosingBracket(jsonFileName); err != nil {
		log.Fatal(err)
	}

	log.Println("All profiles processed and saved to profiles.json")
}


func appendClosingBracket(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("]")
	return err
}
