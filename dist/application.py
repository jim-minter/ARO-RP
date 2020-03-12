from flask import Flask, redirect
app = Flask(__name__)

@app.route('/')
def index():
    return redirect("https://azure.microsoft.com/en-us/services/openshift", 302)

@app.route('/<path:path>')
def static_proxy(path):
    return app.send_static_file(path)

if __name__ == '__main__':
    app.run()
