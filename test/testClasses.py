import requests

from my_unit import TestUnit
from random_data import *


class RegisterTest(TestUnit):
    def __init__(self):
        super().__init__()
        self.name = "register"
        self.url = TestUnit.domain + "/api/auth/register"
        self.method = requests.post

    @staticmethod
    def expect_io():
        raw_data = {"name": random_user_name(), "email": random_email(), "password": random_password()}
        email_data = {**raw_data, "email": ""}
        dup_email_data = {**raw_data, "email": "1"}
        wrn_emp_name_data = {**raw_data, "name": "", "email": random_email()}

        # 输入数据顺序
        given_input = (raw_data, dup_email_data, email_data, wrn_emp_name_data)
        # 期望输出数据
        expected_return = (
            {"code": 200, "msg": "register successful."},
            {"code": 422, "msg": "exist email address!"},
            {"code": 422, "msg": "illegal email address!"},
            {"code": 200, "msg": "no user name,auto generated.", "name": None})
        return zip(given_input, expected_return)

    @staticmethod
    def format(data: dict) -> dict:
        if "name" in data:
            data["name"] = None
        return data


class LoginTest(TestUnit):
    def __init__(self):
        super().__init__()
        self.name = "login"
        self.url = TestUnit.domain + "/api/auth/login"
        self.method = requests.post

    def expect_io(self) -> zip:
        # 注册账号
        legal_data = {"name": random_user_name(), "email": random_email(), "password": random_password()}
        response = RegisterTest().pre_call(legal_data)
        assert response == {"code": 200, "msg": "register successful."}

        raw_data = legal_data
        emp_email_data = {**raw_data, "email": ""}
        err_password_data = {**raw_data, "password": ""}

        # 输入数据顺序
        given_input = (raw_data, emp_email_data, err_password_data)
        # 期望输出数据
        expected_return = ({"code": 200, "data": {"token": None}, "msg": "登录成功"},
                           {"code": 422, "msg": "用户不存在"},
                           {"code": 400, "msg": "用户名与密码不匹配"})

        return zip(given_input, expected_return)

    @staticmethod
    def format(data: dict) -> dict:
        if "data" in data:
            # print("token:", data["data"]["token"])
            data["data"]["token"] = None
        return data


class InfoTest(TestUnit):
    def __init__(self):
        super().__init__()
        self.name = "Info"
        self.url = TestUnit.domain + "/api/auth/info"
        self.method = requests.get

    @staticmethod
    def expect_io():
        # 注册并登录
        legal_data = {"name": random_user_name(), "email": random_email(), "password": random_password()}
        RegisterTest().pre_call(legal_data)
        response: dict = LoginTest().pre_call(legal_data)
        assert response["msg"] == "登录成功"

        dataset = ({}, legal_data)

        del (legal_return := dict(legal_data))["password"]
        out = ({"code": 401, "msg": "权限不足"},
               {"code": 200, "data": legal_return})
        return zip(dataset, out)

    @staticmethod
    def format(response: dict) -> dict:
        if "data" in response and "user" in response["data"]:
            # print("token:", data["data"]["token"])
            response["data"]["user"].pop("ID")
            response["data"]["user"].pop("CreatedAt")
            response["data"]["user"].pop("UpdatedAt")
            response["data"]["user"].pop("DeletedAt")
            response["data"]["user"].pop("Password")
        return response


class EnrollReceiveTest(TestUnit):
    def __init__(self):
        super().__init__()
        self.name = "enroll.receive"
        self.url = TestUnit.domain + "/api/enroll/receive"
        self.method = requests.post

    @staticmethod
    def expect_io():
        legal_data = {
            "name": random_user_name(),
            "student_id": random_stu_id(),
            "major": "major",
            "phone": "phone",
            "grade": "grade",
            "gender": "男",
            "firstChoice": "技术开发部",
            "secondChoice": "",
            "introduction": "introduction",
            "hope": "hope",
            "hobbies": "hobbies",
        }
        dataset = (legal_data,)
        expected_return = ({"code": 200, "msg": "报名成功", },)
        return zip(dataset, expected_return)
