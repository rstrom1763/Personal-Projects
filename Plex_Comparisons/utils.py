def import_csv(csv):

    import re

    csv = (open(csv, 'r',encoding="ISO-8859-1")).readlines()
    headers = csv[0].split(",")
    csv.remove(csv[0])
    list = []

    for header in headers:
        headers[headers.index(header)] = header.replace('"', "")

    for line in csv:

        dict = {}
        line = re.split(r',(?=")', line)

        for header in headers:
            dict[header] = line[headers.index(header)].replace('"', "")

        list.append(dict)
    return(list)
