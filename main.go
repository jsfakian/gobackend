package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"gobackend/api"
	"gobackend/configuration"
	"gobackend/db"
)

func main() {

	//Load router
	router := api.InitRoutes()

	//Load configuration
	configuration.LoadConfiguration()

	//Creating sqlite object
	db.InitDB()

	defer db.GetDB().Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8765" //localhost
	}

	fmt.Println(port)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig(),
	}
	if err := srv.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile(configuration.Conf.Certificate)
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile(configuration.Conf.ServerKey)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}
