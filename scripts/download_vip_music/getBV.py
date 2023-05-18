import requests
from bs4 import BeautifulSoup
import re
import os
import time
headers = {
        "user-agent":
        "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
        "Referer": "https://www.bilibili.com/",
    }
payload = {
    "keyword" : "Main Actor",
    "from_source" : "webtop_search",
    "spm_id_from" : "333.1007",
    "search_source" : "2"
}
url = "https://search.bilibili.com/all"

search = """红尘客栈
明明就
稻香
本草纲目
"""

if __name__ == "__main__":
    keyswords = search.split('\n')
    video_url = []
    with open("bv.txt", "w") as file:
        for keyword in keyswords:
            time.sleep(3)
            payload["keyword"] = keyword
            print("搜索{}".format(keyword))
            res = requests.get(url, params=payload, headers=headers)
            print(res.text)
            soup = BeautifulSoup(res.text, 'html.parser')

            for script in soup.find_all("script"):
                bv = re.search(r'BV[0-9a-zA-Z]*', script.text, 0)
                if bv:
                    file.write(bv.group())
                    file.write("\n")
    print(video_url)


#         result = os.system("you-get https://www.bilibili.com/video/{}".format(video_url))
#         print(result)
