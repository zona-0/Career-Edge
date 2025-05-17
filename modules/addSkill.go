package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)


func AddSkill() {
	var nameSkill string
	var newSkill data.Skill
	
	fmt.Printf("[>>] Please enter a skill\n(Golang, Python): ")
	fmt.Scan(&nameSkill)

	newSkill.Name = nameSkill
	data.Skills[data.SkillCount] = newSkill
	data.SkillCount++
	fmt.Println("[!!] Skill added sucesstully")
}
