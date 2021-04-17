package servidor

import (
	"crud_basico/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func Criarusuario(w http.ResponseWriter, r *http.Request) {
	corporequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falhou"))
		return
	}
	var usuario usuario
	if erro = json.Unmarshal(corporequisicao, &usuario); erro != nil {
		w.Write([]byte("Erroooooooouuuuu"))
	}

	fmt.Println(usuario)
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erroooooooooooooooou"))
		return
	}

	statement, erro := db.Prepare("insert into usuarios (nome, email)values (?,?)")
	if erro != nil {
		w.Write([]byte("Erro ao cruar o statement"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {

		w.Write([]byte("Erro ao executar o statement!"))
		return
	}
	idinserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao executar o statement!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprint("Usuario inserido com sucesso! ID: %d", idinserido)))

	//Insert into usuarios (nome, email)values ("nome", "email")
	//prepare statement
}
