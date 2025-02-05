[English](./README.md) | [中文](./docs/README_CN.md)

# 是什么

mdwrapper是一个用来打包markdown文件的工具,它可以识别markdown文件中对应**图片\文件的引用**,自动将**图片\文件**打包到指定目录中,并替换markdown文件中的引用路径.

# 为什么

1. 文件引用
> 转换为PDF文件可以保存引用的图片文件,但是,无法保存引用的外部资源文件(例如`[file](dir/some_file)`)作为PDF的附件.
> 手动打包费时费力

2. 可编辑性
> 编辑PDF文件并不方便,所以,保留原本的markdown文件,方便后续的修改和分享
> 如果你是typora用户,typora图片的默认存储路径是在自己的数据文件夹下,查找对应的图片不够方便.

# 怎么用

```shell
./bin/mdwrapper /path/to/markdown
```
> 默认输出为.zip文件,文件名与markdown文件名相同

- 使用`-o`参数指定输出.zip文件名字