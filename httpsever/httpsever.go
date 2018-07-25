package httpsever

import (
	"net/http"
	"sync"
	"database/sql"
	"log"
	"strconv"
	"../dbmanager"
)

var (
	g_tname_uuid     = "t_uuid"
	g_tname_qq_login = "qq_login_"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func queryFuncHandler() http.HandlerFunc {
	var mutex sync.Mutex

	fn := func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()

		db := dbmanager.GetConn("root", "localhost", 3307, "temp", "123456")

		var id sql.NullInt64

		r.ParseForm()
		for k, v := range r.Form {
			log.Println("k:" + k)

			for _, v1 := range v {
				log.Print("," + v1)
				if k == "id" {
					id.Int64, _ = strconv.ParseInt(v1, 10, 64)
					break
				}
			}
		}
		//查询出结果返回
		strId := strconv.FormatInt(id.Int64, 10)
		log.Println("select * from " + g_tname_qq_login + strconv.FormatInt(id.Int64%3+1, 10) + " where id=?" + strId)
		row1 := db.Instance.QueryRow("select * from " + g_tname_qq_login + strconv.FormatInt(id.Int64%3+1, 10) + " where id=?", strId)
		var (
			resultId sql.NullInt64
			resultName sql.NullString
			resultLoginTime sql.NullString
		)

		err3 := row1.Scan(&resultId, &resultName, &resultLoginTime);
		switch  {
		case err3 == sql.ErrNoRows:
			log.Printf("No user with that ID.")
			w.Write([] byte("No user with that ID."))
		case err3 != nil:
			log.Fatal(err3)
		default:
			w.Write([] byte("-----id" + strId + "," + resultName.String + "," + resultLoginTime.String))
		}
	}

	return http.HandlerFunc(fn)
}

func insertFuncHandler() http.HandlerFunc {
	var mutex sync.Mutex

	fn := func(w http.ResponseWriter, r *http.Request) {
		mutex.Lock()
		defer mutex.Unlock()

		db := dbmanager.GetConn("root", "localhost", 3307, "temp", "123456")

		// 插入id
		stmt1, err1 := db.Instance.Prepare("insert into " + g_tname_uuid + "(id) values(?)")
		checkErr(err1)

		if _, err := stmt1.Exec(sql.NullInt64{}); err != nil {
			log.Fatal(err)
			return
		}
		defer stmt1.Close()

		//查询最大的uuid
		var id sql.NullInt64
		row := db.Instance.QueryRow("select max(id) from " + g_tname_uuid)
		if err := row.Scan(&id); err != nil {
			log.Println("query QueryRow g_tname_uuid no id")
			return
		}

		str_id := strconv.Itoa(int(id.Int64))
		log.Println("id is-------------:" + str_id)

		//插入到1-3的login表中
		tempString := "insert into " + g_tname_qq_login + strconv.FormatInt(id.Int64%3+1, 10) + " (id, name, login_time) values (?, ?, ?)"
		stmt2, err2 := db.Instance.Prepare(tempString)
		checkErr(err2)

		if _, err := stmt2.Exec(str_id, "liyujiang" + str_id, "123"); err != nil {
			log.Fatal(err)
			return
		}
		defer stmt2.Close()

		//查询出结果返回
		row1 := db.Instance.QueryRow("select * from " + g_tname_qq_login + strconv.FormatInt(id.Int64%3+1, 10) + " where id=?", str_id)
		var (
			resultId sql.NullInt64
			resultName sql.NullString
			resultLoginTime sql.NullString
		)

		err3 := row1.Scan(&resultId, &resultName, &resultLoginTime)
		switch  {
		case err3 == sql.ErrNoRows:
			log.Printf("No user with that ID.")
		case err3 != nil:
			log.Fatal(err3)
		default:
		}

		w.Write([] byte("-----id" + strconv.FormatInt(resultId.Int64, 10) + "," + resultName.String + "," + resultLoginTime.String))
	}

	return http.HandlerFunc(fn)
}

//实现数据库的分表
func StartHttpServer() {
	mux := http.NewServeMux()

	insertFn := insertFuncHandler()
	mux.Handle("/insert", insertFn)

	queryFn := queryFuncHandler()
	mux.Handle("/query", queryFn)

	log.Println("start...")
	http.ListenAndServe(":8888", mux)
}
