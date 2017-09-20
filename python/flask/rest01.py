# From https://sourcedexter.com/python-rest-api-flask/

import datetime
import json
import time

from flask import Flask, request

app = Flask(__name__)

@app.route("/timestamp", methods=["GET"])
def get_timestamp():
    currentTime = time.time()
    timeMillis = int(round(currentTime * 1000))
    return str(timeMillis)

@app.route("/datetime", methods=["POST"])
def get_datetime():
    reqBody = json.loads(request.data)
    millisTime = int(reqBody["millis"])
    dateTime = datetime.datetime.fromtimestamp(millisTime / 1000)
    return str(dateTime)

if __name__ == "__main__":
    app.run(host="0.0.0.0", threaded=True)
