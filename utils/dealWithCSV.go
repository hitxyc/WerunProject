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

func DealWithCSV(filename string, is_graduate bool) error {
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
					if is_graduate {
						transToGraduate(&informationTable, &contents, studentChan, &wq)
					} else {
						transToUndergraduate(&informationTable, &contents, studentChan, &wq)
					}
				}()
			}
		}
		// 所有协程工作完毕后关闭通道
		wq.Wait()
		close(studentChan)
	}()
	// 将数据存入数据库中
	addToDatabase(studentChan)
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

// transToUndergraduate 将csv文件中的信息绑定在entity.Undergraduate结构体上
func transToUndergraduate(informationTable *[]string, contents *[]string, studentChan chan<- entity.Student, wq *sync.WaitGroup) {
	var us = &entity.UndergraduateStudent{}
	// 本科生
	us.Score = make(map[string]float64)
	for index, value := range *informationTable {
		switch value {
		case "学号":
			us.StudentId = (*contents)[index] // 绑定学生学号
		case "考号":
			us.StudentId = (*contents)[index] // 绑定学生考号
		case "姓名":
			us.Name = (*contents)[index] // 绑定学生姓名
		case "性别":
			us.Gender = (*contents)[index] // 绑定学生性别
		case "班级":
			us.Class = (*contents)[index] // 绑定学生班级
		default:
			score, _ := strconv.ParseFloat((*contents)[index], 64)
			us.Score[(*informationTable)[index]] = score
		}
	}
	if us.StudentId != "" {
		studentChan <- us
	}
	wq.Done()
}

func transToGraduate(informationTable *[]string, contents *[]string, studentChan chan<- entity.Student, wq *sync.WaitGroup) {
	var gs = &entity.GraduateStudent{}
	// 研究生
	gs.Score = make(map[string]float64)
	for index, value := range *informationTable {
		switch value {
		case "学号":
			gs.StudentId = (*contents)[index] // 绑定学生学号
		case "考号":
			gs.StudentId = (*contents)[index] // 绑定学生考号
		case "姓名":
			gs.Name = (*contents)[index] // 绑定学生姓名
		case "性别":
			gs.Gender = (*contents)[index] // 绑定学生性别
		case "导师":
			gs.Tutor = (*contents)[index] // 绑定学生班级
		default:
			score, _ := strconv.ParseFloat((*contents)[index], 64)
			gs.Score[(*informationTable)[index]] = score
		}
	}
	if gs.StudentId != "" {
		studentChan <- gs
	}
	wq.Done()
}

// addToDatabase 将学生信息传入数据库中
func addToDatabase(studentChan <-chan entity.Student) {
	var sm mapper.StudentMapper
	for student := range studentChan {
		lock.Lock()
		err := sm.SaveStudent(&student)
		if err != nil {
			log.Println(err)
		}
		lock.Unlock()
	}
}
