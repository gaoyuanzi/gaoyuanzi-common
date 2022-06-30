package pgdb

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

func OpenConnect(ip string, port string, user string, passwd string, database string) (*sql.DB, error) {
	tport, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf(fmt.Sprint(err))
		return nil, err
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		ip, tport, user, passwd, database)
	connsystemdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf(fmt.Sprint(err))
	}
	// defer connsystemdb.Close()
	err = connsystemdb.Ping()
	if err != nil {
		return nil, err
	}
	return connsystemdb, nil
}

func DBQry(db *sql.DB, sqltext string) ([]interface{}, error) {

	retlst := make([]interface{}, 0)

	query, err := db.Query(sqltext)
	if err != nil {
		return retlst, err
	}
	defer query.Close()

	//读出查询出的列字段名
	cols, _ := query.Columns()
	//values是每个列的值，这里获取到byte里
	values := make([][]byte, len(cols))
	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans := make([]interface{}, len(cols))
	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}
	for query.Next() { //循环，让游标往下推
		//query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都>放在values里
		if err := query.Scan(scans...); err != nil {
			return retlst, err
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := cols[k]
			row[key] = string(v)
		}
		retlst = append(retlst, row)
	}
	return retlst, nil
}

func DBExec(db *sql.DB, sqltext string) ([]interface{}, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(sqltext)
	if err != nil {
		return nil, errors.New("sql exec err: " + sqltext + "," + fmt.Sprint(err))
	}
	return nil, nil
}
