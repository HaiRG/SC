package api

import (
    //"fmt"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
    "serve/database/db"
)
func peopleHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        req.ParseForm()
        w.Write([]byte("{\n\"people\":[\n"))
        fail := 0
        i := 0
        for count := 0 ;;count++{
            item := db.GetValue([]byte("people"), []byte(strconv.Itoa(count)))
            if len(item) != 0 {
                fail = 0
                i++
                str := "people/" + strconv.Itoa(count)
                w.Write([]byte("\""))
                w.Write([]byte(str))
                w.Write([]byte("\""))
                if i  >= db.GetCount([]byte("people")) {
                    break
                }
                w.Write([]byte(",\n"))
            }else{
                fail++
                if fail == 5{
                    break;
                }
            }
            
        }
        w.Write([]byte("]\n}"))
    }
}

func peopleIdHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    re := mux.Vars(req)
    _, err := strconv.Atoi(re["id"])
    if err != nil {
        log.Fatal(err)
    }
    data := db.GetValue([]byte("people"), []byte(re["id"]))
    w.Write([]byte(data))
}
func filmsHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        req.ParseForm()
        w.Write([]byte("{\n\"films\":[\n"))
        fail := 0
        i := 0
        for count := 0 ;;count++{
            item := db.GetValue([]byte("films"), []byte(strconv.Itoa(count)))
            if len(item) != 0 {
                fail = 0
                i++
                str := "films/" + strconv.Itoa(count)
                w.Write([]byte("\""))
                w.Write([]byte(str))
                w.Write([]byte("\""))
                if i  >= db.GetCount([]byte("films")) {
                    break
                }
                w.Write([]byte(",\n"))
            }else{
                fail++
                if fail == 5{
                    break;
                }
            }
            
        }
        w.Write([]byte("]\n}"))
    }
}

func filmsIdHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    re := mux.Vars(req)
    _, err := strconv.Atoi(re["id"])
    if err != nil {
        log.Fatal(err)
    }
    data := db.GetValue([]byte("films"), []byte(re["id"]))
    w.Write([]byte(data))
}

func planetsHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        req.ParseForm()
        w.Write([]byte("{\n\"planets\":[\n"))
        fail := 0
        i := 0
        for count := 0 ;;count++{
            item := db.GetValue([]byte("planets"), []byte(strconv.Itoa(count)))
            if len(item) != 0 {
                fail = 0
                i++
                str := "planets/" + strconv.Itoa(count)
                w.Write([]byte("\""))
                w.Write([]byte(str))
                w.Write([]byte("\""))
                if i  >= db.GetCount([]byte("planets")) {
                    break
                }
                w.Write([]byte(",\n"))
            }else{
                fail++
                if fail == 5{
                    break;
                }
            }
            
        }
        w.Write([]byte("]\n}"))
    }
}

func planetsIdHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    re := mux.Vars(req)
    _, err := strconv.Atoi(re["id"])
    if err != nil {
        log.Fatal(err)
    }
    data := db.GetValue([]byte("planets"), []byte(re["id"]))
    w.Write([]byte(data))
}

func speciesHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        req.ParseForm()
        w.Write([]byte("{\n\"species\":[\n"))
        fail := 0
        i := 0
        for count := 0 ;;count++{
            item := db.GetValue([]byte("species"), []byte(strconv.Itoa(count)))
            if len(item) != 0 {
                fail = 0
                i++
                str := "species/" + strconv.Itoa(count)
                w.Write([]byte("\""))
                w.Write([]byte(str))
                w.Write([]byte("\""))
                if i  >= db.GetCount([]byte("species")) {
                    break
                }
                w.Write([]byte(",\n"))
            }else{
                fail++
                if fail == 5{
                    break;
                }
            }
            
        }
        w.Write([]byte("]\n}"))
    }
}

func speciesIdHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    re := mux.Vars(req)
    _, err := strconv.Atoi(re["id"])
    if err != nil {
        log.Fatal(err)
    }
    data := db.GetValue([]byte("species"), []byte(re["id"]))
    w.Write([]byte(data))
}

func starshipsHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        req.ParseForm()
        w.Write([]byte("{\n\"starships\":[\n"))
        fail := 0
        i := 0
        for count := 0 ;;count++{
            item := db.GetValue([]byte("starships"), []byte(strconv.Itoa(count)))
            if len(item) != 0 {
                fail = 0
                i++
                str := "starships/" + strconv.Itoa(count)
                w.Write([]byte("\""))
                w.Write([]byte(str))
                w.Write([]byte("\""))
                if i  >= db.GetCount([]byte("starships")) {
                    break
                }
                w.Write([]byte(",\n"))
            }else{
                fail++
                if fail == 5{
                    break;
                }
            }
            
        }
        w.Write([]byte("]\n}"))
    }
}

func starshipsIdHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    re := mux.Vars(req)
    _, err := strconv.Atoi(re["id"])
    if err != nil {
        log.Fatal(err)
    }
    data := db.GetValue([]byte("starships"), []byte(re["id"]))
    w.Write([]byte(data))
}

func vehiclesHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        req.ParseForm()
        w.Write([]byte("{\n\"vehicles\":[\n"))
        fail := 0
        i := 0
        for count := 0 ;;count++{
            item := db.GetValue([]byte("vehicles"), []byte(strconv.Itoa(count)))
            if len(item) != 0 {
                fail = 0
                i++
                str := "vehicles/" + strconv.Itoa(count)
                w.Write([]byte("\""))
                w.Write([]byte(str))
                w.Write([]byte("\""))
                if i  >= db.GetCount([]byte("vehicles")) {
                    break
                }
                w.Write([]byte(",\n"))
            }else{
                fail++
                if fail == 5{
                    break;
                }
            }
            
        }
        w.Write([]byte("]\n}"))
    }
}

func vehiclesIdHandler(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.WriteHeader(http.StatusOK)
    re := mux.Vars(req)
    _, err := strconv.Atoi(re["id"])
    if err != nil {
        log.Fatal(err)
    }
    data := db.GetValue([]byte("vehicles"), []byte(re["id"]))
    w.Write([]byte(data))
}
