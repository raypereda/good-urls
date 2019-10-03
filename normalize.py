import html2text
import re
import random
import subprocess
import sys
import urllib.request


def normalize(s):
    s = s.lower()
    s = "__label__" + s
    s = s.replace("'", " ' ")
    s = s.replace('"', "")
    s = s.replace(".", " . ")
    s = s.replace("<br \/", " ")
    s = s.replace(",", " , ")
    s = s.replace("(", " () ")
    s = s.replace(")", " ) ")
    s = s.replace("!", " ! ")
    s = s.replace("?", " ? ")
    s = s.replace(";", " ; ")
    s = s.replace(":", " : ")
    s = re.sub(r"\s+", " ", s)
    s = re.sub(r"\d", "@", s)
    return s


def text(url):
    with urllib.request.urlopen(url) as response:
        html = response.read()
    h = html2text.HTML2Text()
    h.ignore_links = True
    return h.handle(html)


def main():
    if len(sys.argv) != 2:
        print("Expecting an input file name.")
        sys.exit(1)

    lines = []
    input_filename = sys.argv[1]
    with open(input_filename) as inputfile:
        for line in inputfile:
            label, url = line.split(" ", 1)
            print("helloooooo")
            print("label:", label, "url:", url)
            l = label + " " + text(url)
            lines.append(l)

    # random.shuffle(lines)

    for l in lines:
        print(l)


req = urllib.request.Request('http://www.voidspace.org.uk')
with urllib.request.urlopen(req) as response:
    the_page = response.read()
print(the_page)

html = the_page

h = html2text.HTML2Text()
h.ignore_links = True
print(h.handle(html))


# url = "https://www.google.com"
# with urllib.request.urlopen(url) as response:
#     html = response.read()
# print(html)

# h = html2text.HTML2Text()
# h.ignore_links = True
# print(h.handle(html))

exit(0)

if __name__ == "__main__":
    main()
