package helper

import (
	"errors"
	"strings"
	"time"

	"github.com/gogf/gf/os/gtime"
)

const DateTimeLayout = "2006-01-02 15:04:05"
const DateLayout = "2006-01-02"

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期字符串格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	date, err := time.Parse(timeFormatTpl[0:len(sdate)], sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl[0:len(edate)], edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	date1Str := date.Format(timeFormatTpl)
	d = append(d, date1Str)
	if date1Str == date2Str {
		return d
	}
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// GetDiffDay 计算两个日期差，默认返回天数
// 参数为日期字符串格式，如：2020-01-01
func GetDiffDay(sdate, edate string) int {
	timeFormatTpl := "2006-01-02 15:04:05"
	date, err := time.Parse(timeFormatTpl[0:len(sdate)], sdate)
	if err != nil {
		// 时间解析，异常
		return 0
	}
	date2, err := time.Parse(timeFormatTpl[0:len(edate)], edate)
	if err != nil {
		// 时间解析，异常
		return 0
	}
	return int(date2.Sub(date).Hours() / 24)
}

// GetWeekDate 根据指定时间和1~7之间的数字（周几），获取对应周几的日期
// date 参数为日期字符串格式，如：2020-01-01
// weekday 参数为数字，如：1表示周一
// format 参数为字符串，如：Y-m-d H:i:s
func GetWeekDate(date string, weekday int, format ...string) string {
	if len(format) == 0 {
		format = []string{"Y-m-d"}
	}
	timeFormatTpl := "2006-01-02 15:04:05"
	d, err := time.Parse(timeFormatTpl[0:len(date)], date)
	if err != nil {
		// 时间解析，异常
		return ""
	}
	dateNew := d.AddDate(0, 0, weekday-int(d.Weekday()))
	return gtime.New(dateNew).Format(format[0])
}

// GetDaysBetweenDate 获取两个时间之间年份、月份，日数据
// 参数为日期字符串格式，如：2020-01-01  2020-03-01
// format 格式：Ymd、Ym、Y
// prefix 格式:dwd_admin_
func GetDaysBetweenDate(sdate, edate, format string, prefix ...string) []string {
	startMonth := ""
	endMonth := ""
	month := []string{}
	plusYears := 0
	plusMonths := 0
	plusDays := 0
	prefixStr := ""
	if len(prefix) > 0 {
		prefixStr = prefix[0]
	}
	if strings.Contains(format, "d") {
		plusDays = 1
	} else if strings.Contains(format, "m") {
		plusMonths = 1
	} else if strings.Contains(format, "Y") {
		plusYears = 1
	}
	if sdate != "" {
		startMonth = gtime.NewFromStr(sdate).Format(format)
		month = append(month, prefixStr+startMonth)
	}
	if edate != "" {
		endMonth = gtime.NewFromStr(edate).Format(format)
		if startMonth != endMonth {
			month = append(month, prefixStr+endMonth)
		}
	}
	if startMonth != "" && endMonth != "" {
		for {
			if plusYears == 0 && plusMonths == 0 && plusDays == 0 {
				break
			}
			startMonth = gtime.NewFromStr(sdate).AddDate(plusYears, plusMonths, plusDays).Format(format)
			if startMonth == endMonth {
				break
			}
			month = append(month, prefixStr+startMonth)
		}
	} else {
		//如果时间参数为空，则默认当前时间
		month = append(month, prefixStr+gtime.Now().Format(format))
	}
	return month
}

// IsTimeCross 判断两组时间是否有交叉
func IsTimeCross(t1 [2]time.Time, t2 [2]time.Time) bool {
	if t2[0].Unix() < t1[1].Unix() && t1[0].Unix() < t2[1].Unix() {
		return true
	}
	return false
}

type TimeRange struct {
	StartTime time.Time
	EndTime   time.Time
}

// 判断时间是否在时间交叉属组里面
func IsTimeInTimeRanges(t TimeRange, timeRanges []TimeRange) bool {
	for _, tr := range timeRanges {
		if t.EndTime.After(tr.StartTime) && t.StartTime.Before(tr.EndTime) ||
			(t.StartTime.Before(tr.EndTime) && t.EndTime.After(tr.EndTime)) ||
			(t.StartTime.After(tr.StartTime) && t.EndTime.Before(tr.EndTime)) ||
			(t.StartTime.Before(tr.StartTime) && t.EndTime.After(tr.EndTime)) ||
			(t.StartTime == tr.StartTime && t.EndTime == tr.EndTime) {
			return true
		}
	}
	return false
}

// ParseSliceTime 解析判断一个时间段
func ParseSliceTime(st, et, layout string) (t [2]time.Time, err error) {
	t[0], err = time.Parse(layout, st)
	if err != nil {
		return
	}
	t[1], err = time.Parse(layout, et)
	if err != nil {
		return
	}
	if t[0].Unix() >= t[1].Unix() {
		err = errors.New("时间段选择错误")
		return
	}
	return
}

// ParseTimeCross 判断两组时间段是否有重叠
func ParseTimeCross(st, et, layout string) (func(st1, st2, layout string) (bool, error), error) {
	t, err := ParseSliceTime(st, et, layout)
	if err != nil {
		return nil, err
	}
	return func(st1, st2, layout string) (bool, error) {
		_t, _err := ParseSliceTime(st1, st2, layout)
		if _err != nil {
			return false, _err
		}
		return IsTimeCross(t, _t), nil
	}, nil
}
