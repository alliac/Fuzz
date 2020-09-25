# Fuzz

![](https://img.shields.io/badge/Language-Go-brightgreen)

基于Beego的Fabric(v1.1 and v2.0)智能合约模糊测试

在此之前请确保已安装：

```
Hyperledger Fabric v1.1
Hyperledger Fabric v2.0
Beego
Go-fuzz
Go v1.12.1
```

如有Bug，欢迎提出

Step 1：

```
cd $GOPATH
mkdir example
git clone https://github.com/doporg/Fuzz.git
```

Step 2：

```
cd server
bee run
```

Step 3：

前端

```
git clone https://github.com/doporg/dop-web.git
```

Step 4：

```
npm install
npm start
```

Step 5：

```
访问http://localhost:3000/#/fuzz
```

### 待测智能合约

待测智能合约需满足以下条件：

```
1.Fabric v1.1 or v2.0
2.package包名为“test”
3.v1.1待测智能合约链码对象名称为“Chaincode”，v2.0为“SimpleChaincode”
```

### 表单格式

**Test Case**

以Json格式给出模糊测试初始语料，key为待测智能合约功能函数名称，value为该功能函数的参数。

如某待测智能合约包含`addIngInfo`和`getIngInfo`两个功能函数，addIngInfo函数包含FoodID、IngID、IngName三个参数，getIngInfo包含FoodID一个参数。则该待测智能合约Test Case格式如下：

```
{"addIngInfo":{"FoodID":"001","IngID":"aa","IngName":"bb"},"getFoodInfo":{"FoodID":"001"},"getIngInfo":{"FoodID":"001"}}
```

**Version**

v1.1版本待测智能合约输入1.1，v2.0版本输入2.0

**Smart Contract**

上传待测智能合约

**Button**

```
start：执行测试
stop：停止测试，返回结果数据（建议测试执行10min以上）
Download：下载测试结果。包含crashers、corpus、suppressions和output.log
```

### Dockerfile

**已打包镜像**

```
docker pull alliac0901/fuzz-image:1.2
docker pull alliac0901/fuzz-image:1.0（仅支持Fabric v2.0版本智能合约）
docker pull registry.dop.clsaa.com/dop/fuzz-image:1.2
```

**本地镜像打包**

依赖包（置于根目录下）

| 文件夹名称 |                        说明                         |
| :--------: | :-------------------------------------------------: |
|  dvyukov   | go-fuzz工具：https://github.com/dvyukov/go-fuzz.git |
| fabric2.0  |               Hyperledger Fabric v2.0               |
|     go     |                     go v1.12.1                      |
|    bin     |        bee、go-fuzz、go-fuzz-build可执行文件        |

Step 1：

```
cd server
docker build -t fuzz-image .
```

Step 2：

```
docker images
//查看是否存在fuzz-image镜像
```

Step 3：

```
docker run -it --name test-instance -p 8080:8080 fuzz-image
```

完