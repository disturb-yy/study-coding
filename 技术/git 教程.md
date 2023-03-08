# 使用 Git



## git clone —— 获取

`git clone [URL]`：用于将远程的仓库克隆到本地



## git status —— 显示状态



## git pull —— 用于从远程存储库获取更新结果



## git add —— 将当前目录下修改的所有代码从工作区添加到暂存区 . 代表当前目录



## git ccommit —— 记录存储库的变化

使用`git commit`命令以后，只是将存储库的变化记录，如果要同步到远程存储库，还需要使用`git push`命令提交更新。

常见的`git commit`命令执行方式如下：

```
git commit -a -m "commit message"
```

其中`-m "message"`用于指定与当前命令伴随的消息，`-a`选项用于通知`git commit`命令自动包含所有的修改后的文件。



## git push —— 提交更新

