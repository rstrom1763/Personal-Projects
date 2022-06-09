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
                author = comment['author']
                if author in output['authors']:
                    author_comments = output['authors'][author]['comments']
                    author_comments.append(comment['body'])
                    output['authors'][author] = {'comment_count':output['authors'][author]['comment_count'] + 1,
                                                'comments':author_comments}
                else:
                    output['authors'][author] = {'comment_count':1,
                                                'comments':[comment['body']]}
                # Process subreddit
                if comment['subreddit'] in output['subreddits']:
                    output['subreddits'][comment['subreddit']] = output['subreddits'][comment['subreddit']] + 1
                else:
                    output['subreddits'][comment['subreddit']] = 1

        save_file = open(os.getcwd() + '/' + 'save.json','w')
        json.dump(output,save_file)
        save_file.close()
parse_reddit_data()
