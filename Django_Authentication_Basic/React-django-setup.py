# $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$            BACKEND SERVER            $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

# process to setup django in backend
# select a destop as a directroy
# pip install virtualenv
# virtualenv env
# env\scripts\activate
# pip install django
# django-admin startproject projectName
# take the virtualenv folder and put it into the projectName folder
# cd projectName
# python manage.py runserver

# python manage.py startapp {NewFolderName inside projectName}

# update the setting file in projectName folder like this in the INSTALLED_APPS part 'NewFolderName.apps.ApiConfig',
# go to the views files in the NewFolderName folder
# create a new file name as urls.py file in the NewFolderName folder
# update the urls.py file in the projectName folder as per your need

# make the routes of array of object in the views folder for the routing in the backend

# setup the django database
# update the models.py file in the NewFolderName Folder

# python manage.py migrate                 {for rebuilt the database in the django database server}{terminal}
# python manage.py makemigrations                                                                  {terminal}
# python manage.py migrate                 {for rebuilt the database in the django database server} {terminal}

# process to get the acess of the database of django
# python manage.py createsuperuser
# python manage.py runserver


# setup a new models in the database of django server
# go to te admin.py file of the NewFolderName folder
# update the code there             {admin.py  import the data from the model.py files}


# SETUP DJANGO REST FRAMEWORK
# pip install djangorestframework
# update the setting file in projectName folder like this in the INSTALLED_APPS part 'rest-framework'
# update vie.py file in the NewFolderName folder  {by importing the package from rest_framework}  {response and api_view}


# make and update views.py and urls.py file in NewFolderName folder as per your need
# you an import the models.py date to any where before using it               {it is use to connect the connecting with database of django}


# create a new file named as  serializers.py in the NewFolderName folder
# import the serializers in tje views.py file of NewFolderName folder

#
#
#
#
#

# AGAIN GO TTHE BACKEND
# python pip install django-cors-headers
# update the setting file in projectName folder like this in the INSTALLED_APPS part 'corsheaders',
# update the setting file in projectName folder like this in the MIDDLEWARE part 'corsheaders.middleware.CorsMiddleware',
# update the setting file in projectName folder like this in the BOTTOM OF THE FILE part 'CORS_ALLOW_ALL_ORIGINS = True',

#
#
#
#
#
# $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$            FRONTEND SERVER            $$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

# Go to the destop and open the ierminal
# create a react app
# make the folder of code
# in the fetch part of react app give the url of the backend directly
# setup a proxy url in the react package.json file                 {  "proxy": "http://127.0.0.1:8000",}

#
#
#
#
#
# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@         INTEGRATING BOTH            @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

# drag and drop the complete frontend end folder to the backend django server
# cd frontend folder
# npm run build

#
# update the setting file in projectName folder like this in the TEMPLATES part               { 'DIRS': [BASE_DIR / 'frontend/build'],}
# update the setting file in projectName folder like this in the STATICFILES_DIRS part    { STATICFILES_DIRS = [BASE_DIR / 'frontend/build/static'] }
# update the urls.py file in projectName folder like this     {from django.views.generic import TemplateView}  { path('', TemplateView.as_view(template_name='index.html')),}

#
#
#
#
#
# @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@         MAKING API RESTFULL            @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
#
# create a utils.py files in the NewFolderName Folder
# update the code in views.py file , utils.py file
