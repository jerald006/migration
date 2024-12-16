package main

import (
	"context"
	"log"
	"mongo_cockroach/models"

	"time"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// MongoDB connection
	// mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin-dev:wUYvZfvoWHFIB2s0@xeniapp-cluster-dev.vswr9.mongodb.net/xeni-db-dev?retryWrites=true&w=majority&appName=xeniapp-cluster-dev"))
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://app-user-dev:wUYvZfvoWHFIB2s0@xeniapp-cluster-dev.vswr9.mongodb.net/xeni-db-dev?retryWrites=true&w=majority&appName=xeniapp-cluster-dev"))
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
	// mongoCollectionUsers := mongoClient.Database("xeni-db-dev").Collection("users")
	// mongoCollectionCompany := mongoClient.Database("agent_details").Collection("company")

	// CockroachDB connection
	// connStr := "postgresql://root@localhost:26257/agent_details?sslmode=disable"
	//connStr := "jdbc:postgresql://xeni-falcon-dev-db-user:jXFjF2LuU2nAW-PgkNALbg@xeni-crdb-falcon-dev-5328.j77.aws-us-west-2.cockroachlabs.cloud:26257/xeni-dev"
	// cockroachDB, err := sql.Open("postgres", connStr)
	// if err != nil {
	//  log.Fatal(err)
	// }

	dsn := "host=xeni-crdb-falcon-dev-5328.j77.aws-us-west-2.cockroachlabs.cloud port=26257 user=xeni-falcon-dev-db-user password=jXFjF2LuU2nAW-PgkNALbg dbname=xeni-dev sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Migrate the schema (optional, but recommended)
	// db.AutoMigrate(&models.CockroachDBAgency{})

	//Function to migrate existing documents
	migrateExistingDocuments := func(mongoCollection *mongo.Collection) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		cursor, err := mongoCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(ctx)

		for cursor.Next(ctx) {
			var document models.MongoAgency
			if err := cursor.Decode(&document); err != nil {
				log.Printf("Error decoding document: %v", err)
				continue
			}

			// Check if agency exists
			var agencyCount int64
			if err := db.Model(&models.CockroachDBAgency{}).
				Where("id = ?", document.Id.Hex()).
				Count(&agencyCount).Error; err != nil {
				log.Printf("Error checking existence for agency ID %s: %v", document.Id.Hex(), err)
				continue
			}

			if agencyCount > 0 {
				log.Printf("Skipping existing agency with ID: %s", document.Id.Hex())
				continue
			}

			// Convert and insert the agency document
			cockroachAgency := document.ConvertMongoToCockroach()
			if err := db.Create(&cockroachAgency).Error; err != nil {
				log.Printf("Failed to insert agency %s: %v", document.Id.Hex(), err)
				continue
			}

			log.Printf("Successfully migrated agency: %s", document.Id.Hex())

			// Now handle the agency policies
			var policyCount int64
			if err := db.Model(&models.CockroachDBAgencyPolicy{}).
				Where("agency_id = ?", document.Id.Hex()).
				Count(&policyCount).Error; err != nil {
				log.Printf("Error checking existence for policy agency ID %s: %v", document.Id.Hex(), err)
				continue
			}

			if policyCount > 0 {
				log.Printf("Skipping existing policy for agency with ID: %s", document.Id.Hex())
				continue
			}

			// Convert and insert the policy document
			cockroachPolicy := document.ConvertMongoToCockroachPolicy()
			if err := db.Create(&cockroachPolicy).Error; err != nil {
				log.Printf("Failed to insert policy for agency %s: %v", document.Id.Hex(), err)
			} else {
				log.Printf("Successfully migrated policy for agency: %s", document.Id.Hex())
			}
		}

		if err := cursor.Err(); err != nil {
			log.Printf("Cursor error: %v", err)
		}
	}

	// Function to migrate existing documents for users
	// migrateExistingUsers := func(mongoCollection *mongo.Collection, tableName string) {
	//  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//  defer cancel()

	//  cursor, err := mongoCollection.Find(ctx, bson.M{})
	//  if err != nil {
	//      log.Fatal(err)
	//  }
	//  defer cursor.Close(ctx)

	//  for cursor.Next(ctx) {
	//      var document models.MongoUser
	//      if err := cursor.Decode(&document); err != nil {
	//          log.Printf("Error decoding document: %v", err)
	//          continue
	//      }

	//      // Check if user exists using a direct count query
	//      var count int64
	//      if err := db.Model(&models.CockroachDBUser{}).
	//          Where("id = ?", document.Id).
	//          Count(&count).Error; err != nil {
	//          log.Printf("Error checking existence for ID %s: %v", document.Id, err)
	//          continue
	//      }

	//      if count > 0 {
	//          log.Printf("Skipping existing user with ID: %s", document.Id)
	//          continue
	//      }

	//      // Convert and insert the document
	//      cockroachUser := document.ConvertMongoToCockroachUser()
	//      if err := db.Create(&cockroachUser).Error; err != nil {
	//          log.Printf("Failed to insert user %s: %v", document.Id, err)
	//      }

	//      log.Printf("Successfully migrated user: %s", document.Id)
	//  }

	//  if err := cursor.Err(); err != nil {
	//      log.Printf("Cursor error: %v", err)
	//  }
	// }

	// Migrate existing documents from company collection

	migrateExistingDocuments(mongoCollectionAgencies)

	// Migrate existing documents from user collection

	// migrateExistingUsers(mongoCollectionUsers, "users")

	select {}
}
