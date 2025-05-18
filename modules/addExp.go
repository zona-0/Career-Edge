package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func AddExperience() {
	var title, company string
	var newExperience data.Experience

	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                          [Add Work Experience]                     ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ Please enter your job experience details below                     ║")
	fmt.Println("║ Use underscore (_) instead of spaces for multi-word input          ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ [>>] Examples of Job Titles:                                       ║")
	fmt.Println("║     - Fullstack_Developer                                          ║")
	fmt.Println("║     - Backend_Engineer                                             ║")
	fmt.Println("║     - Frontend_Developer                                           ║")
	fmt.Println("║     - Data_Analyst                                                 ║")
	fmt.Println("║     - Marketing_Specialist                                         ║")
	fmt.Println("║     - Content_Writer                                               ║")
	fmt.Println("║     - Graphic_Designer                                             ║")
	fmt.Println("║     - Teacher                                                      ║")
	fmt.Println("║     - Sales_Executive                                              ║")
	fmt.Println("║                                                                    ║")
	fmt.Println("║ [>>] Examples of Companies:                                        ║")
	fmt.Println("║     - Telkom_Indonesia                                             ║")
	fmt.Println("║     - Google                                                       ║")
	fmt.Println("║     - Microsoft                                                    ║")
	fmt.Println("║     - Local_High_School                                            ║")
	fmt.Println("║     - Marketing_Firm                                               ║")
	fmt.Println("║                                                                    ║")
	// fmt.Println("║ [>>] Examples of Description (brief):                              ║")
	// fmt.Println("║     - Backend_Engineer_with_5_years_experience                     ║")
	// fmt.Println("║     - Digital_Marketing_and_Social_Media_Manager                   ║")
	// fmt.Println("║     - Created_graphic_designs_for_clients                          ║")
	// fmt.Println("║     - Teaching_English_to_high_school_students                     ║")
	// fmt.Println("║     - Managed_sales_team_and_exceeded_targets                      ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

	fmt.Print("[>>] Job Title: ")
	fmt.Scan(&title)
	title = ToLower(title)

	fmt.Print("[>>] Company Name: ")
	fmt.Scan(&company)
	company = ToLower(company)

	// fmt.Print("[>>] Description (brief): ")
	// fmt.Scan(&desc)
	// desc = ToLower(desc)

	newExperience.Title = title
	newExperience.Company = company
	// newExperience.Description = desc

	data.Experiences[data.ExperienceCount] = newExperience
	data.ExperienceCount += 1

	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                    [!!] Experience added successfully              ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
}
