package modules
import (
	"fmt"
	"CAREER-EDGE/data"
)

func CreateResume() {
	var i int
	var user string
	SuggestionAI()

	fmt.Scan(&user)
	// Clear()
	// TODO: AI resume sec
	if ToLower(user) == "resume"{
		fmt.Println(">> Tentu! untuk membuat resume yang efektif, aku butuh informasi dasar terlebih dahulu. Silahkan isi data berikut: ")
		fmt.Println("[A] Informasi Pribadi: ")
		fmt.Printf("	[X] Nama: \n	[X] Alamat: \n	[X] No. HP: \n	[X] Email: \n")
		fmt.Println("[B] Jelaskan tentang dirimu: ")
		fmt.Println("	[X] About Me: ")
		// TODO: TAKE EDUCATION DATA
		fmt.Println("[C] Pendidikan: ")
		if data.EducationCount == 0 {
			fmt.Println("	[HELP] Anda belum mengisi data pendidikan. Silakan isi melalui menu utama!")
		} else {
			for i = 0; i < data.EducationCount; i++ {
				fmt.Printf("	[V] Sekolah/Universitas: %s\n	[V] Jenjang: %s\n	[V] Tahun Lulus: %d\n",
					data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
			}
		}

		// TODO: TAKE SKILLS DATA
		fmt.Println("[D] Keterampilan: ")
		if data.SkillCount == 0 {
			fmt.Println("	[HELP] Anda belum mengisi data keterampilan. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.SkillCount {
				fmt.Printf("	[V] %d. %s\n", i+1, Capitalize(data.Skills[i].Name))
				i += 1
			}
		}


		// TODO: PENGALAMAN KERJA
		fmt.Println("[E] Pengalaman Kerja: ")
		if data.ExperienceCount == 0 {
			fmt.Println("	[HELP] Anda belum mengisi data pengalaman kerja. Silakan isi melalui menu utama!")
		} else {
			i = 0
			for i < data.ExperienceCount {
				fmt.Printf("	[V] %d. %s %s\n", i+1, Capitalize(data.Experiences[i].Title), Capitalize(data.Experiences[i].Company))
				i += 1
			}
		}

		// TODO: CERTIFICATION SEC
		fmt.Println("[F] Sertifikat: ")

		// TODO: END SEC
		fmt.Println()
		fmt.Printf(">> Silahkan isi informasi di atas atau cukup jawab yang kamu punya sekarang.\n   Setelah itu, aku akan buatkan resume versi teks yang rapi dan siap pakai\n")
		fmt.Println()
		endSec()

		// TODO: USER INPUT SEC
		var nama, alamat, hp, email string

		fmt.Print(">> Nama: ")
		fmt.Scan(&nama)
		fmt.Print(">> Alamat: ")
		fmt.Scan(&alamat)
		fmt.Print(">> No. Hp: ")
		fmt.Scan(&hp)
		fmt.Print(">> Email: ")
		fmt.Scan(&email)
		fmt.Println()

		var kata, aboutme string
		var selesai bool = false

		fmt.Println(">> Tentang Saya: ")
		fmt.Println("   [HELP] Akhiri dengan tanda '.' untuk mengakhiri")
		for selesai == false{
			fmt.Scan(&kata)

			if kata == "."{
				selesai = true
			} else {
				if aboutme == ""{
					aboutme = kata
				} else {
					aboutme = aboutme + " " + kata
				}
			}
		}

		var sertifikat string
		fmt.Print(">> Sertifikat: ")
		fmt.Scan(&sertifikat)


		// TODO: RESUME PRINT SECTION
		Header()
		fmt.Println(" [A] Informasi Pribadi: ")
		fmt.Printf("	[V] Nama: %s\n	[V] Alamat: %s\n	[V] No. HP: %s\n	[V] Email: %s\n", nama, alamat, hp, email)
		fmt.Println(" [B] About Me: ", aboutme)
		// TODO: TAKE EDUCATION DATA
		fmt.Println(" [C] Pendidikan: ")
		fmt.Println("	 [HELP] Jika anda masih melihat pesan ini, maka anda harus mengisi data pada menu aplikasi diawal!")
		for i = 0; i < data.EducationCount; i++ {
			fmt.Printf("	 [V] Sekolah/Universitas: %s\n	[V] Jenjang: %s\n	[V] Tahun Lulus: %d\n",
			data.Educations[i].School, data.Educations[i].Degree, data.Educations[i].Year)
		}
		// TODO: TAKE SKILLS DATA
		fmt.Println(" [D] Keterampilan: ")
		fmt.Println("	 [HELP] Jika anda masih melihat pesan ini, maka anda harus mengisi data pada menu aplikasi diawal!")
		i = 0
		for i < data.SkillCount {
			// var skillName string = Capitalize(data.Skills[i].Name)
			fmt.Printf("	 [V] %d. %s\n", i+1, Capitalize(data.Skills[i].Name))
			i += 1
		}
		// TODO: PENGALAMAN KERJA
		fmt.Println(" [E] Pengalaman Kerja: ")
		fmt.Println("	 [HELP] Jika anda masih melihat pesan ini, maka anda harus mengisi data pada menu aplikasi diawal!")
		i = 0	
		for i < data.ExperienceCount {
			fmt.Printf("	[V] %d. %s %s\n", i+1, Capitalize(data.Experiences[i].Title), Capitalize(data.Experiences[i].Company))
			i += 1
		}
		// TODO: CERTIFICATION SEC
		fmt.Println(" [F] Sertifikat: ", sertifikat)
		endSec()

		//fmt.Println("TEST: ", aboutme)
		// fmt.Println("[A] Informasi Pribadi: ")
		// fmt.Printf("	[X] Nama: %s\n	[X] Alamat: %s\n	[X] No. HP: %s\n	[X] Email: %s\n", nama, alamat, hp, email)
		//ManageEducation()
	}else {
		fmt.Println()
		fmt.Println("[HELP] ketik 'resume' untuk bantuan membuat resume atau ketik 'suratkerja' untuk bantuan membuat surat lamaran kerja")
	}
}
