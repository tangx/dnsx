# DNSx

多后台的 dns 解析配置命令行

+ [x] 使用 `interface` 方法实现多 `backend` 接入
+ [x] 使用 `cobra` 创建子命令
  + [x] `add`
  + [ ] `delete`
  + [ ] `update`
+ [ ] 使用 `cobra` 实现 `config.json` 的配置
  + [ ] 增加
  + [ ] 更新
  + [ ] 优化 `LoadConfig()` 加载位置

+ [ ] 完成 `auto-complete` for zsh: 支持子命令补全， `config.json` 中的域名列表补全

+ 支持多后台
  + [x] `qcloud cns`
  + [x] `aliyun alidns`
  + [ ] `dnspod`

