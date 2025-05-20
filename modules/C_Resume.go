package modules
import "fmt"


func CreateResume() {
	var user string
	SuggestionAI()

	fmt.Scan(&user)
	if user == "resume"{
		fmt.Println(">> Tentu! untuk membuat resume yang efektif, aku butuh informasi dasar terlebih dahulu. Silahkan isi data berikut: ")
		fmt.Println("[A] Informasi Pribadi: ")
		fmt.Printf("	[X] Nama: \n	[X] Alamat: \n	[X] No. HP: \n	[X] Email: \n")
		fmt.Println("[B] Jelaskan tentang dirimu!: ")
		fmt.Println("	[X] About Me: ")
		
	}
}
