package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {

	c := &Cliente{
		UUID:          "uuid",
		Nome:          "Andre",
		Endereco:      "Rua",
		Cadastrado_em: time.Now().String(),
	}

	dirTest := "tmp"
	//Verificar se o nome da pasta ja existe
	//Caso não exita cria com noma passado na
	if _, err := os.Stat(dirTest); err != nil {
		if os.IsNotExist(err) {
			e := os.Mkdir(dirTest, 0755)
			if e != nil {
				t.Errorf(" > Falha err: %s\n", e)
			}
		} else {
			t.Errorf(" > Falha err: %s\n", err)
		}
	}

	//Cria o arquivo json
	newFile := filepath.Join(dirTest, c.UUID) + ".json"
	nf, err := os.Create(newFile)
	if err != nil {
		t.Errorf(" > Falha ao criar o arquivo err: %s\n", err)
	}

	b, err := json.Marshal(c)
	if err != nil {
		t.Errorf(" > Falha  na conversãoerr: %s\n", err)
	}

	//escrevi o arquivo os dados recebido na mensagem
	if _, err := nf.Write([]byte(b)); err != nil {
		t.Errorf(" > Falha  ao escrever no arquivo: %s\n", err)
	}
	nf.Close()

	err = os.RemoveAll(dirTest)
	if err != nil {
		t.Errorf(" > Falha  ao remover novoarquivo: %s\n", err)
	}

}
