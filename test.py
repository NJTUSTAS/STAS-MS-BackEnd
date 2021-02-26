"""
本文件用于执行测试功能。
形如“RegisterTest”的每一个类表示一组针对接口的测试数据。
填写时只要实现__init__()皆可。

其中需要实现expected_io和reshape（可选）

self.expected_io包含各种情况下每一组输入对应的相对输出,应当为zip。
对于输入对应的输出中输出不确定的内容返回值设置为空字符串，并且需要实现下面的函数。
self.reshape是一个函数，在包含不确定返回值的情况下将不确定的内容转换为空字符串。
"""

from util.my_unit import *

# 测试域名
domain = "http://localhost:8080"


class RegisterTest(Unittest):
    # 输入数据编辑
    def __init__(self):
        super().__init__()
        self.url = domain + "/api/auth/register"

        raw_data = \
            {"name": random_user_name(),
             "email": random_email(),
             "password": random_password()}
        email_data = {**raw_data, "email": ""}
        dup_email_data = {**raw_data, "email": "1"}
        wrn_emp_name_data = {**raw_data, "name": "", "email": random_email()}

        # 输入数据顺序
        dataset = (raw_data, dup_email_data, email_data, wrn_emp_name_data)
        # 期望输出数据
        expected_return = (
            {"code": 200, "msg": "register successful."},
            {"code": 422, "msg": "exist email address!"},
            {"code": 422, "msg": "illegal email address!"},
            {"code": 200, "msg": "no user name,auto generated.", "name": ""},
        )

        def remake(data: dict) -> dict:
            if "name" in data:
                data["name"] = ""
            return data

        self.expected_io: zip = zip(dataset, expected_return)
        self.reshape = remake


class LoginTest(Unittest):
    def __init__(self):
        super().__init__()
        self.url = domain + "/api/auth/login"
        legal_data = {"name": random_user_name(), "email": random_email(), "password": random_password()}
        RegisterTest().single_test(legal_data)

        raw_data = legal_data
        emp_email_data = {**raw_data, "email": ""}
        err_password_data = {**raw_data, "password": ""}

        # 输入数据顺序
        dataset = (raw_data, emp_email_data, err_password_data)
        # 期望输出数据
        expected_return = (
            {"code": 200, "data": {"token": ""}, "msg": "登录成功"},
            {"code": 422, "msg": "用户不存在"},
            {"code": 400, "msg": "用户名与密码不匹配"})

        def remake(data: dict) -> dict:
            if "data" in data:
                data["data"]["token"] = ""
            return data

        self.expected_io = zip(dataset, expected_return)
        self.reshape = remake


if __name__ == "__main__":
    repeat = 40
    print()
    print("register", all(RegisterTest().unittest() for i in range(repeat)))
    print("login", all(LoginTest().unittest() for i in range(repeat)))
