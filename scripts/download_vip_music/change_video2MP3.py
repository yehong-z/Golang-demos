import os


def change_video2MP3(filepath, target_path):
    print("将 {}目录下的 mp4 文件转为 mp3文件并移动到 {}目录下".format(filepath, target_path))
    def dfs(relative_path):
        pwd = filepath + relative_path
        target = target_path + relative_path
        filename = os.listdir(pwd)
        for name in filename:
            file_or_dir = pwd + name
            if name.endswith(".mp4"):
                if not os.path.exists(target):
                    os.mkdir(target)
                target_filename = target + name[:-1] + '3'
                command = "ffmpeg -i \"{}\" -vn \"{}\" -y".format(file_or_dir, target_filename)
                print("正在处理 {}".format(relative_path + name))
                result = os.system(command)
                if result == 0:
                    print("成功")
                else:
                    print("失败")
            else :
                if os.path.isdir(file_or_dir):
                    dfs(relative_path + name + "/")
    dfs("/")
    print("全部文件处理完成")

if __name__ == "__main__":
    curr_dir = os.getcwd()
    target_file = "C:/Users/yehong/Music"
    change_video2MP3(curr_dir, target_file)