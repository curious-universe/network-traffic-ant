package elasticsearch

import (
	"context"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client
var host = "http://192.168.5.193:9200"

//初始化
func init() {
	//errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	_, err = client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Elasticsearch version %s\n", esversion)
}
