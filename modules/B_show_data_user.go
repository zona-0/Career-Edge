package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func ShowAllUserData() {
	var i int
	var edu data.Education

	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        [ RINGKASAN DATA USER ]                     ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

	// TODO: Show skill data
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║              [ SKILL USER ]            ║")
	fmt.Println("╠════════════════════════════════════════╣")
	if data.SkillCount == 0 {
		fmt.Println("║      [!] Belum ada data skill          ║")
	} else {
		i = 0
		for i < data.SkillCount {
			fmt.Printf("║  %-1d. %-34s ║\n", i+1, Capitalize(data.Skills[i].Name))
			i = i + 1
		}
	}
	fmt.Println("╚════════════════════════════════════════╝")

	// TODO: Show experience data
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        [ PENGALAMAN KERJA ]                        ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	if data.ExperienceCount == 0 {
		fmt.Println("║              [!] Belum ada pengalaman kerja                        ║")
	} else {
		fmt.Println("║  No.  ║        Jabatan         ║         Perusahaan                ║")
		fmt.Println("╠═══════╬════════════════════════╬═══════════════════════════════════╣")
		i = 0
		for i < data.ExperienceCount {
			fmt.Printf("║  %-4d ║ %-22s ║ %-33s ║\n", i+1,
				Capitalize(data.Experiences[i].Title),
				Capitalize(data.Experiences[i].Company))
			i = i + 1
		}
	}
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")

	// TODO: Show education data
	fmt.Println("╔════════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                        [ RIWAYAT PENDIDIKAN ]                      ║")
	fmt.Println("╠════════════════════════════════════════════════════════════════════╣")
	if data.EducationCount == 0 {
		fmt.Println("║               [!] Belum ada data pendidikan                        ║")
	} else {
		fmt.Println("║  No.  ║       Sekolah/Universitas        ║     Gelar     ║ Tahun   ║")
		fmt.Println("╠═══════╬══════════════════════════════════╬═══════════════╬═════════╣")
		i = 0
		for i < data.EducationCount {
			edu = data.Educations[i]
			fmt.Printf("║  %-4d ║ %-32s ║ %-13s ║ %-7d ║\n", i+1,
				Capitalize(edu.School),
				Capitalize(edu.Degree),
				edu.Year)
			i = i + 1
		}
	}
	fmt.Println("╚════════════════════════════════════════════════════════════════════╝")
}
