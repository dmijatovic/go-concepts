package modb

import (
	context "context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodb string = "mongodb://root:changeme@localhost:27017"
var mc mongo.Client
var db mongo.Collection

// BlogItem object to add
type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func mongoCnn() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongodb))
	if err != nil {
		return client, err
	}
	if err != nil {
		return client, err
	}
	// log.Println("Connected to mongodb...", mongodb)
	return client, nil
}

func mongoDB(mc *mongo.Client) (*mongo.Collection, error) {
	//concet to todo
	// err := mc.Connect(context.TODO())
	// if err != nil {
	// 	return nil, err
	// }
	db := mc.Database("dv4all").Collection("blog")
	return db, nil
}

// Connect will connect to MongoDB
func Connect() {
	// connect to mongo
	mc, err := mongoCnn()
	// log.Println(mc)
	if err != nil || mc == nil {
		// we panic
		log.Panic(err)
	}
	db, e := mongoDB(mc)
	if e != nil || db == nil {
		// we panic
		log.Panic(err)
	}

	log.Println("Connected to MongoDB...", mongodb)
}

// Close will close the connection to MongoDB. Use when closing the app
func Close() {
	//close mongoDB
	err := mc.Disconnect(context.TODO())
	if err != nil {
		log.Panic(err)
	}
	log.Print("Closed MongoDB connection...")
}

// CreateBlog will create new blog in database.
// It returns blog item with ID assigned
func CreateBlog(item interface{}) (BlogItem, error) {

	log.Println("Create blog in database...", item)

	res, err := db.InsertOne(context.Background(), bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Println(err)
	}
	id := res.InsertedID
	log.Println(id)
	// // instert one item to mongodb
	// res, err := db.InsertOne(context.Background(), item)
	// if err != nil {
	// 	log.Println("Failed to add: ", err)
	// 	return BlogItem{}, err
	// }
	// // extract new uuid
	// oid, ok := res.InsertedID.(primitive.ObjectID)
	// if !ok {
	// 	return BlogItem{}, status.Errorf(
	// 		codes.Internal,
	// 		fmt.Sprintf("Failed to convert UID"),
	// 	)
	// }
	// //assign ID
	// // item.ID = oid
	// log.Print("UID: ", oid)
	//return item
	return BlogItem{}, nil
}
