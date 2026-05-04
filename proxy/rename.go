package proxies

import (
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/biter777/countries"
)

var (
	counter     = make(map[string]int)
	counterLock = sync.Mutex{}
)

// emojiToCode 将国旗emoji映射为两字母国家代码
var emojiToCode = map[string]string{
	"🇭🇰": "HK", "🇹🇼": "TW", "🇯🇵": "JP", "🇸🇬": "SG", "🇰🇷": "KR",
	"🇺🇸": "US", "🇬🇧": "GB", "🇩🇪": "DE", "🇫🇷": "FR", "🇦🇺": "AU",
	"🇨🇦": "CA", "🇷🇺": "RU", "🇧🇷": "BR", "🇮🇳": "IN", "🇳🇱": "NL",
	"🇸🇪": "SE", "🇮🇹": "IT", "🇪🇸": "ES", "🇻🇳": "VN", "🇲🇾": "MY",
	"🇹🇭": "TH", "🇮🇩": "ID", "🇵🇭": "PH", "🇹🇷": "TR", "🇦🇪": "AE",
	"🇸🇦": "SA", "🇿🇦": "ZA", "🇦🇷": "AR", "🇲🇽": "MX", "🇪🇬": "EG",
	"🇳🇬": "NG", "🇵🇰": "PK", "🇧🇩": "BD", "🇺🇦": "UA", "🇵🇱": "PL",
	"🇨🇿": "CZ", "🇭🇺": "HU", "🇦🇹": "AT", "🇨🇭": "CH", "🇧🇪": "BE",
	"🇵🇹": "PT", "🇬🇷": "GR", "🇳🇴": "NO", "🇩🇰": "DK", "🇫🇮": "FI",
	"🇮🇪": "IE", "🇳🇿": "NZ", "🇵🇪": "PE", "🇨🇱": "CL", "🇨🇴": "CO",
}

// CountryCodeToChinese 把国家代码转为中文名称
func CountryCodeToChinese(code string) string {
	code = strings.ToUpper(strings.TrimSpace(code))

	// 扩展后的中文映射表（包含更多国家和地区）
	chineseMap := map[string]string{
		"HK": "香港", "TW": "台湾", "JP": "日本", "SG": "新加坡", "KR": "韩国",
		"US": "美国", "GB": "英国", "DE": "德国", "FR": "法国", "AU": "澳大利亚",
		"CA": "加拿大", "RU": "俄罗斯", "BR": "巴西", "IN": "印度", "NL": "荷兰",
		"SE": "瑞典", "IT": "意大利", "ES": "西班牙", "VN": "越南", "MY": "马来西亚",
		"TH": "泰国", "ID": "印度尼西亚", "PH": "菲律宾", "TR": "土耳其", "AE": "阿联酋",
		"SA": "沙特阿拉伯", "ZA": "南非", "AR": "阿根廷", "MX": "墨西哥", "EG": "埃及",
		"NG": "尼日利亚", "PK": "巴基斯坦", "BD": "孟加拉国", "UA": "乌克兰", "PL": "波兰",
		"CZ": "捷克", "HU": "匈牙利", "AT": "奥地利", "CH": "瑞士", "BE": "比利时",
		"PT": "葡萄牙", "GR": "希腊", "NO": "挪威", "DK": "丹麦", "FI": "芬兰",
		"IE": "爱尔兰", "NZ": "新西兰", "PE": "秘鲁", "CL": "智利", "CO": "哥伦比亚",
		// 可继续添加其他常见国家代码
	}

	if ch, ok := chineseMap[code]; ok {
		return ch
	}

	// fallback: 使用 countries 库获取中文名（需要确认库是否支持中文）
	country := countries.ByName(code)
	if country != countries.Unknown {
		// 返回英文名（如果库不支持中文，也可以保持英文）
		return country.String()
	}
	return "其他"
}

// extractCountryCode 从节点名中提取国家代码，支持：
// 1. 国旗 emoji（例如 "🇭🇰" → "HK"）
// 2. 常见两字母代码（边界匹配，避免误匹配）
// 3. 国家中文名（如“香港” → “HK”）
func extractCountryCode(name string) string {
	// 1. 尝试从 emoji 提取
	for emoji, code := range emojiToCode {
		if strings.Contains(name, emoji) {
			return code
		}
	}

	// 2. 尝试匹配两字母国家代码（单词边界）
	// 常用的代码列表（按优先级排序）
	codes := []string{
		"HK", "TW", "JP", "SG", "KR", "US", "GB", "DE", "FR", "AU",
		"CA", "RU", "BR", "IN", "NL", "SE", "IT", "ES", "VN", "MY",
		"TH", "ID", "PH", "TR", "AE", "SA", "ZA", "AR", "MX", "EG",
		"NG", "PK", "BD", "UA", "PL", "CZ", "HU", "AT", "CH", "BE",
		"PT", "GR", "NO", "DK", "FI", "IE", "NZ", "PE", "CL", "CO",
	}
	// 构造正则：匹配独立的大写两字母（前后非字母或字符串边界）
	// 注意：节点名中可能包含非字母分隔符如 _ - 空格等，使用 \b 或自定义边界
	// 为简单起见，用正则 \b[A-Z]{2}\b 匹配完整单词
	// 但需避免 "US" 出现在 "AUS" 中，所以使用边界匹配
	for _, code := range codes {
		// 使用正则匹配独立的两字母代码
		re := regexp.MustCompile(`\b` + code + `\b`)
		if re.MatchString(strings.ToUpper(name)) {
			return code
		}
	}

	// 3. 尝试匹配中文名称（反向映射）
	chineseToCode := map[string]string{
		"香港": "HK", "台湾": "TW", "日本": "JP", "新加坡": "SG", "韩国": "KR",
		"美国": "US", "英国": "GB", "德国": "DE", "法国": "FR", "澳大利亚": "AU",
		"加拿大": "CA", "俄罗斯": "RU", "巴西": "BR", "印度": "IN", "荷兰": "NL",
		"瑞典": "SE", "意大利": "IT", "西班牙": "ES", "越南": "VN", "马来西亚": "MY",
		"泰国": "TH", "印度尼西亚": "ID", "菲律宾": "PH", "土耳其": "TR", "阿联酋": "AE",
		"沙特阿拉伯": "SA", "南非": "ZA", "阿根廷": "AR", "墨西哥": "MX", "埃及": "EG",
	}
	for ch, code := range chineseToCode {
		if strings.Contains(name, ch) {
			return code
		}
	}

	return ""
}

// Rename 主重命名函数（保持不变，但会受益于更准确的 extractCountryCode）
func Rename(name string) string {
	counterLock.Lock()
	defer counterLock.Unlock()

	counter[name]++

	countryCode := extractCountryCode(name)
	chineseCountry := CountryCodeToChinese(countryCode)

	// 最终格式：香港_原始名_序号（可以自定义）
	return chineseCountry + "_" + name + "_" + strconv.Itoa(counter[name])
}

// ResetRenameCounter 重置计数器
func ResetRenameCounter() {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter = make(map[string]int)
}

// CountryCodeToFlag 返回国家代码对应的 emoji 标志（保留备用）
func CountryCodeToFlag(countryCode string) string {
	code := strings.ToUpper(countryCode)
	country := countries.ByName(code)
	if country == countries.Unknown {
		return "❓Other"
	}
	return country.Emoji()
}