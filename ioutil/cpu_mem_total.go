package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"strconv"
	"strings"

	"github.com/toolkits/file"
)

func main() {

	//parseLine(bs)

	//	reader := bufio.NewReader(bytes.NewBuffer(bs))

	//	for {

	//		line, err := file.ReadLine(reader)
	//		if err == io.EOF {
	//			err = nil
	//			break
	//		} else if err != nil {
	//			return
	//		}
	//		parseLine(line)
	//	}
	cpu_num, mem_num := count_cpu_mem()
	fmt.Printf("CPU Total: %v and MEM Total: %v", cpu_num, mem_num)
}

func count_cpu_mem() (cpu_num string, mem_num string) {
	cpu_file := "E:/workspace/yh/OpenBridge-passos-proxy/open-faclon/src/go-program/ioutil/cpuinfo"
	cpu_bs, cpu_err := ioutil.ReadFile(cpu_file)
	if cpu_err != nil {
		return
	}
	fmt.Println("CpuTotal:" + fmt.Sprintf("%d", strings.Count(string(cpu_bs), string("physical id"))))

	mem_file := "E:/workspace/yh/OpenBridge-passos-proxy/open-faclon/src/go-program/ioutil/meminfo"
	mem_bs, mem_err := ioutil.ReadFile(mem_file)
	if mem_err != nil {
		return
	}
	reader := bufio.NewReader(bytes.NewBuffer(mem_bs))
	for {
		line, err := file.ReadLine(reader)
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return
		}

		fields := strings.Fields(string(line))
		fieldName := fields[0]

		if len(fields) == 3 {
			val, numerr := strconv.ParseUint(fields[1], 10, 64)
			if numerr != nil {
				continue
			}
			if fieldName == "MemTotal:" {
				fmt.Println("MemTotal:" + fmt.Sprintf("%d", val/1024) + "M")
				mem_num = fmt.Sprintf("%v", val/1024)
			}

		}
	}
	cpu_num = fmt.Sprintf("%v", strings.Count(string(cpu_bs), string("physical id")))
	return cpu_num, mem_num
}
