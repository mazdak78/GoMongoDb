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
