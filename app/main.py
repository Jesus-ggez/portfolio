from flask import Flask, render_template
import os


app: Flask = Flask(__name__)


def get_css_files() -> list[str]:
    files: list[str] = []

    for file in os.listdir('./static/css/'):
        if os.path.isdir(file): continue

        if not file.endswith('.css'): continue

        file = '/static/css/' + file
        files.append(file)


    return files

@app.route('/')
def index():
    return render_template('index.html', dir_css=get_css_files())


app.run(port=3000, debug=True)
