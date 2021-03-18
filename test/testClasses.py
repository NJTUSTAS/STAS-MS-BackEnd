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
        raw_data = {"Name": random_user_name(), "Email": random_email(), "Password": random_password()}
        email_data = {**raw_data, "Email": ""}
        dup_email_data = {**raw_data, "Email": "1"}
        wrn_emp_name_data = {**raw_data, "Name": "", "Email": random_email()}

        # 输入数据顺序
        given_input = ({"data": _} for _ in
                       (raw_data, dup_email_data, email_data, wrn_emp_name_data))
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
        legal_data = {"Name": random_user_name(), "Email": random_email(), "Password": random_password()}
        response = RegisterTest().pre_call({"data": legal_data})
        assert response == {"code": 200, "msg": "register successful."} or print(f"{response=}")

        raw_data = legal_data
        emp_email_data = {**raw_data, "Email": ""}
        err_password_data = {**raw_data, "Password": ""}

        # 输入数据顺序
        given_input = ({"data": _} for _ in (raw_data, emp_email_data, err_password_data))
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
        user = {"Name": random_user_name(), "Email": random_email(), "Password": random_password()}
        legal_data = {"data": user}
        register_response = RegisterTest().pre_call(legal_data)
        assert register_response["msg"] == 'register successful.' or print(
            "注册失败\n"
            f"{legal_data=}\n"
            f"\n{register_response=}")
        response: dict = LoginTest().pre_call(legal_data)

        assert response["msg"] == "登录成功" or print(
            f"{legal_data=}"
            f"{register_response=}\n"
            f"{response=}")

        token = response["data"]["token"]
        # print(token)

        legal_input = {"headers": {"Authorization": token}}
        dataset = (legal_input, {"headers": {}},)

        del (user_without_password := user)["Password"]
        out = ({"code": 200, "data": {"user": user_without_password}},
               {"code": 401, "msg": "权限不足"},)
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
        expected_return = ({"code": 200, "msg": "报名成功", },)
        return zip(dataset, expected_return)
