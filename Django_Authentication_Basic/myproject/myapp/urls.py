from django.urls import path
from . import views

urlpatterns = [
    path('', views.index, name='index'),
    path('counter', views.counter, name='counter'),
    path('register', views.register, name='register'),
    path('loginUser', views.loginUser, name='loginUser'),
    path('logoutUser', views.logoutUser, name='logoutUser'),
    path('post/<str:pk>', views.postuser, name='postuser')
]
