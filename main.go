// package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"strings"
// 	"time"

// 	_ "github.com/lib/pq"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	// MongoDB connection
// 	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jerald:jeraldjero0602@cluster0.yzki9.mongodb.net/agent_details?retryWrites=true&w=majority&appName=Cluster0"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	err = mongoClient.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer mongoClient.Disconnect(ctx)
// 	mongoCollection := mongoClient.Database("agent_details").Collection("agents")

// 	// CockroachDB connection
// 	connStr := "postgresql://root@localhost:26257/agent_details?sslmode=disable"
// 	cockroachDB, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cockroachDB.Close()

// 	// Set up change stream to listen for inserts
// 	pipeline := mongo.Pipeline{
// 		bson.D{{"$match", bson.D{{"operationType", "insert"}}}},
// 	}
// 	opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
// 	changeStream, err := mongoCollection.Watch(context.Background(), pipeline, opts)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer changeStream.Close(context.Background())

// 	for changeStream.Next(context.Background()) {
// 		var changeEvent struct {
// 			FullDocument bson.M `bson:"fullDocument"`
// 		}
// 		if err := changeStream.Decode(&changeEvent); err != nil {
// 			log.Fatal(err)
// 		}

// 		document := changeEvent.FullDocument

// 		// Check for duplicate email
// 		var count int
// 		err := cockroachDB.QueryRow("SELECT COUNT(*) FROM agent WHERE email = $1", document["email"]).Scan(&count)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if count > 0 {
// 			log.Printf("Duplicate email found, skipping document: %v\n", document["email"])
// 			continue
// 		}

// 		columns := make([]string, 0, len(document))
// 		values := make([]interface{}, 0, len(document))
// 		placeholders := make([]string, 0, len(document))
// 		i := 1
// 		for key, value := range document {
// 			// Quote column names with double quotes
// 			columns = append(columns, fmt.Sprintf("\"%s\"", key))
// 			// Convert primitive.ObjectID to string
// 			if oid, ok := value.(primitive.ObjectID); ok {
// 				value = oid.Hex()
// 			}
// 			values = append(values, value)
// 			placeholders = append(placeholders, fmt.Sprintf("$%d", i))
// 			i++
// 		}
// 		insertQuery := fmt.Sprintf("INSERT INTO agent (%s) VALUES (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
// 		_, err = cockroachDB.Exec(insertQuery, values...)
// 		if err != nil {
// 			log.Printf("Error inserting document %v: %v\n", document["_id"], err)
// 		}
// 	}

// 	if err := changeStream.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }

// package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"strings"
// 	"time"

// 	_ "github.com/lib/pq"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// func main() {
// 	// MongoDB connection
// 	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://jerald:jeraldjero0602@cluster0.yzki9.mongodb.net/agent_details?retryWrites=true&w=majority&appName=Cluster0"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	err = mongoClient.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer mongoClient.Disconnect(ctx)

// 	mongoCollectionAgents := mongoClient.Database("agent_details").Collection("agents")
// 	mongoCollectionCompany := mongoClient.Database("agent_details").Collection("company")

// 	// CockroachDB connection
// 	connStr := "postgresql://root@localhost:26257/agent_details?sslmode=disable"
// 	cockroachDB, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cockroachDB.Close()

// 	// Function to handle migration of documents
// 	migrateDocument := func(mongoCollection *mongo.Collection, tableName string, uniqueColumn string) {
// 		// Set up change stream to listen for inserts
// 		pipeline := mongo.Pipeline{
// 			bson.D{{"$match", bson.D{{"operationType", "insert"}}}},
// 		}
// 		opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
// 		changeStream, err := mongoCollection.Watch(context.Background(), pipeline, opts)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer changeStream.Close(context.Background())

// 		for changeStream.Next(context.Background()) {
// 			var changeEvent struct {
// 				FullDocument bson.M `bson:"fullDocument"`
// 			}
// 			if err := changeStream.Decode(&changeEvent); err != nil {
// 				log.Fatal(err)
// 			}
// 			document := changeEvent.FullDocument

// 			// Convert primitive.ObjectID to string before checking for duplicates
// 			if oid, ok := document[uniqueColumn].(primitive.ObjectID); ok {
// 				document[uniqueColumn] = oid.Hex()
// 			}

// 			// Check for duplicate entry
// 			var count int
// 			err := cockroachDB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1", tableName, uniqueColumn), document[uniqueColumn]).Scan(&count)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			if count > 0 {
// 				log.Printf("Duplicate %s found, skipping document: %v\n", uniqueColumn, document[uniqueColumn])
// 				continue
// 			}

