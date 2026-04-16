package main

import (
	"fmt"
	"strings"
)

func formatMealList(meals []Meal) string {
	if len(meals) == 0 {
		return "-"
	}

	var sb strings.Builder
	for _, m := range meals {
		sb.WriteString(fmt.Sprintf(" · ~%s\n", m.Name))
	}

	return sb.String()
}

func formatMealListAdded(meals []Meal) string {
	if len(meals) == 0 {
		return "—"
	}
	var sb strings.Builder
	for _, m := range meals {
		sb.WriteString(fmt.Sprintf("  • *%s*\n", m.Name))
	}
	return sb.String()
}
