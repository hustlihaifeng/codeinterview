# 安装及helloworld

1. python:3.4

```shell
yum install https://centos7.iuscommunity.org/ius-release.rpm
yum makecache
yum install python34u python34u-pip

alias pip='pip3'
alias py='python'
alias python='python3'
```

2. flask

```shell
python -m venv flask
. venv/bin/activate
pip install Flask
mkdir app
mkdir app/static
mkdir app/templates
mkdir tmp
```

```shell
cat run.py 
#!flask/bin/python
from app import app
app.run(debug = True, host='0.0.0.0', port=5001)  # 注意这里的端口监听
```

```shell
cat app/__init__.py 
from flask import Flask

app = Flask(__name__)
from app import views
```

```shell
cat app/views.py 
from app import app

@app.route('/')
@app.route('/index')
def index():
    return "Hello, World!"
```



## helloworld

- <https://linuxize.com/post/how-to-install-flask-on-centos-7/>

1. 命令

```shell
mkdir my_flask_app
cd my_flask_app
python3 -m venv venv
source venv/bin/activate
pip install Flask
python -m flask --version
vim hello.py
```

- hello.py
```python
from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return 'Hello World!'
```

```shell
export FLASK_APP=hello
flask run
curl http://127.0.0.1:5000

```

# flask-tutorial

```shell
mkdir flask-tutorial
cd flask-tutorial
python -m venv venv
. venv/bin/activate
pip install Flask
mkdir flaskr
```

## `__init__.py`

1. 两个作用

- it will contain the application factory：Any configuration, registration, and other setup the application needs will happen inside the function
- it tells Python that the `flaskr` directory should be treated as a package

## `flaskr/__init__.py`

```python
import os

from flask import Flask


def create_app(test_config=None):
    # create and configure the app
    # __name__是当前python module名字，这里似乎是flaskr结尾的
    # instance_relative_config=True告诉flask配置文件路径是相对于instance folder的，instance folder在flickr目录之外，可以包含一些不能提交到git的数据，如秘钥、数据库文件。
    app = Flask(__name__, instance_relative_config=True)
    app.config.from_mapping(
        # SECRET_KEY这里设置成字母dev，来方便调试。线上需要替换成随机秘钥。
        SECRET_KEY='dev',
        # DATABASE 指定了SQLite数据库文件存储在哪里
        DATABASE=os.path.join(app.instance_path, 'flaskr.sqlite'),
    )

    if test_config is None:
        # load the instance config, if it exists, when not testing
        # 相对于app.instance_path
        app.config.from_pyfile('config.py', silent=True)
    else:
        # load the test config if passed in
        app.config.from_mapping(test_config)

    # ensure the instance folder exists
    try:
        os.makedirs(app.instance_path)
    except OSError:
        pass

    # a simple page that says hello
    @app.route('/hello')
    def hello():
        return 'Hello, World!'

    return app
```

## 运行

```shell
export FLASK_APP=flaskr
# Development mode shows an interactive debugger whenever a page raises an exception, and restarts the server whenever you make changes to the code.
export FLASK_ENV=development
flask run --host=192.168.56.101 --port=5002
curl http://192.168.56.101:5002/hello
```

## 命令行参数

1. 详见<http://flask.pocoo.org/docs/1.0/cli/#custom-commands>

2. 三种形式：

- 直接创建命令,无映射

```python
import click
from flask import Flask

app = Flask(__name__)

@app.cli.command()
@click.argument('name')
def create_user(name):
    ...
```

```shel
flask create_user admin
```

- 直接创建命令，有映射

```python
from flask.cli import with_appcontext
@click.command('init-db')
@with_appcontext
def init_db_command():
    """Clear the existing data and create new tables."""
    init_db()
    click.echo('Initialized the database.')
```

[`click.command()`](http://click.pocoo.org/api/#click.command) defines a command line command called `init-db` that calls the `init_db` function and shows a success message to the user. 

- 创建命令组

```python
import click
from flask import Flask
from flask.cli import AppGroup

app = Flask(__name__)
user_cli = AppGroup('user')

@user_cli.command('create')
@click.argument('name')
def create_user(name):
    ...

app.cli.add_command(user_cli)
```

```shell
flask user create demo
```

- 关于Application Context

  - 使用`@app.cli.command()` decorator创建的命令will be executed with an application context pushed，, so your command and extensions have access to the app and its configuration. 见第一种形式
  - 如果使用`@click.command()` decorator,也可以使用`from flask.cli import with_appcontext`达到相同的效果。如：

  ```python
  import click
  from flask.cli import with_appcontext
  
  @click.command
  @with_appcontext
  def do_work():
      ...
  
  app.cli.add_command(do_work)
  ```

  

## 初始化时注册db和注册清理函数

```python
def init_app(app):
    app.teardown_appcontext(close_db) # tells Flask to call that function when cleaning up after returning the response
    app.cli.add_command(init_db_command) # adds a new command that can be called with the flask command.
```

```shell
. venv/bin/activate
export FLASK_APP=flaskr
export FLASK_ENV=development
flask init-db
Initialized the database.
```

## 打印请求参数

1. 在有`APP = Flask(__name__)`的文件里面，添加如下代码：

```python
from flask import Flask,request

import logging
logging.basicConfig(level=logging.DEBUG,format='%(asctime)s %(filename)s[line:%(lineno)d][%(levelname)s] %(message)s',datefmt='%a, %d %b %Y %H:%M:%S')

@APP.before_request
def before_request():
logging.info("curl -X{} \"{}\" --data '{}' -H 'Content-Type:{}' '{}'".format(request.method,request.url,request.data,request.headers["Content-Type"]),request.headers)

```