// 			columns := make([]string, 0, len(document))
// 			values := make([]interface{}, 0, len(document))
// 			placeholders := make([]string, 0, len(document))
// 			i := 1
// 			for key, value := range document {
// 				columns = append(columns, fmt.Sprintf("\"%s\"", key))
// 				// Convert primitive.ObjectID to string
// 				if oid, ok := value.(primitive.ObjectID); ok {
// 					value = oid.Hex()
// 				}
// 				values = append(values, value)
// 				placeholders = append(placeholders, fmt.Sprintf("$%d", i))
// 				i++
// 			}
// 			insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
// 			_, err = cockroachDB.Exec(insertQuery, values...)
// 			if err != nil {
// 				log.Printf("Error inserting document %v: %v\n", document["_id"], err)
// 			}
// 		}
// 		if err := changeStream.Err(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	// Function to migrate existing documents
// 	migrateExistingDocuments := func(mongoCollection *mongo.Collection, tableName string, uniqueColumn string) {
// 		cursor, err := mongoCollection.Find(context.Background(), bson.M{})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer cursor.Close(context.Background())

// 		for cursor.Next(context.Background()) {
// 			var document bson.M
// 			if err := cursor.Decode(&document); err != nil {
// 				log.Fatal(err)
// 			}

// 			// Convert primitive.ObjectID to string before checking for duplicates
// 			if oid, ok := document[uniqueColumn].(primitive.ObjectID); ok {
// 				document[uniqueColumn] = oid.Hex()
// 			}

// 			// Check for duplicate entry
// 			var count int
// 			err := cockroachDB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1", tableName, uniqueColumn), document[uniqueColumn]).Scan(&count)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			if count > 0 {
// 				log.Printf("Duplicate %s found, skipping document: %v\n", uniqueColumn, document[uniqueColumn])
// 				continue
// 			}

// 			columns := make([]string, 0, len(document))
// 			values := make([]interface{}, 0, len(document))
// 			placeholders := make([]string, 0, len(document))
// 			i := 1
// 			for key, value := range document {
// 				columns = append(columns, fmt.Sprintf("\"%s\"", key))
// 				// Convert primitive.ObjectID to string
// 				if oid, ok := value.(primitive.ObjectID); ok {
// 					value = oid.Hex()
// 				}
// 				values = append(values, value)
// 				placeholders = append(placeholders, fmt.Sprintf("$%d", i))
// 				i++
// 			}
// 			insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
// 			_, err = cockroachDB.Exec(insertQuery, values...)
// 			if err != nil {
// 				log.Printf("Error inserting document %v: %v\n", document["_id"], err)
// 			}
// 		}
// 	}

// 	// Migrate existing documents from company collection
// 	migrateExistingDocuments(mongoCollectionCompany, "company", "_id")

// 	// Migrate agents collection
// 	go migrateDocument(mongoCollectionAgents, "agent", "email")

// 	// Migrate company collection
// 	go migrateDocument(mongoCollectionCompany, "company", "_id")

