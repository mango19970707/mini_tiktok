package dal

import (
	"context"
	"log"
	"time"

	bson "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Video struct {
	VideoId       int64     `bson:"video_id"`
	UserId        int64     `bson:"user_id"`
	PlayUrl       string    `bson:"play_url"`
	CoverUrl      string    `bson:"cover_url"`
	FavoriteCount int64     `bson:"favorite_count"`
	Favorites     []int64   `bson:"favorites"`
	CommentCount  int64     `bson:"comment_count"`
	Comments      []Comment `bson:"comments, inline"`
	PublishDate   time.Time `bson:"publish_date"`
}

type Comment struct {
	CommentId  int64     `bson:"comment_id"`
	UserId     int64     `bson:"user_id"`
	Content    string    `bson:"content"`
	CreateDate time.Time `bson:"create_date"`
}

func QueryVideosByTime(t time.Time)([]*Video, error){
	// 指定获取要操作的数据集
	collection := MongoCli.Database("video").Collection("videos")
	findOptions := options.Find()
	findOptions.SetLimit(30)//设置一次获取的最大视频数
	results := []*Video{}
	cur, err := collection.Find(context.TODO(), bson.M{"publish_date": bson.M{"$gte": t}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Video
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// 完成后关闭游标
	cur.Close(context.TODO())
	return results, nil
}
