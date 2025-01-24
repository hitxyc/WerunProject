package utils

import (
	"encoding/csv"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"studentManagementSystem/entity"
	"studentManagementSystem/mapper"
	"sync"
)

var lock sync.Mutex

func DealWithCSV(filename string) error {
	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// 读取csv文件
	reader := csv.NewReader(file)
	flag := true                  // 标记第一行
	var informationTable []string // 表头信息
	studentChan := make(chan entity.Student, 100)
	var wq sync.WaitGroup
	// 进行文件的批量保存
	go func() {
		for {
			records, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
				continue
			}
			decoder := simplifiedchinese.GBK.NewDecoder()
			for _, record := range records {
				// 处理第一行表头信息
				if flag {
					// 转为GBK编码, 将一行内容分割
					contents := transToGBK(&record, decoder)
					// 提取第一行信息
					informationTable = getInformation(&contents)
					flag = false
					continue
				}
				go func() {
					wq.Add(1)
					// 转为GBK编码, 将一行内容分割
					contents := transToGBK(&record, decoder)
					// 将csv文件中的内容转为entity.Student
					transToStudent(&informationTable, &contents, studentChan, &wq)
				}()
			}
		}
		// 所有协程工作完毕后关闭通道
		wq.Wait()
		close(studentChan)
	}()
	// 将数据存入数据库中
	AddToDatabase(studentChan)
	return nil
}

// transToGBK 转码函数, 并将一行内容分割
func transToGBK(record *string, decoder *encoding.Decoder) []string {
	converted_reader := transform.NewReader(strings.NewReader(*record), decoder)
	content, _ := ioutil.ReadAll(converted_reader)
	contents := strings.Split(string(content), "\t")
	return contents
}

// getInformation 从第一行提取信息, 主要提取考试科目信息
func getInformation(content *[]string) []string {
	informationTable := make([]string, len(*content))
	for i := 0; i < len(*content); i++ {
		informationTable[i] = (*content)[i]
	}
	return informationTable
}

// transToStudent 将csv文件中的信息绑定在entity.Student结构体上
func transToStudent(informationTable *[]string, contents *[]string, studentChan chan<- entity.Student, wq *sync.WaitGroup) {
	var student entity.Student
	student.Score = make(map[string]float64)
	for index, value := range *informationTable {
		switch value {
		case "考号":
			student.StudentId = (*contents)[index] // 绑定学生学号
		case "姓名":
			student.Name = (*contents)[index] // 绑定学生姓名
		case "性别":
			student.Gender = (*contents)[index] // 绑定学生性别
		case "班级":
			student.Class = (*contents)[index] // 绑定学生班级
		default:
			score, _ := strconv.ParseFloat((*contents)[index], 64)
			student.Score[(*informationTable)[index]] = score
		}
	}
	studentChan <- student
	wq.Done()
}

// AddToDatabase 将学生信息传入数据库中
func AddToDatabase(studentChan <-chan entity.Student) {
	var sm mapper.StudentMapper
	for student := range studentChan {
		lock.Lock()
		sm.SaveStudent(&student)
		lock.Unlock()
	}
}
