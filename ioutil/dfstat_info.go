package main

import (
	"bufio"
	"bytes"
	"github.com/toolkits/file"
	"io"
	"io/ioutil"
	"strings"
	"fmt"
)

func main(){
	ListMountPoint();
}

// return: [][$fs_spec, $fs_file, $fs_vfstype]
func ListMountPoint()  ([][3]string, error) {
	contents, _ := ioutil.ReadFile("./mounts")
	//contents, _ := ioutil.ReadFile("E:/workspace/yh/OpenBridge-passos-proxy/open-faclon/src/go-program/ioutil/mounts")

	ret := make([][3]string, 0)

	reader := bufio.NewReader(bytes.NewBuffer(contents))
	for {
		line, err := file.ReadLine(reader)
		fmt.Printf("%v \n",string(line))
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			return nil, err
		}

		fields := strings.Fields(string(line))
		// Docs come from the fstab(5)
		// fs_spec     # Mounted block special device or remote filesystem e.g. /dev/sda1
		// fs_file     # Mount point e.g. /data
		// fs_vfstype  # File system type e.g. ext4
		// fs_mntops   # Mount options
		// fs_freq     # Dump(8) utility flags
		// fs_passno   # Order in which filesystem checks are done at reboot time

		fs_spec := fields[0]
		fs_file := fields[1]
		fs_vfstype := fields[2]

		if strings.Contains(fs_file,"kubelet"){
			continue
		}

		// keep /dev/xxx device with shorter fs_file (remove mount binds)
		if strings.HasPrefix(fs_spec, "/dev") {
			deviceFound := false
			for idx := range ret {
				if ret[idx][0] == fs_spec {
					deviceFound = true
					if len(fs_file) < len(ret[idx][1]) {
						ret[idx][1] = fs_file
					}
					break
				}
			}
			if !deviceFound {
				ret = append(ret, [3]string{fs_spec, fs_file, fs_vfstype})
			}
		} else {
			ret = append(ret, [3]string{fs_spec, fs_file, fs_vfstype})
		}
	}
	fmt.Println(ret)
	return ret, nil
}
