package scraper

type Profile struct {
	Name                     string   `json:"name"`
	Email                    string   `json:"email"`
	ProfileURL               string   `json:"profile_url"`
	BadgesCount              int      `json:"badges_count"`
	IncompleteAssignmentsCount int    `json:"incomplete_assignments_count"`
	CompletedAssignmentsCount int     `json:"completed_assignments_count"`
	Badges                   []string `json:"badges"`
	IncompleteAssignments     []string `json:"incomplete_assignments"`
	CompletedAssignments      []string `json:"completed_assignments"`
}
