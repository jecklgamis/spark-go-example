package main

import (
	"context"
	"flag"
	"github.com/apache/spark-connect-go/v40/spark/sql"
	"github.com/apache/spark-connect-go/v40/spark/sql/utils"
	"log"
)

var (
	remote = flag.String("remote", "sc://localhost:15002",
		"Spark connect server")
)

func main() {
	flag.Parse()
	println("Submitting job to", *remote)
	ctx := context.Background()
	spark, err := sql.NewSessionBuilder().Remote(*remote).Build(ctx)
	if err != nil {
		log.Fatalf("Failed: %s", err)
	}
	defer utils.WarnOnError(spark.Stop, func(err error) {})

	df, err := spark.Sql(ctx, "select id from range(10)")
	if err != nil {
		log.Fatalf("Failed: %s", err)
	}
	df.Show(ctx, 100, false)
}
