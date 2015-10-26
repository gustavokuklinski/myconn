# MyConn
Small lib to connect MySQL using go-sql-driver/mysql

### Example

```go
package main

import (
	"log"

	"github.com/gustavokuklinski/myconn"
)

func main() {
  // Open connection to your MySQL database
	db := myconn.DbConn("YOUR USERNAME", "YOUR PASSWORD", "YOUR DATABASE")

  // Build a Query
	_, err := db.Query("YOUR SQL QUERY")

  // Verify your query
	if err != nil {
		panic(err.Error())
	}

	log.Println("Succefull conected")

  // Close connection
	defer db.Close()
}
```
