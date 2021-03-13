# encoding =utf-8
import random


def random_hex_str(n: int):
    return "".join(random.sample("0123456789abcdef", n))


def random_user_name():
    return random_hex_str(8)


def random_stu_id():
    return f"{random.randint(2016, 2021)}{random.randint(0, 9999999)}"


def random_email():
    address = random_hex_str(random.randint(4, 10))
    domain = random_hex_str(random.randint(1, 4)) + random.choice([".edu", ".com", ".cn"])
    return f"{address}@{domain}"


def random_password():
    return random_hex_str(8)
