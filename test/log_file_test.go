package test

import (
	"log"
	"os"
	"testing"
)

func TestLogFile(t *testing.T) {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		t.Fatal("存储日志文件失败:" + err.Error())
	}

	log.SetOutput(file)

	defer func() {
		err := file.Close()
		if err != nil {
			t.Fatal("关闭文件流失败:" + err.Error())
		}
	}()

	t.Log("存储log日志成功")
}
