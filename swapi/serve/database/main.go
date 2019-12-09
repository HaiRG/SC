package main

import (
    "fmt"
    "log"
    "strconv"
    "encoding/json"
    "serve/database/db"
    "serve/database/swapi"
)
func main() {
    dbase := "db/swa.db"
 //  db.Init(dbase)
    fmt.Println("init")
    db.Start(dbase)
fmt.Println("start")
    getData()
    db.Stop()
fmt.Println("finish")
}

func getData() {
   // getFilm()
   //getPerson()
   //getPlanet()
   //getSpecies()
   // getVehicle()
   getStarship()
}

func getFilm() {
    c := swapi.DefaultClient
    fail := 0
    for index := 1; ; index++ {
	  fmt.Println(index)
        data,_ := c.Film(index)
        jsonStr ,_ := json.MarshalIndent(data, "", "  ")
        indexStr := strconv.Itoa(index)
        if len(db.GetValue([]byte("films"), []byte(indexStr))) == 0 {
            if !store([]byte("films"), []byte(indexStr), jsonStr) {
                fail ++
            fmt.Println(fail)
                if fail == 5{
                    break
                }
            }else{
                fail = 0
            }
        }else{
        fmt.Println("same")
       }
    }
}

func getPerson() {
    c := swapi.DefaultClient
    fail := 0
    for index := 1; ; index++ {
        fmt.Println(index)
        data,_  := c.Person(index)
        jsonStr ,_ := json.MarshalIndent(data, "", "  ")
        indexStr := strconv.Itoa(index)
        if len(db.GetValue([]byte("people"), []byte(indexStr))) == 0 {
            if !store([]byte("people"), []byte(indexStr), jsonStr) {
                fail ++
                if fail == 5{
                    break
                }
            }else{
                fail = 0
            }
        }
    }
}

func getPlanet() {
    c := swapi.DefaultClient
    fail := 0
    for index := 1; ; index++ {
	  fmt.Println(index)
        data,_  := c.Planet(index)
        jsonStr ,_ := json.MarshalIndent(data, "", "  ")
        indexStr := strconv.Itoa(index)
        if len(db.GetValue([]byte("planets"), []byte(indexStr))) == 0 {
            if !store([]byte("planets"), []byte(indexStr), jsonStr) {
                fail ++
                if fail == 5{
                    break
                }
            }else{
                fail = 0
            }
        }
    }
}

func getSpecies() {
    c := swapi.DefaultClient
    fail := 0
    for index := 1; ; index++ {
	  fmt.Println(index)
        data,_  := c.Species(index)
        jsonStr ,_ := json.MarshalIndent(data, "", "  ")
        indexStr := strconv.Itoa(index)
        if len(db.GetValue([]byte("species"), []byte(indexStr))) == 0 {
            if !store([]byte("species"), []byte(indexStr), jsonStr) {
                fail ++
                if fail == 5{
                    break
                }
            }else{
                fail = 0
            }
        }
    }
}

func getStarship() {
    c := swapi.DefaultClient
    fail := 0
    for index := 1; ; index++ {
	  fmt.Println(index)
        data,_  := c.Starship(index)
        jsonStr ,_ := json.MarshalIndent(data, "", "  ")
        indexStr := strconv.Itoa(index)
        if len(db.GetValue([]byte("starships"), []byte(indexStr))) == 0 {
            if !store([]byte("starships"), []byte(indexStr), jsonStr) {
                fail ++
                if fail == 5{
                    break
                }
            }else{
                fail = 0
            }
        }
    }
}

func getVehicle() {
    c := swapi.DefaultClient
    fail := 0
    for index := 1; ; index++ {
        //var data interface{}
	  fmt.Println(index)
        data,_  := c.Vehicle(index)
        jsonStr ,_ := json.MarshalIndent(data, "", "  ")
        indexStr := strconv.Itoa(index)
        if len(db.GetValue([]byte("vehicles"), []byte(indexStr))) == 0 {
            if !store([]byte("vehicles"), []byte(indexStr), jsonStr) {
                fail ++
                if fail == 5{
                    break
                }
            }
        }
    }
}
func store(bucket []byte, key []byte, value []byte) bool {
    stb := &swapi.Film{}
    err := json.Unmarshal(value, &stb)
    if err != nil {
        log.Fatal(err)
        return false
    } else if len(stb.URL) == 0 {
        return false
    }
    db.Update(bucket, key, value)
    return true
}
