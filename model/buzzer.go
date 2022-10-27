package model

type Buzzer struct {
	Estado          int `json:"estado"`
	TempoSegundos   int `json:"tempo_segundos"`
	QuantidadeVezes int `json:"quantidade_vezes"`
}
