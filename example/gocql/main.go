package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gocql/gocql/otelgocql"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

var tracer = otel.Tracer("app_or_package_name")

const keyspace = "gocql_example"

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	if err := initDB(); err != nil {
		log.Fatal(err)
	}

	cluster := newCassandraCluster(keyspace)
	session, err := otelgocql.NewSessionWithTracing(ctx, cluster)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	traceGocqlQueries(ctx, session)

	if err := truncateTable(ctx, session); err != nil {
		log.Fatal(err)
	}
}

func traceGocqlQueries(ctx context.Context, session *gocql.Session) {
	ctx, span := tracer.Start(ctx, "gocql-queries")
	defer span.End()

	insertBooks(ctx, session)
	bookID := selectBook(ctx, session)
	updateBook(ctx, session, bookID)
	deleteBook(ctx, session, bookID)

	fmt.Println("trace", otelplay.TraceURL(span))
}

func newCassandraCluster(keyspace string) *gocql.ClusterConfig {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.LocalQuorum
	cluster.ProtoVersion = 3
	cluster.Timeout = 10 * time.Second
	return cluster
}

func insertBooks(ctx context.Context, session *gocql.Session) {
	batch := session.NewBatch(gocql.LoggedBatch)
	for i := 0; i < 5; i++ {
		batch.Query(
			"INSERT INTO book (id, title, author_first_name, author_last_name) VALUES (?, ?, ?, ?)",
			gocql.TimeUUID(),
			fmt.Sprintf("Example Book %d", i),
			fmt.Sprintf("author_last_name %d", i),
			fmt.Sprintf("author_last_name %d", i),
		)
	}
	if err := session.ExecuteBatch(batch.WithContext(ctx)); err != nil {
		trace.SpanFromContext(ctx).RecordError(err)
	}
}

func selectBook(ctx context.Context, session *gocql.Session) string {
	res := session.
		Query(
			"SELECT id from book WHERE author_last_name = ?",
			"author_last_name 1",
		).
		WithContext(ctx).
		Iter()

	var bookID string
	for res.Scan(&bookID) {
		res.Scan(&bookID)
	}

	res.Close()

	return bookID
}

func updateBook(ctx context.Context, session *gocql.Session, bookID string) {
	if err := session.
		Query(
			"UPDATE book SET title = ? WHERE id = ?",
			"Example Book 1 (republished)", bookID,
		).
		WithContext(ctx).
		Exec(); err != nil {
		trace.SpanFromContext(ctx).RecordError(err)
	}
}

func deleteBook(ctx context.Context, session *gocql.Session, bookID string) {
	if err := session.
		Query("DELETE FROM book WHERE id = ?", bookID).
		WithContext(ctx).
		Exec(); err != nil {
		trace.SpanFromContext(ctx).RecordError(err)
	}
}

func initDB() error {
	cluster := newCassandraCluster("system")
	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}

	stmt := fmt.Sprintf(
		"CREATE KEYSPACE IF NOT EXISTS %s WITH replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }",
		keyspace,
	)
	if err := session.Query(stmt).Exec(); err != nil {
		return err
	}

	session.Close()

	cluster = newCassandraCluster(keyspace)
	session, err = cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	stmt = "CREATE table IF NOT EXISTS book(id UUID, title text, author_first_name text, author_last_name text, PRIMARY KEY(id))"
	if err = session.Query(stmt).Exec(); err != nil {
		return err
	}

	if err := session.Query("CREATE INDEX IF NOT EXISTS ON book(author_last_name)").Exec(); err != nil {
		return err
	}

	return nil
}

func truncateTable(ctx context.Context, session *gocql.Session) error {
	if err := session.Query("TRUNCATE TABLE book").WithContext(ctx).Exec(); err != nil {
		return err
	}

	return nil
}
