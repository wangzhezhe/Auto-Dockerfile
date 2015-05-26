package controllers

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"github.com/gorilla/websocket"
)

func Filecompress(tw *tar.Writer, dir string, fi os.FileInfo) {

	//打开文件 open当中是 目录名称/文件名称 构成的组合
	fmt.Println(dir + "/" + fi.Name())
	fr, err := os.Open(dir + "/" + fi.Name())
	fmt.Println(fr.Name())
	if err != nil {
		panic(err)
	}
	defer fr.Close()

	hdr, err := tar.FileInfoHeader(fi, "")

	hdr.Name = fr.Name()
	if err = tw.WriteHeader(hdr); err != nil {
		panic(err)
	}

	_, err = io.Copy(tw, fr)
	if err != nil {
		panic(err)
	}
	//打印文件名称
	fmt.Println("add the file: " + fi.Name())

}

func Dirtotar(sourcedir string, tardir string) {
	//file write 在tardir目录下创建
	fw, err := os.Create(tardir + "/" + "deployments.tar.gz")
	//type of fw is *os.File
	//fmt.Println(reflect.TypeOf(fw))
	if err != nil {
		panic(err)

	}
	defer fw.Close()

	//gzip writer
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	//tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()
	//fmt.Println(reflect.TypeOf(tw))
	//add the deployments contens
	//Dircompress(tw, "deployments/")

	//write into the dockerfile
	filepath := sourcedir + "/" + "Dockerfile"
	fmt.Println(filepath)
	fileinfo, err := os.Stat(filepath)

	if err != nil {
		panic(err)

	}

	//fmt.Println(reflect.TypeOf(os.FileInfo(fileinfo)))
	//dockerfile要单独放在根目录下 和其他archivefile并列
	//tardir and fileinfo decide the position of file which compress into the tar
	Filecompress(tw, tardir, fileinfo)
	fmt.Println("tar.gz packaging OK")

}

//return a tar stream
func SourceTar(filename string) *os.File {
	//"tardir/deployments.tar.gz"
	fw, _ := os.Open(filename)
	//fmt.Println(reflect.TypeOf(fw))
	return fw

}

//递归删除文件夹
func Cleandir(dirname string) {

	//打开文件夹
	dirhandle, err := os.Open(dirname)
	//fmt.Println(dirname)
	//fmt.Println(reflect.TypeOf(dir))
	if err != nil {
		panic(nil)
	}
	defer dirhandle.Close()

	//fis, err := ioutil.ReadDir(dir)
	fis, err := dirhandle.Readdir(0)
	//fis的类型为 []os.FileInfo
	//fmt.Println(reflect.TypeOf(fis))
	if err != nil {
		panic(err)
	}

	//遍历文件列表 每一个文件到要写入一个新的*tar.Header
	//var fi os.FileInfo
	for _, fi := range fis {
		if fi.IsDir() {
			newname := dirname + "/" + fi.Name()
			//fmt.Println("using dir")
			//fmt.Println(newname)
			//这个样直接continue就将所有文件写入到了一起 没有层级结构了
			//Filecompress(tw, dir, fi)
			Cleandir(newname)

		} else {
			//如果是普通文件 直接写入 dir 后面已经有了 /
			filename := dirname + "/" + fi.Name()
			fmt.Println(filename)
			err := os.Remove(filename)
			if err != nil {
				panic(err)
			}
			fmt.Println("delete " + filename)
		}

	}
	//递归结束 删除当前文件夹
	err = os.Remove(dirname)
	fmt.Println("delete " + dirname)
	if err != nil {
		panic(err)
	}

}

func Tartoimage(imagename string, uploadtar string) *http.Response {

	//dockerhub的认证信息
	auth := AuthConfiguration{
	//	Username:      "wangzhe",
	//	Password:      "3.1415",
	//	Email:         "w_hessen@126.com",
	//	ServerAddress: "https://10.211.55.5",
	}

	client := &http.Client{}
	url := "http://0.0.0.0:2376/build?dockerfile=" + imagename + "/Dockerfile&q=true&t=" + imagename
	body, err := ioutil.ReadFile(imagename + "/deployments.tar.gz")
	if err != nil {
		panic(err)
	}
	reqest, err := http.NewRequest("POST", url, strings.NewReader(string(body)))
	if err != nil {
		panic(err)
	}

	reqest.Header.Set("Content-Type", "application/tar")
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(auth)
	reqest.Header.Set("X-Registry-Config", base64.URLEncoding.EncodeToString(buf.Bytes()))
	response, _ := client.Do(reqest)
	return response

}

var count = 0

func Getname() string {
	count++
	num := count % 100
	dirname := "temp_test" + strconv.Itoa(num)
	return dirname
}

// Operations about object
type BuildController struct {
	beego.Controller
}

func writeCmdOutput(ws *websocket.Conn,res http.ResponseWriter, pipeReader *io.PipeReader) {

	buffer := make([]byte, 20)
	for {
		n, err := pipeReader.Read(buffer)
		if err != nil {
			pipeReader.Close()
			break
		}

		data := buffer[0:n]
		fmt.Print(string(data))
		res.Write(data)
		//attention : add the Flush method in beego/router.go
		//res.(http.Flusher).Flush()
        ws.WriteMessage(websocket.TextMessage, data)
		
		//reset buffer
		for i := 0; i < n; i++ {
			buffer[i] = 0
		}
	}
}


// @Title testBuild
// @Description : input json file output the stream
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (o *BuildController) Post() {
	   fmt.Println("test post")
	   o.TplNames="test.html"
	}

// @Title testBuild
// @Description : input json file output the stream
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (o *BuildController) Get() {

	ws, err := websocket.Upgrade(o.Ctx.ResponseWriter, o.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(o.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	
	dirname := Getname()
	err = os.Mkdir(dirname, 0777)
	if err != nil {
		panic(err)
	}
    //the file is puted in the root dir
	

	for {

		_, body, err := ws.ReadMessage()
		if err != nil {
			return
		}
	println(string(body))
    f, err := os.Create(dirname + "/" + "Dockerfile")
	f.Write([]byte(body))
	f.Close()
	Dirtotar(dirname, dirname)	
	//send the seployments.tar.gz file to the docker deamon
	docker_response := Tartoimage(dirname, dirname+"/"+"deployments.tar.gz")
	
		//fmt.Println(docker_response.StatusCode)
	res := o.Ctx.ResponseWriter
	//write contents into the pipeWriter and read the contends from the pipReader
	pipeReader, pipeWriter := io.Pipe()
	read := bufio.NewReader(docker_response.Body)

	go func() {
		for {
			line, _, err := read.ReadLine()
			if err == io.EOF {
				break
			}
			// the byte number of \n is 10
			// the type of line is a slice
			line = append(line, 10)
			fmt.Print(string(line))
			pipeWriter.Write(line)
		}
		//io.Copy(pipeWriter, read)
		defer pipeWriter.Close()
	}()

	fmt.Println("the output type:", reflect.TypeOf(o.Ctx.Output))
	//pipeReader = res
	writeCmdOutput(ws,res, pipeReader)

   	defer Cleandir(dirname)
	
	
    }
	
}
