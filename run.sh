#!/bin/sh
./subs-check -f config.yaml &
PID=$!

echo "subs-check 已启动，PID: $PID"

# 每20秒检查一次输出文件
while true; do
  sleep 20
  if [ -f "output/all.yaml" ] && [ -s "output/all.yaml" ]; then
    echo "检测完成，输出文件已生成，等待Gist写入..."
    sleep 15
    kill $PID 2>/dev/null || true
    echo "已退出"
    exit 0
  fi
  echo "等待检测完成..."
done
