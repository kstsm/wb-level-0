package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	k "github.com/kstsm/wb-level-0/producer/internal/kafka"
	"github.com/kstsm/wb-level-0/producer/models"
	"log"
	"time"
)

func Run() {
	p, err := k.NewProducer()
	if err != nil {
		log.Fatalf("failed to create mq p: %v", err)
	}
	defer p.Close()

	for i := 0; ; i++ {
		order := models.Order{
			OrderUID:    uuid.New(),
			TrackNumber: "WBILMTESTTRACK",
			Entry:       "WBIL",
			Delivery: models.Delivery{
				Name:    "Test Testov",
				Phone:   "+9720000000",
				Zip:     "2639809",
				City:    "Kiryat Mozkin",
				Address: "Ploshad Mira 15",
				Region:  "Kraiot",
				Email:   "test@gmail.com",
			},
			Payment: models.Payment{
				Transaction:  "b563feb7b2b84b6test",
				RequestID:    "",
				Currency:     "USD",
				Provider:     "wbpay",
				Amount:       1817,
				PaymentDT:    1637907727,
				Bank:         "alpha",
				DeliveryCost: 1500,
				GoodsTotal:   317,
				CustomFee:    0,
			},
			Items: []models.Item{
				{
					ChrtID:      9934930,
					TrackNumber: "WBILMTESTTRACK",
					Price:       453,
					Rid:         "ab4219087a764ae0btest",
					Name:        "Mascaras",
					Sale:        30,
					Size:        "0",
					TotalPrice:  317,
					NmID:        2389212,
					Brand:       "Vivienne Sabo",
					Status:      202,
				},
			},
			Locale:            "en",
			InternalSignature: "",
			CustomerID:        "test",
			DeliveryService:   "meest",
			ShardKey:          "9",
			SmID:              99,
			DateCreated:       time.Now().UTC(),
			OofShard:          "1",
		}

		data, err := json.Marshal(order)
		if err != nil {
			log.Fatalf("failed to marshal json: %v", err)
		}

		key := fmt.Sprintf("%d", i%3)

		if err = p.Produce(data, []byte(key), "orders"); err != nil {
			fmt.Println("err")
		}
		time.Sleep(time.Second * 5)
	}
}
