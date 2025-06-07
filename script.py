from kafka import KafkaProducer
import json
import time
from datetime import datetime

# Kafka setup
producer = KafkaProducer(
    bootstrap_servers="localhost:9092",
    value_serializer=lambda v: json.dumps(v).encode("utf-8"),
    batch_size=16384,         # Default: 16KB; tune as needed
    linger_ms=10              # Wait a few ms to batch more
)

# Config
topic = "notifications"
total_messages = 100_000
batch_size = 1000

# Track overall timing
start_all = time.time()

for batch_start in range(0, total_messages, batch_size):
    messages = []

    # Capture timestamp for the whole batch
    batch_timestamp = datetime.utcnow().isoformat() + "Z"

    for i in range(batch_start, batch_start + batch_size):
        msg = {
            "user_id": str(i % 1000),
            "title": f"Message {i}",
            "message": f"This is message number {i}"
        }

        messages.append(msg)

    # Send all in batch
    batch_start_time = time.time()
    for msg in messages:
        producer.send(topic, msg)
    producer.flush()  # Ensure all are sent

    batch_duration = time.time() - batch_start_time
    print(f"✅ Sent batch {batch_start}-{batch_start + batch_size - 1} in {batch_duration:.2f}s")

total_duration = time.time() - start_all
print(f"🎉 Done sending {total_messages} messages in {total_duration:.2f}s")
