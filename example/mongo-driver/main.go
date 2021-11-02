package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
			{"item", "canvas"},
			{"qty", 100},
			{"tags", bson.A{"cotton"}},
			{"size", bson.D{
				{"h", 28},
				{"w", 35.5},
				{"uom", "cm"},
			}},
		})
	if err != nil {
		return err
	}

	_, err = coll.Find(
		ctx,
		bson.D{{"item", "canvas"}},
	)
	if err != nil {
		return err
	}

	_, err = coll.InsertMany(
		ctx,
		[]interface{}{
			bson.D{
				{"item", "journal"},
				{"qty", int32(25)},
				{"tags", bson.A{"blank", "red"}},
				{"size", bson.D{
					{"h", 14},
					{"w", 21},
					{"uom", "cm"},
				}},
			},
			bson.D{
				{"item", "mat"},
				{"qty", int32(25)},
				{"tags", bson.A{"gray"}},
				{"size", bson.D{
					{"h", 27.9},
					{"w", 35.5},
					{"uom", "cm"},
				}},
			},
			bson.D{
				{"item", "mousepad"},
				{"qty", 25},
				{"tags", bson.A{"gel", "blue"}},
				{"size", bson.D{
					{"h", 19},
					{"w", 22.85},
					{"uom", "cm"},
				}},
			},
		})
	if err != nil {
		return err
	}

	return nil
}
