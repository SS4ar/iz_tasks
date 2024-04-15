from flask import Flask, redirect, render_template, render_template_string, request, session, url_for

app = Flask(__name__)


@app.route("/notfound")
def notfound():
    route = request.args.get("route")

    template = '''{%% extends "notFoundPage.html" %%}
    {%% block content %%}
        <div class="text">
            <h1>404</h1>  
        <h2>Couldn't launch :(</h2>
            <h3>Page %s Not Found</h3> 
        </div>
    {%% endblock %%}''' % (route)
    return render_template_string(template), 404


if __name__ == "__main__":
    app.run(host="0.0.0.0", port="5000", debug=False)
