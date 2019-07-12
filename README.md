# mysqlgogenerator

## 功能

连接MySQL数据库，生成所有表的  go model struct  还有 SELECT  INSERT 语句





## 快速开始

#### 0、git clone 

#### 1、修改 config.toml 中数据库连接 MysqlAddr  名称到具体数据库

#### 2、windows 下 需要使用git bash，参考以下步骤

```
# pwd
xxxx/mysqlgogenerator
#ls
app/  cmd/  config/  config.toml  out/  query.tpl  README.md  run.sh  vendor/
# sh run.sh
```

#### 3、Linux环境下 运行 run.sh 

#### 4、输出结果在 /out 文件夹下





### 编译tips

首先将源码目录放在 **GOPATH** 以下的目录

查看本地go version

1、go version 低于 1.11 使用 go vendor 

2、go version 高于 1.11 使用 go module进行包管理 

go mod 可能需要设置GO111MODULE=on

![https://github.com/elegantm/mysqlgogenerator/pic/2019-07-12_14-50-23.png]()

