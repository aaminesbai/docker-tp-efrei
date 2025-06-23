from http.server import BaseHTTPRequestHandler, HTTPServer

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == '/':
            self.send_response(200)
            self.end_headers()
            self.wfile.write(b'Hello from custom app!')
        else:
            self.send_error(404)

if __name__ == '__main__':
    server = HTTPServer(('', 8000), Handler)
    server.serve_forever()
