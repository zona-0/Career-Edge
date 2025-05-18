package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func AddEducation() {
	var newEducation data.Education
	var school, degree, major string
	var year int

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                  [Education Input - AI Assistant]            ║")
	fmt.Println("╠══════════════════════════════════════════════════════════════╣")
	fmt.Println("║ Please enter your education details below.                   ║")
	fmt.Println("║ Use underscore (_) instead of spaces. Example:               ║")
	fmt.Println("║  - School/University: Telkom_University                      ║")
	fmt.Println("║  - Degree: D1/D2/D3/D4 || S1/S2/S3                           ║")
	fmt.Println("║  - Major: Informatika, Akuntansi, Teknik_Industri            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")

	fmt.Print("[>>] Input School/University\n>>> ")
	fmt.Scanln(&school)

	fmt.Print("[>>] Degree\n>>> ")
	fmt.Scanln(&degree)

	fmt.Print("[>>] Major\n>>> ")
	fmt.Scanln(&major)

	fmt.Print("[>>] Year of graduation (Example: 2024)\n>>> ")
	fmt.Scanln(&year)

	newEducation.School = school
	//newEducation.Degree = degree
	newEducation.Degree = degree + "_" + major
	newEducation.Year = year

	data.Educations[data.EducationCount] = newEducation
	data.EducationCount = data.EducationCount + 1

	fmt.Println("╔══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                [!!] Education added successfully            ║")
	fmt.Println("╚══════════════════════════════════════════════════════════════╝")
}
