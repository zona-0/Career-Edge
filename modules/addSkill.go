package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func AddSkill() {
	var nameSkill string
	var newSkill data.Skill
	var count int = 0

	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                     [Skill Input - AI Assistant]                  ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	fmt.Println("║ Please enter your skills one by one (max 10)                      ║")
	fmt.Println("║ Use underscore (_) for multiword skills. Example:                 ║")
	fmt.Println("║  - Golang, Backend, HTML, CSS, SQL                                ║")
	fmt.Println("║  - CreativeWriting, Public_Speaking, Marketing, Photoshop         ║")
	fmt.Println("║ Capitalization doesn't matter. Type 'done' when you're finished   ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

	for count < 10 {
		fmt.Print(">>> ")
		fmt.Scan(&nameSkill)
		nameSkill = ToLower(nameSkill)
		
		if nameSkill == "done" {
			break
		}

		newSkill.Name = nameSkill
		data.Skills[data.SkillCount] = newSkill
		data.SkillCount += 1

		count += 1
	}

	fmt.Println("[!!] Skill added successfully")
}
