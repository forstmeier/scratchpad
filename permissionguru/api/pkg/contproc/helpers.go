package contproc

import "github.com/schollz/closestmatch"

func matchTexts(sources []string, target string) string {
	cm := closestmatch.New(sources, []int{3})

	return cm.Closest(target)
}
