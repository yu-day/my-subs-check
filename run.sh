#!/bin/sh
./subs-check -f config.yaml &
PID=$!
echo "subs-check 已启动，PID: $PID"

MAX_WAIT=1800
ELAPSED=0
while [ $ELAPSED -lt $MAX_WAIT ]; do
  sleep 20
  ELAPSED=$((ELAPSED + 20))
  if [ -f "output/all.yaml" ] && [ -s "output/all.yaml" ]; then
    echo "检测完成，等待Gist写入..."
    sleep 15
    kill $PID 2>/dev/null || true
    echo "已正常退出，耗时约 ${ELAPSED} 秒"
    exit 0
  fi
  echo "等待检测完成... (${ELAPSED}/${MAX_WAIT}秒)"
done

echo "超时退出"
kill $PID 2>/dev/null || true
exit 1
