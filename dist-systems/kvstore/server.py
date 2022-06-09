from http.server import BaseHTTPRequestHandler, HTTPServer

from typing import Optional
import csv
from urllib.parse import urlparse
import urllib.parse

FILE = "store.csv"


def load() -> dict[str, str]:
    try:
        file = open(FILE)
        return {k: v for k, v in csv.reader(file)}
    except OSError:
        return {}


kv = load()


def set_value(key: str, value: str):
    kv[key] = value


def get_value(key: str) -> Optional[str]:
    return kv.get(key)


def flush():
    print("flushing!")
    w = csv.writer(open(FILE, "w"))
    for k, v in kv.items():
        w.writerow([k, v])


class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        query = urlparse(self.path).query
        key = urllib.parse.parse_qs(query).get('key')[0]

        value = kv.get(key)
        if value is None:
            self.send_error(404)
        else:
            self.send_response(200)
            self.send_header('Content-type', 'text/html')
            self.end_headers()
            self.wfile.write(bytes(value, "utf8"))

    def do_PUT(self):
        query = urlparse(self.path).query
        parsed = urllib.parse.parse_qs(query)
        key = parsed['key'][0]
        value = parsed['value'][0]
        kv[key] = value
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()


with HTTPServer(('', 8000), handler) as server:
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        server.shutdown()
        flush()
    print("good night!")
