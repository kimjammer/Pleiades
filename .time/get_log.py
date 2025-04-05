import subprocess
import re

def process_git_log():
    cmd = ['git', 'log', '--author=' + '"' + subprocess.getoutput('git config user.name') + '"',
           '--pretty=format:"%ad %h %s"', '--date=short']
    print(' '.join(cmd))
    output = subprocess.getoutput(' '.join(cmd))
    for line in output.split("\n"):
        extracted = re.findall(r'\(([^)]+)\)', line)
        cleaned = re.sub(r'\s*\([^)]+\)', '', line)
        try:
            time = extracted[0]
        except IndexError:
            time = "\t"
        [date, git_hash, *msg] = cleaned.split(" ")
        print(date, git_hash, time, " ".join(msg))

if __name__ == "__main__":
    process_git_log()