import json
import time

from flask import Flask, request

app = Flask(__name__)

@app.route("/", methods=["GET", "POST"])
def processRoot():
    if request.method == "GET":
        return "ROOT!\n"

    if request.method == "POST":
        result = "Request body start\n"
        result += request.get_data(as_text=True) + "\n"
        result += "Request body end\n"

        try:
            reqBodyDict = json.loads(request.get_data().decode('utf-8'))
        except:
            result += "ERROR: Failed to parse request body JSON!\n"
            return result, 400

        if "test" not in reqBodyDict.keys():
            result += "ERROR: Request body should include \"test\"!\n"
            return result, 400
        result += "Test: %s\n" % reqBodyDict['test']

        if "anotherTest" not in reqBodyDict.keys():
            result += "ERROR: Request body should include \"anotherTest\"!\n"
            return result, 400
        result += "anotherTest: %s\n" % reqBodyDict['anotherTest']

        return result, 200

@app.route("/post/<postId>", methods=["GET"])
def showPost(postId):
    if request.method == "GET":
        return "Post ID: %s\n" % postId

if __name__ == "__main__":
    app.run(host="127.0.0.1", debug=True, threaded=True)
