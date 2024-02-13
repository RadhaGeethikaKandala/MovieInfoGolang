package testing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/RadhaGeethikaKandala/MovieRental/internal/app/config"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	dStub "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDatabase struct {
	DbInstance *sql.DB
	container  testcontainers.Container
}

const configFileRelativePath = "../config"
const postgresDatasource = "postgres://%s:%s@%s:%s/%s?sslmode=%s"

func SetupTestDatabase() *TestDatabase {

	// setup db container
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	container, dbInstance, err := createContainer(ctx)
	if err != nil {
		log.Fatal("failed to setup test", err)
	}

	// migrate db schema
	err = migrateDb(dbInstance)
	if err != nil {
		log.Fatal("failed to perform db migration", err)
	}
	cancel()

	return &TestDatabase{
		container:  container,
		DbInstance: dbInstance,
	}
}

func (tdb *TestDatabase) TearDown() {
	tdb.DbInstance.Close()
	// remove test container
	_ = tdb.container.Terminate(context.Background())
}

func createContainer(ctx context.Context) (testcontainers.Container, *sql.DB, error) {
	databaseConf := config.ReadConfig(configFileRelativePath).DatabaseTest

	var env = map[string]string{
		"POSTGRES_PASSWORD": databaseConf.Password,
		"POSTGRES_USER":     databaseConf.Username,
		"POSTGRES_DB":       databaseConf.Dbname,
	}
	var port = fmt.Sprintf(databaseConf.Port + "/tcp")

	fmt.Println("port is", port)

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres:14-alpine",
			ExposedPorts: []string{port},
			Env:          env,
			WaitingFor:   wait.ForLog("database system is ready to accept connections"),
		},
		Started: true,
	}
	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return container, nil, fmt.Errorf("failed to start container: %v", err)
	}

	p, err := container.MappedPort(ctx, nat.Port(databaseConf.Port))
	if err != nil {
		return container, nil, fmt.Errorf("failed to get container external port: %v", err)
	}

	log.Println("postgres container ready and running at port: ", p.Port())

	time.Sleep(time.Second)

	databaseURL :=
		fmt.Sprintf(
			postgresDatasource,
			databaseConf.Username,
			databaseConf.Password,
			databaseConf.Host,
			p.Port(),
			databaseConf.Dbname,
			databaseConf.Sslmode,
		)

	fmt.Println("database url in creation ", databaseURL)
	db, err := sql.Open("postgres",
		databaseURL)

	if err != nil {
		return container, db, fmt.Errorf("failed to establish database connection: %v", err)
	}

	return container, db, nil
}

func migrateDb(dbInstance *sql.DB) error {
	databaseConf := config.ReadConfig("../config").DatabaseTest
	// get location of test
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get path")
	}

	pathToMigrationFiles := filepath.Dir(path) + "/../database/migrations/"
	fmt.Println("path is ", pathToMigrationFiles)

	instance, err := dStub.WithInstance(dbInstance, &dStub.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file:%s", pathToMigrationFiles), databaseConf.Dbname, instance)
	if err != nil {
		fmt.Println("error", err)
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	log.Println("migration done")

	return nil
}
