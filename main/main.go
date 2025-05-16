package main
import "fmt"

type Skill struct {
	Name string
}
type Experience struct {
	Title       string
	Company     string
	Description string
}

var skills[50]Skill
var skillCount int = 0
var experiences [50]Experience
var experienceCount int = 0

func main(){
	var menu int
	for {
		fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
		fmt.Println("║                     C A R E E R  E D G E                      ║")
		fmt.Println("║     AI Assistant for Smart Resumes & Cover Letter Creation    ║")
		fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
		fmt.Println("║ [1] Tambah Skill                      | Add Skill             ║")
		fmt.Println("║ [2] Tambah Pengalaman                 | Add Experience        ║")
		fmt.Println("║ [3] Tambah Pendidikan                 | Add Education         ║")
		fmt.Println("║ [4] Evaluasi Resume (OgerAI)          | Evaluate Resume       ║")
		fmt.Println("║ [5] Tampilkan Data Pengguna           | Show User Data        ║")
		fmt.Println("║ [6] Buat Surat Lamaran Otomatis       | Generate Cover Letter ║")
		fmt.Println("║ [7] Tambah Lowongan Pekerjaan         | Add Job Listing       ║")
		fmt.Println("║ [8] Cari Lowongan Pekerjaan           | Search Job Listing    ║")
		fmt.Println("║ [9] Urutkan Gaji (Selection Sort)     | Sort by Salary        ║")
		fmt.Println("║ [10] Urutkan Kata Kunci (Insertion)   | Sort by Keyword       ║")
		fmt.Println("║ [0] Keluar                            | Exit                  ║")
		fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
		fmt.Print("[!!!] Pilih menu / Choose option: ")	
		fmt.Scanln(&menu)

		if menu == 1 {
			addSkill()
			backToMenu()
		} else if menu == 2 {
			addExperience()
			backToMenu()
		} else if menu == 0{
			fmt.Println("Ty")
			break
		}
		fmt.Println()
	}
}

func backToMenu(){
	var backMenu string
	fmt.Print("[XX] Type 'pback' to return to the main menu")
	for {
		fmt.Scanln(&backMenu)
		if backMenu == "pback"{
			break
		} else {
			fmt.Println()
		}
	}
}

func addSkill(){
	var nameSkill string
	var newSkill Skill
	fmt.Println("[>>] Please enter a skill (e.g., Java, Python): ")
	fmt.Scan(&nameSkill)

	newSkill.Name = nameSkill
	skills[skillCount] = newSkill
	skillCount++
	fmt.Println("[!!] Skill added sucesstully")
}


func addExperience() {
    var title, company, desc string
	var newExperience Experience

	fmt.Print("[>>] Job title (e.g., Software Engineer): ")
    fmt.Scanln(&title)
    fmt.Print("[>>] Company (e.g., ABC Corp): ")
    fmt.Scanln(&company)
    fmt.Print("[>>] Brief description (e.g., Web application development): ")
    fmt.Scanln(&desc)

    newExperience.Title = title
    newExperience.Company = company
    newExperience.Description = desc

    experiences[experienceCount] = newExperience
    experienceCount++

    fmt.Println("[>>] Work experience added successfully.")
}