package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type WeatherService struct {
	weatherClient *WeatherClient
	mailClient    *MailClient
}

func New() *WeatherService {

	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		panic("cannot lookup api key")
	}

	sender, ok := os.LookupEnv("SENDER_ADDRESS")
	if !ok {
		panic("cannot lookup SENDER_MAIL")
	}

	pass, ok := os.LookupEnv("SENDER_PASS")
	if !ok {
		panic("cannot lookup SENDER_PASS")
	}

	host, ok := os.LookupEnv("SENDER_SMTP_HOST")
	if !ok {
		panic("cannot lookup SENDER_SMTP_HOST")
	}

	port, ok := os.LookupEnv("SENDER_SMTP_PORT")
	if !ok {
		panic("cannot lookup SENDER_SMTP_PORT")
	}
	intport, err := strconv.Atoi(port)
	if err != nil {
		panic("cannot lookup SENDER_SMTP_PORT")
	}
	receiver, ok := os.LookupEnv("RECEIVER_ADDRESS")
	if !ok {
		panic("cannot lookup RECEIVER_ADDRESS")
	}

	return &WeatherService{
		weatherClient: NewWeatherClient(apiKey),
		mailClient: NewMailClient(Config{
			sender:   sender,
			receiver: receiver,
			password: pass,
			smtphost: host,
			smtpport: intport,
		}),
	}
}

func (p *WeatherService) Default() string {
	return "Hello!"
}

type Request struct {
	Location string
}

type Response struct {
	Response string
}

// example of JSON api method
// test with:
//   mantil invoke wheater/get --data '{"location":"Milan"}'
func (s *WeatherService) Get(ctx context.Context, req Request) (Response, error) {

	var r Response
	log.Printf("Received request %s", req)

	weather, err := s.weatherClient.FetchWheater(req.Location)

	if err != nil {
		return r, err
	}

	marshaled, err := json.Marshal(weather)

	if err != nil {
		return r, err
	}

	err = s.mailClient.SendMail(string(marshaled))

	if err != nil {
		return r, err
	}

	return Response{Response: fmt.Sprintf("Weather Details for station %s sent to %s", req.Location, s.mailClient.config.receiver)}, nil
}
