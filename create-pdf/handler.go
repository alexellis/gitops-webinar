package function

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/phpdave11/gofpdf"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		http.Error(w, "give a ?msg= value for your PDF", http.StatusBadRequest)
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Times", "B", 16)
	// pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, msg)

	pdfPath := path.Join(os.TempDir(), "hello.pdf")
	err := pdf.OutputFileAndClose(pdfPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	res, err := ioutil.ReadFile(pdfPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment; filename="content.pdf"`)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
