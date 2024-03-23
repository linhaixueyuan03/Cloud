# Cloud
> 轻量级云盘  
> 基于go-zero、xorm实现

## 实现功能
- [x] 用户登录  
- [x] 用户详情  
- [ ] 邮箱注册  
- [ ] 文件关联存储  
- [ ] 文件夹层级管理  
- [ ] 文件分享

## 项目规划


## 系统依赖
| 名称     | 版本    |
|--------|-------|
| golang | 1.21  |
| xorm   | 1.3.9 |

## 项目运行
```
# 启动服务  
go run core.go -f etc/core-api.yaml
# 使用api生成代码
goctl api go -api core.api -dir . -style go_zero

```
