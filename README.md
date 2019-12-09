# MongoDB Go driver  Query Samples 
 Samples in Golang for MongoDb driver - Various Find query examples
 
 The original CRUD examples are taken from here:
 https://dev.to/eduardohitek/mongodb-golang-driver-tutorial-49e5
 
It was very hard to found bson samples for Golang, so during time I started to gather and create this sample collection.

 ## Query examples:
  
 
 Retrieve by int field "age" greater than and equal to 30
 	
 	filter := bson.M{"age": bson. M{"$gte":30}}
 
 	
 Retrieve by int field "age" less than and equal to 30
 
    filter = bson.M{"age": bson. M{"$lte":30}}
    

Retrieve by string field "name" field equal to "Mazdak" (Case sensitive)
	
	filter = bson.M{"name": "Mazdak"}


 Retrieve by string field "name" equal to "Mazdak" (Case In-sensitive) - Regex way
 	
 	name := "^mazdak"
 	regex := bson.M{"$regex": primitive.Regex{Pattern: name, Options : "i"}}
 
 	filter = bson.M{"name": regex}


Like operator. Retrieve by string field "name"

	name = "mazdak"
	regex = bson.M{"$regex": primitive.Regex{Pattern: name, Options : "i"}}

	filter = bson.M{"name": regex}

	
Get documents when a specific does not exists
	
	filter = bson.M{"age": bson. M{"$exists": false}}

	
And operator: "age" greater/equal to "30" and "name" equal to "Anai"

	filter = bson.M{
		"$and": []bson.M{
			bson.M{"name": "Anai"},
			bson. M{"age": bson. M{"$gte":30}},
		},
	}


Or operator: "age" greater/equal to 30 or "name" equal to "Anai"

	filter = bson.M{
		"$or": []bson.M{
			bson.M{"name": "Anai"},
			bson. M{"age": bson. M{"$gte":30}},
		},
	}
	
And-Or operator together 1. "age" less than 10 OR greater than 30, AND "signed" false

	filter = bson.M{
		"$and": []bson.M{
				bson.M{
					"$or": []bson.M{
						bson.M{"age": bson. M{"$lt":10}},
						bson.M{"age": bson. M{"$gt":30}},
					},
				},
				bson.M{"signed": false},
			},
		}

And-Or operator together 2. "age" more than 30 AND "alias" equal "Mazoo" , OR "signed" false

	filter = bson.M{
		"$or": []bson.M{
			bson.M{
				"$and": []bson.M{
					bson.M{"alias": "Mazoo"},
					bson.M{"age": bson. M{"$gt":30}},
				},
			},
			bson.M{"signed": false},
		},
	}
	
Get one Record by ObjectId

    docID, _ := primitive.ObjectIDFromHex("5de30185e4fabe4778f0ffdf")

	hero := ReturnOneHero(c, bson.M{"_id" : docID})

Get documents created from yesterday

	filter = bson. M{"createdDate": bson. M{"$gte": time.Now().AddDate(0,0,-1)}}
	
Get documents updated since last 5 days based on timestamp type field
	
	filter = bson. M{"lastUpdate": bson. M{"$gte": primitive.Timestamp{T:uint32(time.Now().AddDate(0,0,-10).Unix())} }}

Create Piplines:

Stage 1: Join to city document.
Stage 2: Select specific fields with $project
Stage 3: Group on city field and count number of each city on the document.

	stage1 := bson.M{
		"$lookup": bson.M{
			"from": "cities",
			"localField": "cityId",
			"foreignField": "_id",
			"as" : "city",
		},
	}

	stage2 := bson.M{
		"$project": bson.M{
			"_id": 0,
			"name": 1,
			"age": 1,
			"city": bson.M{ "$arrayElemAt" : []interface{}{"$city.name",0} },
		},
	}

	stage3 := bson.M{
		"$group": bson.M{
			"_id" : "$city",
			"count": bson.M{
				"$sum": 1,
			},
		},}

Gain the same result with bson.UnmarshalExtJSON. With this technique, we can directly take guery from MongoDb Compass  (Json stirng), and this method will convert it to proper format.

	pipeline := make([]bson.M, 0)
	err := bson.UnmarshalExtJSON([]byte(strings.TrimSpace(`
		   [{
			"$lookup": {
						"from" : "cities",
						"localField": "cityId",
						"foreignField": "_id",
						"as" : "city"
					}}
			, {
			"$project": {
			  "_id": 0,
			  "name": 1,
			  "age": 1,
			  "city": { "$arrayElemAt" : ["$city.name", 0]}
			}},
			{
			"$group": {
			  "_id" : "$city",
			  "count": {
				"$sum": 1
			  }
			}}]
			`)), true, &pipeline)
