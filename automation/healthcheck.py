import argparse
import json
import sys
import time
from urllib.error import HTTPError, URLError
from urllib.request import urlopen


def check_health(url, retries=5, delay=2):
    for attempt in range(1, retries + 1):
        try:
            with urlopen(url, timeout=5) as response:
                if response.status != 200:
                    raise RuntimeError(
                        f"Unexpected HTTP {response.status}"
                    )
                data = json.load(response)

            if data.get("status") == "healthy":
                print("PASS: Application is healthy")
                return True

            print(f"Attempt {attempt}: unhealthy response")

        except (
            HTTPError, URLError, TimeoutError,
            ValueError, RuntimeError
        ) as error:
            print(f"Attempt {attempt}: {error}")

        if attempt < retries:
            time.sleep(delay)

    return False


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument(
        "--url",
        default="http://localhost:8081/health"
    )
    args = parser.parse_args()

    if not check_health(args.url):
        print("FAIL: Application health check")
        return 1

    return 0


if __name__ == "__main__":
    sys.exit(main())
