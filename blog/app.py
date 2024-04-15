from functools import wraps
import os
import pymongo
import requests
from bson.objectid import ObjectId
from dataclasses import dataclass
from flask import Flask, redirect, render_template, render_template_string, request, session, url_for


def mongoConnect():
    db_client = pymongo.MongoClient(
        host=os.getenv("MONGO_HOST"),
        username=os.getenv("MONGO_USER"),
        password=os.getenv("MONGO_PASS"),
    )
    current_db = db_client["blog"]
    return current_db


app = Flask(__name__)
app.secret_key = 'g8y348f3h4f34jf93ij4g3u49gh343k40f9k34gj34uif@87fh34fj8347hfg3487fh348jf34hf837fg'

current_db = mongoConnect()
posts_collection = current_db["posts"]
users_collection = current_db["users"]
comments_collection = current_db["comments"]


def auth(f):
    @wraps(f)
    def decorated_function(*args, **kwargs):
        if 'user' not in session:
            session['user'] = False
        return f(*args, **kwargs)
    return decorated_function



@app.route("/login", methods=["GET", "POST"])
def index():
    if request.method == 'POST':
        current_db = mongoConnect()
        users_collection = current_db["users"]

        action = request.form['action']
        login = request.form['login']
        password = request.form['passwd']

        if action == 'signup':
            if users_collection.find_one({"login": login}):
                error_text = "This user is already exist"
                return render_template("index.html", error_text=error_text)

            new_user = {"login": login, "password": password}
            ins_result = users_collection.insert_one(new_user)
        elif action == 'signin':
            if ((users_collection.find_one({"$where": "this.login == '" + login + "' && this.password == '" + password + "'"}))):
                session['user'] = login
                return redirect(url_for('blog'))
            

    return render_template("index.html")

@auth
@app.route("/", methods=["GET", "POST"])
def blog():
    is_user = session.get('user', False)
    if not (is_user):
        return redirect(url_for('index'))

    posts = posts_collection.find()
    return render_template("blog.html", posts=posts)



@auth
@app.route("/profile")
def profile():
    is_user = session.get('user', False)
    if not (is_user):
        return redirect(url_for('index'))
    try:
        username = users_collection.find_one({"login": session['user']})['login']
    except Exception:
        return redirect(url_for('index'))


    flag_1 = ''
    if username == 'admin':
        flag_1 = os.getenv('FLAG')


    return render_template("profile.html", username=username, flag=flag_1)

@auth
@app.route("/delete")
def delete():
    users_collection.delete_one({"login": session['user']})
    session['user'] = ""

    return redirect(url_for(f'index'))


@auth
@app.route("/logout")
def logout():
    session['user'] = ""

    return redirect(url_for(f'index'))


@auth
@app.route("/post", methods=["GET"])
def view():
    is_user = session.get('user', False)
    if not (is_user):
        return redirect(url_for('index'))
    
    id = int(request.args.get("id"))
    post = posts_collection.find_one({"id": id})
    comms = current_db.comments.find({"id": id})

    return render_template("view.html", post=post, comments=comms)



@auth
@app.route("/comment", methods=["POST"])
def comment():
    is_user = session.get('user', False)
    if not (is_user):
        return redirect(url_for('index'))
    try:
        request.data.decode('utf-8')
    except Exception:
        return "Only english text", 500

    r = requests.post(f'http://comment:5000/serialize', data=request.data.decode('utf-8'))
    try:
        comm = r.json()
    except Exception:
        return r.text, 500
    comments_collection.insert_one(comm)
    return "success", 200


@app.route("/images")
def imgs():
    filename = request.args.get('file')
    r = requests.get(f'http://images:5000/images?file={filename}')
    
    return r.content



@app.errorhandler(404)
def page_not_found(e):
    route = request.path.replace("/", "")
    r = requests.get(f'http://notfound:5000/notfound?route={route}')

    return render_template_string(r.text), 404



if __name__ == "__main__":
    app.run(host="0.0.0.0", port="5000", debug=False)
