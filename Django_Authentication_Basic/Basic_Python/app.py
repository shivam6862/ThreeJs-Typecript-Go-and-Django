# Lecture 1 (Print)
from new import Student
from math import *
print("Hi , I am shivam kumar", 21)

# Lecture 2 (concate)
name = "shivam"
age = 21
print(name + " Kumar")
# TypeError: can only concatenate str (not "int") to str   print("age" + age)
print("age", age)

# Lecture 3 (String)
print("Hi\n\'everyone\'")
print(name[0])
print(name.upper())
print(name.lower())
print(name.islower())
print(len(name))
print(name.index('a'))
print(name.replace('m', ''))

# Lecture 4 (Number)
print(21117119)
print(6+7)
print(str(21117119) + " shivam")
print(abs(-10))
print(max(5, 7))
print(bin(9))

print(sqrt(9))

# Lecture 5 {user inout}
name = input("Input your name ")
print(name + " Kumar")
age = int(input("Input your age"))
print(type(age))

# Lecture 6 (Word replacement )
print(name.replace("s", "b"))

# Lecture 7  (lIst)
character = ["aa", "bb", "cc", "dd"]
print(character[0][0])
print(character[1:])
string = list(("qq", "ww", 3, True))
print(string)

# Lecture 8 (List Method)
character.extend(string)
print(character)
print(character.count("aa"))

# Lecture 9 (Tuples) {unmutable data struture}
three_numer = (1, 2, 3, 8, "l")
three_numer_1 = tuple((1, 2, 3, 8, "l"))
print(three_numer)
print(three_numer_1)

# Lecture 10 (Function)


def Helper(name):
    print("Hello ", name)


Helper("Shivam")


def Helper_1(*name):
    print("Hello ", name)


Helper_1("Shivam", "Roshan")

# Lecture 11 (Return)


def helper_2():
    return 9+5


print(helper_2())

# Lecture 12 (if else )
if 2 == 3 or 5 == 6:
    print(True)
elif 3 == 3 and 4 == 4:
    print(4)
else:
    print(False)

# Lecture 13 (Dictionaries)
data = {
    "hotal": "By Hyderabad Biryani Hotel",
    "star": 3.7,
    "time": "Opens next at 10 am, tomorrow",
    "name": "Chicken Biryani",
    "price": 70,
    "para": "A classic dish made by layering rice over slow cooked chicken gavy and juicy pieces to keep you satisfied!",
    "image": "/swiggey/Search/Big/1a.jpg",
    "Dtime": "23 MINS",
    "link": "/Restaurants/ifc",
    "itemId": 76543
}
print(data["star"])

# Lecture 14 (while loop)
i = 0
while i < 4:
    print(i)
    i = i+1

# Lecture 15 (for loop)
for letter in "Shiv":
    print(letter)
    if letter == 'h':
        break

for i in range(5):
    print(i+1)

# Nested Loop
# Basic Calculator

# lecture 16 (try and except)
try:
    a = int(input("Enter a number: "))
    print(a)
except:
    print("Value is not a integer")

try:
    a = int(input("Enter a number: "))
    print(a/0)
except ValueError:
    print("Something went , Wrong!")

try:
    a = int(input("Enter a number: "))
    print(a)
except ValueError:
    print("Something went , Wrong!")
else:
    print("Nothing wenrt Wrong")
finally:
    print("Nothing wenrt Wrong")

# Lecture 17 (reading and writing ina file)
state_file = open('state.txt', 'r')
state_file_cut = open('state_copy.txt', 'w')
# new_py = open('new.py','w')
# new_py.write('print(\'This is a new file\')')
print(state_file.readable())
# print(state_file.readlines())
# print(state_file.readlines()[5])
# print(state_file.readline())
n = 65
a = 0
for state in state_file.readlines():
    if state[0] == chr(n) and a < 3:
        state_file_cut.write(state)
        a += 1
        if a == 3:
            a = 0
            n += 1
state_file.close()

# Lecture 18 (class)


class Person:
    def __init__(self, name, age):
        self.name = name
        self.age = age
        # pass continue to pass the error


name = input("Enter your name: ")
age = input("Enter your age: ")
person_1 = Person(name, age)
print(person_1.name)
print(person_1.age)
del person_1.age

# Lecture 19 (Inheritance)


class Person(Student):
    pass


p1 = Person()
print(p1.name)


# pip install virtualenvwrapper-win
#  mkvirtualenv myapp
# pip install django
# django-admin startproject myproject
# deactivate
# workon myapp


# -m venv myapp                           {This is used to create the virtual enviroment}
# myapp\Scripts\activate.bat              {This is used to activate that virtuak enviroment}
# pip install Django                      {Install django in this virtual enviroment}
# django-admin startproject myproject     {Create the starting templete in django}
# deactivate                              {decativate the ongong virtual enviroment}
# myapp\Scripts\activate.bat              {reactivate the vvirtual enviroment in the local enviroment}


# myapp\Scripts\activate.bat               {for run the server first start the local django}
# cd myproject                             {cd to the file in which you want to rin the server}
# python manage.py startapp myapp          {write this thing ti the command promt to create the myapp folder inside the folder}

# make urls.py file inside the myproject/myapp/                                     {folder}
# setup urls.py  {in myapp folder}                                                  {name and function name miust be same}
# setup views.py {in myappfolder}
# python manage.py runserver {command promts write}
# setup urls.py {in myproject folder}
# create a templates folder in the main dir
# setup settings {in myproject folder}
# create a index.html in templates folder

# control + c                              {to killl the server}
# deactivate                               {deactivate the django local server}

# myapp\Scripts\activate.bat               {open the local django}
# cd myproject                             {change the dir}
# python manage.py runserver               {rerun the sever}

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
