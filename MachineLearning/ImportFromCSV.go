package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

/**
pclass          Passenger Class
                (1 = 1st; 2 = 2nd; 3 = 3rd)
survival        Survival
                (0 = No; 1 = Yes)
name            Name
sex             Sex
age             Age
sibsp           Number of Siblings/Spouses Aboard
parch           Number of Parents/Children Aboard
ticket          Ticket Number
fare            Passenger Fare
cabin           Cabin
embarked        Port of Embarkation
                (C = Cherbourg; Q = Queenstown; S = Southampton)
boat            Lifeboat
body            Body Identification Number
home.dest       Home/Destination
*/
type SurvivalPassenger struct {
	Pclass   string `json:"clasePasajero"`
	Survival string `json:"sobrevivio"`
	Name     string `json:"nombre"`
	Sex      string `json:"sexo"`
	Age      string `json:"edad"`
	Sibsp    string `json:"nroAbordoEsposas"`
	Parch    string `json:"nroAbordoParientes"`
	Ticket   string `json:"nroTicket"`
	Fare     string `json:"tarifa"`
	Cabin    string `json:"nroCamarote"`
	Embarked string `json:"puertoEmbarcacion"`
	Boat     string `json:"nroSalvavida"`
	Body     string `json:"nroIdCuerpo"`
	HomeDest string `json:"destino"`
}

var (
	sp             SurvivalPassenger
	sobrevivientes []SurvivalPassenger
)

func main() {
	csvFile, _ := os.Open("titanic3.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		//		log.Println(line)
		sp.Pclass = line[0]
		sp.Survival = line[1]
		sp.Name = line[2]
		sp.Sex = line[3]
		sp.Age = line[4]
		sp.Sibsp = line[5]
		sp.Parch = line[6]
		sp.Ticket = line[7]
		sp.Fare = line[8]
		sp.Cabin = line[9]
		sp.Embarked = line[10]
		sp.Boat = line[11]
		sp.Body = line[12]
		sp.HomeDest = line[13]

		sobrevivientes = append(sobrevivientes, sp)
	}
	sobrevivientesJson, _ := json.Marshal(sobrevivientes)
	fmt.Println(string(sobrevivientesJson))
}
