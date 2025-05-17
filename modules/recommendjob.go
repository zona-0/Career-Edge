package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func RecommendJob() {
	var i, j int
	var backend, frontend, analyst, writer, teacher, marketing, designer int
	var total int
	var educationScore int
	var name, title, school, degree string

	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║                 AI CAREER SUGGESTION ENGINE        ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║ [SYSTEM] Analyzing profile: Skills, Experience,    ║")
	fmt.Println("║         and Education to generate recommendations  ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	for i = 0; i < data.SkillCount; i++ {
		name = ToLower(data.Skills[i].Name)

		if name == "go" || name == "golang" || name == "backend" || name == "java" || name == "python" || name == "csharp" || name == "nodejs" || name == "restapi" {
			backend += 1
		}

		if name == "html" || name == "css" || name == "javascript" || name == "react" || name == "vuejs" || name == "angular" || name == "typescript" {
			frontend += 1
		}

		if name == "sql" || name == "excel" || name == "analysis" || name == "machine_learning" || name == "data_mining" || name == "statistics" || name == "r" || name == "python_data" {
			analyst += 1
		}

		if name == "writing" || name == "creativewriting" || name == "content_creation" || name == "copywriting" || name == "editing" {
			writer += 1
		}

		if name == "teaching" || name == "publicspeaking" || name == "curriculum_design" || name == "education" {
			teacher += 1
		}

		if name == "marketing" || name == "sales" || name == "digital_marketing" || name == "seo" || name == "social_media" || name == "brand_management" {
			marketing += 1
		}

		if name == "design" || name == "photoshop" || name == "illustrator" || name == "graphic_design" || name == "ux" || name == "ui" {
			designer += 1
		}
	}

	for i = 0; i < data.ExperienceCount; i++ {
		title = ToLower(data.Experiences[i].Title)

		if title == "backend_developer" || title == "backend_engineer" {
			backend = backend + 1
		}
		if title == "frontend_developer" || title == "frontend_engineer" {
			frontend = frontend + 1
		}
		if title == "data_analyst" || title == "business_analyst" {
			analyst = analyst + 1
		}
		if title == "writer" || title == "content_writer" || title == "editor" {
			writer = writer + 1
		}
		if title == "teacher" || title == "lecturer" || title == "instructor" {
			teacher = teacher + 1
		}
		if title == "marketing_specialist" || title == "sales_executive" || title == "digital_marketer" {
			marketing = marketing + 1
		}
		if title == "graphic_designer" || title == "ui_designer" || title == "ux_designer" {
			designer = designer + 1
		}
	}

	for j = 0; j < data.EducationCount; j++ {
		school = ToLower(data.Educations[j].School)
		degree = ToLower(data.Educations[j].Degree)

		if school == "telkom_university" {
			educationScore = educationScore + 5
		} else if school == "itb" {
			educationScore = educationScore + 4
		} else if school == "ui" {
			educationScore = educationScore + 3
		} else if school == "binus_university" {
			educationScore = educationScore + 2
		} else if school == "universitas_negeri" {
			educationScore = educationScore + 1
		} else {
			educationScore = educationScore + 0
		}

		if degree == "s3_computer_science" || degree == "s3_informatics" {
			educationScore = educationScore + 3
			backend = backend + 1
			frontend = frontend + 1
			analyst = analyst + 1
		} else if degree == "s2_computer_science" || degree == "s2_informatics" {
			educationScore = educationScore + 2
			backend = backend + 1
			frontend = frontend + 1
			analyst = analyst + 1
		} else if degree == "s1_computer_science" || degree == "s1_informatics" || degree == "s1_informatika" {
			educationScore = educationScore + 1
			backend = backend + 1
			frontend = frontend + 1
			analyst = analyst + 1
		} else if degree == "s1_english_literature" || degree == "s1_literature" {
			writer = writer + 1
		} else if degree == "s1_education" {
			teacher = teacher + 1
		} else if degree == "s1_marketing" || degree == "s1_business_administration" {
			marketing = marketing + 1
		} else if degree == "s1_design" {
			designer = designer + 1
		}
	}

	fmt.Println("\n═════════════════════════════════════════════════════")
	fmt.Println(">> AI Suggestion Summary:")
	fmt.Println("═════════════════════════════════════════════════════")

	total = backend + frontend + analyst + writer + teacher + marketing + designer + educationScore

	if backend > 0 {
		fmt.Printf("Backend Developer\t\t(score: %d)\n", backend)
	}
	if frontend > 0 {
		fmt.Printf("Frontend Developer\t\t(score: %d)\n", frontend)
	}
	if analyst > 0 {
		fmt.Printf("Data Analyst\t\t\t(score: %d)\n", analyst)
	}
	if writer > 0 {
		fmt.Printf("Content Writer / Editor\t(score: %d)\n", writer)
	}
	if teacher > 0 {
		fmt.Printf("Teacher / Lecturer\t\t(score: %d)\n", teacher)
	}
	if marketing > 0 {
		fmt.Printf("Marketing Specialist\t\t(score: %d)\n", marketing)
	}
	if designer > 0 {
		fmt.Printf("Graphic Designer\t\t(score: %d)\n", designer)
	}

	if total == 0 {
		fmt.Println("No relevant suggestions found. Please input more profile data.")
	}

	fmt.Println("═════════════════════════════════════════════════════")
	fmt.Println("[INFO] These are rule-based suggestions. As you update your")
	fmt.Println("       profile with more accurate data, AI will give better")
	fmt.Println("       career alignment.")
	fmt.Println("═════════════════════════════════════════════════════")
}
