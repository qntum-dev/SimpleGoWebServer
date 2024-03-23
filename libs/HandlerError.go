package libs

import "net/http"

func HandleErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something Went Wrong")
}
