import requests
import json

URL = 'https://safebrowsing.googleapis.com/v4/threatMatches:find?key='
KEY = 'trololo'
FILENAME = '/tmp/check.txt'

def get_json(url, data):
    json_data = json.dumps(data)
    headers = {
        'Content-Type': 'application/json'
    }
    r = requests.post(url, json_data, headers=headers)
    return json.loads(r.text.encode('UTF-8'))

valid_url =  URL+KEY
data = {
        "client": {
                    "clientId":      "my_project",
                    "clientVersion": "1.5.2"
                },
        "threatInfo": {
                    "threatTypes":      ["MALWARE", "SOCIAL_ENGINEERING", "POTENTIALLY_HARMFUL_APPLICATION", "UNWANTED_SOFTWARE"],
                    "platformTypes":    ["ANY_PLATFORM"],
                    "threatEntryTypes": ["URL"],
                    "threatEntries": [
                                {"url": "http://malware.testing.google.test/testing/malware/"},
                                {"url": "http://ianfette.org"},
                            ]
                }
}

output = get_json(valid_url, data)

#print unicode(output)
if len(output) > 0:
    list_data =  output.get('matches'.encode('UTF-8')) #.get('threatType'.encode('UTF-8'))
    with open(FILENAME, 'a') as f:
        for elem in list_data:
            f.write(elem.get('threat'.encode('UTF-8')).get('url'.encode('UTF-8')) + '\n')
    f.close()
else:
    open(FILENAME, 'w').close()
    #print "OK"
