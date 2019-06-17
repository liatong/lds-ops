package model 
import (
	"strconv"
	"fmt"
)


type SQLQuery struct{
	TableName string
	SQLField string
	sqlLimit string 
	sqlLimitStatus bool
	sqlCondition string
	sqlConditionStats bool 
}

func (s *SQLQuery)SetWhere(fliedname string,fliedvalue string){
	s.sqlConditionStats = true
	spacer := " "
	if s.sqlCondition != "" {
		//当多个查询条件时，默认多个条件时“and" 关系。可有优化为自己设定关联条件。
		spacer = " and "
	}
	//s.sqlCondition = s.sqlCondition + " " + filename + '=' + '"' + filevalue + '"'
	s.sqlCondition = s.sqlCondition + spacer + fliedname + "=" + "\"" + fliedvalue + "\""
}
func (s *SQLQuery)SetLimit(page int, pagesize int){
	s.sqlLimitStatus = true
	var start int 
	if page <= 1  {
		start = 0 
	} else {
	    start = (page -1 ) * pagesize
	}
	s.sqlLimit = " limit "+strconv.Itoa(start)+","+strconv.Itoa(pagesize)

}
func (s *SQLQuery)GetQuery()(sqlScript string){

	sqlScript = "select " + s.SQLField + " from " + s.TableName 
	if s.sqlConditionStats {
		sqlScript = sqlScript + " where " + s.sqlCondition
	}
	if s.sqlLimitStatus {
		sqlScript = sqlScript + " " + s.sqlLimit
	}
	fmt.Print(sqlScript)
	return sqlScript
}