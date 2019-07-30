package utils

import "time"

const (
	FormatTime     = "15:04:05"            //时间格式
	FormatDate     = "2006-01-02"          //日期格式
	FormatDateTime = "2006-01-02 15:04:05" //完整时间格式
	FormatDateTime2 = "2006-01-02 15:04" //完整时间格式
)

func GetToday(format string) string {
	today := time.Now().Format(format)
	return today
}

//获取今天剩余秒数
func GetTodayLastSecond() time.Duration {
	today := GetToday(FormatDate) + " 23:59:59"
	end, _ := time.ParseInLocation(FormatDateTime, today, time.Local)
	return time.Duration(end.Unix()-time.Now().Local().Unix()) * time.Second
}

// 处理出生日期函数
func GetBrithDate(idcard string) string {
	l := len(idcard)
	var s string
	if l == 15 {
		s = "19" + idcard[6:8] + "-" + idcard[8:10] + "-" + idcard[10:12]
		return s
	}
	if l == 18 {
		s = idcard[6:10] + "-" + idcard[10:12] + "-" + idcard[12:14]
		return s
	}
	return GetToday(FormatDate)
}

//获取相差时间-秒
func GetSecondDifferByTime(start_time, end_time time.Time) int64 {
	diff := end_time.Unix() - start_time.Unix()
	return diff
}

//获取相差时间-天数
func GetDayDiffer(start_time, end_time time.Time) int {
	t1, _ := time.Parse("2006-01-02", start_time.Format("2006-01-02"))
	t2, _ := time.Parse("2006-01-02", end_time.Format("2006-01-02"))
	hours := int(t2.Sub(t1).Hours())

	return hours / 24
}

func GetNowDayTime(paramTime time.Time) (has bool) {
	nowTime := time.Now().Format("2006-01-02")
	nowDay, _ := time.ParseInLocation("2006-01-02", nowTime, time.Local)
	has = true
	if nowDay.After(paramTime) {
		has = false
	}
	return
}

func GetYearDiffer(start_time, end_time string) int {
	t1, _ := time.ParseInLocation("2006-01-02", start_time, time.Local)
	t2, _ := time.ParseInLocation("2006-01-02", end_time, time.Local)
	age := t2.Year() - t1.Year()
	if t2.Month() < t1.Month() || (t2.Month() == t1.Month() && t2.Day() < t1.Day()) {
		age--
	}

	return age

}
