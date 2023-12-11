import pandas as pd
import json


with open("results.json") as f:
    json_data = json.load(f)


df = pd.read_json(json_data)

print(df)