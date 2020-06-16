type EftelingData struct {
	LastUpdate      string `json:"lastUpdate"`
	MaintenanceInfo struct {
		Monsieurcannibale struct {
			From          string `json:"from"`
			End           string `json:"end"`
			OpenInWeekend bool   `json:"openInWeekend"`
		} `json:"monsieurcannibale"`
	} `json:"maintenanceInfo"`
	Park struct {
		OpenFrom       string `json:"openFrom"`
		OpenUntil      string `json:"openUntil"`
		BusyIndication string `json:"busyIndication"`
	} `json:"park"`
	Rides struct {
		Harmonievandrie struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"harmonievandrie"`
		Ontmoeteftelingbewoners struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"ontmoeteftelingbewoners"`
		Sneeuwprinsesjelkaenvuurprinskendrik struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"sneeuwprinsesjelkaenvuurprinskendrik"`
		Eenwonderlijkeontmoetingmetruiterthomas struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"eenwonderlijkeontmoetingmetruiterthomas"`
		Pardoesdetovernar struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"pardoesdetovernar"`
		Fotopuntsymbolica struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"fotopuntsymbolica"`
		Lichtpuntjessteenbokplein struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"lichtpuntjessteenbokplein"`
		Sprookjesboomzingendansmee struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"sprookjesboomzingendansmee"`
		Jrgenfreilich struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"jrgenfreilich"`
		Eftelingkidsradiolive struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"eftelingkidsradiolive"`
		Wintereftelingkinderkoor struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"wintereftelingkinderkoor"`
		Dezeszwanenrondvaart struct {
			Type              string `json:"type"`
			WaitingTime       int    `json:"waitingTime"`
			WaitingPercentage int    `json:"waitingPercentage"`
			AttractionType    string `json:"attractionType"`
			State             string `json:"State"`
		} `json:"dezeszwanenrondvaart"`
		Maxenmoritzmusiker struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"maxenmoritzmusiker"`
		Desmaeckmaker struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"desmaeckmaker"`
		Restaurantfabula struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"restaurantfabula"`
		Fabulasouvenirs struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"fabulasouvenirs"`
		Lichtpuntjesoudennieuw struct {
			Type string `json:"type"`
		} `json:"lichtpuntjesoudennieuw"`
		Harmonievandrieoudennieuw struct {
			Type string `json:"type"`
		} `json:"harmonievandrieoudennieuw"`
		Pardoesoudennieuw struct {
			Type string `json:"type"`
		} `json:"pardoesoudennieuw"`
		Vertelkabouters struct {
			Type string `json:"type"`
		} `json:"vertelkabouters"`
		Debounties struct {
			Type string `json:"type"`
		} `json:"debounties"`
		Thisisbeethoven struct {
			Type string `json:"type"`
		} `json:"thisisbeethoven"`
		Eftelingmuzikantenoudennieuw struct {
			Type string `json:"type"`
		} `json:"eftelingmuzikantenoudennieuw"`
		Muzikaalentertainment struct {
			Type string `json:"type"`
		} `json:"muzikaalentertainment"`
		Silentdisco struct {
			Type string `json:"type"`
		} `json:"silentdisco"`
		Vuurwerkplaatsblauw struct {
			Type string `json:"type"`
		} `json:"vuurwerkplaatsblauw"`
		Vuurwerkplaatsrood struct {
			Type string `json:"type"`
		} `json:"vuurwerkplaatsrood"`
		Caro struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"caro"`
		Hetwapenvanraveleijn struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"hetwapenvanraveleijn"`
		Hetwapenvanraveleijnlunch struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"hetwapenvanraveleijnlunch"`
		Kleyneklaroen struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"kleyneklaroen"`
		Kogeloog struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"kogeloog"`
		Denguldengaarde struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"denguldengaarde"`
		Eigenheymertonvandevenplein struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"eigenheymertonvandevenplein"`
		Unoxkraamtonvandevenplein struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"unoxkraamtonvandevenplein"`
		Hoorndesovervloeds struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"hoorndesovervloeds"`
		Smulpaap struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"smulpaap"`
		Kinderspoor struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"kinderspoor"`
		Jokieswereld struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"jokieswereld"`
		Sprookjesboomshow struct {
			Type      string `json:"type"`
			ShowTimes []struct {
				ShowDateTime string `json:"ShowDateTime"`
			} `json:"showTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Showduration   string   `json:"showduration"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"sprookjesboomshow"`
		Wafelsalondensuykerbuyk struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"wafelsalondensuykerbuyk"`
		Foodlaan struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"foodlaan"`
		Pizzachalet struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"pizzachalet"`
		Unoxkraamvogelrokplein struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"unoxkraamvogelrokplein"`
		Hetwittepaard struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"hetwittepaard"`
		Happinessstationwp struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"happinessstationwp"`
		Tpoffertje struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"tpoffertje"`
		Gelaarsdekat struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"gelaarsdekat"`
		Grootmoederskeuken struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"grootmoederskeuken"`
		Gamegallery struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Properties   []string `json:"properties"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"gamegallery"`
		Halvemaen struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"halvemaen"`
		Sprookjessprokkelaar struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"sprookjessprokkelaar"`
		Deverleiding struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Properties   []string `json:"properties"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"deverleiding"`
		Hetsuykerhuys struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"hetsuykerhuys"`
		Panoramalacarte struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"panoramalacarte"`
		Panoramaselfservice struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"panoramaselfservice"`
		Panoramakiosk struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"panoramakiosk"`
		Korfje struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"korfje"`
		Doudetuffer struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"doudetuffer"`
		Polleskeuken struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"polleskeuken"`
		Rondjevandemolen struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"rondjevandemolen"`
		Degebrandeboon struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"degebrandeboon"`
		Happinessstationpk struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"happinessstationpk"`
		Tokopagode struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"tokopagode"`
		Stationdeoost struct {
			Type           string `json:"type"`
			OpeningTimes   string `json:"openingTimes"`
			AttractionType string `json:"attractionType"`
		} `json:"stationdeoost"`
		Debazaar struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"debazaar"`
		Stoomcarrousel struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"stoomcarrousel"`
		Wachtruimteeersteklas struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"wachtruimteeersteklas"`
		Derustendereiziger struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"derustendereiziger"`
		Debrutaleaap struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"debrutaleaap"`
		Deverseoogst struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"deverseoogst"`
		Dehongerigemachinist struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"dehongerigemachinist"`
		Polkamarina struct {
			Type              string   `json:"type"`
			WaitingTime       int      `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"polkamarina"`
		Eigenheymerbijstationdeoost struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"eigenheymerbijstationdeoost"`
		Delikkebaerd struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"delikkebaerd"`
		Souvenirafhaalpunt struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"souvenirafhaalpunt"`
		Stoomtreinr struct {
			Type              string   `json:"type"`
			WaitingTime       int      `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"stoomtreinr"`
		Efteldingen struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"efteldingen"`
		Dekombuys struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"dekombuys"`
		Unoxkraamvliegendehollanderplein struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"unoxkraamvliegendehollanderplein"`
		Demeermin struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"demeermin"`
		Stoomtreinm struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"stoomtreinm"`
		Flierefluiter struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"flierefluiter"`
		Melkhuysje struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"melkhuysje"`
		Yoghurtbarbijbaron1898 struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"yoghurtbarbijbaron1898"`
		Hetseylendfregat struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"hetseylendfregat"`
		Unoxkraampiraaplein struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"unoxkraampiraaplein"`
		Wittewalvis struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"wittewalvis"`
		Indenoudenmarskramer struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"indenoudenmarskramer"`
		Python struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			AlternateID       string   `json:"alternateId"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"python"`
		Casacaracol struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"casacaracol"`
		Piranabar struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"piranabar"`
		Pythonsinglerider struct {
			Type              string `json:"type"`
			WaitingTime       int    `json:"waitingTime"`
			WaitingPercentage int    `json:"waitingPercentage"`
			AttractionType    string `json:"attractionType"`
			State             string `json:"State"`
		} `json:"pythonsinglerider"`
		Vreugdevuurruigrijk struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"vreugdevuurruigrijk"`
		Vreugdevuurmarerijk struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"vreugdevuurmarerijk"`
		Vreugdevuurreizenrijk struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"vreugdevuurreizenrijk"`
		Droomvlucht struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"droomvlucht"`
		Virtueledroomvlucht struct {
			Type              string `json:"type"`
			WaitingTime       int    `json:"waitingTime"`
			WaitingPercentage int    `json:"waitingPercentage"`
			AttractionType    string `json:"attractionType"`
			State             string `json:"State"`
		} `json:"virtueledroomvlucht"`
		Burgerbackerij struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"burgerbackerij"`
		Stroopwafelchalet struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"stroopwafelchalet"`
		Eigenheymersteenbokplein struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"eigenheymersteenbokplein"`
		Braadworstchalet struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"braadworstchalet"`
		Oase struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"oase"`
		Raveleijn struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"raveleijn"`
		Devliegendehollander struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			AlternateID       string   `json:"alternateId"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"devliegendehollander"`
		Dromerijen struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"dromerijen"`
		Devliegendehollandersinglerider struct {
			Type              string `json:"type"`
			WaitingTime       int    `json:"waitingTime"`
			WaitingPercentage int    `json:"waitingPercentage"`
			AttractionType    string `json:"attractionType"`
			State             string `json:"State"`
		} `json:"devliegendehollandersinglerider"`
		Octopus struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"octopus"`
		Hollandsegebakskraam struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"hollandsegebakskraam"`
		Devrolijkenoot struct {
			Type           string   `json:"type"`
			OpeningTimes   string   `json:"openingTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"devrolijkenoot"`
		Jorisendedraak struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			AlternateID       string   `json:"alternateId"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"jorisendedraak"`
		Loetiek struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"loetiek"`
		Jorisendedraaksinglerider struct {
			Type              string `json:"type"`
			WaitingTime       int    `json:"waitingTime"`
			WaitingPercentage int    `json:"waitingPercentage"`
			AttractionType    string `json:"attractionType"`
			State             string `json:"State"`
		} `json:"jorisendedraaksinglerider"`
		Villavolta struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"villavolta"`
		Baron1898 struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			AlternateID       string   `json:"alternateId"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"baron1898"`
		Baron1898Singlerider struct {
			Type              string `json:"type"`
			WaitingTime       int    `json:"waitingTime"`
			WaitingPercentage int    `json:"waitingPercentage"`
			AttractionType    string `json:"attractionType"`
			State             string `json:"State"`
		} `json:"baron1898singlerider"`
		Carnavalfestival struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"carnavalfestival"`
		Volkvanlaafmonorail struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"volkvanlaafmonorail"`
		Jokieenjet struct {
			Type      string `json:"type"`
			ShowTimes []struct {
				ShowDateTime string `json:"ShowDateTime"`
			} `json:"showTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Showduration   string   `json:"showduration"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"jokieenjet"`
		Volkvanlaaf struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"volkvanlaaf"`
		Carrouselsantonpieckplein struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"carrouselsantonpieckplein"`
		Monsieurcannibale struct {
			Type              string   `json:"type"`
			WaitingTime       int      `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
			Maintenance       struct {
				From          string `json:"from"`
				End           string `json:"end"`
				OpenInWeekend bool   `json:"openInWeekend"`
			} `json:"maintenance"`
		} `json:"monsieurcannibale"`
		Vogelrok struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"vogelrok"`
		Fotopuntpiraa struct {
			Type string `json:"type"`
		} `json:"fotopuntpiraa"`
		Fotopuntdevliegendehollander struct {
			Type string `json:"type"`
		} `json:"fotopuntdevliegendehollander"`
		Fotopuntjorisendedraak struct {
			Type string `json:"type"`
		} `json:"fotopuntjorisendedraak"`
		Fotopuntpython struct {
			Type string `json:"type"`
		} `json:"fotopuntpython"`
		Gondoletta struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"gondoletta"`
		Pagode struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"pagode"`
		Fotopuntbaron1898 struct {
			Type string `json:"type"`
		} `json:"fotopuntbaron1898"`
		Fabula struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"fabula"`
		Spookslot struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"spookslot"`
		Pirana struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"pirana"`
		Fatamorgana struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"fatamorgana"`
		Sprookjesboomerwaseens struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"sprookjesboomerwaseens"`
		Symbolica struct {
			Type              string   `json:"type"`
			WaitingTime       string   `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			AlternateID       string   `json:"alternateId"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Properties        []string `json:"properties"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"symbolica"`
		Symbolicasingleriders struct {
			Type              string `json:"type"`
			WaitingTime       int    `json:"waitingTime"`
			WaitingPercentage int    `json:"waitingPercentage"`
			AttractionType    string `json:"attractionType"`
			State             string `json:"State"`
		} `json:"symbolicasingleriders"`
		Aquanura struct {
			Type      string `json:"type"`
			ShowTimes []struct {
				ShowDateTime string `json:"ShowDateTime"`
				Edition      string `json:"Edition"`
			} `json:"showTimes"`
			AttractionType string   `json:"attractionType"`
			TextNl         string   `json:"text_nl"`
			DetailTextNl   string   `json:"detail_text_nl"`
			Area           string   `json:"area"`
			Properties     []string `json:"properties"`
			Showduration   string   `json:"showduration"`
			Location       string   `json:"location"`
			Photos         []string `json:"photos"`
			TextEn         string   `json:"text_en"`
			DetailTextEn   string   `json:"detail_text_en"`
		} `json:"aquanura"`
		Eftelingmuziekanten struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"eftelingmuziekanten"`
		Lichtpuntjes struct {
			Type           string        `json:"type"`
			ShowTimes      []interface{} `json:"showTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"lichtpuntjes"`
		Sprookjesboom struct {
			Type              string   `json:"type"`
			WaitingTime       int      `json:"waitingTime"`
			WaitingPercentage int      `json:"waitingPercentage"`
			AttractionType    string   `json:"attractionType"`
			State             string   `json:"State"`
			TextNl            string   `json:"text_nl"`
			DetailTextNl      string   `json:"detail_text_nl"`
			Area              string   `json:"area"`
			Location          string   `json:"location"`
			Photos            []string `json:"photos"`
			TextEn            string   `json:"text_en"`
			DetailTextEn      string   `json:"detail_text_en"`
		} `json:"sprookjesboom"`
		Theaterrestaurantapplaus struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"theaterrestaurantapplaus"`
		Carrouselpaleiss struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
		} `json:"carrouselpaleiss"`
		Carrouselbar struct {
			Type           string        `json:"type"`
			OpeningTimes   []interface{} `json:"openingTimes"`
			AttractionType string        `json:"attractionType"`
			TextNl         string        `json:"text_nl"`
			DetailTextNl   string        `json:"detail_text_nl"`
			Area           string        `json:"area"`
			Properties     []string      `json:"properties"`
			Location       string        `json:"location"`
			Photos         []string      `json:"photos"`
			TextEn         string        `json:"text_en"`
			DetailTextEn   string        `json:"detail_text_en"`
		} `json:"carrouselbar"`
		Desoeteinval struct {
			Type         string   `json:"type"`
			TextNl       string   `json:"text_nl"`
			DetailTextNl string   `json:"detail_text_nl"`
			Area         string   `json:"area"`
			Properties   []string `json:"properties"`
			Location     string   `json:"location"`
			Photos       []string `json:"photos"`
			TextEn       string   `json:"text_en"`
			DetailTextEn string   `json:"detail_text_en"`
		} `json:"desoeteinval"`
	} `json:"rides"`
	Tags []string `json:"tags"`
}
