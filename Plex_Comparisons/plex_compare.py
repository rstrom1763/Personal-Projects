import os
os.system("cls")

def compare_db(db1, db2):

    import utils

    db1 = utils.import_csv(db1)
    db2 = utils.import_csv(db2)

    dict1 = {}
    dict2 = {}
    list = []

    for dict in db1:
        dict1[dict["Sort title"]] = dict
    for dict in db2:
        dict2[dict["Sort title"]] = dict
    
    for movie in dict1:
        if movie not in dict2:
            list.append(dict1[movie])

    file = open("C:/strom/test.txt","w")
    file.write(str(list))

compare_db("C:\Strom\KentLibrary.csv", "C:/Strom/ryanlibrary.csv")
