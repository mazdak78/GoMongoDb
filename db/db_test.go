package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestReturnAllHeroes(t *testing.T) {

	c := GetClient()


	//get age greater than and equal to 30
	filter := bson. M{"age": bson. M{"$gte":30}}

	log.Println("Get documents with age greater than and equal to 30:")
	heroes := ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")

	//get age less than and equal to 30
	filter = bson. M{"age": bson. M{"$lte":30}}

	log.Println("Get documents with age less than and equal to 30:")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")


	//Get documents name equal to Mazdak(Case sensitive)
	filter = bson.M{"name": "Mazdak"}

	log.Println("Get documents name equal to Mazdak(Case sensitive)")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")

	//Get documents name equal to Mazdak(Case In-sensitive) - Regex way
	name := "^mazdak"
	regex := bson.M{"$regex": primitive.Regex{Pattern: name, Options : "i"}}

	filter = bson. M{"name": regex}

	log.Println("Get documents name equal to Mazdak(Case In-sensitive) - Regex way")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")

	//Get documents: name like 'mazdak'
	name = "mazdak"
	regex = bson.M{"$regex": primitive.Regex{Pattern: name, Options : "i"}}

	filter = bson. M{"name": regex}

	log.Println("Get documents: name like 'mazdak' ")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")


	//Get documents: age field does not exists
	filter = bson. M{"age": bson. M{"$exists": false}}

	log.Println("Get documents: age field does not exists")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}


	log.Println("--------------------------")

	//And operator: age greater/equal to 30 and name equal to Anai
	filter = bson.M{
		"$and": []bson.M{
			bson.M{"name": "Anai"},
			bson. M{"age": bson. M{"$gte":30}},
		},
	}

	log.Println("And operator: age greater/equal to 30 and name equal to Anai")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")

	//Or operator: age greater/equal to 30 or name equal to Anai
	filter = bson.M{
		"$or": []bson.M{
			bson.M{"name": "Anai"},
			bson. M{"age": bson. M{"$gte":30}},
		},
	}

	log.Println("Or operator: age greater/equal to 30 or name equal to Anai")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}


	log.Println("--------------------------")

	//And-Or operator together 1
	// age less than 10 OR greater than 30, AND "signed" false
	filter = bson.M{
		"$and": []bson.M{
				bson.M{
					"$or": []bson.M{
						bson. M{"age": bson. M{"$lt":10}},
						bson. M{"age": bson. M{"$gt":30}},
					},
				},
				bson.M{"signed": false},
			},
		}

	log.Println("And-Or operator together 1")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")

	//And-Or operator together 2
	// age more than 30 AND alias equal Mazoo , OR "signed" false
	filter = bson.M{
		"$or": []bson.M{
			bson.M{
				"$and": []bson.M{
					bson. M{"alias": "Mazoo"},
					bson. M{"age": bson. M{"$gt":30}},
				},
			},
			bson.M{"signed": false},
		},
	}

	log.Println("And-Or operator together 2")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")


	//time
	// get docs created from yesterday
	filter = bson. M{"createdDate": bson. M{"$gte": time.Now().AddDate(0,0,-1)}}

	log.Println("Get docs created from yesterday")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")

	//timestamp
	// get docs updated since last 5 days
	filter = bson. M{"lastUpdate": bson. M{"$gte": primitive.Timestamp{T:uint32(time.Now().AddDate(0,0,-5).Unix())} }}

	log.Println("Get docs created from yesterday")
	heroes = ReturnAllHeroes(c, filter)
	for _, hero := range heroes {
		log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.Id)
	}

	log.Println("--------------------------")

}

func TestReturnOneHero(t *testing.T) {
	c := GetClient()

	docID, _ := primitive.ObjectIDFromHex("5de30166e4fabe4778f0ffde")

	hero := ReturnOneHero(c, bson.M{"_id" : docID})
	log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age, hero.LastUpdate)
	log.Println("--------------------------")
}

func TestInsertNewHero(t *testing.T) {
	c := GetClient()

	hero := Hero{Name: "Mansouri", Alias: "Mazoo 2", Signed: true, Age: 36}
	insertedID := InsertNewHero(c, hero)
	log.Println(insertedID)
	hero = ReturnOneHero(c, bson.M{"alias": "Mazoo"})
	log.Println(hero.Name, hero.Alias, hero.Signed)
	log.Println("--------------------------")
}


func TestRemoveOneHero(t *testing.T) {
	c := GetClient()

	heroesRemoved := RemoveOneHero(c, bson.M{"alias": "Mazdak Mansouri"})
	log.Println("Heroes removed count:", heroesRemoved)
	hero := ReturnOneHero(c, bson.M{"alias": "Doctor Strange"})
	log.Println("Is Hero empty?", hero == Hero{ })
	log.Println("--------------------------")
}

func TestUpdateHero(t *testing.T) {
	c := GetClient()

	//docID, _ := primitive.ObjectIDFromHex("5de30185e4fabe4778f0ffdf")

	heroes := ReturnAllHeroes(c, bson.M{})

	for _, h := range heroes{

		random := rand.Intn(30)

		heroesUpdated:= UpdateHero(c, bson.M{/*"createdDate": time.Now(),*/ "lastUpdate": primitive.Timestamp{T:uint32(time.Now().AddDate(0,0, -random).Unix())} }, bson.M{"_id": h.Id})
		log.Println("Heroes updated count: ", heroesUpdated)

	}


	log.Println("--------------------------")
}
