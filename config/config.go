package config

import _ "embed"

type Config struct {
    // 测速通过 SpeedTestUrl="" 控制，媒体通过 MediaCheck: false 控制
    // 新增：低配设备专用选项
    DisableWebUI        bool   `yaml:"disable-web-ui"`         // 彻底关闭 Web UI（节省内存）
    DisableSubStore     bool   `yaml:"disable-sub-store"`      // 彻底关闭 sub-store 服务（省内存/CPU）
    MemoryOptimize      bool   `yaml:"memory-optimize"`        // 开启后每轮检测完强制GC
    MaxProxiesPerCheck  int    `yaml:"max-proxies-per-check"`  // 每轮最多检测N个节点，0=不限（低配限流用）
	
	PrintProgress        bool     `yaml:"print-progress"`
	Concurrent           int      `yaml:"concurrent"`
	SpeedConcurrent      int      `yaml:"speed-concurrent"`
	MediaConcurrent      int      `yaml:"media-concurrent"`
	CheckInterval        int      `yaml:"check-interval"`
	CronExpression       string   `yaml:"cron-expression"`
	AliveTestUrl         string   `yaml:"alive-test-url"`
	SpeedTestUrl         string   `yaml:"speed-test-url"`
	DownloadTimeout      int      `yaml:"download-timeout"`
	DownloadMB           int      `yaml:"download-mb"`
	TotalSpeedLimit      int      `yaml:"total-speed-limit"`
	MinSpeed             int      `yaml:"min-speed"`
	Timeout              int      `yaml:"timeout"`
	MediaCheckTimeout    int      `yaml:"media-check-timeout"`
	FilterRegex          string   `yaml:"filter-regex"`
	SaveMethod           string   `yaml:"save-method"`
	WebDAVURL            string   `yaml:"webdav-url"`
	WebDAVUsername       string   `yaml:"webdav-username"`
	WebDAVPassword       string   `yaml:"webdav-password"`
	GithubToken          string   `yaml:"github-token"`
	GithubGistID         string   `yaml:"github-gist-id"`
	GithubAPIMirror      string   `yaml:"github-api-mirror"`
	WorkerURL            string   `yaml:"worker-url"`
	WorkerToken          string   `yaml:"worker-token"`
	S3Endpoint           string   `yaml:"s3-endpoint"`
	S3AccessID           string   `yaml:"s3-access-id"`
	S3SecretKey          string   `yaml:"s3-secret-key"`
	S3Bucket             string   `yaml:"s3-bucket"`
	S3UseSSL             bool     `yaml:"s3-use-ssl"`
	S3BucketLookup       string   `yaml:"s3-bucket-lookup"`
	SubUrlsReTry         int      `yaml:"sub-urls-retry"`
	SubUrlsRetryInterval int      `yaml:"sub-urls-retry-interval"`
	SubUrlsTimeout       int      `yaml:"sub-urls-timeout"`
	SubUrlsConcurrent    int      `yaml:"sub-urls-concurrent"`
	SubUrlsGetUA         string   `yaml:"sub-urls-get-ua"`
	SubUrlsRemote        []string `yaml:"sub-urls-remote"`
	SubUrls              []string `yaml:"sub-urls"`
	SuccessRate          float32  `yaml:"success-rate"`
	MihomoApiUrl         string   `yaml:"mihomo-api-url"`
	MihomoApiSecret      string   `yaml:"mihomo-api-secret"`
	ListenPort           string   `yaml:"listen-port"`
	RenameNode           bool     `yaml:"rename-node"`
	OutputDir            string   `yaml:"output-dir"`
	AppriseApiServer     string   `yaml:"apprise-api-server"`
	RecipientUrl         []string `yaml:"recipient-url"`
	NotifyTitle          string   `yaml:"notify-title"`
	SubStorePort         string   `yaml:"sub-store-port"`
	SubStorePath         string   `yaml:"sub-store-path"`
	SubStoreSyncCron     string   `yaml:"sub-store-sync-cron"`
	SubStorePushService  string   `yaml:"sub-store-push-service"`
	SubStoreProduceCron  string   `yaml:"sub-store-produce-cron"`
	MihomoOverwriteUrl   string   `yaml:"mihomo-overwrite-url"`
	MediaCheck           bool     `yaml:"media-check"`
	Platforms            []string `yaml:"platforms"`
	SuccessLimit         int32    `yaml:"success-limit"`
	NodePrefix           string   `yaml:"node-prefix"`
	NodeType             []string `yaml:"node-type"`
	EnableWebUI          bool     `yaml:"enable-web-ui"`
	APIKey               string   `yaml:"api-key"`
	GithubProxy          string   `yaml:"github-proxy"`
	Proxy                string   `yaml:"proxy"`
	CallbackScript       string   `yaml:"callback-script"`
	Filter               []string `yaml:"filter"`
	KeepDays             int      `yaml:"keep-days"`
	DNS                  DNSConfig `yaml:"dns"`
}

// DNSConfig controls mihomo's global resolver used by every proxy probe.
// Leaving Enable=false keeps the historical behavior (mihomo SystemResolver, v4 only).
type DNSConfig struct {
	Enable                bool     `yaml:"enable"`
	IPv6                  bool     `yaml:"ipv6"`
	Nameserver            []string `yaml:"nameserver"`
	ProxyServerNameserver []string `yaml:"proxy-server-nameserver"`
	DefaultNameserver     []string `yaml:"default-nameserver"`
}

var GlobalConfig = &Config{
	// 新增配置，给未更改配置文件的用户一个默认值
	ListenPort:         ":8199",
	NotifyTitle:        "🔔 节点状态更新",
	MihomoOverwriteUrl: "http://127.0.0.1:8199/sub/ACL4SSR_Online_Full.yaml",
	MediaCheckTimeout:  10,
	Platforms:          []string{"openai", "youtube", "netflix", "disney", "gemini", "iprisk"},
	DownloadMB:         20,
	AliveTestUrl:       "http://gstatic.com/generate_204",
	SubUrlsGetUA:       "clash.meta (https://github.com/yu-day/my-subs-check)",
	SubUrlsReTry:       3,
	SubUrlsConcurrent:  20,
	
	//新增 低配默认值
    MemoryOptimize:     false,
    MaxProxiesPerCheck: 0,
}

//go:embed config.example.yaml
var DefaultConfigTemplate []byte

var GlobalProxies []map[string]any
