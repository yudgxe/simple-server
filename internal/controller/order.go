package controller

import (
	"encoding/json"
	"net/http"

	"github.com/yudgxe/simple-server/internal/model"
)

func HandleOrder(cache *model.Order) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if cache != nil {
			orderID := r.URL.Path[1:]
			if orderID == cache.UID {
				res, err := json.Marshal(cache)
				if err != nil {
					http.Error(w, err.Error(), 400)
					return
				}
				w.Write(res)
			}
		}
		res, err := json.Marshal(&model.Order{})
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Write(res)
	}
}
