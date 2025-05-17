package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func AddEducation(){
	var newEducation data.Education
	var school, degree string
	var year int

	fmt.Printf("[>>] Input School/University\n(Telkom University): ")
	fmt.Scanln(&school)
	fmt.Printf("[>>] Degree/Major\n(S1 Computer Engineering): ")
	fmt.Scanln(&degree)
	fmt.Printf("[>>] Year of graduation: ")
	fmt.Scanln(&year)

	newEducation.School = school
	newEducation.Degree = degree
	newEducation.Year = year

	data.Educations[data.EducationCount] = newEducation
	data.EducationCount++

	fmt.Println("[!!] Education added successfully")
}