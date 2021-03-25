import requests

from my_unit import TestUnit
from random_data import *


def get_token(data):
    register_response = get_registered(data)
    response: dict = LoginTest().pre_call(data)
    assert response["msg"] == "登录成功" or print(
        f"{data=}"
        f"{register_response=}\n"
        f"{response=}")
    token = response["data"]["token"]
    return token


def get_registered(data):
    register_response = RegisterTest().pre_call(data)
    assert register_response.get("msg") == 'register successful.' or print(
        "注册失败\n"
        f"{data=}\n"
        f"\n{register_response=}")
    return register_response


class RegisterTest(TestUnit):
    def __init__(self):
        super().__init__()
        self.name = "register"
        self.url = TestUnit.domain + "/api/auth/register"
        self.method = requests.post

    @staticmethod
    def expect_io():
        raw_data = {"Name": random_user_name(), "Email": random_email(), "Password": random_password()}
        email_data = {**raw_data, "Email": ""}
        dup_email_data = {**raw_data, "Email": "1"}
        wrn_emp_name_data = {**raw_data, "Name": "", "Email": random_email()}

        # 输入数据顺序
        given_input = ({"data": _} for _ in
                       (raw_data, dup_email_data, email_data, wrn_emp_name_data))
        # 期望输出数据
        expected_return = (
            {"code": 200, "data": None, "msg": "register successful."},
            {"code": 422, "data": None, "msg": "exist email address!"},
            {"code": 422, "data": None, "msg": "illegal email address!"},
            {"code": 200, "data": {"name": None}, "msg": "no user name,auto generated."})
        return zip(given_input, expected_return)

    @staticmethod
    def format(data: dict) -> dict:
        if data["data"] is not None:
            data["data"]["name"] = None
        return data


class LoginTest(TestUnit):
    def __init__(self):
        super().__init__()
        self.name = "login"
        self.url = TestUnit.domain + "/api/auth/login"
        self.method = requests.post

    def expect_io(self) -> zip:
        # 注册账号
        legal_data = {"Name": random_user_name(), "Email": random_email(), "Password": random_password()}

        response = get_registered({"data": legal_data})

        raw_data = legal_data
        emp_email_data = {**raw_data, "Email": ""}
        err_password_data = {**raw_data, "Password": ""}

        # 输入数据顺序
        given_input = ({"data": _} for _ in (raw_data, emp_email_data, err_password_data))
        # 期望输出数据
        expected_return = ({"code": 200, "data": {"token": None}, "msg": "登录成功"},
                           {"code": 422, "data": None, "msg": "用户不存在"},
                           {"code": 400, "data": None, "msg": "用户名与密码不匹配"})

        return zip(given_input, expected_return)

    @staticmethod
    def format(data: dict) -> dict:
        if data["data"] is not None:
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
        user = {"Name": random_user_name(), "Email": random_email(), "Password": random_password()}
        token = get_token({"data": user})

        # print(token)

        legal_input = {"headers": {"Authorization": token}}
        dataset = (legal_input, {"headers": {}},)

        del (user_without_password := user)["Password"]
        out = ({"code": 200, "data": {"user": user_without_password}, "msg": ""},
               {"code": 401, "data": None, "msg": "权限不足"},)
        return zip(dataset, out)


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
        dataset = ({"data": _} for _ in (legal_data,))
        expected_return = ({"code": 200, "data": None, "msg": "报名成功", },)
        return zip(dataset, expected_return)


class NoteSendTest(TestUnit):

    def __init__(self):
        super().__init__()
        self.name = "note.send"
        self.url = TestUnit.domain + "/api/note"
        self.method = requests.post

    @staticmethod
    def expect_io():
        legal_data = {"data": {
            "name": "test_01",
            "data": "test note_01"
        }},
        expected_return ={'code': 200, 'data': None, 'msg': 'submitted'},
        return zip(legal_data, expected_return)


if __name__ == "__main__":
    TestUnit.domain = "http://localhost:8080"
    NoteSendTest().test()
