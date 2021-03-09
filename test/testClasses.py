# from test.my_unit import *
from my_unit import *
import requests

testClasses: list[Unittest] = []

domain = "http://localhost:8080"
# 已经在远程服务器上部署
# domain = "http://202.119.245.31:80"


# 这是一个装饰器。他抓取一个类，对这个类做一些事情，然后返回这个类给调用装饰器的东西
def add_to_test_list(testcases:Unittest)->Unittest:
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

        raw_data = {
            "name": random_user_name(),
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

        self.set_expect_io(legal_data)

    def set_expect_io(self,legal_data):
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
        self.expected_io = zip(dataset, expected_return)

    def remake(data: dict) -> dict:
            if "data" in data:
                print("token:", data["data"]["token"])
                data["data"]["token"] = ""
            return data

    # def single_test():
    #     pass
        

# @add_to_test_list
class InfoTest(Unittest):
    def __init__(self):
        super().__init__()
        self.name = "Info"
        self.method=requests.get

        legal_data = {"name": random_user_name(), "email": random_email(), "password": random_password()}
        RegisterTest().single_test(legal_data)
        responce=LoginTest().single_test(legal_data)
        res_dat:dict=json.loads(responce.content)
        print (res_dat)
        self.token=res_dat


        self.set_expect_io(legal_data)
    
    def set_expect_io(legal_data):
        dataset=({},legal_data)

        legal_output=dict(legal_data)
        legal_output.pop("Password")
        out=(
            {"code": 401, "msg": "权限不足"},
            {"code":200,"data":legal_data}
        )
    
    def remake(response: dict) -> dict:
        if "data" in response and "user" in response["data"]:
                # print("token:", data["data"]["token"])
                response["data"]["user"].pop("ID")
                response["data"]["user"].pop("CreatedAt")
                response["data"]["user"].pop("UpdatedAt")
                response["data"]["user"].pop("DeletedAt")
                response["data"]["user"].pop("Password")
        return response
        




@add_to_test_list
class EnrollReceiveTest(Unittest):
    def __init__(self):
        super().__init__()
        self.name = "enroll.receive"
        self.url = domain + "/enroll/receive"
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
        self.expected_io = zip(dataset, expected_return)

