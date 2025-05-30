---
mapped_pages:
  - https://www.elastic.co/guide/en/elasticsearch/client/go-api/current/runningqueries.html
---

# Running queries [runningqueries]

## Request structures [_request_structures]

Each endpoint comes  with a Request type that represents the body of its request. For example, a simple search request for a term "Foo" in the `name` field could be written like this:

```go
search.Request{
    Query: &types.Query{
        Term: map[string]types.TermQuery{
            "name": {Value: "Foo"},
        },
    },
}
```


## Raw JSON [_raw_json]

Lastly if you want to use your own pre-baked JSON queries using templates or even a specific encoder, you can pass the body directly to the `Raw` method of the endpoint:

```go
es.Search().Raw([]byte(`{
  "query": {
    "term": {
      "user.id": {
        "value": "kimchy",
        "boost": 1.0
      }
    }
  }
}`))
```

No further validation or serialization is done on what is sent through this method, setting a payload with this takes precedence over any request structure you may submit before running the query.


