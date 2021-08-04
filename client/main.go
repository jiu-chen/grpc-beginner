package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"hi-grpc/services"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, _ := tls.LoadX509KeyPair("../ca/client.pem", "../ca/client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("../ca/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ServerName:   "localhost",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId: 20})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
