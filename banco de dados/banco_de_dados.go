package main

import (
 "database/sql"
 "fmt"
 "log"
 _"github.com/jackc/pgx "
  _"github.com/pkg/errors"
  _"golang.org/x/crypto "
  _"golang.org/x/text "
)

func main()  {
 //Estrutura do banco //urlconexao := "usuario:senha@/banco"
   urlconexao := ":Saopaulo3@/bancodedados"
   db, erro := sql.Open("postgre", urlconexao)
   if erro!=nil{
    log.Fatal(erro)
   }
   fmt.Println(db)

   //como fazer querys
   //linhas, erro := //db.Query("select * from usuarios")

}


//Para fazer conexão sql //go get github.com/go-sql-driver/mysql
//github.com/jackc/pgx  para postgre