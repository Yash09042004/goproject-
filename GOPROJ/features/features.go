package features

import (
	"bufio"
	"fmt"
	"sync"

	//	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name string
	PRN string
	Marks int

}
// capital names import in main automatically 
type Class struct {
	Students []Student
}



//c *class ---> is associates with Class attributes 
func(c *Class) NewClass(){


	mystudents := []Student{}
	file,err:=os.Open("db.txt")
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line," ")
		
		marks,err:=strconv.Atoi(fields[2])
		if err!=nil{
			panic(err)
		}
		obj := Student{
			Name: fields[0],
			PRN: fields[1],
			Marks: marks,
		}
		mystudents = append(mystudents,obj)
	}

	// fmt.Println(mystudents)

	c.Students = mystudents


}

func (c *Class) Writetofile(wg *sync.WaitGroup){

	defer wg.Done()
	var data []string
	for _,v := range c.Students{
		line := fmt.Sprintf("%v %v %v",v.Name,v.PRN,v.Marks)
		fmt.Println(line)
		data = append(data,line)
	}
	err:=os.WriteFile("db.txt",[]byte(strings.Join(data,"\n")),0664)
	if err!=nil{
		panic(err)
	}
	fmt.Println("data written to file ")
}

func (c *Class) ShowStudents(){
	for _,student := range c.Students{
		println(student.Name,student.PRN,student.Marks)
	}
}

func (c *Class) AddStudent(){
	var wg sync.WaitGroup

	fmt.Println("enter name : ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("enter prn : ")
	var prn string
	fmt.Scanln(&prn)
	fmt.Println("enter marks : ")
	var marks int
	fmt.Scanln(&marks)

	obj := Student{
		Name: name,
		PRN: prn,
		Marks: marks,

	}

	c.Students = append(c.Students,obj)
	wg.Add(1)
	go c.Writetofile(&wg)
	wg.Wait()
	//writing data to the file 
}

func (c *Class) UpdateStudent(){
	fmt.Println("enter prn of student to update : ")
	var prn string
	fmt.Scanln(&prn)

	for i,val:=range c.Students{
		// name := val.Name
		Prn := val.PRN
				if prn == Prn{
					fmt.Println("enter name : ")
			var name string
			fmt.Scanln(&name)
			fmt.Println("enter prn : ")
			var PRN string
			fmt.Scanln(&prn)
			fmt.Println("enter marks : ")
			var marks int
			fmt.Scanln(&marks)

			c.Students[i] = Student{Name:name,PRN:PRN,Marks:marks}

		}

	}
	fmt.Println("data updated ")
}

func (c *Class) DeleteStudent(){
	var wg sync.WaitGroup
	fmt.Println("enter PRN: ")
	var prn string 
	fmt.Scanln(&prn )

	var idx int 
	for i, val := range c.Students {
		if val.PRN == prn {
			idx = i
			break
		}
	}
	wg.Add(1)
	go c.Writetofile(&wg)
	wg.Wait()
	tempARR := c.Students
	tempARR = append(tempARR[:idx],tempARR[idx+1:]...)
	c.Students = tempARR
	fmt.Println("data deleted ")

}

func (c *Class) GetStats(){
	//sum
	sum:=c.sum()
	//avg
	avg:=c.avg()
	//min
	min:=c.mini()
	//max
	max:=c.maxi()

	fmt.Printf("Sum: %v\n Avg: %v\n Min: %v\n Max: %v\n",sum,avg,min,max)
	
}

func (c *Class) sum() int {
	TotalSu := 0

	for _, val := range c.Students {
		TotalSu += val.Marks
	}

	return TotalSu
}
func (c *Class) mini() int {
	mini := 1000

	for _, val := range c.Students {
		mini = min(mini, val.Marks)
	}

	return mini
}
func (c *Class) maxi() int {
	maxi := 0

	for _, val := range c.Students {
		maxi = max(maxi, val.Marks)
	}

	return maxi
}



func (c *Class) avg() float64 {
	TotalSu := 0
	TotalSu = c.sum()
	return float64(TotalSu) / float64(len(c.Students))
}