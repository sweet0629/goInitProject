package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法：goInitProject [项目名称]")
		fmt.Println("示例：goInitProject myproject")
		os.Exit(1)
	}

	projectName := os.Args[1]
	if projectName == "." {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Printf("错误：获取当前目录失败：%v\n", err)
			os.Exit(1)
		}
		projectName = filepath.Base(wd)
	}

	if err := createProject(projectName); err != nil {
		fmt.Printf("错误：%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ 项目 %s 创建成功!\n", projectName)
}

func createProject(name string) error {
	dirs := []string{
		filepath.Join(name, "cmd", "server"),
		filepath.Join(name, "cmd", "worker"),
		filepath.Join(name, "internal", "config"),
		filepath.Join(name, "internal", "handler"),
		filepath.Join(name, "internal", "middleware"),
		filepath.Join(name, "internal", "model"),
		filepath.Join(name, "internal", "repository"),
		filepath.Join(name, "internal", "service"),
		filepath.Join(name, "pkg", "cache"),
		filepath.Join(name, "pkg", "database"),
		filepath.Join(name, "pkg", "logger"),
		filepath.Join(name, "pkg", "utils"),
		filepath.Join(name, "api"),
		filepath.Join(name, "configs"),
		filepath.Join(name, "scripts"),
		filepath.Join(name, "deployments"),
		filepath.Join(name, "docs"),
		filepath.Join(name, "tests"),
		filepath.Join(name, "third_party"),
	}

	for _, dir := range dirs {
		if err := createDirectory(dir); err != nil {
			return fmt.Errorf("创建目录 %s 失败：%w", dir, err)
		}
	}

	if err := writeFile(filepath.Join(name, "go.mod"), generateGoMod(name)); err != nil {
		return fmt.Errorf("创建 go.mod 失败：%w", err)
	}

	return nil
}

func createDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func generateGoMod(name string) string {
	return fmt.Sprintf(`module github.com/sweet0629/%s

go 1.21
`, name)
}
