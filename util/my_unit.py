import json
import random

# 生成随机字符串
import requests


def random_hex_str(n: int):
    return "".join(random.sample("0123456789abcdef", n))


# 8位用户名
def random_user_name():
    return random_hex_str(8)


# 随机邮箱，测试用
def random_email():
    address = random_hex_str(random.randint(4, 10))
    domain = random_hex_str(random.randint(1, 4)) + \
             random.choice([".edu", ".com", ".cn"])
    return f"{address}@{domain}"


# 随机密码
def random_password():
    return random_hex_str(8)


class Unittest:
    def __init__(self):
        self.reshape = lambda x: x
        self.expected_io = ""
        self.url = ""

    #    指定数据测试
    def single_test(self, data):
        response = requests.post(self.url, data)
        # print(response.text)

    # 批量测试
    def unittest(self):
        passed = True
        for data, expect in self.expected_io:
            response = requests.post(self.url, data)
            try:
                ret = json.loads(response.text)
            except json.decoder.JSONDecodeError:
                ret = {}
            if self.reshape(ret) != expect:
                print(False, f"\ninput:{data}\nexpect:{expect}\nresponse:{response.text}")
                passed = False
        return passed

        # print(ret == expect)

    def reshape(self, _):
        pass
