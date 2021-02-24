from util.my_unit import *

# 测试域名
domain = "http://localhost:8080"


class Unittest:
    def __init__(self):
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

        self.reshape = remake

        self.expected_io = zip(dataset, expected_return)


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

        self.reshape = remake
        self.expected_io = zip(dataset, expected_return)


if __name__ == "__main__":
    print()
    print("register", all(RegisterTest().unittest() for i in range(50)))
    print("login", all(LoginTest().unittest() for i in range(50)))
