package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "sreenath15"
    dbname   = "workdb"
)
 
func main() {
    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
 
    db, err := sql.Open("postgres", psqlconn)
    CheckError(err)
 
    defer db.Close()
 
    // insert
    // hardcoded
    insertStmt := `insert into "student"("id","name", "roll") values(1,'John', 1)`
    _, e := db.Exec(insertStmt)
    CheckError(e)
 
    // dynamic
    insertDynStmt := `insert into "student"("id","name", "roll") values(2,$1, $2)`
    _, e = db.Exec(insertDynStmt, "Jane", 2)
    CheckError(e)
    updateStmt := `update "student" set "name"=$1, "roll"=$2 where "id"=$3`
    _, e = db.Exec(updateStmt, "Mary", 3, 2)
    CheckError(e)
    deleteStmt := `delete from "student" where id=$1`
    _, e = db.Exec(deleteStmt, 1)
    CheckError(e)
    rows, err := db.Query(`SELECT "name", "roll" FROM "student"`)
    CheckError(err)
 
    defer rows.Close()
    for rows.Next() {
    var name string
    var roll int
 
    err = rows.Scan(&name, &roll)
    CheckError(err)
 
    fmt.Println(name, roll)
}
 
CheckError(err)
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}