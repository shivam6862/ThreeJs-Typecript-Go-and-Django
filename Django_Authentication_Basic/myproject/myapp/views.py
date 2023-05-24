from django.shortcuts import render, redirect
from django.http import HttpResponse
from .models import Features

from django.contrib.auth.models import User, auth
from django.contrib.auth import authenticate, login, logout
from django.contrib import messages

# Create your views here.


# def index(request):
#     return HttpResponse('<h1>Hi, Shivam kumar, welcome</h1>')
#     name = 'Shivam Kumar'
#     context = {
#         'name': "Shivam Kumar",
#         'enrollement': 21117119
#     }
#     return render(request, 'index.html', context)


# def index(request):
#     return render(request, 'index.html')


# def index(request):
#     features1 = Features()
#     features1.id = 0
#     features1.name = "Book"
#     features1.details = "I love is life"
#     # features1.is_true = True
#     features = [features1, features1, features1, features1, features1]
#     return render(request, 'index.html', {'features': features})


def index(request):
    features = Features.objects.all()
    return render(request, 'index.html', {'features': features})


def counter(request):
    # text = request.GET['text']                {NOT GET METHOD DO NOT REQUIRED CSRF_TOKEN}
    text = request.POST['text']
    amount_of_word = len(text.split())
    context = {
        "amount_of_word": amount_of_word,
        "text": text
    }
    return render(request, 'counter.html', context)


def register(request):
    if request.method == 'POST':
        username = request.POST['username']
        email = request.POST['email']
        password = request.POST['password']
        password2 = request.POST['password2']

        if password == '' or email == '' or password2 == '' or username == '':
            messages.info(request, "All the value must be filled")
            return redirect('register')
        elif password == password2:
            if User.objects.filter(email=email).exists():
                messages.info(request, 'Email Already Used')
                return redirect('register')
            elif User.objects.filter(username=username).exists():
                messages.info(request, 'Username Already Used')
                return redirect('register')
            else:
                user = User.objects.create_user(
                    username=username, email=email, password=password)
                user.save()
                return redirect('login')
        else:
            messages.info(request, "Password must be same")
            return redirect('register')
    else:
        return render(request, 'register.html')


def loginUser(request):
    if request.method == 'POST':
        username = request.POST['username']
        email = request.POST['email']
        password = request.POST['password']

        if password == '' or email == '' or username == '':
            messages.info(request, "All the value must be filled")
            return redirect('loginUser')

        user = authenticate(username=username, password=password)

        if user is not None:
            login(request, user)
            return redirect('/')
        else:
            messages.info(request, 'Credientials Invalids')
            return redirect('/loginUser')
    else:
        return render(request, 'loginUser.html')


def logoutUser(request):
    logout(request)
    return redirect('/')


def postuser(request, pk):
    return render(request, 'postuser.html', {"pk": pk})
