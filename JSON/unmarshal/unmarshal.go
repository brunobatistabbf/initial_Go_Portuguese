package unmarshal

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	fmt.Println("USO DA UNMARSHAL")
	cachorroJSON := `{"nome": 
"Rex": 
"raca": 
"Dalmata": 
"idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
}

//Usei a função marshal no arquivo main da pasta json
//Agora dando continuidade - uso do unmarshal
