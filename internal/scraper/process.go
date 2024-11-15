package scraper

func getCompletedAssignments(badges []string, assignments []string, arcades []string) []string {
	var completed []string
	for _, assignment := range assignments {
		for _, badge := range badges {
			if badge == assignment {
				completed = append(completed, badge)
				break
			}
		}
	}

	found := false
	for _, arcade := range arcades {
		for _, badge := range badges {
			if badge == arcade {
				completed = append(completed, badge)
				found = true
			}
		}

		if found {
			break
		}
	}

	return completed
}

func getIncompleteAssignments(completed []string, assignments []string, arcades []string) []string {
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

	for _, arcade := range arcades {
		found := false
		for _, badge := range completed {
			if arcade == badge {
				found = true
				break
			}
		}

		if !found {
			incomplete = append(incomplete, arcade)
		}
	}

	return incomplete
}
