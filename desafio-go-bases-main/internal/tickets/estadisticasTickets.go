package tickets

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumeroViajes struct {
	viajesTotales   int
	viajesMadrugada int
	viajesMañana    int
	viajesTarde     int
	ViajesNoche     int
}

type DatosTickets struct {
	tickets []Ticket
	total   int
	NumeroViajes
}

func ReadTicketsCSV() ([]Ticket, error) {
	file, err := os.Open("../tickets.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var tickets []Ticket

	for _, record := range records {
		if len(record) < 6 {
			continue
		}

		ticket := Ticket{
			ID:      record[0],
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Time:    record[4],
			Value:   record[5],
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func GetTotalTickets(tickets []Ticket, destination string) (int, error) {

	count := 0
	for _, ticket := range tickets {
		if ticket.Country == destination {
			count++
		}
	}

	return count, nil
}

func GetContadorPorPeriodo(time string) (int, error) {
	tickets, err := ReadTicketsCSV()
	if err != nil {
		return 0, err
	}

	count := 0
	for _, ticket := range tickets {
		if strings.HasPrefix(ticket.Time, time) {
			count++
		}
	}

	return count, nil
}

func AverageDestination(destination string, total int) (int, error) {
	tickets, err := ReadTicketsCSV()
	if err != nil {
		return 0, err
	}

	count := 0
	totalValue := 0
	for _, ticket := range tickets {
		if ticket.Country == destination {
			value, err := strconv.Atoi(ticket.Value)
			if err != nil {
				return 0, err
			}
			count++
			totalValue += value
		}
	}

	if count == 0 {
		return 0, fmt.Errorf("No hay tickets para el destino %s", destination)
	}

	average := totalValue / count
	return average, nil
}
