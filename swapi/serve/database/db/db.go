package db

import (
    "log"
    "os"
    "time"
    "github.com/boltdb/bolt"
)

var db *bolt.DB

func Start(str string) {
    var dbName = "swa.db"
    var err error
    dbName = str
    db, err = bolt.Open(dbName, 0666, &bolt.Options{Timeout: 1 * time.Second})
    if err != nil {
        log.Fatal(err)
        return
    }
}

func Stop(){
    if err := db.Close(); err != nil {
        log.Fatal(err)
    }
}

func Init(str string) {
     _,err := os.Open(str)  
    if err == nil{
        return
    }
    Start(str)
    if err := db.Update(func(tx *bolt.Tx) error {
        tx.CreateBucket([]byte("users"))
        tx.CreateBucket([]byte("films"))
        tx.CreateBucket([]byte("people"))
        tx.CreateBucket([]byte("planets"))
        tx.CreateBucket([]byte("species"))
        tx.CreateBucket([]byte("starships"))
        tx.CreateBucket([]byte("vehicles"))
        return nil
    }); err != nil {
        log.Fatal(err)
    }
    Stop()
}


func Update(bucket[]byte, key []byte, value []byte) {
    if err := db.Update(func(tx *bolt.Tx) error {
        if err := tx.Bucket(bucket).Put(key, value); err != nil {
            return err
        }
        return nil
    }); err != nil {
        log.Fatal(err)
    }
}

func GetValue(bucket []byte, key []byte) string {
    var res []byte
    if err := db.View(func(tx *bolt.Tx) error {
        byteLen := len(tx.Bucket([]byte(bucket)).Get(key))
        res = make([]byte, byteLen)
        copy(res[:], tx.Bucket([]byte(bucket)).Get(key)[:])
        return nil
    }); err != nil {
        log.Fatal(err)
    }
    return string(res)
}

func GetCount(bucket []byte) int {
    count := 0
    if err := db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        c := b.Cursor()
        for k, _ := c.First(); k != nil; k, _ = c.Next() {
            count ++
        }
        return nil
    }); err != nil {
        log.Fatal(err)
    }
    return count
}




