package handle

import (
	"io"
	"net/http"
	"os"
	"text/template"
)

type Upload struct {
	PortServer string
}

func NewUpload(port string) *Upload {
	return &Upload{
		PortServer: port,
	}
}

var templates = template.Must(template.ParseFiles("./server/public/upload.html"))

var dstDir = "./example/file/"

func (u *Upload) UploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error Retrieving the File"))
		return
	}
	defer file.Close()

	dst, err := os.Create(dstDir + handler.Filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Creating the File"))
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error Saving the File"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File Uploaded Successfully"))

}

func (u *Upload) PageUpload(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "upload.html", u)
}
