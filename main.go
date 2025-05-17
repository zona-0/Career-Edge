package main

import (
	"CAREER-EDGE/modules"
	"fmt"
)

func main() {
	var menu int

	fmt.Println("===============================================================")
	fmt.Println(">>> Initializing Systems...")
	fmt.Println(">>> Loading modules: Resume Engine, CoverLetter Generator...")
	fmt.Println(">>> Welcome to your personal career assistant")
	fmt.Println("===============================================================")
	fmt.Println("[!!!] NOTICE:")
	fmt.Println(">>> This system is a specialized assistant designed only to help you build")
	fmt.Println(">>> Smart resumes and job specific cover letters")
	fmt.Println(">>> This AI is NOT a general chatbot and does not support open conversation")
	fmt.Println(">>> Please use the menu options to build your career profile effectively")
	fmt.Println("===============================================================")
	fmt.Println()

	for {
		fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
		fmt.Println("║                     C A R E E R  E D G E                      ║")
		fmt.Println("║     AI Assistant for Smart Resumes & Cover Letter Creation    ║")
		fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
		fmt.Println("║ [1] Tambah Skill                      | Add Skill             ║")
		fmt.Println("║ [2] Tambah Pengalaman                 | Add Experience        ║")
		fmt.Println("║ [3] Tambah Pendidikan                 | Add Education         ║")
		// fmt.Println("║ [4] Evaluasi Resume <AI Assistance>   | Evaluate Resume       ║")
		// fmt.Println("║ [5] Tampilkan Data Pengguna           | Show User Data        ║")
		// fmt.Println("║ [6] Buat Surat Lamaran Otomatis       | Generate Cover Letter ║")
		// fmt.Println("║ [7] Tambah Lowongan Pekerjaan         | Add Job Listing       ║")
		// fmt.Println("║ [8] Cari Lowongan Pekerjaan           | Search Job Listing    ║")
		// fmt.Println("║ [9] Urutkan Gaji <Selection Sort>     | Sort by Salary        ║")
		// fmt.Println("║ [10] Urutkan Kata Kunci <Insertion>   | Sort by Keyword       ║")
		fmt.Println("║ [0] Keluar                            | Exit                  ║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
		fmt.Print("[!!!] Pilih menu / Choose option: ")
		fmt.Scanln(&menu)

		if menu == 1 {
			modules.AddSkill()
			modules.BackToMenu()
		} else if menu == 2 {
			modules.AddExperience()
			modules.BackToMenu()
		} else if menu == 3 {
			modules.AddEducation()
			modules.BackToMenu()
		} else if menu == 0 {
			fmt.Println("Ty")
			break
		}
		fmt.Println()
	}
}
