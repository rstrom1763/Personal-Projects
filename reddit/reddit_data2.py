def parse_reddit_data():
    import json
    import os

    output = {'authors': {}, 'subreddits': {}, 'words': {}}
    data_dir = os.getcwd() + '/data'
    data_files = os.listdir(data_dir)

    for file in data_files:
        test = data_dir + "/" + file
        with open(test) as comments:
            while True:
                comment = comments.readline()
                if not comment:
                    break
                comment = json.loads(comment)

                # Process author
                if comment['author'] in output['authors']:
                    output['authors'][comment['author']] = output['authors'][comment['author']] + 1
                else:
                    output['authors'][comment['author']] = 1

                # Process subreddit
                if comment['subreddit'] in output['subreddits']:
                    output['subreddits'][comment['subreddit']] = output['subreddits'][comment['subreddit']] + 1
                else:
                    output['subreddits'][comment['subreddit']] = 1

        save_file = open(os.getcwd() + '/' + 'save.json','w')
        json.dump(output,save_file)
        save_file.close()
parse_reddit_data()
