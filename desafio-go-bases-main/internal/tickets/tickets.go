package tickets

import "errors"

type Ticket struct {
	ID      string
	Name    string
	Email   string
	Country string
	Time    string
	Value   string
}

func (t Ticket) GetHourTime() (int, error) {
	return ObtenerHoraPeriodo(t.Time)
}
func (t Ticket) GetStringTime() (string, error) {
	tiempo, err := t.GetHourTime()
	if err != nil {
		return "", errors.New("error al obtener la hora del ticket")
	}
	return ObtenerPeriodo(tiempo)
}
