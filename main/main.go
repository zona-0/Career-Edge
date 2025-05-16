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
type Education struct {
	School, Degree string
	Year   int
}

var skills[100]Skill
var skillCount int = 0
var experiences [100]Experience
var experienceCount int = 0
var educations [100]Education
var educationCount int = 0

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
			addSkill()
			backToMenu()
		} else if menu == 2 {
			addExperience()
			backToMenu()
		} else if menu == 3{
			addEductaion()
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
	fmt.Println("[XX] Type 'pback' to return to the main menu")
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
	fmt.Printf("[>>] Please enter a skill\n(Golang, Python): ")
	fmt.Scan(&nameSkill)

	newSkill.Name = nameSkill
	skills[skillCount] = newSkill
	skillCount++
	fmt.Println("[!!] Skill added sucesstully")
}

func addExperience() {
    var title, company, desc string
	var newExperience Experience

	fmt.Printf("[>>] Job title\n(Fullstack Developer, Software Engineer): ")
    fmt.Scanln(&title)
    fmt.Printf("[>>] Company\n(Telkom Indonesia): ")
    fmt.Scanln(&company)
    fmt.Printf("[>>] About Me\n(Description): ")
    fmt.Scanln(&desc)

    newExperience.Title = title
    newExperience.Company = company
    newExperience.Description = desc

    experiences[experienceCount] = newExperience
    experienceCount++

    fmt.Println("[>>] Work experience added successfully")
}

func addEductaion() {
	var newEducation Education
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

	educations[educationCount] = newEducation
	educationCount++

	fmt.Println("[!!] Education added successfully")
}
