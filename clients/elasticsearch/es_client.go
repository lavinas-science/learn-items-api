package elasticsearch

import (
	"github.com/olivere/elastic"
	"log"
	"time"
	"os"
	"net/http"

)

func getClient() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elastic.SetHeaders(http.Header{
		  "X-Caller-Id": []string{"..."},
		}),
	  )
}