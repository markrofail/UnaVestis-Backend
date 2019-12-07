package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"hello\":\"world\"}"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

//func main() {
//	kinds := hm.GetAllKinds()
//	chosenKind := kinds[0]
//	fmt.Println(chosenKind)
//
//	typesArr := hm.GetItemTypes(chosenKind)
//	chosenType := typesArr[0]
//	fmt.Println(chosenType)
//
//	requestUrl := buildUrl(chosenKind, chosenType)
//	testPage := hm.GetPage(chosenKind, chosenType)
//	fmt.Println(testPage)
//
//	rawJson := getJson(requestUrl)
//	requestData, err := extractData(rawJson)
//	if err != nil {
//		log.Println(err)
//	}
//
//	products := getItems(requestData)
//
//	//jsonResponse, err := json.Marshal(products)
//	//if err != nil {
//	//	log.Println(err)
//	//}
//
//	pretty.Println(products)
//}
