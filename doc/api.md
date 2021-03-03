# API documentation

## 注册

路径：/api/auth/register

读取字段：[name],email,password 返回值：

返回值：

- 空邮箱：{"code": 422,  "msg":"illegal email address!"}
- 重复邮箱：{"code": 422,"msg":"exist email address!"}
- 缺用户名：{"code": 200,"msg":  "no username,auto generated.","name": name}
  - 缺失用户名会自动生成16位随机用户名
- 成功：{"code": 200,"msg":  "register successful."}

## 登录

路径：/api/auth/login

读取字段：email,password

返回值：
- 用户不存在：{"code": 422, "msg": "用户不存在"}
- 密码错误：{"code": 400, "msg": "用户名与密码不匹配"}
- token生成错误：{"code": 500, "msg": "系统异常"}
- 成功：{"code": 200,"data":{"token": token},"msg":"登录成功"}

## 接受招新简历

路径：/enroll/receive

读取字段：

- name：姓名，最多20个字
- gender：性别，最多2个字
- phone：手机，11位
- studentId:学号
- major：专业，最多10个字
- grade：入学年份
- firstChoice：第一志愿，最多五个字
- secondChoice：第二志愿，最多五个字
- introduction：自我介绍，最多250字
- hope：期望，最多250字
- hobbies：兴趣，最多50字

返回值：
- 成功：{"code": 200,"msg":"报名成功"}
        


