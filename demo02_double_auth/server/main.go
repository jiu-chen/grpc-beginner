package main

import (
	"crypto/tls"
	"crypto/x509"
	"demo02/services"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//tls 是 ssl 的升级版
	cert, _ := tls.LoadX509KeyPair("../ca/server.pem", "../ca/server.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("../ca/ca.pem")
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, //服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	log.Println("grpc server start...")
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}
