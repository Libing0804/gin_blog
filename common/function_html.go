package common

import "time"

func IsODD(num int) bool {
	return num%2 == 0
}
func GetNextName(strs []string, index int) string {
	return strs[index+1] //html中的意思是去获得导航条的路径

}

//格式化时间
func Date(layout string) string {
	return time.Now().Format(layout)

}
func DateDay(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}
