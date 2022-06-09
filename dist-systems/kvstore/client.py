from typing import Optional
import urllib.request
import urllib.parse
from urllib.error import HTTPError

import common

"""Makes HTTP requests to the kv server."""

URL = f"http://{common.HOST}:{common.SERVER_PORT}"


def set_value(key: str, value: str):
    query_string = urllib.parse.urlencode({"key": key, "value": value})
    url = f"{URL}?{query_string}"
    req = urllib.request.Request(url, method='PUT')
    urllib.request.urlopen(req)


def get_value(key: str) -> Optional[str]:
    query_string = urllib.parse.urlencode({"key": key})
    url = f"{URL}?{query_string}"
    try:
        with urllib.request.urlopen(url) as response:
            return response.read().decode('utf8')

    except HTTPError as err:
        if err.code == 404:
            return None
        else:
            raise
