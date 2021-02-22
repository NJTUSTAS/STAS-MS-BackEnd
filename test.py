from my_unit import *

# 测试域名
domin = "http://localhost:8080"

# 输入数据编辑
raw_data = \
    {"name": random_user_name(),
     "email": random_email(),
     "password": random_password()}
err_email_data = {**raw_data, "email": ""}
err_dupemail_data = {**raw_data, "email": "1"}
warn_name_data = {**raw_data, "name": "", "email": random_email()}

# 输入数据顺序
dataset = (raw_data, err_dupemail_data, err_email_data, warn_name_data)
# 期望输出数据
expected_return = (
    {"code": 200, "msg": "register successful."},
    {"code": 422, "msg": "exist email address!"},
    {"code": 422, "msg": "illegal email address!"},
    {"code": 200, "msg": "no user name,auto generated.", "name": ""},
)

expected_io = zip(dataset, expected_return)


def test():
    unittest(expected_io, domin)
    print("\n")


if __name__ == "__main__":
    for i in range(10):
        test()