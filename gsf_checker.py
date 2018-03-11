import requests
import json
import argparse
import sys

URL = 'https://safebrowsing.googleapis.com/v4/threatMatches:find?key='
FILENAME = '/tmp/check.txt'

def get_url(file):
    """
    return safebrowsing url with key;
    accept path to key file
    """
    try:
        with open(file, 'r') as f:
            return URL + f.readlines()[0].rstrip()
    except IOError as err:
        print("Failed with error: %s" % err)
        sys.exit(1)

def get_json(url, data):
    json_data = json.dumps(data)
    headers = {
        'Content-Type': 'application/json'
    }
    try:
        r = requests.post(url, json_data, headers=headers, timeout=5)
    except requests.exceptions.RequestException as err:
        print("Request error: %s" % err)
        sys.exit(1)
    except requests.exceptions.Timeout as err:
        print("Request timeout 5 seconds with error: %s" % err)
        sys.exit(1)
    return r.json()

data = {
        "client": {
                    "clientId":      "myproject",
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

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Google Safe browser checker')
    parser.add_argument('--key', '-k', required=True, help='specify the path to  key file')
    path_to_key = parser.parse_args().key

    output = get_json(get_url(path_to_key), data)
    #print(output)

    if len(output) > 0:
        list_data =  output.get('matches')
        with open(FILENAME, 'a') as f:
            for elem in list_data:
                f.writelines(elem.get('threat').get('url') + '\n')
    else:
        open(FILENAME, 'w').close()
        #print "OK"
