from flask import Flask, redirect, render_template, render_template_string, request, session, url_for
import yaml

app = Flask(__name__)


@app.route("/serialize", methods=['POST'])
def serialize():
    text = request.data.decode('utf-8')
    try:
        data = yaml.unsafe_load(text)
    except Exception:
        return "error parse text (only english is allowed)"
    try:
        id = int(data['comment']['id'])
    except Exception:
        return "Id is not int"
    try:
        content = data['comment']['text']
    except Exception:
        return "error parse text"
    comment = {"id":id, "text":content }
    return comment


if __name__ == "__main__":
    app.run(host="0.0.0.0", port="5000", debug=False)
