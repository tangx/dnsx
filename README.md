# DNSx

多后台的 dns 解析配置命令行

+ [x] 使用 `interface` 方法实现多 `backend` 接入
+ [x] 使用 `cobra` 创建子命令
  + [x] `add`
  + [x] `delete`: 交互确认。
  + [ ] `update`
  + [ ] `enable`, `disable`
  + [x] `search`
  + [x] `configure`
+ [x] 使用 `cobra` 实现 `config.json` 的配置
  + [x] 增加
  + [x] 更新
  + [x] 优化 `LoadConfig()` 加载位置

+ [ ] 完成 `auto-complete` for zsh: 支持子命令补全， `config.json` 中的域名列表补全

+ 支持多后台
  + [x] `qcloud cns`
  + [x] `aliyun alidns`
  + [ ] `dnspod`

