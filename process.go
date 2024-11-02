package main

import "os"

func appendClosingBracket(filename string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("]")
	return err
}

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
