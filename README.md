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
