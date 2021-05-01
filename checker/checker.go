package checker

import (
	"errors"
	"log"
	"time"

	"github.com/toxinu/dekiteru/services"
)

// Run runs the given service
func Run(service string, interval int, retries int, parameters map[string]interface{}) error {
	var (
		delta time.Duration
		start time.Time
		err   error
	)
	s, ok := services.Services[service]
	if !ok {
		return errors.New("this service does not exist")
	}
	for t := 1; t <= retries; t++ {
		start = time.Now()

		err := s.Run(parameters)
		switch err.(type) {
		case *services.HardError:
			return err
		case *services.SoftError:
			log.Println(err)
		case nil:
			return err
		}

		if t+1 > retries {
			return err
		}

		delta = time.Now().Sub(start)
		if delta < time.Duration(interval)*time.Second {
			time.Sleep(time.Duration(interval)*time.Second - delta)
		}
	}
	return err
}
