#Ryan Strom

#exports as a CSV what movies are in the first DB and not in the second
#inputs are the file paths to each of the CSV files.
def plex_compare(db1, db2, file, exclude_file=""):

    import utils

    space_needed = 0
    if exclude_file != "":
        exclude_file = (open(exclude_file,'r',encoding='utf-8')).readlines()
    for line in exclude_file:
        exclude_file[exclude_file.index(line)] = line.replace("\n","")

    #Import both CSV's into list of dictionaries
    db1 = utils.import_csv(db1)
    db2 = utils.import_csv(db2)

    dict1 = {}
    dict2 = {}
    list = []

    for dict in db1:
        dict1[dict["Title"]] = dict
    for dict in db2:
        dict2[dict["Title"]] = dict
    
    for movie in dict1:
        if movie not in dict2 and movie not in exclude_file:
            list.append(dict1[movie])

    #Iterated through the dictionaries to print required file space
    for movie in list:
        if movie["Part Size as Bytes"].isdecimal():
            space_needed += int(movie["Part Size as Bytes"])
    print(utils.human_readable(space_needed))

    utils.export_csv(list, file)


#Examples:
#plex_compare("C:/Strom/KentLibrary.csv", "C:/Strom/ryanlibrary.csv", "C:/strom/ryan_no_have.csv", exclude_file="C:/strom/test.txt")
#plex_compare("C:/Strom/ryanlibrary.csv", "C:/strom/kentlibrary.csv", "C:/strom/kent_no_have.csv")
