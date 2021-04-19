package servidor

import (
	"crud_basico/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
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

//Buscarusuarios no banco de dados
func Buscarusuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco"))
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario"))
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao escanear o usuario"))
		return
	}
}

//Buscarusuario especifico no banco de dados
func Buscarusuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		fmt.Println(ID)
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))

	}
	linha, erro := db.Query("select * from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Erro ao buscar o usuario"))
		return
	}
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o  usuario"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para JSON!"))
		return
	}
}

//Atualizarusuario usuario no banco de dados
func Atualizarusuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}

	corporequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Erro ao ler o corpo da requisião"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corporequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar ao banco"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuario set nome = ?, email = ?, where = ?")
	if erro != nil {
		w.Write([]byte("Erro ao criar statement"))
		return
	}
	defer statement.Close()
	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Erro ao atualizar usuario!"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
