from abc import abstractmethod

import requests


class UnsetMethodError(Exception):
    def __init__(self, name):
        self.name = name

    def __str__(self):
        return f"{self.name}测试类没有设置调用接口类型（get/post/etc）"


class UnsetAddressError(Exception):
    def __init__(self, name):
        self.name = name

    def __str__(self):
        # self.reason = "测试类没有设置接口相对路径"
        return f"测试类: {self.name} 没有设置接口相对路径"


class TestUnit:
    # 本地测试
    domain = "http://localhost:8080"

    # 已经在远程服务器上部署
    # domain = "http://202.119.245.31:80"

    @abstractmethod
    def __init__(self):
        """子类需要指定相对路径和调用方法"""
        self.name: str = "未命名异常类"
        self.url = ""
        self.method = None

    @staticmethod
    @abstractmethod
    def expect_io():
        """生成一组需要进行测试的数据"""
        pass

    @staticmethod
    def format(data: dict) -> dict:
        """后端返回值中有些不能事先确定的，如登录token,通过这个方法处理掉"""
        return data

    def pre_call(self, data) -> dict:
        """单次调用，给登录需要注册之类需要前置的接口来调用前置"""
        try:
            return self.method(self.url, data).json()
        except TypeError as e:
            if "can only concatenate str (not \"NoneType\") to str" in e.args:
                raise UnsetAddressError(self.name)
            if "'NoneType' object is not callable" in e.args:
                raise UnsetMethodError(self.name)

    def single_test(self, input_data: dict, expect: dict) -> bool:
        try:
            response = self.method(self.url, input_data).json()
            ret = self.format(response) == expect
        except TypeError as e:
            raise UnsetMethodError(self.name)
        except requests.exceptions.MissingSchema:
            raise UnsetAddressError(self.name)
        if not ret:
            print(f"{self.name}未通过测试。"
                  "未通过测试点:\n"
                  f"\t{input_data=}\n"
                  f"\t{expect=}\n"
                  f"\t{response=}\n")
        return ret

    def group_test(self) -> bool:
        """测试一整组数据"""
        ret = all(self.single_test(data, expect) for data, expect in self.expect_io())
        if __debug__ and ret:
            print(f"{self.name}已经通过一组测试")
        return ret

    def test(self, count):
        ret = all(self.group_test() for _ in range(count))
        if ret:
            print(f"{self.name}已经通过测试")
