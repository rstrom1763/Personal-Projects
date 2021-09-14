#Ryan Strom

#exports as a CSV what movies are in the first DB and not in the second
#inputs are the file paths to each of the CSV files.
def plex_compare(db1, db2, file):

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

    utils.export_csv(list, file)


#Examples:
#plex_compare("C:/Strom/KentLibrary.csv", "C:/Strom/ryanlibrary.csv", "C:/strom/ryan_no_have.csv")
#plex_compare("C:/Strom/ryanlibrary.csv", "C:/strom/kentlibrary.csv", "C:/strom/kent_no_have.csv")
