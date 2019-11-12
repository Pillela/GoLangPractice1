package Math_Functions

import (
	"fmt"
	"net/http"
	)

func Index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Testing basic http!")
}
