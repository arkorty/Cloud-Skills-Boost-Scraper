package scraper

func getCompletedAssignments(badges []string, assignments []string) []string {
	var completed []string
	for _, assignment := range assignments {
		for _, badge := range badges {
			if badge == assignment {
				completed = append(completed, badge)
				break
			}
		}
	}
	return completed
}

func getIncompleteAssignments(completed []string, assignments []string) []string {
	var incomplete []string
	for _, assignment := range assignments {
		found := false
		for _, badge := range completed {
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
