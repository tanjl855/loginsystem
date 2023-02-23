package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"loginsystem/Log"
	"time"
)

func main() {
	Log.Info = Log.InitLogLog()     //初始化info日志系统
	Log.ErrorLog = Log.InitErrLog() //初始化err日志系统
	// 设置客户端连接配置
	virtualHost := "172.10.2.244"
	user := "cny"
	password := "123"
	// localHost := "localhost"
	// host := localHost
	host := virtualHost
	port := "27017"
	mongoURI := "mongodb://" + user + ":" + password + "@" + host + ":" + port
	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	// 连接到MongoDB
	client, err := mongo.Connect(ctx, clientOptions) //ctx -> new(emptyCtx)  emptyCtx -> int
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	Log.Info.Println("Connected to MongoDB!")
	collection := client.Database("MongoDB_test").Collection("student")
	// Find all documents in which the "name" field is "tanjl".
	// Specify the Sort option to sort the returned documents by age in ascending order.
	opts := options.Find().SetSort(bson.M{"_id": 1})
	cursor, err := collection.Find(ctx, bson.M{ /*"name": "tanjl"*/ }, opts) //bson.D{{"name", "tjl"}}
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	var results []bson.D
	if err = cursor.All(ctx, &results); err != nil {
		cursor.Close(ctx)
		Log.ErrorLog.Fatal(err)
	}
	defer cursor.Close(ctx)
	for _, result := range results {
		Log.Info.Println(result)
	}

	// CRUD
	collection = client.Database("MongoDB_test").Collection("student")
	type Student struct {
		Name string
		Age  int
	}
	s1 := Student{"小红", 12}
	s2 := Student{"小兰", 10}
	s3 := Student{"小黄", 11}
	// 批量创建文档
	students := []interface{}{s1, s2, s3}
	insertManyResult, err := collection.InsertMany(ctx, students)
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	Log.Info.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	// 更新文档
	filter := bson.D{{"name", "小兰"}}
	update := bson.D{
		{"$inc", bson.D{
			{"age", 1},
		}},
	}
	updateResult, err := collection.UpdateMany(ctx, filter, update)
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	Log.Info.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	// 查找文档
	// 查找单个文档
	/*
		可以用
		bson.D{{
				"name",
				bson.D{
					{
						"$in",
						bson.A{"张三", "李四"},
					},
				},
			}} 来查找name字段与张三和李四匹配的文档
	*/
	var result Student
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	Log.Info.Printf("Found a single document: %+v\n", result)
	// 查找多个文档
	// 此方法返回一个游标。游标提供了一个文档流，你可以通过它一次迭代和解码一个文档。当游标用完之后，应该关闭游标。
	findOptions := options.Find()
	findOptions.SetLimit(2)
	var resultSlice []*Student
	cur, err := collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	for cur.Next(ctx) {
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			Log.ErrorLog.Fatal(err)
		}
		resultSlice = append(resultSlice, &elem)
	}
	if err := cur.Err(); err != nil {
		Log.ErrorLog.Fatal(err)
	}
	cur.Close(ctx)
	fmt.Printf("Found multiple documents (array of pointers):%#v\n", resultSlice)
	// 删除文档
	// 删除名字是小黄的那个
	collection = client.Database("MongoDB_test").Collection("tjl")
	deleteResult, err := collection.DeleteOne(ctx, bson.D{{}})
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
	// 删除所有
	collection = client.Database("MongoDB_test").Collection("student")
	deleteResult1, err := collection.DeleteMany(ctx, bson.D{{"name", "小黄"}})
	if err != nil {
		Log.ErrorLog.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult1.DeletedCount)
	Log.Info.Println("Connection to MongoDB closed.")
}

// ConnectToDB 连接池模式
func ConnectToDB(uri, name string, timeout time.Duration, num uint64) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	o := options.Client().ApplyURI(uri)
	o.SetMaxPoolSize(num)
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		return nil, err
	}

	return client.Database(name), nil
}
