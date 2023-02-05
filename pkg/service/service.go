package service

import (
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"

	// load postgres driver
	_ "github.com/lib/pq"

	"atlassian.carcgl.com/bitbucket/ls/lms/db"
	"atlassian.carcgl.com/bitbucket/ls/lms/pkg/api/proto/v1"
)

// Local database connection parameters
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

// LMS gRPC server
type LMSService struct {
	gRPCServerAddr string
	pgDatabase     *sqlx.DB
}

// New - creates a new LMS gRPC server
func New(grpcAddress string) *LMSService {
	return &LMSService{
		gRPCServerAddr: grpcAddress,
		pgDatabase:     nil,
	}
}

// Run - runs the gRPC server
func (lms *LMSService) Run() {
	log.Println("Running LMS service")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	pgsql, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer pgsql.Close()

	err = pgsql.Ping()
	if err != nil {
		pgsql.Close()
		panic(err)
	}

	log.Println("Successfully connected to postgres db!")

	lms.pgDatabase = pgsql

	e := db.CreateSchema(lms.pgDatabase) // Create Database Tables
	if e != nil {
		log.Fatalf("can not create schema  %v", e)
	}

	//Start a grpc server
	lms.startGRPC()
}

// startGRPC - starts the gRPC server
func (lms *LMSService) startGRPC() error {
	listener, err := net.Listen("tcp", lms.gRPCServerAddr)

	if err != nil {
		return err
	}
	gs := grpc.NewServer()

	proto.RegisterLMSRecordServiceServer(gs, lms)

	// Start gRPC server
	log.Printf("starting gRPC interface %s", lms.gRPCServerAddr)
	err = gs.Serve(listener)

	return err
}
