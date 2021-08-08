package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

type clientElastic struct {
	client *elasticsearch.Client
}

// NewClient ..
func NewClient(cfg elasticsearch.Config) (Client, error) {
	x := &clientElastic{}
	var err error
	x.client, err = x.connect(cfg)
	return x, err
}

// Connect ...
func (e *clientElastic) connect(cfg elasticsearch.Config) (*elasticsearch.Client, error) {
	return elasticsearch.NewClient(cfg)
}

func (e *clientElastic) Version() (map[string]interface{}, error) {
	var result map[string]interface{}
	res, err := e.client.Info()
	if err != nil {
		return nil, err
	}

	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e *clientElastic) Search(ctx context.Context, index string, query map[string]interface{}) (interface{}, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}
	res, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex(index),
		e.client.Search.WithBody(&buf),
		e.client.Search.WithTrackTotalHits(true),
		e.client.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e *clientElastic) Mapping(ctx context.Context, index string, field interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(field)
	if err != nil {
		return nil, err
	}
	res, err := e.client.Indices.Create(
		index,
		e.client.Indices.Create.WithBody(strings.NewReader(string(data))),
	)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil

}

func (e *clientElastic) Insert(ctx context.Context, index, id string, field interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(field)
	if err != nil {
		return nil, err
	}
	res, err := e.client.Index(
		index,
		strings.NewReader(string(data)),
		e.client.Index.WithContext(ctx),
		e.client.Index.WithDocumentID(id),
	)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil

}

func (e *clientElastic) Delete(ctx context.Context, index, id string) (map[string]interface{}, error) {
	res, err := e.client.Delete(index, id)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
