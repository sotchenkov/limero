package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Msg struct {
	Value map[string]interface{} `json:"value"`
}

func createQueue(brokerAddr string, qname string, presize string) error {
	req, err := http.NewRequest(http.MethodPut, "http://"+brokerAddr+"/queue?name="+qname+"&size="+presize, nil)
	if err != nil {
		log.Print("queue creation error")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("request error")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("status code is not 201: %v", resp.StatusCode)
		return errors.New("wrong status code")
	}

	log.Print("queue has been created successfully")
	return nil
}

func sendMsg(value string, brokerAddr string, qname string) error {
	body := []byte(fmt.Sprintf(`{"value": {"message": "%s"}}`, value))
	log.Printf("Sending message payload: %s", body)

	req, err := http.NewRequest(http.MethodPost, "http://"+brokerAddr+"/msg?qname="+qname, bytes.NewBuffer(body))
	if err != nil {
		log.Print("couldn't send a message")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	httpClient := http.Client{Timeout: 5 * time.Second}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Print("couldn't send a message")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Printf("status code is not 201: %v", resp.StatusCode)
		return errors.New("wrong status code")
	}

	log.Print("the message has been sent successfully")
	return nil
}

func getMsg(brokerAddr string, qname string) (error, *Msg) {
	httpClient := http.Client{Timeout: 5 * time.Second}
	r, err := httpClient.Get("http://" + brokerAddr + "/msg?qname=" + qname)
	if err != nil {
		log.Print("request to receive a message from the broker could not be executed")
		return err, &Msg{Value: map[string]interface{}{}}
	}

	defer r.Body.Close()

	var msg *Msg
	err = json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		log.Print("message could not be decoded")
		return err, &Msg{Value: map[string]interface{}{}}
	}

	return nil, msg
}

func deleteQueue(brokerAddr string, qname string) {
	req, err := http.NewRequest(http.MethodDelete, "http://"+brokerAddr+"/queue?name="+qname, nil)
	if err != nil {
		log.Print("queue deletion error")
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("request error")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("status code is not 200: %v", resp.StatusCode)
	}

	log.Print("queue has been deleted successfully")
}

// NewRandomString generates random string with given size.
func newRandomString(size int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")

	b := make([]rune, size)
	for i := range b {
		b[i] = chars[rnd.Intn(len(chars))]
	}

	return string(b)
}

func BrokerStress() {
	// Setup
	const (
		queueName       = "HightloadBroker"
		brokerAddr      = ":7920"
		presize         = "1000"
		sendGorutines   = 10000
		reciveGorutines = 10000
		sendTimeout     = 1 // milliseconds
		getTimeout      = 1 // milliseconds
		msgLen          = 5 // characters
	)

	err := createQueue(brokerAddr, queueName, presize)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	defer deleteQueue(brokerAddr, queueName)

	sendTicker := time.NewTicker(sendTimeout * time.Millisecond)
	done := make(chan bool)

	for i := 0; i < sendGorutines; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				case <-sendTicker.C:
					err := sendMsg(newRandomString(msgLen), brokerAddr, queueName)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}()
	}

	getTicker := time.NewTicker(getTimeout * time.Millisecond)
	for i := 0; i < reciveGorutines; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				case <-getTicker.C:
					err, msg := getMsg(brokerAddr, queueName)
					if err != nil {
						log.Fatal(err)
					}
					log.Printf("received msg: %s", msg.Value["message"])
				}
			}
		}()
	}

	// Реализуем graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	s := <-stop

	log.Print("stopping stress...", zap.String("signal", s.String()))

	done <- true

	log.Print("deleting a queue")
	deleteQueue(brokerAddr, queueName)
	log.Print("queue has been deleted")

	log.Print("stress stopped")
}

func main() {
	BrokerStress()
}
