package modules

import (
	"fmt"
	"CAREER-EDGE/data"
)

func RecommendJob() {
	var i, j int
	var backend, frontend, analyst, writer, teacher, marketing, designer, devops, mobile, security, management, finance, customer_service int
	var game, video, law, health, admin, logistics, hr, psychology, chef, events, cashier, operator, fnb, creator, total int
	var fullstack, it, companyScore int
	var educationScore int
	var name, title, school, degree, major, company string

	fmt.Println("\n╔════════════════════════════════════════════════════╗")
	fmt.Println("║                 AI CAREER SUGGESTION ENGINE        ║")
	fmt.Println("╠════════════════════════════════════════════════════╣")
	fmt.Println("║ [SYSTEM] Analyzing profile: Skills, Experience,    ║")
	fmt.Println("║         and Education to generate recommendations  ║")
	fmt.Println("╚════════════════════════════════════════════════════╝")

	for i = 0; i < data.SkillCount; i++ {
		name = ToLower(data.Skills[i].Name)
		
		if name == "go" || name == "golang" || name == "backend" || name == "java" || name == "python" || name == "csharp" || name == "nodejs" || name == "restapi" || name == "spring" || name == "laravel" || name == "pengembang_backend" {
			backend += 1
		}

		if name == "html" || name == "css" || name == "javascript" || name == "react" || name == "vuejs" || name == "angular" || name == "typescript" || name == "nextjs" || name == "tailwind" || name == "pengembang_frontend" {
			frontend += 1
		}

		if name == "sql" || name == "excel" || name == "analysis" || name == "machine_learning" || name == "data_mining" || name == "statistics" || name == "r" || name == "python_data" || name == "powerbi" || name == "tableau" || name == "analisis_data" {
			analyst += 1
		}

		if name == "writing" || name == "creativewriting" || name == "content_creation" || name == "copywriting" || name == "editing" || name == "menulis" || name == "penulis" || name == "konten" {
			writer += 1
		}

		if name == "teaching" || name == "publicspeaking" || name == "curriculum_design" || name == "education" || name == "training" || name == "mentoring" || name == "mengajar" || name == "guru" {
			teacher += 1
		}

		if name == "marketing" || name == "sales" || name == "digital_marketing" || name == "seo" || name == "social_media" || name == "brand_management" || name == "ads" || name == "market_research" || name == "pemasaran" || name == "penjualan" {
			marketing += 1
		}

		if name == "design" || name == "photoshop" || name == "illustrator" || name == "graphic_design" || name == "ux" || name == "ui" || name == "figma" || name == "adobe_xd" || name == "desain_grafis" {
			designer += 1
		}

		if name == "docker" || name == "kubernetes" || name == "linux" || name == "ci_cd" || name == "jenkins" || name == "devops" || name == "infrastructure" || name == "terraform" {
			devops += 1
		}

		if name == "android" || name == "ios" || name == "flutter" || name == "kotlin" || name == "swift" || name == "react_native" || name == "mobile_dev" || name == "pengembang_mobile" {
			mobile += 1
		}

		if name == "cybersecurity" || name == "network_security" || name == "penetration_testing" || name == "ethical_hacking" || name == "firewall" || name == "encryption" || name == "keamanan_siber" {
			security += 1
		}

		if name == "project_management" || name == "scrum" || name == "agile" || name == "kanban" || name == "pmp" || name == "jira" || name == "planning" || name == "manajemen_proyek" {
			management += 1
		}

		if name == "finance" || name == "accounting" || name == "budgeting" || name == "financial_analysis" || name == "tax" || name == "bookkeeping" || name == "akuntansi" || name == "keuangan" {
			finance += 1
		}

		if name == "customer_service" || name == "support" || name == "call_center" || name == "crm" || name == "client_relations" || name == "layanan_pelanggan" {
			customer_service += 1
		}

		if name == "game_dev" || name == "unity" || name == "unreal" || name == "gamemaker" || name == "godot" || name == "game_design" || name == "pengembang_game" {
			game += 1
		}

		if name == "video_editing" || name == "premiere" || name == "after_effects" || name == "davinci_resolve" || name == "final_cut" || name == "motion_graphics" || name == "editor_video" {
			video += 1
		}

		if name == "law" || name == "legal" || name == "paralegal" || name == "contracts" || name == "litigation" || name == "compliance" || name == "legal_research" || name == "hukum" || name == "pengacara" {
			law += 1
		}

		if name == "nursing" || name == "healthcare" || name == "patient_care" || name == "clinical" || name == "public_health" || name == "medical_assistant" || name == "perawat" || name == "kesehatan" {
			health += 1
		}

		if name == "admin" || name == "administration" || name == "data_entry" || name == "secretary" || name == "office_management" || name == "filing" || name == "admin_staf" || name == "sekretaris" {
			admin += 1
		}

		if name == "logistics" || name == "supply_chain" || name == "warehouse" || name == "inventory" || name == "transportation" || name == "procurement" || name == "logistik" || name == "gudang" || name == "pengiriman" {
			logistics += 1
		}

		if name == "human_resources" || name == "recruitment" || name == "talent_acquisition" || name == "payroll" || name == "hr_policy" || name == "training_development" || name == "hrd" || name == "sdm" {
			hr += 1
		}

		if name == "psychology" || name == "counseling" || name == "mental_health" || name == "therapy" || name == "social_work" || name == "psikologi" || name == "konselor" {
			psychology += 1
		}

		if name == "chef" || name == "cooking" || name == "kitchen" || name == "food" || name == "baking" || name == "catering" || name == "juru_masak" || name == "koki" {
			chef += 1
		}

		if name == "event_planning" || name == "event_management" || name == "wedding_planning" || name == "coordination" || name == "acara" || name == "event_organizer" {
			events += 1
		}

		if name == "kasir" || name == "cashier" || name == "penjaga_toko" {
			cashier += 1
		}

		if name == "operator_produksi" || name == "produksi" || name == "pekerja_pabrik" {
			operator += 1
		}

		if name == "barista" || name == "kopi" || name == "waiter" || name == "waitress" || name == "pelayan" {
			fnb += 1
		}

		if name == "content_creator" || name == "influencer" || name == "youtuber" || name == "tiktok" || name == "pembuat_konten" {
			creator += 1
		}

	}

	for i = 0; i < data.ExperienceCount; i++ {
		title = ToLower(data.Experiences[i].Title)
		company = ToLower(data.Experiences[i].Company)
		if company == "telkom_indonesia" || company == "bukalapak" || company == "tokopedia" || company == "shopee" || company == "gojek" ||
			company == "pertamina" || company == "bank_mandiri" || company == "bca" || company == "bni" || company == "indosat" || company == "xl" ||
			company == "astra_international" || company == "unilever_indonesia" || company == "indomie" || company == "indofood" || company == "mayora" ||
			company == "djarum" || company == "kalbe_farmina" || company == "mandiri_sekuritas" || company == "bpjs_kesehatan" || company == "bpjs_ketenagakerjaan" ||
			company == "bank_bri" || company == "pegadaian" || company == "wings_group" || company == "kimia_farma" || company == "blue_bird" || 
			company == "traveloka" || company == "tiket_com" || company == "pln" || company == "pelni" || company == "garuda_indonesia" || 
			company == "citilink" || company == "angkasapura" || company == "kai" || company == "pertani" || company == "bumn" || 
			company == "antarakita" || company == "sampoerna" || company == "gudang_garam" || company == "lippo_group" || company == "sinarmas" ||
			company == "pakuwon_group" || company == "agung_podomoro" || company == "summarecon" || company == "maspion" || company == "adaro" || 
			company == "freeport_indonesia" || company == "mnc_group" || company == "transcorp" || company == "net_tv" || company == "tvri" || 
			company == "kompas_gramedia" || company == "detik_com" || company == "tribunnews" || company == "kumparan" || company == "liputan6" {
			companyScore += 5
		} else if company != "" {
			companyScore += 1
		} else {
			companyScore += 0
		}

		if title == "backend_developer" || title == "backend_engineer" || title == "server_engineer" || title == "api_developer" {
			backend += 1
		}

		if title == "frontend_developer" || title == "frontend_engineer" || title == "web_developer" || title == "react_developer" || title == "vue_developer" {
			frontend += 1
		}

		if title == "fullstack_developer" || title == "fullstack_engineer" {
			fullstack += 1
		}

		if title == "mobile_developer" || title == "android_developer" || title == "ios_developer" || title == "flutter_developer" {
			mobile += 1
		}

		if title == "data_analyst" || title == "business_analyst" || title == "data_scientist" || title == "data_engineer" {
			analyst += 1
		}

		if title == "writer" || title == "content_writer" || title == "editor" || title == "copywriter" || title == "creative_writer" {
			writer += 1
		}

		if title == "teacher" || title == "lecturer" || title == "instructor" || title == "guru" || title == "pengajar" || title == "tutor" {
			teacher += 1
		}

		if title == "marketing_specialist" || title == "sales_executive" || title == "digital_marketer" || title == "seo_specialist" || title == "social_media_specialist" {
			marketing += 1
		}

		if title == "graphic_designer" || title == "ui_designer" || title == "ux_designer" || title == "visual_designer" || title == "motion_graphic_designer" {
			designer += 1
		}

		if title == "devops_engineer" || title == "site_reliability_engineer" || title == "infrastructure_engineer" {
			devops += 1
		}

		if title == "system_administrator" || title == "network_engineer" || title == "it_support" {
			it += 1
		}

		if title == "product_manager" || title == "project_manager" || title == "scrum_master" {
			management += 1
		}

		if title == "finance_staff" || title == "accountant" || title == "financial_analyst" || title == "auditor" {
			finance += 1
		}

		if title == "customer_service" || title == "call_center_agent" || title == "client_support" {
			customer_service += 1
		}

		if title == "hr_specialist" || title == "recruiter" || title == "hr_manager" {
			hr += 1
		}

		if title == "lawyer" || title == "legal_staff" || title == "legal_advisor" {
			law += 1
		}

		if title == "nurse" || title == "doctor" || title == "medical_assistant" || title == "healthcare_worker" {
			health += 1
		}

		if title == "admin_staff" || title == "administrative_assistant" || title == "secretary" {
			admin += 1
		}

		if title == "warehouse_staff" || title == "logistics_officer" || title == "supply_chain_specialist" || title == "delivery_staff" {
			logistics += 1
		}

		if title == "event_organizer" || title == "event_planner" || title == "event_coordinator" {
			events += 1
		}

		if title == "cashier" || title == "storekeeper" || title == "shop_assistant" {
			cashier += 1
		}

		if title == "production_operator" || title == "factory_worker" || title == "machine_operator" {
			operator += 1
		}

		if title == "barista" || title == "waiter" || title == "kitchen_staff" || title == "cook" || title == "chef" {
			fnb += 1
		}

		if title == "content_creator" || title == "influencer" || title == "youtuber" || title == "streamer" {
			creator += 1
		}

		if title == "psychologist" || title == "counselor" || title == "therapist" {
			psychology += 1
		}

		if title == "video_editor" || title == "videographer" || title == "motion_editor" {
			video += 1
		}

		if title == "game_developer" || title == "game_designer" || title == "game_programmer" {
			game += 1
		}
	}

	for j = 0; j < data.EducationCount; j++ {
		school = ToLower(data.Educations[j].School)
		degree = ToLower(data.Educations[j].Degree)
		major = ToLower(data.Educations[j].Major)

		if school == "universitas_indonesia" || school == "ui" ||
		school == "gadjah_mada_university" || school == "universitas_gadjah_mada" || school == "ugm" ||
		school == "bandung_institute_of_technology" || school == "institut_teknologi_bandung" || school == "itb" ||
		school == "airlangga_university" || school == "universitas_airlangga" || school == "unair" ||
		school == "ipb_university" || school == "bogor_agricultural_university" || school == "ipb" {
			educationScore += 5
		} else if school == "institut_teknologi_sepuluh_nopember" || school == "its" ||
				school == "universitas_padjadjaran" || school == "unpad" ||
				school == "diponegoro_university" || school == "universitas_diponegoro" || school == "undip" ||
				school == "universitas_brawijaya" || school == "ub" ||
				school == "bina_nusantara_university" || school == "binus" {
			educationScore += 3
		} else if school == "telkom_university" || school == "telkom" ||
				school == "universitas_hasanuddin" || school == "unhas" ||
				school == "universitas_sebelas_maret" || school == "uns" ||
				school == "atma_jaya_catholic_university_jakarta" || school == "atma_jaya" ||
				school == "universitas_islam_indonesia" || school == "uii" ||
				school == "universitas_muhammadiyah_yogyakarta" || school == "umy" ||
				school == "universitas_pendidikan_indonesia" || school == "upi" ||
				school == "universitas_sumatera_utara" || school == "usu" ||
				school == "yogyakarta_state_university" || school == "universitas_negeri_yogyakarta" || school == "uny" {
			educationScore += 2
		} else if school == "state_university_of_malang" || school == "universitas_negeri_malang" || school == "um" ||
				school == "udayana_university" || school == "universitas_udayana" ||
				school == "universitas_andalas" || school == "unand" ||
				school == "universitas_kristen_petra" || school == "ukp" ||
				school == "universitas_muhammadiyah_surakarta" || school == "ums" ||
				school == "university_of_lampung" || school == "universitas_lampung" || school == "unila" ||
				school == "university_of_mataram" || school == "universitas_mataram" || school == "unram" {
			educationScore += 1
		} else {
			educationScore += 0
		}

		if degree == "d1" || degree == "d2" || degree == "d3" {
			educationScore += 1
		} else if degree == "d4" || degree == "s1" {
			educationScore += 3
		} else if degree == "s2" {
			educationScore += 4
		} else if degree == "s3" {
			educationScore += 5
		} else {
			educationScore += 0
		}

		if major == "pendidikan" {
   			teacher += 1    
		} else if major == "hukum" {
			law += 1      
		} else if major == "ekonomi" {
			finance += 1    
		} else if major == "akuntansi" {
			finance += 1    
		} else if major == "manajemen" {
			management += 1 
		} else if major == "komputer" {
			backend += 1
			frontend += 1
			fullstack += 1
			analyst += 1    
		} else if major == "teknik" {
			operator += 1
			designer += 1  
		} else if major == "sains" {
			analyst += 1   
		} else if major == "sosial" {
			analyst += 1    
		} else if major == "seni" {
			designer += 1
			writer += 1     
		} else if major == "sastra" {
			writer += 1    
		} 
	}

	fmt.Println("\n═════════════════════════════════════════════════════════════════════════════")
fmt.Println(">> AI Suggestion Summary:")
fmt.Println("═════════════════════════════════════════════════════════════════════════════")

total = backend + frontend + analyst + writer + teacher + marketing + designer + devops + mobile +
    security + management + finance + customer_service + game + video + law + health + admin +
    logistics + hr + psychology + chef + events + cashier + operator + fnb + creator + educationScore

	if total == 0 {
		fmt.Println("[!!] No relevant suggestions found. Please input more profile data.")
	} else {
		fmt.Println("[!!] AI has analyzed your profile. Based on your skills, experience, and education level,")
		fmt.Println("[!!] here are the best career paths and recommended companies for you:\n")

		if backend >= 5 {
			fmt.Println("[Big Company] Backend Developer at Tokopedia, Gojek, Bukalapak")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", backend)
			fmt.Println("[••] Suggestion: Highlight your experience with scalable systems and microservices.\n")
		} else if backend >= 3 {
			fmt.Println("[Medium Company] Backend Developer at Ruangguru, Qasir, Mamikos")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", backend)
			fmt.Println("[••] Suggestion: Highlight your experience with scalable systems and microservices.\n")
		} else if backend > 0 {
			fmt.Println("[Startup] Backend Developer at LocalTech, NextDev")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", backend)
			fmt.Println("[••] Suggestion: Highlight your experience with scalable systems and microservices.\n")
		}

		if frontend >= 5 {
			fmt.Println("[Big Company] Frontend Developer at Traveloka, Blibli, Gojek")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", frontend)
			fmt.Println("[••] Suggestion: Improve UI/UX skills and use React/Vue in your portfolio.\n")
		} else if frontend >= 3 {
			fmt.Println("[Medium Company] Frontend Developer at KlikDokter, Sayurbox")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", frontend)
			fmt.Println("[••] Suggestion: Improve UI/UX skills and use React/Vue in your portfolio.\n")
		} else if frontend > 0 {
			fmt.Println("[Startup] Frontend Developer at CodeCraft, DevStart")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", frontend)
			fmt.Println("[••] Suggestion: Improve UI/UX skills and use React/Vue in your portfolio.\n")
		}

		// TODO: Data Analyst
		if analyst >= 5 {
			fmt.Println("[Big Company] Data Analyst at Telkom Indonesia, Shopee, Bank BCA")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", analyst)
			fmt.Println("[••] Suggestion: Practice SQL and dashboarding with Power BI or Tableau.\n")
		} else if analyst >= 3 {
			fmt.Println("[Medium Company] Data Analyst at Xendit, Sociolla")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", analyst)
			fmt.Println("[••] Suggestion: Practice SQL and dashboarding with Power BI or Tableau.\n")
		} else if analyst > 0 {
			fmt.Println("[Startup] Data Analyst at AnalytIQ, StatLab")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", analyst)
			fmt.Println("[••] Suggestion: Practice SQL and dashboarding with Power BI or Tableau.\n")
		}

		// TODO: Content Writer / Editor
		if writer >= 5 {
			fmt.Println("[Big Company] Content Writer / Editor at Kumparan, IDN Media, Zenius")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", writer)
			fmt.Println("[••] Suggestion: Build a portfolio and strengthen SEO writing.\n")
		} else if writer >= 3 {
			fmt.Println("[Medium Company] Content Writer / Editor at Hipwee, FemaleDaily")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", writer)
			fmt.Println("[••] Suggestion: Build a portfolio and strengthen SEO writing.\n")
		} else if writer > 0 {
			fmt.Println("[Startup] Content Writer / Editor at BlogHive, LocalStory")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", writer)
			fmt.Println("[••] Suggestion: Build a portfolio and strengthen SEO writing.\n")
		}

		// TODO: Teacher / Lecturer
		if teacher >= 5 {
			fmt.Println("[Big Company] Teacher / Lecturer at Binus, Universitas Indonesia, Ruangguru")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", teacher)
			fmt.Println("[••] Suggestion: Showcase your teaching materials and get certified.\n")
		} else if teacher >= 3 {
			fmt.Println("[Medium Company] Teacher / Lecturer at HarukaEdu, Zenius")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", teacher)
			fmt.Println("[••] Suggestion: Showcase your teaching materials and get certified.\n")
		} else if teacher > 0 {
			fmt.Println("[Startup] Teacher / Lecturer at KelasKita, EduStart")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", teacher)
			fmt.Println("[••] Suggestion: Showcase your teaching materials and get certified.\n")
		}

		// TODO: Marketing Specialist
		if marketing >= 5 {
			fmt.Println("[Big Company] Marketing Specialist at Grab, Shopee, TikTok Indonesia")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", marketing)
			fmt.Println("[••] Suggestion: Demonstrate social media campaigns and analytics knowledge.\n")
		} else if marketing >= 3 {
			fmt.Println("[Medium Company] Marketing Specialist at Evermos, Sirclo")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", marketing)
			fmt.Println("[••] Suggestion: Demonstrate social media campaigns and analytics knowledge.\n")
		} else if marketing > 0 {
			fmt.Println("[Startup] Marketing Specialist at PromoPilot, StartSell")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", marketing)
			fmt.Println("[••] Suggestion: Demonstrate social media campaigns and analytics knowledge.\n")
		}

		// TODO: Graphic Designer
		if designer >= 5 {
			fmt.Println("[Big Company] Graphic Designer at Tokopedia, Bukalapak, Netflix ID")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", designer)
			fmt.Println("[••] Suggestion: Polish your Behance/Dribbble portfolio.\n")
		} else if designer >= 3 {
			fmt.Println("[Medium Company] Graphic Designer at Jakmall, Warung Pintar")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", designer)
			fmt.Println("[••] Suggestion: Polish your Behance/Dribbble portfolio.\n")
		} else if designer > 0 {
			fmt.Println("[Startup] Graphic Designer at PixelCraft, LokalDesign")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", designer)
			fmt.Println("[••] Suggestion: Polish your Behance/Dribbble portfolio.\n")
		}

		// TODO: DevOps Engineer
		if devops >= 5 {
			fmt.Println("[Big Company] DevOps Engineer at Blibli, Gojek, JNE")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", devops)
			fmt.Println("[••] Suggestion: Demonstrate CI/CD pipelines and infrastructure as code.\n")
		} else if devops >= 3 {
			fmt.Println("[Medium Company] DevOps Engineer at Alterra, Nodeflux")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", devops)
			fmt.Println("[••] Suggestion: Demonstrate CI/CD pipelines and infrastructure as code.\n")
		} else if devops > 0 {
			fmt.Println("[Startup] DevOps Engineer at InfraEdge, LiteOps")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", devops)
			fmt.Println("[••] Suggestion: Demonstrate CI/CD pipelines and infrastructure as code.\n")
		}

		// TODO: Mobile App Developer
		if mobile >= 5 {
			fmt.Println("[Big Company] Mobile App Developer at Traveloka, TikTok, Tokopedia")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", mobile)
			fmt.Println("[••] Suggestion: Build Flutter/React Native projects.\n")
		} else if mobile >= 3 {
			fmt.Println("[Medium Company] Mobile App Developer at Kulina, MySkill")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", mobile)
			fmt.Println("[••] Suggestion: Build Flutter/React Native projects.\n")
		} else if mobile > 0 {
			fmt.Println("[Startup] Mobile App Developer at AppHero, NextMob")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", mobile)
			fmt.Println("[••] Suggestion: Build Flutter/React Native projects.\n")
		}

		// TODO: Cybersecurity Specialist
		if security >= 5 {
			fmt.Println("[Big Company] Cybersecurity Specialist at Bank Indonesia, Tokopedia, Kominfo")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", security)
			fmt.Println("[••] Suggestion: Earn cybersecurity certifications and join CTFs.\n")
		} else if security >= 3 {
			fmt.Println("[Medium Company] Cybersecurity Specialist at Vaksincom, Privy ID")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", security)
			fmt.Println("[••] Suggestion: Earn cybersecurity certifications and join CTFs.\n")
		} else if security > 0 {
			fmt.Println("[Startup] Cybersecurity Specialist at SecuStart, ByteSecure")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", security)
			fmt.Println("[••] Suggestion: Earn cybersecurity certifications and join CTFs.\n")
		}

		// TODO: Project Manager
		if management >= 5 {
			fmt.Println("[Big Company] Project Manager at Telkomsel, Pertamina, BRI")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", management)
			fmt.Println("[••] Suggestion: Highlight project planning, leadership, and agile methods.\n")
		} else if management >= 3 {
			fmt.Println("[Medium Company] Project Manager at HappyFresh, Pinhome")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", management)
			fmt.Println("[••] Suggestion: Highlight project planning, leadership, and agile methods.\n")
		} else if management > 0 {
			fmt.Println("[Startup] Project Manager at ManageOne, AgileStudio")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", management)
			fmt.Println("[••] Suggestion: Highlight project planning, leadership, and agile methods.\n")
		}

		// TODO: Financial Analyst
		if finance >= 5 {
			fmt.Println("[Big Company] Financial Analyst at Bank Mandiri, OJK, PwC")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", finance)
			fmt.Println("[••] Suggestion: Sharpen Excel and financial modeling skills\n")
		} else if finance >= 3 {
			fmt.Println("[Medium Company] Financial Analyst at Bibit, TernakUang")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", finance)
			fmt.Println("[••] Suggestion: Sharpen Excel and financial modeling skills\n")
		} else if finance > 0 {
			fmt.Println("[Startup] Financial Analyst at FinPilot, BudgetBuddy")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", finance)
			fmt.Println("[••] Suggestion: Sharpen Excel and financial modeling skills\n")
		}

		// TODO: Customer Service
		if customer_service >= 5 {
			fmt.Println("[Big Company] Customer Service at Shopee, Tokopedia, Traveloka")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", customer_service)
			fmt.Println("[••] Suggestion: Improve communication skills and empathy.\n")
		} else if customer_service >= 3 {
			fmt.Println("[Medium Company] Customer Service at Sirclo, Evermos")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", customer_service)
			fmt.Println("[••] Suggestion: Improve communication skills and empathy.\n")
		} else if customer_service > 0 {
			fmt.Println("[Startup] Customer Service at HelpLine, FastDesk")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", customer_service)
			fmt.Println("[••] Suggestion: Improve communication skills and empathy.\n")
		}

		// TODO: Game Developer
		if game >= 5 {
			fmt.Println("[Big Company] Game Developer at ArenaNet ID, Agate Studio, Garena")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", game)
			fmt.Println("[••] Suggestion: Master Unity or Unreal and build small games.\n")
		} else if game >= 3 {
			fmt.Println("[Medium Company] Game Developer at Toge Productions, Digital Happiness")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", game)
			fmt.Println("[••] Suggestion: Master Unity or Unreal and build small games.\n")
		} else if game > 0 {
			fmt.Println("[Startup] Game Developer at PixelPlay, IndieDev")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", game)
			fmt.Println("[••] Suggestion: Master Unity or Unreal and build small games.\n")
		}

		//TODO:  Video Editor
		if video >= 5 {
			fmt.Println("[Big Company] Video Editor at NET TV, Narasi, IDN Media")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", video)
			fmt.Println("[••] Suggestion: Create a showreel and learn motion graphics tools.\n")
		} else if video >= 3 {
			fmt.Println("[Medium Company] Video Editor at Kreator.ID, LokalTV")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", video)
			fmt.Println("[••] Suggestion: Create a showreel and learn motion graphics tools.\n")
		} else if video > 0 {
			fmt.Println("[Startup] Video Editor at VideoVibe, ClipStart")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", video)
			fmt.Println("[••] Suggestion: Create a showreel and learn motion graphics tools.\n")
		}

		//TODO:  Legal Advisor
		if law >= 5 {
			fmt.Println("[Big Company] Legal Advisor at Mahkamah Agung, Kantor Hukum Ternama")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", law)
			fmt.Println("[••] Suggestion: Deepen knowledge of regulations and legal writing.\n")
		} else if law >= 3 {
			fmt.Println("[Medium Company] Legal Advisor at Startup Hukum Lokal")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", law)
			fmt.Println("[••] Suggestion: Deepen knowledge of regulations and legal writing.\n")
		} else if law > 0 {
			fmt.Println("[Startup] Legal Advisor at LegalEase, LawEntry")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", law)
			fmt.Println("[••] Suggestion: Deepen knowledge of regulations and legal writing.\n")
		}

		//TODO: Healthcare Worker
		if health >= 5 {
			fmt.Println("[Big Company] Healthcare Worker at RSUPN, Siloam, Halodoc")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", health)
			fmt.Println("[••] Suggestion: Showcase certifications and patient care experience.\n")
		} else if health >= 3 {
			fmt.Println("[Medium Company] Healthcare Worker at Good Doctor, ProSehat")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", health)
			fmt.Println("[••] Suggestion: Showcase certifications and patient care experience.\n")
		} else if health > 0 {
			fmt.Println("[Startup] Healthcare Worker at HealthStart, KlinikKita")
			fmt.Printf("[!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", health)
			fmt.Println("[••] Suggestion: Showcase certifications and patient care experience.\n")
		}

		//TODO: Administrative Staff
		if admin >= 5 {
			fmt.Println("[Big Company] Administrative Staff at Pertamina, PLN, Bank Mandiri")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", admin)
		fmt.Println("[••] Suggestion: Strengthen MS Office and organizational skills\n")
		} else if admin >= 3 {
		fmt.Println("[Medium Company] Administrative Staff at DANA, OVO")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", admin)
		fmt.Println("[••] Suggestion: Strengthen MS Office and organizational skills\n")
		} else if admin > 0 {
		fmt.Println("[Startup] Administrative Staff at OfficeHub, AdminPro")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", admin)
		fmt.Println("[••] Suggestion: Strengthen MS Office and organizational skills\n")
		}

		//TODO: Logistics Coordinator
		if logistics >= 5 {
		fmt.Println("[Big Company] Logistics Coordinator at JNE, TIKI, Gojek")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", logistics)
		fmt.Println("[••] Suggestion: Focus on supply chain management skills\n")
		} else if logistics >= 3 {
		fmt.Println("[Medium Company] Logistics Coordinator at Deliveree, Lalamove")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", logistics)
		fmt.Println("[••] Suggestion: Focus on supply chain management skills\n")
		} else if logistics > 0 {
		fmt.Println("[Startup] Logistics Coordinator at FastLog, QuickShip")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", logistics)
		fmt.Println("[••] Suggestion: Focus on supply chain management skills\n")
		}

		//TODO:  HR Specialist
		if hr >= 5 {
		fmt.Println("[Big Company] HR Specialist at Unilever, Telkom Indonesia, Danone")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", hr)
		fmt.Println("[••] Suggestion: Develop interview and employee relations skills\n")
		} else if hr >= 3 {
		fmt.Println("[Medium Company] HR Specialist at Modalku, Kredivo")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", hr)
		fmt.Println("[••] Suggestion: Develop interview and employee relations skills\n")
		} else if hr > 0 {
		fmt.Println("[Startup] HR Specialist at PeopleOps, HRStart")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", hr)
		fmt.Println("[••] Suggestion: Develop interview and employee relations skills\n")
		}

		//TODO:  Psychologist / Counselor
		if psychology >= 5 {
		fmt.Println("[Big Company] Psychologist / Counselor at RSJ, Klinik Pratama, Sekolah")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", psychology)
		fmt.Println("[••] Suggestion: Highlight counseling certifications and experience.\n")
		} else if psychology >= 3 {
		fmt.Println("[Medium Company] Psychologist / Counselor at Klinik Prima, Sekolah Alam")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", psychology)
		fmt.Println("[••] Suggestion: Highlight counseling certifications and experience.\n")
		} else if psychology > 0 {
		fmt.Println("[Startup] Psychologist / Counselor at MindCare, SafeTalk")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", psychology)
		fmt.Println("[••] Suggestion: Highlight counseling certifications and experience.\n")
		}

		//TODO: Chef / Cook
		if chef >= 5 {
		fmt.Println("[Big Company] Chef / Cook at Hotel Mulia, Ritz Carlton, Local Restaurant Chain")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", chef)
		fmt.Println("[••] Suggestion: Showcase culinary certifications and special dishes.\n")
		} else if chef >= 3 {
		fmt.Println("[Medium Company] Chef / Cook at Hotel Santika, Local Bistro")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", chef)
		fmt.Println("[••] Suggestion: Showcase culinary certifications and special dishes.\n")
		} else if chef > 0 {
		fmt.Println("[Startup] Chef / Cook at FoodieStart, HomeKitchen")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", chef)
		fmt.Println("[••] Suggestion: Showcase culinary certifications and special dishes.\n")
		}

		//TODO:  Event Organizer
		if events >= 5 {
		fmt.Println("[Big Company] Event Organizer at Dyandra, Rajawali Corpora")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", events)
		fmt.Println("[••] Suggestion: Show event portfolio and negotiation skills\n")
		} else if events >= 3 {
		fmt.Println("[Medium Company] Event Organizer at EO Lokal, PartyUp")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", events)
		fmt.Println("[••] Suggestion: Show event portfolio and negotiation skills\n")
		} else if events > 0 {
		fmt.Println("[Startup] Event Organizer at SmallEvents, FestStart")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", events)
		fmt.Println("[••] Suggestion: Show event portfolio and negotiation skills\n")
		}

		//TODO: Cashier
		if cashier >= 5 {
		fmt.Println("[Big Company] Cashier at Indomaret, Alfamart, Circle K")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", cashier)
		fmt.Println("[••] Suggestion: Focus on accuracy and customer service skills\n")
		} else if cashier >= 3 {
		fmt.Println("[Medium Company] Cashier at Minimart Lokal, Toko Sembako")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", cashier)
		fmt.Println("[••] Suggestion: Focus on accuracy and customer service skills\n")
		} else if cashier > 0 {
		fmt.Println("[Startup] Cashier at LocalStore, QuickMart")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", cashier)
		fmt.Println("[••] Suggestion: Focus on accuracy and customer service skills\n")
		}

		//TODO: Machine Operator
		if operator >= 5 {
		fmt.Println("[Big Company] Machine Operator at Indofood, Unilever, Astra")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", operator)
		fmt.Println("[••] Suggestion: Emphasize machine handling certifications.\n")
		} else if operator >= 3 {
		fmt.Println("[Medium Company] Machine Operator at LocalFactory, PlantX")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", operator)
		fmt.Println("[••] Suggestion: Emphasize machine handling certifications.\n")
		} else if operator > 0 {
		fmt.Println("[Startup] Machine Operator at SmallPlant, MachPro")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", operator)
		fmt.Println("[••] Suggestion: Emphasize machine handling certifications.\n")
		}

		//TODO: Food & Beverage Staff section
		if fnb >= 5 {
		fmt.Println("[Big Company] Food & Beverage Staff at Hotel Indonesia Kempinski, Marriott")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", fnb)
		fmt.Println("[••] Suggestion: Highlight service skills and certifications.\n")
		} else if fnb >= 3 {
		fmt.Println("[Medium Company] Food & Beverage Staff at Local Cafe, Bistro X")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", fnb)
		fmt.Println("[••] Suggestion: Highlight service skills and certifications.\n")
		} else if fnb > 0 {
		fmt.Println("[Startup] Food & Beverage Staff at QuickBite, FoodStart")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", fnb)
		fmt.Println("[••] Suggestion: Highlight service skills and certifications.\n")
		}

		//TODO: Content Creator / Influencer section
		if creator >= 5 {
		fmt.Println("[Big Company] Content Creator / Influencer at YouTube, TikTok, Instagram")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 75%% - 100%%\n", creator)
		fmt.Println("[••] Suggestion: Build consistent content and follower base.\n")
		} else if creator >= 3 {
		fmt.Println("[Medium Company] Content Creator / Influencer at Local Media, Blog")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 50%% - 75%%\n", creator)
		fmt.Println("[••] Suggestion: Build consistent content and follower base.\n")
		} else if creator > 0 {
		fmt.Println("[Startup] Content Creator / Influencer at IndieVlog, MiniBlog")
		fmt.Printf("     [!!!] Match Score: %d/10 | Acceptance Chance: 15%% - 50%%\n", creator)
		fmt.Println("[••] Suggestion: Build consistent content and follower base.\n")
		}

		fmt.Println("═════════════════════════════════════════════════════════════════════════════")
		fmt.Println("[INFO] These are AI-driven suggestions. Keep improving your profile,")
		fmt.Println("       and Career-Edge will update your opportunities in real-time!")
		fmt.Println("═════════════════════════════════════════════════════════════════════════════")
	}
}
