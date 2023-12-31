package messages

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"words-bot/bot"
	"words-bot/db"
	"words-bot/dictionary"
	"words-bot/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SendRandomWord() {
	tgbot := bot.GetBot()
	var wg sync.WaitGroup
	collection, _ := db.GetCollection("users")

	cursor, err := collection.Find(context.TODO(), bson.D{}, options.Find())
	if err != nil {
		fmt.Println(err)
	}

	for cursor.Next(context.TODO()) {
		wg.Add(1)

		var user users.User
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println(err)
		}

		if len(user.Words) > 1 {
			go func() {
				randWordId := user.Words[rand.Intn(len(user.Words))]
				randWord, err := dictionary.GetWordById(randWordId)
				if err != nil {
					fmt.Println(err)
				}

				msg := CardWithActions(randWord, user.TgID)
				tgbot.Send(msg)

			}()
		}
	}
	if err := cursor.Err(); err != nil {
		fmt.Println(err)
	}
	cursor.Close(context.TODO())
}
