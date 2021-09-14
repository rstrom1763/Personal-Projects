class Movie():
    def __init__(self,title,length,rating,genre):
        self.title = title
        self.length = length
        self.rating = rating
        self.genre = genre

    def to_json(self):
        import json
        return(json.dumps(self.__dict__))
    
