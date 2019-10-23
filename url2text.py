from bs4 import BeautifulSoup
from bs4.element import Comment
import re
import requests
import sys


def tag_visible(element):
    if element.parent.name in ['style', 'script', 'head', 'title', 'meta', '[document]']:
        return False
    if isinstance(element, Comment):
        return False
    return True


def text_from_html(body):
    soup = BeautifulSoup(body, 'html.parser')
    texts = soup.findAll(text=True)
    visible_texts = filter(tag_visible, texts)
    return u" ".join(t.strip() for t in visible_texts)


def url2text(url):
    # url = 'http://' + url
    try:
        page = requests.get(url)  # to extract page from website
        html_code = page.content  # to extract html code from page
        readable_text = text_from_html(html_code)
    except Exception as e:
        # print(e)
        return str(e)
    return readable_text


# good_urls = ['2appstudio.com', 'Blurb.com']
url = 'https://epublisher.world/en/'
print(url2text(url))

for line in sys.stdin:
    label, url = line.split(",", 1)
    # print("label", label)
    # print("url", url)
    # print("text", url2text(url))
    text = url2text(url)
    text = re.sub('\s+', ' ', text).strip()
    print(label, " ", text)
