package handlers

import (
	"html"
	"io"
	"net/http"
	"strconv"

	"github.com/jackc/pgx"
)

type yakstakRow struct {
	id       int64
	publicID string
	name     string
}

type YakstakIndex struct {
	DB *pgx.ConnPool
}

func (action *YakstakIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var yakstakRows []yakstakRow

	rows, err := action.DB.Query("select id, public_id, name from yakstak")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var r yakstakRow
		err := rows.Scan(&r.id, &r.publicID, &r.name)
		if err != nil {
			panic(err)
		}
		yakstakRows = append(yakstakRows, r)
	}

	YakstakIndexHtml(w, yakstakRows)
}

func YakstakIndexHtml(writer io.Writer, yakstakRows []yakstakRow) (err error) {
	io.WriteString(writer, `<table>`)
	for _, r := range yakstakRows {
		io.WriteString(writer, `<tr>`)
		io.WriteString(writer, `<td>`)
		io.WriteString(writer, html.EscapeString(strconv.FormatInt(r.id, 10)))
		io.WriteString(writer, `</td>`)
		io.WriteString(writer, `<td>`)
		io.WriteString(writer, html.EscapeString(r.publicID))
		io.WriteString(writer, `</td>`)
		io.WriteString(writer, `<td>`)
		io.WriteString(writer, html.EscapeString(r.name))
		io.WriteString(writer, `</td>`)
		io.WriteString(writer, `</tr>`)
	}
	io.WriteString(writer, `</table>`)
	io.WriteString(writer, html.EscapeString("<Jack>"))

	return
}
