package repository

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ObjectId string

const (
	ObjectIdEmpty ObjectId = ""
)

type MongoDBRepositoryModelPointer[T any] interface {
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
	*T
}

type MongoDBRepository[T any, P MongoDBRepositoryModelPointer[T]] struct {
	collection *mongo.Collection
}

func (r *MongoDBRepository[T, P]) InsertOne(ctx context.Context, model P) (ObjectId, bool) {
	t := time.Now().UTC()
	model.SetCreatedAt(t)
	model.SetUpdatedAt(t)
	result, err := r.collection.InsertOne(ctx, model)
	if err != nil {
		log.Error("MongoDB：InsertOne（", model, "）写入集合：", r.collection.Name(), " 失败：", err)
		return ObjectIdEmpty, false
	}
	if objectId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return ObjectId(objectId.Hex()), true
	}
	return ObjectIdEmpty, false
}

func (r *MongoDBRepository[T, P]) UpdateOneById(ctx context.Context, id ObjectId, fields bson.D) (int64, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("MongoDB：UpdateOneById，标识：", id, "，集合：", r.collection.Name(), " 失败：", err)
		return 0, false
	}

	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now().UTC()})
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: fields}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("MongoDB：UpdateOneById，标识：", id, "，集合：", r.collection.Name(), " 失败：", err)
		return 0, false
	}
	return result.MatchedCount, true
}

func (r *MongoDBRepository[T, P]) UpdateOne(ctx context.Context, filter bson.D, fields bson.D) (int64, bool) {
	fields = append(fields, bson.E{Key: "updated_at", Value: time.Now().UTC()})
	update := bson.D{{Key: "$set", Value: fields}}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Error("MongoDB：UpdateOne，筛选条件：", filter, "，集合：", r.collection.Name(), " 失败：", err)
		return 0, false
	}
	return result.MatchedCount, true
}

func (r *MongoDBRepository[T, P]) FindOneById(ctx context.Context, id ObjectId, accumulator P) (P, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("MongoDB：FindOneById，标识：", id, "，集合：", r.collection.Name(), " 失败：", err)
		return nil, false
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	err = r.collection.FindOne(ctx, filter).Decode(accumulator)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Error("MongoDB：FindOneById，标识：", id, "，集合：", r.collection.Name(), " 失败：", err)
		}
		return nil, false
	}
	return accumulator, true
}

func (r *MongoDBRepository[T, P]) FindOne(ctx context.Context, filter bson.D, accumulator P) (P, bool) {
	err := r.collection.FindOne(ctx, filter).Decode(accumulator)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, true
		}
		log.Error("MongoDB：FindOne，筛选条件：", filter, "，集合：", r.collection.Name(), " 失败：", err)
		return nil, false
	}
	return accumulator, true
}

func (r *MongoDBRepository[T, P]) DeleteOneById(ctx context.Context, id ObjectId) (int64, bool) {
	objectId, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		log.Error("MongoDB：DeleteOneById，标识：", id, "，从集合删除：", r.collection.Name(), " 失败：", err)
		return 0, false
	}

	filter := bson.D{{Key: "_id", Value: objectId}}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Error("MongoDB：DeleteOneById，标识：", id, "，从集合删除：", r.collection.Name(), " 失败：", err)
		return 0, false
	}
	return result.DeletedCount, true
}
