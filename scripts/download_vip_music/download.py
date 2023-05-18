import os
import time
if __name__ == "__main__":
    content = []
    with open("bv.txt", "r") as file:
        content = file.read().split('\n')
    downloadfail = []
    for bv in content:
        command = "you-get https://www.bilibili.com/video/{}".format(bv)
        print(command)
        res = os.system(command)
        if res != 0:
            downloadfail.append(bv)
    with open("fail.txt", "w") as file:
        for bv in downloadfail:
            file.write(bv)
            file.write("\n")
