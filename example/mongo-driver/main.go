package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
	"go.opentelemetry.io/otel"

	"github.com/uptrace/opentelemetry-go-extra/otelplay"
)

var tracer = otel.Tracer("app_or_package_name")

func main() {
	ctx := context.Background()

	shutdown := otelplay.ConfigureOpentelemetry(ctx)
	defer shutdown()

	opt := options.Client()
	opt.Monitor = otelmongo.NewMonitor()
	opt.ApplyURI("mongodb://localhost:27017")

	mdb, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err := mdb.Ping(ctx, nil); err != nil {
		log.Fatal(err)
		return
	}

	ctx, span := tracer.Start(ctx, "mongodb-main-span")
	defer span.End()

	if err := run(ctx, mdb.Database("example")); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("trace", otelplay.TraceURL(span))
}

// Copyright (C) MongoDB, Inc. 2017-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

func run(ctx context.Context, db *mongo.Database) error {
	coll := db.Collection("inventory")

	_, err := coll.InsertOne(
		ctx,
		bson.D{
			primitive.E{Key: "item", Value: "canvas"},
			primitive.E{Key: "qty", Value: 100},
			primitive.E{Key: "tags", Value: bson.A{"cotton"}},
			primitive.E{Key: "size", Value: bson.D{
				primitive.E{Key: "h", Value: 28},
				primitive.E{Key: "w", Value: 35.5},
				primitive.E{Key: "uom", Value: "cm"},
			}},
		})
	if err != nil {
		return err
	}

	_, err = coll.Find(
		ctx,
		bson.D{primitive.E{Key: "item", Value: "canvas"}},
	)
	if err != nil {
		return err
	}

	_, err = coll.InsertMany(
		ctx,
		[]interface{}{
			bson.D{
				primitive.E{Key: "item", Value: "journal"},
				primitive.E{Key: "qty", Value: int32(25)},
				primitive.E{Key: "tags", Value: bson.A{"blank", "red"}},
				primitive.E{Key: "size", Value: bson.D{
					primitive.E{Key: "h", Value: 14},
					primitive.E{Key: "w", Value: 21},
					primitive.E{Key: "uom", Value: "cm"},
				}},
			},
			bson.D{
				primitive.E{Key: "item", Value: "mat"},
				primitive.E{Key: "qty", Value: int32(25)},
				primitive.E{Key: "tags", Value: bson.A{"gray"}},
				primitive.E{Key: "size", Value: bson.D{
					primitive.E{Key: "h", Value: 27.9},
					primitive.E{Key: "w", Value: 35.5},
					primitive.E{Key: "uom", Value: "cm"},
				}},
			},
			bson.D{
				primitive.E{Key: "item", Value: "mousepad"},
				primitive.E{Key: "qty", Value: 25},
				primitive.E{Key: "tags", Value: bson.A{"gel", "blue"}},
				primitive.E{Key: "size", Value: bson.D{
					primitive.E{Key: "h", Value: 19},
					primitive.E{Key: "w", Value: 22.85},
					primitive.E{Key: "uom", Value: "cm"},
				}},
			},
		})
	if err != nil {
		return err
	}

	return nil
}
