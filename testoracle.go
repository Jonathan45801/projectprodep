package main

import (
    "fmt"
    "database/sql"
    _ "github.com/godror/godror"
)
  
func testoracle(){
  
    db, err := sql.Open("godror", "<your username>/<your password>@service_name")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
      
      
    rows,err := db.Query("select sysdate from dual")
    if err != nil {
        fmt.Println("Error running query")
        fmt.Println(err)
        return
    }
    defer rows.Close()
  
    var thedate string
    for rows.Next() {
		
        rows.Scan(&thedate)
    }
    fmt.Printf("The date is: %s\n", thedate)
}