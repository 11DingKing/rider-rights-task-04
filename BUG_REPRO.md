# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

人工录入的诉求类别有时带着首尾空格，或大小写和分办规则的写法不一致；明明是同一类别，事项却没有分到对应部门。请修复类别匹配的输入整理，同时保留关键词规则和默认兜底的职责。改动请留在生产代码，测试文件不要改。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-04
- 仓库地址：https://github.com/11DingKing/rider-rights-task-04.git
- parent SHA：604eac06bc54b98bd9b020f95ab5e9503d5347b5

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-04.git bug-repro
cd bug-repro
git checkout --detach 604eac06bc54b98bd9b020f95ab5e9503d5347b5
go test ./internal/domain -run "^TestRuleCategoryMatchingNormalizesInput$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestRuleCategoryMatchingNormalizesInput$" -count=1
--- FAIL: TestRuleCategoryMatchingNormalizesInput (0.00s)
    task04_test.go:9: category matching ignored case or surrounding whitespace
FAIL
FAIL	riderguard/internal/domain	0.038s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestRuleCategoryMatchingNormalizesInput$" -count=1
--- FAIL: TestRuleCategoryMatchingNormalizesInput (0.00s)
    task04_test.go:9: category matching ignored case or surrounding whitespace
FAIL
FAIL	riderguard/internal/domain	0.001s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，题面中的触发条件应得到正确业务结果，原始异常不再出现；定向验证、相关包测试和仓库全量回归测试必须通过，不得通过删除或跳过测试、降低断言强度或绕过目标逻辑使验证转绿。
