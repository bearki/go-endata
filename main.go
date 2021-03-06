package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bearki/go-endata/tool"
	"github.com/urfave/cli/v2"
)

// 入口
func main() {
	log.Println("工具运行中···")

	// 获取可执行文件的名称
	execFileName := filepath.Base(os.Args[0])

	// 创建一个命令行程序
	app := &cli.App{
		// 填写程序信息
		Name:      "AssetsToGoCode",
		Usage:     "Encode static resources into Golang code",
		UsageText: execFileName + " command [command options]",
		Version:   "v0.0.3",

		// 命令列表
		Commands: []*cli.Command{
			// 将资源编码成GO文件
			{
				Name:      "create",
				Usage:     "Encode static resources into Golang file",
				UsageText: execFileName + " create [command options]",
				Flags: []cli.Flag{
					// 需要编码的文件或文件夹
					&cli.StringFlag{
						Name: "src",
						Aliases: []string{
							"s",
						},
						Usage:    "The path of the file or folder that needs to be encoded",
						Required: true,
					},

					// 编码后的输出文件夹
					&cli.StringFlag{
						Name: "out",
						Aliases: []string{
							"o",
						},
						Value:    "out",
						Usage:    "Encoded Go file output folder",
						Required: false,
					},

					// 文件作者（默认为空）
					&cli.StringFlag{
						Name: "author",
						Aliases: []string{
							"a",
						},
						Value:    "",
						Usage:    "The author of the file",
						Required: false,
					},

					// 生成的Gin路由文件使用到的包名前缀（默认为空）
					&cli.StringFlag{
						Name: "pack-prefix",
						Aliases: []string{
							"p",
						},
						Value:    "",
						Usage:    "Package prefix of the generated Go file",
						Required: false,
					},

					// 是否显示详细信息
					&cli.BoolFlag{
						Name: "view",
						Aliases: []string{
							"v",
						},
						Value:    false,
						Usage:    "whether to show detailed information",
						Required: false,
					},
				},
				Action: func(c *cli.Context) error {
					// 执行Go文件生成
					return tool.RunTool(tool.InputParams{
						SrcPath:    c.String("src"),
						OutDirPath: c.String("out"),
						Author:     c.String("author"),
						PackPrefix: c.String("pack-prefix"),
						ShowDetail: c.Bool("view"),
					})
				},
			},

			// 生成Ext Map ContentType JSON文件
			{
				Name:      "extmap",
				Usage:     "Create a JSON file with file suffix mapping Content-Type",
				UsageText: "assets_to_gocode.exe extmap [command options]",
				Action: func(c *cli.Context) error {
					return tool.CreateExtMapJson()
				},
			},
		},
	}

	// 运行程序
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("工具运行错误: %s", err.Error())
		return
	}
	log.Println("工具运行成功")
}
