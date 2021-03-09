package elasticsearch

import (
	"context"
	"github.com/curious-universe/network-traffic-ant/zaplog"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client
var host = "http://elasticsearch:9200"

//初始化
func init() {
	//errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	//这个地方有个小坑 不加上elastic.SetSniff(false) 会连接不上
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	zaplog.S().Infof("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	zaplog.S().Infof("Elasticsearch version %s\n", esversion)
}

func Create(eIndex string, eBody interface{}) {
	put1, err := client.Index().
		Index(eIndex).
		BodyJson(eBody).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	zaplog.S().Info(put1)
	zaplog.S().Infof("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
