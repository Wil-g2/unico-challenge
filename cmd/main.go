package main

import (
	"encoding/csv"
	"os"
	"runtime"
	"strconv"

	"log"

	validatorv10 "github.com/go-playground/validator/v10"
	"github.com/wil-g2/unico-challenge/domain/fair"
	"github.com/wil-g2/unico-challenge/infra/config"
	"github.com/wil-g2/unico-challenge/infra/database"
	"github.com/wil-g2/unico-challenge/infra/repositories"
	"github.com/wil-g2/unico-challenge/infra/validator"
)

func main() {
	envConfig := config.SetupEnvFile()

	mysql := database.InitMysql(envConfig)

	validator := validator.NewValidator(validatorv10.New())
	repo := repositories.NewFairRepository(mysql)
	service := fair.NewService(repo, validator)

	csvFile, err := os.Open(envConfig.PathCSV)
	if err != nil {
		log.Println("Error opening CSV file", err)

	}
	defer csvFile.Close()

	lines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Println("Error reading CSV file", err)
	}
	jobsNumber := len(lines) - 1
	tasks := make(chan *fair.CreateDTO, jobsNumber)
	results := make(chan bool, jobsNumber)

	for w := 1; w <= runtime.NumCPU(); w++ {
		go worker(tasks, results, service)
	}

	for _, line := range lines {
		Long, _ := strconv.ParseFloat(line[1], 64)
		Lat, _ := strconv.ParseFloat(line[2], 64)
		CodDist, _ := strconv.Atoi(line[5])
		CodSubPref, _ := strconv.Atoi(line[7])
		fair := &fair.CreateDTO{
			Long:         Long,
			Lat:          Lat,
			Setcens:      line[3],
			Areap:        line[4],
			CodDist:      CodDist,
			District:     line[6],
			CodSubPref:   CodSubPref,
			Region5:      line[9],
			Region8:      line[10],
			FairName:     line[11],
			Record:       line[12],
			Street:       line[13],
			Number:       line[14],
			Neighborhood: line[15],
			Reference:    line[16],
		}
		tasks <- fair
	}

	close(tasks)
	for i := 1; i <= jobsNumber; i++ {
		<-results
		log.Println("reading results:" + strconv.Itoa(i))
	}
}

func worker(tasks <-chan *fair.CreateDTO, results chan<- bool, service fair.Service) {
	for task := range tasks {
		err := service.Create(task)
		if err != nil {
			results <- false
		}
		results <- true
	}
}
