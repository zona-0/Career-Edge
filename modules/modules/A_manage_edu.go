package modules
import (
	"fmt"
	"CAREER-EDGE/data"
)

func ManageEducation() {
	var command, index, i, year int
	var school, degree, major string
	var newEducation data.Education
	var running bool = true

	for running {
		fmt.Println("╔══════════════════════════════════════════════════════════════╗")
		fmt.Println("║                      [Education Manager]                     ║")
		fmt.Println("╠══════════════════════════════════════════════════════════════╣")
		fmt.Println("║ 1. Add    - Tambah pendidikan baru                           ║")
		fmt.Println("║ 2. Edit   - Ubah data pendidikan yang sudah dimasukkan       ║")
		fmt.Println("║ 3. List   - Lihat semua pendidikan                           ║")
		fmt.Println("║ 4. Delete - Hapus salah satu data pendidikan                 ║")
		fmt.Println("║ 5. Done   - Selesai dan simpan                               ║")
		fmt.Println("╚══════════════════════════════════════════════════════════════╝")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&command)

		if command == 1 {
			if data.EducationCount >= 10 {
				fmt.Println("[!] Pendidikan sudah mencapai batas maksimum (10)")
			} else {
				SuggestionEducation()
				fmt.Print("[>>] School/University: ")
				fmt.Scan(&school)
				fmt.Print("[>>] Degree: ")
				fmt.Scan(&degree)
				fmt.Print("[>>] Major: ")
				fmt.Scan(&major)
				fmt.Print("[>>] Year of Graduation (contoh: 2024): ")
				fmt.Scan(&year)

				newEducation.School = school
				newEducation.Degree = degree + "_" + major
				newEducation.Year = year

				data.Educations[data.EducationCount] = newEducation
				data.EducationCount += 1

				fmt.Println("[V] Pendidikan berhasil ditambahkan")
				Clear()
			}
		} else if command == 2 {
			if data.EducationCount == 0 {
				fmt.Println("[!] Belum ada data pendidikan untuk diubah")
			} else {
				fmt.Println("╔═══════╦════════════════════════════════════╦══════════════╦══════╗")
				fmt.Println("║  No.  ║           School/University        ║    Degree    ║ Year ║")
				fmt.Println("╠═══════╬════════════════════════════════════╬══════════════╬══════╣")
				for i = 0; i < data.EducationCount; i++ {
					fmt.Printf("║  %-4d ║ %-34s ║ %-12s ║ %-4d ║\n",
						i+1, data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
				}
				fmt.Println("╚═══════╩════════════════════════════════════╩══════════════╩══════╝")

				fmt.Printf("Masukkan nomor data pendidikan yang ingin diubah (1-%d): ", data.EducationCount)
				fmt.Scan(&index)

				if index < 1 || index > data.EducationCount {
					fmt.Println("[!] Nomor tidak valid")
				} else {
					SuggestionEducation()
					fmt.Print("[>>] School/University Baru: ")
					fmt.Scan(&school)
					fmt.Print("[>>] Degree Baru (D1/D2/D3/D4/S1/S2/S3): ")
					fmt.Scan(&degree)
					fmt.Print("[>>] Major Baru: ")
					fmt.Scan(&major)
					fmt.Print("[>>] Year of Graduation Baru: ")
					fmt.Scan(&year)

					data.Educations[index-1].School = school
					data.Educations[index-1].Degree = degree + "_" + major
					data.Educations[index-1].Year = year

					fmt.Println("[V] Data pendidikan berhasil diubah")
					Clear()
				}
			}
		} else if command == 3 {
			if data.EducationCount == 0 {
				fmt.Println("[!] Belum ada data pendidikan yang dimasukkan")
			} else {
				fmt.Println("╔═══════╦════════════════════════════════════╦══════════════╦══════╗")
				fmt.Println("║  No.  ║           School/University        ║    Degree    ║ Year ║")
				fmt.Println("╠═══════╬════════════════════════════════════╬══════════════╬══════╣")
				for i = 0; i < data.EducationCount; i++ {
					fmt.Printf("║  %-4d ║ %-34s ║ %-12s ║ %-4d ║\n",
						i+1, data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
				}
				fmt.Println("╚═══════╩════════════════════════════════════╩══════════════╩══════╝")
			}

		} else if command == 4 {
			if data.EducationCount == 0 {
				fmt.Println("[!] Belum ada data pendidikan untuk dihapus")
			} else {
				fmt.Println("╔═══════╦════════════════════════════════════╦══════════════╦══════╗")
				fmt.Println("║  No.  ║           School/University        ║    Degree    ║ Year ║")
				fmt.Println("╠═══════╬════════════════════════════════════╬══════════════╬══════╣")
				for i = 0; i < data.EducationCount; i++ {
					fmt.Printf("║  %-4d ║ %-34s ║ %-12s ║ %-4d ║\n",
						i+1, data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
				}
				fmt.Println("╚═══════╩════════════════════════════════════╩══════════════╩══════╝")

				fmt.Printf("Masukkan nomor data pendidikan yang ingin dihapus (1-%d): ", data.EducationCount)
				fmt.Scan(&index)

				if index < 1 || index > data.EducationCount {
					fmt.Println("[!] Nomor tidak valid")
				} else {
					for i = index - 1; i < data.EducationCount-1; i++ {
						data.Educations[i] = data.Educations[i+1]
					}
					data.EducationCount -= 1
					fmt.Println("[V] Data pendidikan berhasil dihapus")
				}
				Clear()
			}
		} else if command == 5 {
			fmt.Println("[V] Sesi input pendidikan selesai")
			running = false
			
		} else {
			fmt.Println("[!] Pilihan tidak dikenali. Masukkan angka <1 - 5>")
		}
	}
}
