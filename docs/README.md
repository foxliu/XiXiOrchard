## 西西小果园 用于个人使用的量化交易系统
### 项目目录结构
```
XiXiOrchard/
│
├── cmd/                              # 主程序入口
│   └── main.go                       # 启动程序，初始化各模块
│
├── config/                           # 配置文件
│   ├── config.go                     # 配置结构体定义及加载
│   └── config.yaml                   # 配置文件（如API Key、数据库连接等）
│
├── internal/                         # 内部应用逻辑（不会被外部引用）
│   ├── data/                         # 数据获取与处理模块
│   │   ├── market_data.go            # 获取实时和历史数据的逻辑
│   │   └── storage.go                # 数据存储和管理
│   │
│   ├── strategy/                     # 策略开发与回测模块
│   │   ├── strategy.go               # 策略接口定义
│   │   ├── simple_strategy.go        # 实现简单的策略（如均线交叉）
│   │   └── backtest.go               # 回测框架及实现
│   │
│   ├── trade/                        # 交易执行模块
│   │   ├── order.go                  # 订单管理（创建、修改、撤销）
│   │   ├── execution.go              # 交易执行逻辑（API接口调用）
│   │   └── mock_execution.go         # 模拟交易执行（便于测试）
│   │
│   ├── risk/                         # 风险管理模块
│   │   └── risk_management.go        # 风控逻辑（止损、止盈、仓位管理）
│   │
│   ├── logger/                       # 日志与监控模块
│   │   └── logger.go                 # 日志记录器
│   │
│   └── utils/                        # 公共工具类
│       └── utils.go                  # 通用的帮助函数
│
├── tests/                            # 单元测试
│   ├── data_test.go                  # 数据模块测试
│   ├── strategy_test.go              # 策略与回测模块测试
│   ├── trade_test.go                 # 交易执行模块测试
│   └── risk_test.go                  # 风控模块测试
│
├── scripts/                          # 辅助脚本
│   └── data_download.py              # 下载数据的Python脚本（如果需要）
│
├── docs/                             # 文档
│   ├── README.md                     # 项目简介和说明
│   ├── DESIGN.md                     # 设计文档，说明系统架构及模块设计
│   └── API_REFERENCE.md              # API接口说明文档
│
├── frontend/                           # WPF 前端项目
│   ├── MainWindow.xaml                 # 主窗口
│   ├── MainWindow.xaml.cs              # 主窗口的逻辑代码
│   ├── ViewModels/                     # MVVM 中的 ViewModel
│   ├── Models/                         # MVVM 中的 Model
│   ├── Services/                       # 前端的服务类，可能包括与 gRPC 的通信
│   ├── Resources/                      # 静态资源，如图片、样式等
│   └── App.xaml                        # 应用程序入口
│
├── api/                                # gRPC API 服务
│   ├── proto/                          # Protocol Buffers 定义
│   │   └── healthcheck.proto           # 示例 gRPC 接口定义
│   ├── server.go                       # gRPC 服务器
│   ├── handlers/                       # gRPC 请求处理逻辑
│   │   ├── data_handler.go
│   │   ├── strategy_handler.go
│   │   ├── trade_handler.go
│   │   └── risk_handler.go
│   └── middleware/                     # 中间件
│
├── Makefile                          # 自动化任务脚本
├── go.mod                            # Go modules文件（依赖管理）
└── go.sum                            # 依赖版本锁定文件
```

### 目录结构说明

1. **cmd/**: 主程序入口，通常包含一个`main.go`文件，用于初始化和启动整个系统。

2. **config/**: 配置文件和配置结构体的定义，方便管理系统中的各种配置选项。

3. **internal/**: 核心业务逻辑，按照模块进行分离，确保各个模块职责清晰。
    - **data/**: 处理数据采集与存储的逻辑。
    - **strategy/**: 策略开发与回测功能的实现。
    - **trade/**: 交易执行模块，处理订单和执行交易。
    - **risk/**: 风险管理模块，管理风控逻辑。
    - **logger/**: 日志模块，记录系统运行时的各种日志信息。
    - **utils/**: 通用工具类模块，存放公共函数和工具类。

4. **tests/**: 各模块的单元测试文件，确保每个模块的功能可以独立测试。

5. **scripts/**: 存放辅助脚本文件，如数据下载或批量处理脚本。

6. **docs/**: 项目的文档目录，包括设计文档、API文档、README等。

7. **Makefile**: 用于定义常用的自动化任务，如编译、测试、部署等。

8. **go.mod / go.sum**: Go modules 相关文件，用于管理项目依赖。

### 未来扩展
- **扩展策略**：可以在`strategy/`目录中添加更多的策略文件，并在`strategy.go`中定义统一接口。
- **数据源扩展**：可以在`data/`目录中添加新的数据源（如其他交易所的API接口）。
- **交易执行**：在`trade/`目录中添加支持新的交易所或经纪商API的执行逻辑。
- **监控与报警**：未来可以增加监控模块，用于监控系统状态和策略运行情况，提供实时报警。

先从 **配置文件加载模块 (`config/config.go`)** 开始，然后逐步实现其他模块。

### 计划步骤：
1. **配置文件加载模块 (`config/config.go`)**
   - 定义配置结构体。
   - 实现配置文件的加载逻辑，支持从 `config.yaml` 读取配置。

2. **数据获取模块 (`internal/data/market_data.go`)**
   - 实现从交易所获取实时市场数据的功能。

3. **策略模块 (`internal/strategy/`)**
   - 定义策略接口。
   - 实现简单策略和回测框架。

4. **交易执行模块 (`internal/trade/`)**
   - 实现订单管理和交易执行逻辑。

5. **风险管理模块 (`internal/risk/`)**
   - 实现止损、止盈和仓位管理逻辑。

6. **日志模块 (`internal/logger/`)**
   - 实现系统日志记录功能。

7. **工具类 (`internal/utils/`)**
   - 实现通用工具函数。

8. **单元测试 (`tests/`)**
   - 为每个模块编写单元测试。

9. **gRPC API 服务 (`api/`)**
   - 定义和实现 gRPC 服务。

10. **前端项目 (`frontend/`)**
   - 使用 WPF 实现前端 UI。

