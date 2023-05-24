from django.db import models

# Create your models here.


# class Features:
#     id: int
#     name: str
#     details: str
#     is_true: bool


class Features(models.Model):
    name = models.CharField(max_length=100)
    details = models.CharField(max_length=100)
