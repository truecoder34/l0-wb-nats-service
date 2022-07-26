package messageConsumer

import (
	"encoding/json"
	"log"

	"github.com/truecoder34/l0-wb-nats-service/service/models"
)

func ParseTransactionsMessage(messageData []byte) (models.Transaction, error) {
	var tr models.Transaction
	err := json.Unmarshal(messageData, &tr)
	if err != nil {
		// TODO : ADD CHECK TO PREVENT UNEXPECTED DATA PROCESSING
		log.Print(err)
	}

	//tr.CreatedNestedTransaction(server.DB)

	return tr, err
}

func AddTransactionMessageData(tr models.Transaction) {

}
