package rank

import (
	"encoding/json"
	"log"
	"net/http"
	"ranking/github.com/gorilla/mux"
	"ranking/pods"
	"ranking/sorts"
)

// fa un sort dei nuovi podcast
//
// Si puo usare nel contesto del User
//    per prendere dei podcast nuovi
//    (non necessariamente popolari)
//	  che rientrano nelle categorie preferite
//    e si possono mettere in ordine sul profilo
func rankNewPodcasts(w http.ResponseWriter, r *http.Request)  {
	var podcast []pods.Pod

	// lista di podcast presa dal body del request
	json.NewDecoder(r.Body).Decode(&podcast)

	sorts.RadixSort(podcast, len(podcast)) // si sortano

	json.NewEncoder(w).Encode(podcast) // si manda indietro la lista
}

// inserisce i nuovi podcast (sortati) nella classifica (globale o del user)
func updateRankings(w http.ResponseWriter, r *http.Request)  {
	var reqBody [][]pods.Pod
	var currentRankings []pods.Pod
	var podcastToUpdate []pods.Pod
	var newRanking []pods.Pod

	json.NewDecoder(r.Body).Decode(&reqBody)

	currentRankings = reqBody[0]
	podcastToUpdate = reqBody[1]

	newRanking = sorts.MergeSort(currentRankings, podcastToUpdate)

	json.NewEncoder(w).Encode(newRanking)
}

func RunRankService() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/rankNewEntries", rankNewPodcasts)
	router.HandleFunc("/rankUpdate", updateRankings)
	log.Fatal(http.ListenAndServe(":3000", router))
}