import json
from pprint import pprint
from ast import literal_eval as make_tuple





MorseTemplate = {
    "version" : "international",
    "characters": []
}

CharSet = []


with open('scratch.txt','r') as f:
    for item in f:

        item = make_tuple(item)
        character = {
            "name" : item[0],
            "code" : item[1],
            "value" : "n"
        }

        MorseTemplate['characters'].append(json.dumps(character))

with open('intMorse.json','w') as f:
    f.write(json.dumps(MorseTemplate))    