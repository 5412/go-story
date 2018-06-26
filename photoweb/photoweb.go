package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "io/ioutil"
    "html/template"
    "path"
)

const (
    UPLOAD_PATH = "./uploads"
    TEMPLATE_DIR = "./views"
)

var templates = make(map[string]*template.Template)

func init() {
    fileInfoArr, err := ioutil.ReadDir(TEMPLATE_DIR)
    if err != nil {
        panic(err)
        return
    }

    var templateName, templatePath string
    for _, fileInfo := range fileInfoArr {
        templateName = fileInfo.Name()
        if ext := path.Ext(templateName); ext != ".html" {
            continue
        }
        templatePath = TEMPLATE_DIR + "/" + templateName
        log.Println("Loading template:", templatePath)
        t := template.Must(template.ParseFiles(templatePath))
        templates[templateName] = t
    } 
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        err := renderHtml(w, "upload", nil)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

        return
    }

    if r.Method == "POST" {
        f, h, err := r.FormFile("image")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        filename := h.Filename
        defer f.Close()

        t, err := os.Create(UPLOAD_PATH + "/" + filename)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        defer t.Close()
        if _, err := io.Copy(t, f); err != nil { 
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        http.Redirect(w, r, "/view?id=" + filename, http.StatusFound)
    }
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    imageID := r.FormValue("id")
    imagePath := UPLOAD_PATH + "/" + imageID
    if exists := isExists(imagePath); !exists {
        http.NotFound(w, r)
        return
    }
    w.Header().Set("Content-Type", "image")
    http.ServeFile(w, r, imagePath)
}

func isExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }

    return os.IsExist(err)
}

func listHandle(w http.ResponseWriter, r *http.Request) {
    //w.Header().Set("Content-Type", "text/html;charset=UTF-8")
    fileInfoArr, err := ioutil.ReadDir(UPLOAD_PATH)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    locals := make(map[string]interface{})
    images := []string{}


    for _, fileInfo := range fileInfoArr {
        imgid := fileInfo.Name()
        images = append(images, imgid)
    }

    locals["images"] = images

    if err = renderHtml(w, "list", locals); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}

func renderHtml(w http.ResponseWriter, tmpl string, locals map[string]interface{}) error {
   /*t, err := template.ParseFiles(tmpl + ".html")
    if err != nil {
        return err
    }
    err = t.Execute(w, locals)*/
    tmpl = tmpl + ".html"
    err := templates[tmpl].Execute(w, locals)
    return err
}

func check(err error) {
    if err != nil {
        panic(err)
    r
}

func main() {
    http.HandleFunc("/", listHandle)
    http.HandleFunc("/upload", uploadHandler)
    http.HandleFunc("/view", viewHandler)
    err := http.ListenAndServe(":8088", nil)
    if err != nil {
        log.Fatal("LisenAndServe failed.", err.Error())
    }
}

