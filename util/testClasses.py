from util.my_unit import Unittest, random_user_name, random_email, random_password

testClasses: list[Unittest] = []

domain = "http://localhost:8080"


# 这是一个装饰器。他抓取一个类，对这个类做一些事情，然后返回这个类给调用装饰器的东西
def add_to_test_list(testcases):
    testClasses.append(testcases)
    return testcases


# 这个装饰器会把下面定义的类加入testClasses
@add_to_test_list
class RegisterTest(Unittest):
    # 输入数据编辑
    def __init__(self):
        super().__init__()
        self.name = "Register"
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


@add_to_test_list
class LoginTest(Unittest):
    def __init__(self):
        super().__init__()
        self.name = "Login"
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
                # print("token:", data["data"]["token"])
                data["data"]["token"] = ""
            return data

        self.expected_io = zip(dataset, expected_return)
        self.reshape = remake


@add_to_test_list
class EnrollReceiveTest(Unittest):
    def __init__(self):
        super().__init__()
        self.name = "enroll.receive"
        self.url = domain + "/enroll/receive"
        legal_data = {
            "name": random_user_name(),
            "student_id": "202021100000",
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
        self.expected_io = zip(dataset, expected_return)
