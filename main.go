package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/netutil"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/", getData)
	httpServer := &http.Server{
		Addr:        ":8080",
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
		IdleTimeout: 1 * time.Second,
	}
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Listen: %v", err)
	}
	netutil.LimitListener(l, 100)
	log.Printf("Starting server for testing HTTP POST...\n")

	go func() {
		if err := httpServer.Serve(l); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	}

	cancel()

	defer os.Exit(0)
	return

}

func getData(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		sendErrorAnswer(w, 404, "Wrong path.")
		return
	}

	type reqBody struct {
		Data []string `json:"data"`
	}

	var (
		data   reqBody
		result = make(map[string]string)
	)

	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			sendErrorAnswer(w, 400, "Unmarshal error"+err.Error())
			return
		}
		err = json.Unmarshal(body, &data)
		if err != nil {
			sendErrorAnswer(w, 400, "Unmarshal error"+err.Error())
			return
		}

		if len(data.Data) > 20 {
			sendErrorAnswer(w, 400, "Exceeded the maximum number of incoming URLs. Max urls is 20.")
			return
		}

		for _, url := range data.Data {
			func() {
				var gg interface{}
				resp, err := http.Get(url)
				if err != nil {
					sendErrorAnswer(w, 400, "Request error "+err.Error())
					return
				}
				defer resp.Body.Close()
				s, err := ioutil.ReadAll(resp.Body)
				err = json.Unmarshal(s, &gg)

				res, err := json.Marshal(gg)
				if err != nil {
					sendErrorAnswer(w, 400, "Marshaling result error "+err.Error())
					return
				}

				result[url] = string(res)
			}()

		}
		sendSuccessfulAnswer(w, 200, result)

	default:
		log.Println(w, "Wrong method, only POST methods are supported.")
		sendErrorAnswer(w, 404, "Wrong method, only POST methods are supported.")
		return
	}
}
