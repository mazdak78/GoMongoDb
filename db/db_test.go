package db

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
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

}

func TestReturnOneHero(t *testing.T) {
	c := GetClient()

	hero := ReturnOneHero(c, bson.M{"name": "Vision"})
	log.Println(hero.Name, hero.Alias, hero.Signed, hero.Age)
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

	heroesUpdated := UpdateHero(c, bson.M{"age": 39}, bson.M{"alias": "Mazoo"})
	log.Println("Heroes updated count:", heroesUpdated)

	log.Println("--------------------------")
}
