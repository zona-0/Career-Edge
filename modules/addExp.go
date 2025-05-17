package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func AddExperience() {
	var title, company, desc string
	var newExperience data.Experience

	fmt.Printf("[>>] Job title\n(Fullstack Developer, Software Engineer): ")
	fmt.Scanln(&title)
	fmt.Printf("[>>] Company\n(Telkom Indonesia): ")
	fmt.Scanln(&company)
	fmt.Printf("[>>] About Me\n(Description): ")
	fmt.Scanln(&desc)

	newExperience.Title = title
	newExperience.Company = company
	newExperience.Description = desc

	data.Experiences[data.ExperienceCount] = newExperience
	data.ExperienceCount++

	fmt.Println("[>>] Work experience added successfully")
}