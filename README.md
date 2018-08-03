# deag

**D**istributed **E**volutionary **A**lgorithms in **G**olang

## Introduction

An Implement of deap ([https://github.com/DEAP/deap](https://github.com/DEAP/deap)) which full name is "Distributed Evolutionary Algorithms in Python"

If you want to run Evolutionary Calculation experiment in python, deap is a good choice.

The project is still in its infancy, and interested parties are welcome to participate in the development.

The old repositoy of this project: https://gitee.com/sineatos/deag

## Architecture

```
deag
|-algorithms        // Algorithm implemented by deag
|  |-de             // Differential Evolution
|  |-pso            // Particle Swarm Optimization
├─base              // basic structure
├─benchmarks        // benchmark function
├─tools             // tools
|  |-constraint     // constraint
│  ├─crossover      // common cross operation
|  |-emo            // multi-objective operation
│  ├─inits          // common initialization operation
│  ├─mutation       // common mutation operation
│  ├─selection      // common selection operation
│  └─support        // common auxiliary structures, such as statistics, etc.
└─utility           // utilities
```

## Installation

1. This project has not yet planned to use any third-party libraries. The plan is that core functions allow users to use third-party libraries to enhance their functions.
2. `go get https://github.com/sineatos/deag`

## How to use

1. Because of the documentation of deag has not completed, please read the tests files in each package (especially the `deag/algorithm`) and learn how to use.

## TODO

Feature | Description | progress rate | test(coverage)
---|---|---|---
base | | | 74.1%
base.Fitness | Fitness | Finish | Almost Done
base.Individual | Individual | Finish | Almost Done
base.Individuals | Population | Finish | Finish
tools.inits | Initialization | Finish | 100%
tools.mutation | Mutation | Finish | 96.9%
tools.crossover | Crossover | Finish | 93.1%
tools.selection | Selection | Finish | 83.0%
tools.support | | | 87.9%
tools.support.statistics | Statistics | Finish | Almost Done
tools.support.logbook | Logbook | Finish | Finish
tools.support.halloffame | HOF | Finish | Finish
tools.support.constraint | Constraint | Finish | Finish
tools.support.pareto_front | Pareto Front | Finish | 0%
benchmarks | | | 87.7%
benchmarks.binary | Binary benchmark | Finish | Finish
benchmarks.single_objective | Single objective benchmark | Finish | Almost Done
benchmarks.gp | GP benchmark | 0% | 
benchmarks.movingpeaks | Moving peaks | 0% | 
benchmarks.multi_objectives | Multi-objectives benchmark | Finish | 0%
benchmarks.btools | Tools using with benchmark | 0% | 

### Others
1. Implement concurrent with deag by using goroutine
2. modify comment and document

----

## 项目介绍

**D**istributed **E**volutionary **A**lgorithms in **G**olang

借鉴deap([https://github.com/DEAP/deap](https://github.com/DEAP/deap))，编写的一个Golang下的进化计算框架

本项目还在初级阶段，希望感兴趣的人参与开发

项目原仓库地址：https://gitee.com/sineatos/deag

## 软件架构
软件架构说明

```
deag
|-algorithms        // 采用deag实现的算法
|  |-de             // 差分进化算法
|  |-pso            // 粒子群算法
├─base              // 基础结构
├─benchmarks        // 基准函数
├─tools             // 工具
|  |-constraint     // 约束
│  ├─crossover      // 常用交叉操作
|  |-emo            // 多目标操作(目前只有NSGA2的选择)
│  ├─inits          // 常用初始化操作
│  ├─mutation       // 常用变异操作
│  ├─selection      // 常用选择操作
│  └─support        // 常用辅助结构，如统计量等
└─utility           // 实用工具
```

## 安装教程

1. 本项目暂时还没打算使用任何第三方库，计划是核心功能可以不使用第三方库，但允许用户使用第三方库加强其功能

## 使用说明

1. 参考deag/algorithms中各个子包的测试文件。

## 项目

Golang作为强类型的静态编译性语言，加上Golang的语法糖不多，因此无法做到很多Python等语言的“骚操作”，所以虽然借鉴了deap，但是还是无法实现很多deap中功能。
另一方面，由于考虑到框架的效率问题，deap中的一些功能虽然可以借鉴`interface{}`或者`reflection`等方式实现，但是还是没有选择使用这些方式实现。

实际上golang下也有进化计算框架，如evo，但是看了一下这个框架的大体结构以后，我还是希望能实现一个不同的框架，最好能够兼容多目标算法，因此借鉴deap对多目标问题的处理。

本项目由于是使用Golang实现的，因此计划引入一些Golang的强项特性，如`goroutine`等，但计划在基础功能实现完以后再加入

### 目前借鉴了deap的设计

设计 | 对应实现
---|---
适应值 | base.Fitness
个体 | base.Individual
统计量 | tools.support.statistics
记录 | tools.support.logbook
HallOfFame | tools.support.halloffame
约束 | tools.support.constraint

以上结构除了适应值是结构体以外，个体、统计量和记录都是借口，同时也给出了一个默认的实现(以`Default`命名开头)，其中默认实现为deap的实现，这样设计的目的一是使得整个项目更像一个工程，二是抽象出每一个结构应该实现的功能，三是使得用户可以有自己的实现。

### 移植了deap中的功能模块
模块 | 对应实现
---|---
初始化 | tools.inits
变异 | tools.mutation
交叉 | tools.crossover
选择 | tools.selection
基准函数 | benchmarks

## TODO

这里只写完成基本功能，详细功能需要基本功能能用以后再实现

功能 | 说明 | 实现 | 测试(覆盖率)
---|---|---|---
base | | | 74.1%
base.Fitness | 适应值 | 完成 | 基本完成
base.Individual | 个体 | 完成 | 基本完成
base.Individuals | 种群 | 完成 | 完成
tools.inits | 初始化操作 | 完成 | 100%
tools.mutation | 变异操作 | 完成 | 96.9%
tools.crossover | 交叉操作 | 完成 | 93.1%
tools.selection | 选择操作 | 完成 | 83.0%
tools.support | | | 87.9%
tools.support.statistics | 统计量 | 完成 | 基本完成
tools.support.logbook | 记录 | 完成 | 完成
tools.support.halloffame | HOF | 完成 | 完成
tools.support.constraint | 约束 | 完成 | 完成
tools.support.pareto_front | 帕累托前沿 | 完成 | 0%
benchmarks | | | 87.7%
benchmarks.binary | 二进制benchmark | 完成 | 完成
benchmarks.single_objective | 单目标benchmark | 完成 | 基本完成
benchmarks.gp | GP benchmark | 0% | 
benchmarks.movingpeaks | 移动峰 | 0% | 
benchmarks.multi_objectives | 多目标benchmark | 完成 | 0%
benchmarks.btools | 基准函数使用到的工具 | 0% | 

### 其他TODO

1. 结合goroutine实现并发
2. 完善注释和文档
