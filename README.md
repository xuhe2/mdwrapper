[English](./README.md) | [中文](./docs/README_CN.md)

# What is it?

mdwrapper is a tool for packaging markdown files. It can recognize the references of **images and files** in markdown files, automatically package the **images and files** into the specified directory, and replace the reference paths in the markdown file.

# Why?

1. File References
> Converting to a PDF file can save the referenced image files, but it cannot save the referenced external resource files (e.g., `[file](dir/some_file)`) as attachments to the PDF.
> Manually packaging is time-consuming and laborious.

2. Editability
> Editing PDF files is not convenient, so keep the original markdown files for later modification and sharing.
> If you are a typora user, the default storage path for typora images is in its own data folder, and finding the corresponding images is not convenient enough.

# How to install

## Go

```shell
go install github.com/xuhe2/mdwrapper@latest
```

> Make sure you have installed the GO environment

# How to use?

```shell
./bin/mdwrapper /path/to/markdown
```
> The default output is a .zip file, with the same filename as the markdown file.
- Use the -o parameter to specify the output .zip file name.