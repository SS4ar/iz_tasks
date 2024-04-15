from flask import Flask, redirect, render_template, render_template_string, request, session, url_for
import os
import io

app = Flask(__name__)


@app.route("/images")
def imgs():
    filename = request.args.get('file')
    filename = "static/" + filename
    if os.path.exists(filename) and os.path.isfile(filename):
        with io.open(filename, 'rb') as f:
            file = f.read()
    else:
        file = f"File {filename} is not found"

    return file


if __name__ == "__main__":
    app.run(host="0.0.0.0", port="5000", debug=False)
