def similar(a, b):
    from difflib import SequenceMatcher
    return SequenceMatcher(None, a, b).ratio()


def compare_db(db1, db2):

    import pandas as pd
    import csv

    db1 = pd.read_csv(db1)
    db2 = pd.read_csv(db2)

    dict1 = {}
    dict2 = {}

    for movie in db1:
        dict1[movie["Title"]] = movie["Part Size as Bytes"]
    for movie in db2:
        dict1[movie["Title"]] = movie["Part Size as Bytes"]
    
    print(dict1.keys)



compare_db("C:\Strom\KentLibrary.csv", "C:\Strom\RyanLibrary.csv")