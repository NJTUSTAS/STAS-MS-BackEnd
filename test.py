from typing import Dict

import requests

domin = "http://localhost:8080"

if __name__ == "__main__":
    url = domin + "/api/auth/register"
    dict().update()
    raw_data: Dict[str, str] = {"name": "hello", "email": "123451678903@foo.com", "password": "122345"}
    # err_name_data = {"name": ""}|=raw_data
    # err_psd_data = {**raw_data, "password": ""}
    err_email_data = {**raw_data, "email": ""}
    err_dupemail_data = {**raw_data, "email": "1"}
    err_name_data = {**raw_data, "name": ""}
    # dataset = (raw_data, err_name_data, err_psd_data, err_email_data)
    dataset = (raw_data, err_dupemail_data, err_email_data,err_name_data)
    for data in dataset:
        ret = requests.post(url, data)
        print(data,ret.text)
