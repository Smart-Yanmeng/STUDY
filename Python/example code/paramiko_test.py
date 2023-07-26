# 导入paramiko，（导入前需要先在环境里安装该模块）
import paramiko
import sys


# 定义函数ssh,把操作内容写到函数里,函数里接收参数（写在括号里）,其中port=是设置一个默认值如果没传就用默认
def sshExeCMD(ip, username, password, cmd, port=22):
    # 定义一个变量ssh_clint使用SSHClient类用来后边调用
    ssh_client = paramiko.SSHClient()
    # 通过这个set_missing_host_key_policy方法用于实现登录是需要确认输入yes，否则保存
    ssh_client.set_missing_host_key_policy(paramiko.AutoAddPolicy())
    # 使用try做异常捕获
    try:
        # 使用cnnect类来连接服务器
        ssh_client.connect(hostname=ip, port=port, username=username, password=password)
    # 如果上边命令报错吧报错信息定义到err变量，并输出。
    except Exception as err:
        print("服务器链接失败！！！", ip)
        print(err)
        # 如果报错使用sys的exit退出脚本
        sys.exit()
    # 使用exec_command方法执行命令，并使用变量接收命令的返回值并用print输出
    # 这里也可以把命令做成参数传进来
    stdin, stdout, stderr = ssh_client.exec_command(cmd[0])
    # print(stdout.read().decode("utf-8"))
    # stdin, stdout, stderr = ssh_client.exec_command("ls")
    # 使用decode方法可以指定编码
    print(stdout.read().decode("utf-8"))


# 通过判断模块名运行上边函数
if __name__ == '__main__':
    # 定义一个字典，写服务器信息
    servers = {
        # 以服务器IP为键,值为服务器的用户密码端口定义的字典
        "39.105.219.78": {
            "username": "root",
            "password": "12345678Lc",
            "port": 22,
            "cmd": ["python3 ./server.py"]
        }
    }
    # 使用items方法遍历，使用ip 和info把字典里的键和值遍历赋值，传给函数sshExeCMD
    for ip, info in servers.items():
        # 这里的info也就是上边定义的ip对应值的字典，使用get获取这个字典里对应username键对应的值，赋值给变量username传给函数中使用
        sshExeCMD(
            ip=ip,
            username=info.get("username"),
            password=info.get("password"),
            cmd=info.get("cmd"),
            port=info.get("port")
        )
