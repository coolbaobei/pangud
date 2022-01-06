Pangud开发者文档

PANGUD采用DDD方法，使用依赖反转的改良分层架构，代码依赖顺序自下往上依次为领域层、应用层、接口层、基础设施层。接口层和基础设施层对应六边形架构的适配器，主要职责为：提供对外REST接口，连接数据库，对接第三方系统等。 代码结构：

```bash
pangud
  |--cmd                #应用入口
  |--docs
  |--pkg
  |--|--infrastructure
  |--|--interface
  |--|--|--rest
  |--|--domain
  |--|--|--model         #模型
  |--|--|--repositroy    #存储库
  |--|--|--service       #领域服务
  |--|--|--event         #领域事件
  |--|--shared           #共享代码 各层通用
  |--|--|--util          #工具
  |--|--|--config        #配置
  |--docker              #Dockerfile
  |--config.example.yml  #示例配置文件
```
