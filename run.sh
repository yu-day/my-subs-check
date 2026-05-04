#!/bin/sh
./subs-check -f config.yaml &
PID=$!

# 每隔20秒检查输出是否已写入
while true; do
  sleep 20
  if [ -f "output/all.yaml" ] && [ -s "output/all.yaml" ]; then
    echo "输出文件已生成，等待Gist写入完成..."
    sleep 10   # 再等10秒确保Gist写入完毕
    kill $PID
    exit 0
  fi
done
