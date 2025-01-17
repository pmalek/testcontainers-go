package mariadb_test

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mariadb"
)

func ExampleRunContainer() {
	// runMariaDBContainer {
	ctx := context.Background()

	mariadbContainer, err := mariadb.RunContainer(ctx,
		testcontainers.WithImage("mariadb:11.0.3"),
		mariadb.WithConfigFile(filepath.Join("testdata", "my.cnf")),
		mariadb.WithScripts(filepath.Join("testdata", "schema.sql")),
		mariadb.WithDatabase("foo"),
		mariadb.WithUsername("root"),
		mariadb.WithPassword(""),
	)
	if err != nil {
		panic(err)
	}

	// Clean up the container
	defer func() {
		if err := mariadbContainer.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
	// }

	state, err := mariadbContainer.State(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(state.Running)

	// Output:
	// true
}
