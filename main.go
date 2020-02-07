package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bitbucket.com/hexa/common/database/mysql"
	"bitbucket.com/hexa/movie"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		panic("DB_HOST is empty")
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		panic("DB_PORT is empty")
	}
	dbName := os.Getenv("DB_DATABASE")
	if dbName == "" {
		panic("DB_DATABASE is empty")
	}
	dbUserName := os.Getenv("DB_USERNAME")
	if dbUserName == "" {
		panic("DB_USERNAME is empty")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	// if dbPassword == "" {
	// 	panic("DB_PASSWORD is empty")
	// }

	ctx := context.Background()
	pool, err := GetMYSQLConnection(dbHost, dbUserName, dbPassword, dbName, dbPort, ctx)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	mysqlRepo := mysql.NewMysqlMovieRepository(pool, ctx)

	movieService := movie.NewMoviesService(mysqlRepo)
	movieHandler := movie.NewMovieHandler(movieService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/allmovies", movieHandler.GetAllMovies).Methods("GET")
	router.UseEncodedPath()
	http.Handle("/", AccessControl(router))

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port :3000")
		errs <- http.ListenAndServe(":3000", nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("terminated %s", <-errs)

}

func GetMYSQLConnection(
	host string,
	username string,
	password string,
	database string,
	port string,
	ctx context.Context,
) (*sql.DB, error) {

	dsn := fmt.Sprintf("%v:%s@tcp(%v:%v)/%v", username, password, host, port, database)
	pool, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	ctx, stop := context.WithCancel(ctx)
	defer stop()
	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)
	go func() {
		select {
		case <-appSignal:
			stop()
		}
	}()
	var cancel func()
	ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		return nil, err
	}
	return pool, nil
}

func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
