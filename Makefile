.PHONY: help install dev build clean test

# 默认目标
help:
	@echo "可用命令:"
	@echo "  make install         - 安装所有子项目的依赖"
	@echo "  make dev             - 启动所有服务的开发模式"
	@echo "  make build           - 构建所有子项目"
	@echo "  make clean           - 清理所有构建产物"
	@echo "  make test            - 运行所有测试"
	@echo ""
	@echo "单独服务命令:"
	@echo "  make dev-admin       - 仅启动前端开发服务器"
	@echo "  make dev-server      - 仅启动后端服务"
	@echo "  make dev-proxy       - 仅启动代理服务"
	@echo "  make dev-log-service - 仅启动日志服务"
	@echo "  make build-admin     - 仅构建前端"
	@echo "  make build-server    - 仅构建后端（本地）"
	@echo "  make build-server-linux - 构建后端（Linux 静态链接）"
	@echo "  make build-proxy     - 仅构建代理服务"
	@echo "  make build-log-service - 仅构建日志服务"

# 安装所有依赖
install:
	@echo "安装前端依赖..."
	cd admin && pnpm install
	@echo "安装后端依赖..."
	cd server && go mod download
	@echo "安装代理服务依赖..."
	cd proxy && go mod download
	@echo "安装日志服务依赖..."
	cd log-service && go mod download
	@echo "✅ 所有依赖安装完成"

# 开发模式 - 启动所有服务
dev:
	@echo "⚠️  请在不同的终端窗口中分别运行以下命令:"
	@echo "  make dev-admin       # 前端 (端口 6806)"
	@echo "  make dev-server      # 后端 (端口 6808)"
	@echo "  make dev-proxy       # 代理 (端口 6800)"
	@echo "  make dev-log-service # 日志服务 (端口 6809)"

# 仅启动前端
dev-admin:
	@echo "启动前端开发服务器..."
	cd admin && pnpm dev

# 仅启动后端
dev-server:
	@echo "启动后端服务..."
	cd server && go run cmd/server/main.go

# 仅启动代理
dev-proxy:
	@echo "启动代理服务..."
	cd proxy && go run main.go

# 仅启动日志服务
dev-log-service:
	@echo "启动日志服务..."
	cd log-service && go run cmd/server/main.go

# 构建所有项目
build:
	@echo "构建前端..."
	cd admin && pnpm build
	@echo "构建后端..."
	cd server && go build -o bin/server cmd/server/main.go
	@echo "构建代理服务..."
	cd proxy && go build -o bin/proxy main.go
	@echo "构建日志服务..."
	cd log-service && go build -o bin/log-service cmd/server/main.go
	@echo "✅ 所有项目构建完成"

# 构建前端
build-admin:
	cd admin && pnpm build

# 构建后端
build-server:
	cd server && go build -o bin/server cmd/server/main.go

# 构建后端 Linux 版本（静态链接）
build-server-linux:
	cd server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/server cmd/server/main.go

# 构建代理
build-proxy:
	cd proxy && go build -o bin/proxy main.go

# 构建日志服务
build-log-service:
	cd log-service && go build -o bin/log-service cmd/server/main.go

# 清理构建产物
clean:
	@echo "清理前端构建产物..."
	rm -rf admin/dist
	@echo "清理后端构建产物..."
	rm -rf server/bin
	@echo "清理代理服务构建产物..."
	rm -rf proxy/bin
	@echo "清理日志服务构建产物..."
	rm -rf log-service/bin
	@echo "✅ 清理完成"

# 运行测试
test:
	@echo "运行后端测试..."
	cd server && go test ./...
	@echo "运行代理服务测试..."
	cd proxy && go test ./...
	@echo "运行日志服务测试..."
	cd log-service && go test ./...
	@echo "✅ 测试完成"

