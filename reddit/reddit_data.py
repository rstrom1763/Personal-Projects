def create_listing(comment, dict):
    import re

    attribute = dict['type']

    if comment[attribute] in dict:
        dict[comment[attribute]] = dict[comment[attribute]] + 1
    else:
        dict[comment[attribute]] = 1
    return dict


def create_subreddit_list(file, save_path):
    import json

    types = ['author', 'subreddit']
    authors = {'type': 'author'}
    subreddits = {'type': 'subreddit'}
    dicts = [authors, subreddits]

    for comment in open(file, 'r').readlines():
        comment = json.loads(comment)

        index = 0
        for dict in dicts:
            dicts[index] = create_listing(comment, dict)
            index += 1

    index = 0
    for attribute in types:
        save_file = save_path + '/' + attribute + '.json'
        save_file = open(save_file, 'w')
        json.dump(dicts, save_file)
        save_file.close()
        index += 1


create_subreddit_list(
    "C:/Strom/ml_testing/reddit/data/RC_2011-06.json", 'C:/strom/ml_testing/reddit')
