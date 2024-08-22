# cloud-functions-monorepo

## About

This repository is a sample for managing multiple functions of Cloud Functions (2nd Gen.) in the same monorepo.

This is just a sample and is not suitable for actual operation.

## Local test

1. Copy `.env.template` to `.env` and fill in the values.

2. Copy `.secret.template` to `.secret` and fill in the values.

3. Run the following command to start the local environment.

```bash
docker compose up
```

## How to call

- Function A

```bash
curl localhost:8080/function_a \
  -X POST \
  -H "Content-Type: application/json" \
  -H "ce-id: 123451234512345" \
  -H "ce-specversion: 1.0" \
  -H "ce-time: 2020-01-02T12:34:56.789Z" \
  -H "ce-type: google.cloud.pubsub.topic.v1.messagePublished" \
  -H "ce-source: //pubsub.googleapis.com/projects/test-project/topics/topic_a" \
  -d '{
        "id": "d29ybGQ=func_a",
        "subscription": "projects/test-project/subscriptions/function_a"
      }'
```

- Function B

```bash
curl localhost:8081/function_b \
  -X POST \
  -H "Content-Type: application/json" \
  -H "ce-id: 123451234512345" \
  -H "ce-specversion: 1.0" \
  -H "ce-time: 2020-01-02T12:34:56.789Z" \
  -H "ce-type: google.cloud.pubsub.topic.v1.messagePublished" \
  -H "ce-source: //pubsub.googleapis.com/projects/test-project/topics/topic_b" \
  -d '{
        "id": "d29ybGQ=func_b",
        "subscription": "projects/test-project/subscriptions/function_b"
      }'
```
