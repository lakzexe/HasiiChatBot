from pymongo import MongoClient

import config

hasiidb = MongoClient(config.MONGO_URL)
hasii = hasiidb["HasiiDb"]["Hasii"]


from .chats import *
from .users import *
