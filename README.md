# MongoDB Go driver  Query Samples 
 Samples in Golang for MongoDb driver - Various Find query examples
 
 The original CRUD examples are taken from here:
 https://dev.to/eduardohitek/mongodb-golang-driver-tutorial-49e5
 
It was very hard to found bson samples for Golang, so during time I started to gather and create this sample collection.

 ## Query examples:
  
 
 Retrieve by int field "age" greater than and equal to 30
 	
 	filter := bson. M{"age": bson. M{"$gte":30}}
 
 	heroes := ReturnAllHeroes(c, filter)
 	for _, hero := range heroes {
 		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
 	}
 	
 Retrieve by int field "age" less than and equal to 30
 
    filter = bson. M{"age": bson. M{"$lte":30}}
    
    heroes = ReturnAllHeroes(c, filter)
    for _, hero := range heroes {
    	log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
    }

Retrieve by string field "name" field equal to "Mazdak" (Case sensitive)
	
	filter = bson.M{"name": "Mazdak"}

	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

 Retrieve by string field "name" equal to "Mazdak" (Case In-sensitive) - Regex way
 	
 	name := "^mazdak"
 	regex := bson.M{"$regex": primitive.Regex{Pattern: name, Options : "i"}}
 
 	filter = bson. M{"name": regex}
 
 	heroes = ReturnAllHeroes(c, filter)
 	for _, hero := range heroes {
 		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
 	}

Like operator. Retrieve by string field "name"

	name = "mazdak"
	regex = bson.M{"$regex": primitive.Regex{Pattern: name, Options : "i"}}

	filter = bson. M{"name": regex}

	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}
	
Get documents when a specific does not exists
	
	filter = bson. M{"age": bson. M{"$exists": false}}

	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}
