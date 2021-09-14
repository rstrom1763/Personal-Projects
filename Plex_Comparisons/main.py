from Movie import *
import os
import json

def init_data():
    #Load the JSON strings into memory as objects of corresponding type
    for file in os.listdir(os.getcwd()):
        print(file)

def menu():
    print("1. Create new entry\n2. Delete entry\n3. Change data value\n")
    choice = input("Please choose an option: ")


while True:
    menu()