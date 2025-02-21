package main

import (
	"context"
	"fmt"
	"sync"

	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	startTime := time.Now()
	//RunMultipleConnection(300)
	//RunMultipleConnectionWithConnPool(500, 10)
	RunWithFixedSetOfConnections(1000, 10)
	fmt.Println(time.Now().Sub(startTime))
}

/*
This method creates multiple concurrent connections and is expected to end up in too many connections error
*/
func RunMultipleConnection(n int) {

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			helper()
		}()
	}
	wg.Wait()
}

/*
This method creates multiple concurrent connections, but relies on connection pool.
Hence it might take time but executes without error.
*/
func RunMultipleConnectionWithConnPool(n int, poolSize int) {
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	for i := 0; i < poolSize; i++ {
		ch <- 1
	}

	for i := 0; i < n; i++ {
		<-ch
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				ch <- 1
			}()

			helper()
		}()
	}
	wg.Wait()
}

func RunWithFixedSetOfConnections(n int, poolSize int) {
	var wg sync.WaitGroup
	ch := GetConnPool(poolSize)

	for i := 0; i < n; i++ {
		db := <-ch
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				ch <- db
			}()

			time.Sleep(time.Millisecond * 100)
		}()
	}
	wg.Wait()

}

func helper() {
	connString := "postgres://keerthanasmac:@localhost:5432/postgres"

	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
	time.Sleep(time.Millisecond * 100)
	db.Close(context.Background())
}

func GetConnPool(n int) chan *pgx.Conn {
	connPool := make(chan *pgx.Conn, n)

	connString := "postgres://keerthanasmac:@localhost:5432/postgres"

	for i := 0; i < n; i++ {
		db, err := pgx.Connect(context.Background(), connString)

		if err != nil {
			fmt.Println("Failed to connect to database:", err)
			return nil
		}

		connPool <- db
	}

	return connPool
}
