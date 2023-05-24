# pip install virtualenvwrapper-win
#  mkvirtualenv myapp
# pip install django
# django-admin startproject myproject
# deactivate
# workon myapp


# -m venv myapp     {This is used to create the virtual enviroment}
# myapp\Scripts\activate.bat  {This is used to activate that virtuak enviroment}
# pip install Django    {Install django in this virtual enviroment}
# django-admin startproject myproject     {Create the starting templete in django}
# deactivate   {decativate the ongong virtual enviroment}
# myapp\Scripts\activate.bat   {reactivate the vvirtual enviroment in the local enviroment}

# myapp\Scripts\activate.bat
# cd myproject
# python manage.py startapp myapp

# make urls.py file inside the myproject/myapp/          {folder}
# setup urls.py  {in myapp folder}                                                  {name and function name miust be same}
# setup views.py {in myappfolder}
# python manage.py runserver {command promts write}
# setup urls.py {in myproject folder}
# create a templates folder in the main dir
# setup settings {in myproject folder}
# create a index.html in templates folder

# control + c
# deactivate

# myapp\Scripts\activate.bat    {open the local django}
# cd myproject                  {change the dir}
# python manage.py runserver    {rerun the sever}

# create a static folde for css file
# go to the setting of myproject and set the static folder setting
# lern how to limk the static file


# models.py file
# view.py file same folder
# index.html

# set up the database
# db.sqlite3 file in the main dir
# change the model.py file
# change in the setting.py file


# open a new cmd terminal
# got to the myproject dir

# type         python manage.py makemigrations     {the is use to update any change made in the models file}
# teh type    python manage.py migrate             {Note this both command we hab=ve to always type whwn ever we change the models.py file}
#                                                  {This will affect the db.sqlite3 file}


# python manage.py createsuperuser                  {This is use to create the super user for login}
# change the file admin.py file
# login to the django web page
# add the data in fearture create in webpage of django
# update the data in your created FEATURE folder
# go to the views.py file and change the code there


# making register page
# urls.py setup
# views.py setup
# create a register.html file
# create a register.css file
# write the function id views.py files

# making login page
# urls.py setup
# views.py setup
# create a login.html file
# create a login.css file
# write the function id views.py files

# setting anchor tag for login , logout
# write the function for logout

# setting dynamic url


# download postgresql
# download pgadmin
# mkae a database in pgadmin
# setup the database in setting.py file {DATABASE PART}
# pip install psycopg2
# pip install pillow
# python manage.py makemigrations
# python manage.py migrate
