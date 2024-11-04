package scraper

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchProfile(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func ScrapeProfile(url string, assignments []string) (Profile, error) {
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

	profile.CompletedAssignments = getCompletedAssignments(profile.Badges, assignments)
	profile.CompletedAssignmentsCount = len(profile.CompletedAssignments)

	profile.IncompleteAssignments = getIncompleteAssignments(profile.CompletedAssignments, assignments)
	profile.IncompleteAssignmentsCount = len(profile.IncompleteAssignments)

	return profile, nil
}
