import random
from datetime import datetime, timedelta

ips = [
    f"192.168.1.{i}" for i in range(1, 255)
]

methods = ["GET", "POST", "PUT", "DELETE"]
paths = [
    "/home", "/login", "/api/users",
    "/api/products", "/dashboard", "/checkout"
]
statuses = [200, 200, 200, 201, 301, 400, 403, 404, 500, 502]


timestamp = datetime.now()

target_size = 1024 * 1024 * 1024
written = 0

with open("access.log", "w", buffering=1024 * 1024) as f:
    while written < target_size:
        timestamp += timedelta(seconds=1)

        log = (
            f'{random.choice(ips)} - - '
            f'[{timestamp.strftime("%d/%b/%Y:%H:%M:%S")} +0000] '
            f'"{random.choice(methods)} {random.choice(paths)} HTTP/1.1" '
            f'{random.choice(statuses)} '
            f'{random.randint(100, 10000)}\n'
        )

        f.write(log)
        written += len(log.encode("utf-8"))

print(f"Generated {written / (1024**3):.3f} GB")