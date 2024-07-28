# memq

Multi-writer single consumer in-memory queue (fast and simple).

## Run

```bash
SUB=https://api.example.com/msg KEY=your_api_key memq
```

## Settings

Should be passed as env vars.

```text
SUB  = https://api.example.com/msg  # consumer's address
KEY  = your_api_key                 # authentication token
ADDR = localhost:5359               # local binding address
QCAP = 1024                         # queue capacity
```

## Sending Messages

```http
POST https://queue.example.com/msg
Content-Type: ...

Your message in the body
```

The `Content-Type` is going to be forwarded to the consumer, so be mindful of
that. When sending bytes, use `application/octet-stream`, for JSON data - use
`application/json`.

Anyways, I'm not here to teach you - in any uncertain situation - Google your
way out of it like you always do 😊
