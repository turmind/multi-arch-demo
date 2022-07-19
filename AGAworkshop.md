# AWS Global Accelerator/Custom routing

## 创建测试实例

进入EC2 Instances页面，点击Launch instances启动新实例

![创建实例](img/Screen%20Shot%202022-07-18%20at%201.47.21%20PM.png)

输入实例名字并选择Arm的AMI

![选择arm](img/Screen%20Shot%202022-07-18%20at%201.50.48%20PM.png)

选择实例类型t4g.medium

![选择类型](img/Screen%20Shot%202022-07-18%20at%202.19.47%20PM.png)

点击创建实例

## 对实例附加EC2角色

对EC2增加角色，满足后续使用aws cli访问GA信息及登录的请求

EC2附加角色

![EC2附加角色](img/Screen%20Shot%202022-07-18%20at%202.27.31%20PM.png)

在新页面中选择创建角色

![创建角色](img/Screen%20Shot%202022-07-18%20at%202.31.15%20PM.png)

附加AdministratorAccess的权限 *仅测试使用才这样配置*

![附加权限](img/Screen%20Shot%202022-07-18%20at%202.34.10%20PM.png)

在EC2中选择刚才创建的角色

![附加角色](img/Screen%20Shot%202022-07-18%20at%202.44.55%20PM.png)

## 连接实例

使用Session Manager连接实例 *在更改角色后需要等待一段时间才可以使用*

![连接实例](img/Screen%20Shot%202022-07-18%20at%202.50.22%20PM.png)

## 运行tcp echo server

切换root用户 *仅为了方便切换，在生产环境谨慎使用*

```linux
sudo -i
```

安装git客户端 *为下载相应的测试包*

```linux
yum install -y git
```

下载tcp echo server，地址[server](https://github.com/turmind/multi-arch-demo)

```linux
git clone https://github.com/turmind/multi-arch-demo.git
```

启动服务

```linux
nohup ./multi-arch-demo/bin/echo-server_arm &
```

## EC2安全组设置

通过EC2中的安全选中对应的安全组

![选择修改的安全组](img/Screen%20Shot%202022-07-18%20at%202.57.15%20PM.png)

修改进入规则

![修改进入规则](img/Screen%20Shot%202022-07-18%20at%202.58.38%20PM.png)

增加4000-5000的全网访问

![增加访问端口](img/Screen%20Shot%202022-07-18%20at%202.59.45%20PM.png)

使用telnet进行测试访问，确定服务部署正确及安全组设置正确

```linux
telnet 54.196.211.185 4000
```

## 创建GA

输入GA名称并选择Custom routing

![选择Custom routing](img/Screen%20Shot%202022-07-18%20at%208.22.26%20PM.png)

增加监听端口范围

监听的端口数必须满足子网IP数及每台EC2的端口乘积的总数,如示例中，子网掩码为20，每台EC2监听两个端口，端口总数需要达到4096 * 2 = 8192。

![输入端口范围](img/Screen%20Shot%202022-07-18%20at%2010.56.42%20PM.png)

增加endpoint Groups *选择实例所在region，端口范围从4000-4001*

![endpoint groups](img/Screen%20Shot%202022-07-18%20at%208.25.42%20PM.png)

增加endpoint

![endpoint](img/Screen%20Shot%202022-07-18%20at%208.31.33%20PM.png)

## 查看对应端口的映射

更新aws cli，参考地址：[aws cli update](https://docs.aws.amazon.com/zh_cn/cli/latest/userguide/getting-started-install.html)

```linux
curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
```

命令行查看对应映射的端口，参考地址：[aws globalaccelerator help](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/globalaccelerator/list-custom-routing-port-mappings-by-destination.html)

- region us-west-2为必须项，需要在us-west-2中访问ga的配置信息
- endpoint-id 子网id
- destination-address EC2私有地址

```linux
/usr/local/bin/aws globalaccelerator list-custom-routing-port-mappings-by-destination --endpoint-id subnet-03dce6e53d824d53b --destination-address 172.31.31.134 --region us-west-2
```

## 通过客户端进行测试

- 需要在海外的客户端或者服务端进行部署测试
- 在国内或在AWS服务器中测试，GA的访问速度可能更慢

```linux
git clone https://github.com/turmind/echo-cli.git
./echo-cli_arm -h 54.196.211.185 -p 4000
./echo-cli_arm -h 15.197.176.249 -p 8964
```
