package middleware

import (
	"errors"
	"net/http"
	"strings"
	"zapys/src/db"
	"zapys/src/respostas"
)

func ADMMIDDLEWARE(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		useradm := r.URL.Query().Get("useradm")
		admpass := r.URL.Query().Get("admpass")
		useradm = strings.ToLower(useradm)
		admpass = strings.ToLower(admpass)

		if useradm == "" || admpass == "" {
			respostas.ERROR(w, http.StatusBadRequest, errors.New("Usuário ou senha não fornecidos"))
			return
		}

		dbConn, err := db.DBADM()
		if err != nil {
			respostas.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		defer dbConn.Close()

		query := "SELECT adminpass FROM admin WHERE adminname = ? AND adminpass = ? LIMIT 1"
		row := dbConn.QueryRow(query, useradm, admpass)

		var dbPass string
		err = row.Scan(&dbPass)
		if err != nil {
			respostas.ERROR(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
