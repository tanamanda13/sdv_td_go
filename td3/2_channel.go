package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	ch1 := make(chan Reponse)
	ch2 := make(chan Reponse)

	go callServer("http://localhost:4001/?id=5", ch1)
	go callServer("http://localhost:4001/?id=131", ch2)

	listenCh1, listenCh2 := <-ch1, <-ch2
	fmt.Println("listening on ", listenCh1, listenCh2)

	// select {
	// case <-ch1:
	// 	fmt.Println("received channel1")
	// case <-ch2:
	// 	fmt.Println("received channel2")
	// default:
	// 	fmt.Println("no activity")
	// }
	for i := 0; i < 2; i++ {
		select {
		case <-ch1:
			fmt.Println("received channel1")
		case <-ch2:
			fmt.Println("received channel2")
		// default:
		// 	fmt.Println("no activity")
		}

	}
}

type Reponse struct {
	RespText string
	Err      error
}

func callServer(addr string, resChan chan Reponse) {
	resp, err := http.Get(addr)
	if err != nil {
		// log.Fatalln(err)
		e := Reponse{Err: err}
		resChan <- e
		return
	}
	if resp.StatusCode != 200 {
		e2 := Reponse{Err: errors.New("Le code retourné par le serveur indique une erreur: " + strconv.Itoa(resp.StatusCode))}
		fmt.Println(e2)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// log.Fatalln(err)
		eBody := Reponse{Err: err}
		resChan <- eBody
	} else {
		//Convert the body to type string
		sb := string(body)
		log.Printf(sb)
		goodResp := Reponse{RespText: sb}
		resChan <- goodResp
	}
	defer resp.Body.Close()

	// http.HandleFunc(addr, ServeHTTP)
}
