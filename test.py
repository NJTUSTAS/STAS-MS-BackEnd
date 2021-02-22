import random
from typing import Dict

import requests

domin = "http://localhost:8080"


def random_hex_str(n: int):
    return "".join(random.sample("0123456789abcdef", n))


def random_user_name():
    return random_hex_str(8)


def random_email():
    address = random_hex_str(random.randint(4, 10))
    domain = random_hex_str(random.randint(1, 4)) + \
             random.choice([".edu", ".com", ".cn"])
    return f"{address}@{domain}"

def random_password():
    return random_hex_str(8)

if __name__ == "__main__":
    url = domin + "/api/auth/register"
    dict().update()
    raw_data: Dict[str, str] = \
        {"name": random_user_name(),
         "email": random_email(),
         "password": random_password()}
    # err_name_data = {"name": ""}|=raw_data
    # err_psd_data = {**raw_data, "password": ""}
    err_email_data = {**raw_data, "email": ""}
    err_dupemail_data = {**raw_data, "email": "1"}
    err_name_data = {**raw_data, "name": ""}
    # dataset = (raw_data, err_name_data, err_psd_data, err_email_data)
    dataset = (err_dupemail_data, err_email_data, err_name_data, raw_data)
    for data in dataset:
        ret = requests.post(url, data)
        print(data, ret.text)
