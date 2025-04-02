import json
import subprocess
from datetime import datetime

# For use with the vscode `S-Mitterlehner.time-tracker-vscode` extension

TIME_FILE = "../.vscode/times.json"

def get_git_commits(start, end):
    """Retrieve commit hashes between start and end timestamps."""
    cmd = [
        "git", "log", "--pretty=%H", 
        f"--since={start}", f"--until={end}"
    ]
    result = subprocess.run(cmd, capture_output=True, text=True)
    return result.stdout.strip().split("\n") if result.stdout else []

def process_time_entries(data):
    """Compute duration and add commit hashes."""
    for project in data["projects"].values():
        for entry in project["times"]:
            start, end = entry["from"], entry["till"]
            fmt = "%Y-%m-%d %H:%M:%S"
            
            duration = (datetime.strptime(end, fmt) - datetime.strptime(start, fmt)).total_seconds()
            entry["duration"] = round(duration / 3600, 2)
            entry["hashes"] = get_git_commits(start, end)
    
    return data

# Load JSON file
with open(TIME_FILE) as f:
    data = json.load(f)

# Process and update data
data = process_time_entries(data)

# Save the updated JSON
with open(TIME_FILE) as f:
    json.dump(data, f, indent=2)

print("Updated data saved to times.json")