package handlers

// func Artist(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		ErrorHandler(w, http.StatusMethodNotAllowed)
// 		return
// 	}
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil {
// 		ErrorHandler(w, http.StatusBadRequest)
// 		return
// 	}
// 	if !valid_url(id) {
// 		ErrorHandler(w, http.StatusNotFound)
// 		return
// 	}
// 	// re := regexp.MustCompile("/artist/(\\d+)$")
// 	// matches := re.FindStringSubmatch(r.URL.Path)
// 	tmpl, err := template.ParseFiles("templates/artist.html")
// 	if err != nil {
// 		log.Println("Template parsing error")
// 		ErrorHandler(w, http.StatusInternalServerError)
// 		return
// 	}
// 	err = tmpl.Execute(w, datas[id-1])
// 	if err != nil {
// 		ErrorHandler(w, http.StatusInternalServerError)
// 		return
// 	}
// }

// func valid_url(id int) bool {
// 	if !(id > 0 && id <= len(datas)) {
// 		return false
// 	}
// 	return true
// }

// func get_raw(id int, w http.ResponseWriter, r *http.Response) {
// 	jsonData, err := json.Marshal(datas[id-1])
// 	if err != nil {
// 		log.Println("Marshal error")
// 		ErrorHandler(w, http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Content-Disposition", "inline;")
// 	_, err = w.Write(jsonData)
// 	if err != nil {
// 		fmt.Fprintf(w, "Something went wrong!")
// 		return
// 	}
// }
