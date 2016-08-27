package metawater_edu_proejct

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func init() {
	http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "%s", fmt.Sprintf("%s %s", r.Method, r.URL.Path))
	w.Write([]byte(fmt.Sprintf("%s %s", r.Method, r.URL.Path)))
}