// 	// Wait indefinitely
// 	select {}
// }

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://app-user-dev:wUYvZfvoWHFIB2s0@xeniapp-cluster-dev.vswr9.mongodb.net/?retryWrites=true&w=majority&appName=xeniapp-cluster-dev"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = mongoClient.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	mongoCollectionAgencies := mongoClient.Database("xeni-db-dev").Collection("agencies")
	//mongoCollectionCompany := mongoClient.Database("agent_details").Collection("company")

	// CockroachDB connection
	connStr := "postgresql://root@localhost:26257/agent_details?sslmode=disable"
	cockroachDB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer cockroachDB.Close()

	// Define schema creation SQL statements
	createAgentsTable := `
    CREATE TABLE IF NOT EXISTS agents (
        _id UUID PRIMARY KEY,
        name STRING NOT NULL,
        email STRING UNIQUE NOT NULL,
        phone STRING,
        address STRING,
        created_at TIMESTAMPTZ DEFAULT now(),
        updated_at TIMESTAMPTZ DEFAULT now()
    );`

	createCompanyTable := `
    CREATE TABLE IF NOT EXISTS company (
        _id UUID PRIMARY KEY,
        name STRING NOT NULL,
        email STRING UNIQUE NOT NULL,
        address STRING,
        phone STRING,
        website STRING,
        created_at TIMESTAMPTZ DEFAULT now(),
        updated_at TIMESTAMPTZ DEFAULT now()
    );`

	// Create tables if they do not exist
	if _, err := cockroachDB.Exec(createAgentsTable); err != nil {
		log.Fatal("Error creating agents table: ", err)
	}
	if _, err := cockroachDB.Exec(createCompanyTable); err != nil {
		log.Fatal("Error creating company table: ", err)
	}

	// Function to handle migration of documents
	migrateDocument := func(mongoCollection *mongo.Collection, tableName string, uniqueColumn string) {
		// Set up change stream to listen for inserts
		pipeline := mongo.Pipeline{
			bson.D{{"$match", bson.D{{"operationType", "insert"}}}},
		}
		opts := options.ChangeStream().SetFullDocument(options.UpdateLookup)
		changeStream, err := mongoCollection.Watch(context.Background(), pipeline, opts)
		if err != nil {
			log.Fatal(err)
		}
		defer changeStream.Close(context.Background())

		for changeStream.Next(context.Background()) {
			var changeEvent struct {
				FullDocument bson.M `bson:"fullDocument"`
			}
			if err := changeStream.Decode(&changeEvent); err != nil {
				log.Fatal(err)
			}
			document := changeEvent.FullDocument

			// Convert primitive.ObjectID to string before checking for duplicates
			if oid, ok := document[uniqueColumn].(primitive.ObjectID); ok {
				document[uniqueColumn] = oid.Hex()
			}

			// Check for duplicate entry
			var count int
			err := cockroachDB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1", tableName, uniqueColumn), document[uniqueColumn]).Scan(&count)
			if err != nil {
				log.Fatal(err)
			}
			if count > 0 {
				log.Printf("Duplicate %s found, skipping document: %v\n", uniqueColumn, document[uniqueColumn])
				continue
			}

			columns := make([]string, 0, len(document))
			values := make([]interface{}, 0, len(document))
			placeholders := make([]string, 0, len(document))
			i := 1
			for key, value := range document {
				columns = append(columns, fmt.Sprintf("\"%s\"", key))
				// Convert primitive.ObjectID to string
				if oid, ok := value.(primitive.ObjectID); ok {
					value = oid.Hex()
				}
				values = append(values, value)
				placeholders = append(placeholders, fmt.Sprintf("$%d", i))
				i++
			}
			insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
			_, err = cockroachDB.Exec(insertQuery, values...)
			if err != nil {
				log.Printf("Error inserting document %v: %v\n", document["_id"], err)
			}
		}
		if err := changeStream.Err(); err != nil {
			log.Fatal(err)
		}
	}

	// Function to migrate existing documents
	migrateExistingDocuments := func(mongoCollection *mongo.Collection, tableName string, uniqueColumn string) {
		cursor, err := mongoCollection.Find(context.Background(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.Background())

		for cursor.Next(context.Background()) {
			var document bson.M
			if err := cursor.Decode(&document); err != nil {
				log.Fatal(err)
			}

			// Convert primitive.ObjectID to string before checking for duplicates
			if oid, ok := document[uniqueColumn].(primitive.ObjectID); ok {
				document[uniqueColumn] = oid.Hex()
			}

			// Check for duplicate entry
			var count int
			err := cockroachDB.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s = $1", tableName, uniqueColumn), document[uniqueColumn]).Scan(&count)
			if err != nil {
				log.Fatal(err)
			}
			if count > 0 {
				log.Printf("Duplicate %s found, skipping document: %v\n", uniqueColumn, document[uniqueColumn])
				continue
			}

			columns := make([]string, 0, len(document))
			values := make([]interface{}, 0, len(document))
			placeholders := make([]string, 0, len(document))
			i := 1
			for key, value := range document {
				columns = append(columns, fmt.Sprintf("\"%s\"", key))
				// Convert primitive.ObjectID to string
				if oid, ok := value.(primitive.ObjectID); ok {
					value = oid.Hex()
				}
				values = append(values, value)
				placeholders = append(placeholders, fmt.Sprintf("$%d", i))
				i++
			}
			insertQuery := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))
			_, err = cockroachDB.Exec(insertQuery, values...)
			if err != nil {
				log.Printf("Error inserting document %v: %v\n", document["_id"], err)
			}
		}
	}

	// Migrate existing documents from company collection
	migrateExistingDocuments(mongoCollectionAgencies, "agents", "_id")

	// Migrate agents collection
	// go migrateDocument(mongoCollectionAgents, "agent", "email")

	// Migrate company collection
	// go migrateDocument(mongoCollectionCompany, "company", "_id")

	// Wait indefinitely
	select {}
}
