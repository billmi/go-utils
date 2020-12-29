package helastic

import (
	"fmt"
	"github.com/olivere/elastic"
	Mi "helastic/myelastic"
	"log"
	"time"
	"testing"
)



// Tweet is a structure used for serializing/deserializing data in Elasticsearch.
type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []int                 `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
	SS       float64               `json:"ss"`
	WW       []byte                `json:"ww"`
}

const mapping = `
{
    "settings":{
        "number_of_shards": 1,
        "number_of_replicas": 0
    },
    "mappings":{
        "tweet":{
            "properties":{
                "user":{
                    "type":"keyword"
                },
                "message":{
                    "type":"text",
                    "store": true,
                    "fielddata": true
                },
                "image":{
                    "type":"keyword"
                },
                "created":{
                    "type":"date"
                },
                "tags":{
                    "type":"keyword"
                },
                "location":{
                    "type":"geo_point"
                },
                "suggest_field":{
                    "type":"completion"
                }
            }
        }
    }
}`


//Demo code
func Test_test(t *testing.T) {

	var vHost = "http://127.0.0.1:9200/"

	es := Mi.OnInitES(vHost)
	if es.Err != nil {
		log.Println(es.Err)
		panic(es.Err)
	}

	var (
		index = "my_index"
		_type = "my_type"
	)

	// Create by mapping json
	res := es.CreateIndex("xxj", mapping)
	//fmt.Println(b)
	var s1 = Tweet{
		Created:  time.Now(),
		Location: "ssss",
		SS:       123.11243,
		WW:       []byte("xxjwxc"),
		Tags:     []int{12, 22},
	}

	res = es.Add(index, _type, "10", s1)
	if !res {
		fmt.Println(es.Err)
	}
	fmt.Println(res)

	var s2 = Tweet{
		Created:  time.Now(),
		Location: "qqqqq",
		SS:       723.12,
		WW:       []byte("ycccc"),
		Tags:     []int{123334, 1214},
	}

	es.UpdateValues(index, _type, "18", s2)
	var ws []map[string]interface{}

	res = es.SearchMap(index, _type, `{"query":{"match_all":{}}}`, &ws)
	if res {
		for _, v := range ws {
			fmt.Println(v)
		}
	}

	es.DelIndex(index)
}