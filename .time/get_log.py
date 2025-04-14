import subprocess
import re
from datetime import datetime, timedelta

def last_monday():
    today = datetime.today()
    offset = today.weekday() or 7  # Monday is 0, Sunday is 6
    return (today - timedelta(days=offset)).strftime('%Y-%m-%d')

def process_git_log():
    author = subprocess.getoutput('git config user.name')
    since_date = last_monday()
    cmd = [
        'git', 'log',
        f'--author="{author}"',
        '--pretty=format:"%ad %h %s"',
        '--date=short',
        f'--since={since_date}'
    ]
    print(' '.join(cmd))

    total_time = 0
    output = subprocess.getoutput(' '.join(cmd))

    for line in reversed(output.splitlines()):
        extracted = re.findall(r'\(([^)]+)\)', line)
        cleaned = re.sub(r'\s*\([^)]+\)', '', line)
        time = extracted[0] if extracted else "\t"

        try:
            total_time += float(re.sub(r'[^\d.]', '', time))
        except ValueError:
            pass

        date, git_hash, *msg = cleaned.split()
        print(date, git_hash, time, ' '.join(msg))

    print('\ntotal time:', round(total_time, 2))

if __name__ == "__main__":
    process_git_log()
