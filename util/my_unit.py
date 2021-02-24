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

