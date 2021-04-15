package main

//pacote para trabalhar com json
import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

//json.Marshal()
//json.Unmarshal()
//Conversao de json pra struct e vice versa

//Use crase nos campos
type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Betão", "Pitbull", 3}
	fmt.Println(c)

	cachorroemJSON, erro := json.Marshal(c)

	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroemJSON)
	fmt.Println(bytes.NewBuffer(cachorroemJSON))
	//BYTES.NEWBUFFER para ler o JSON
}
