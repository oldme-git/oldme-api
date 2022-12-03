package ufile

import (
	"encoding/hex"
	"io"
	"strings"
	"sync"
)

var tMap sync.Map

// 注册文件类型
func init() {
	tMap.Store("ffd8ffe000104a464946", "jpg")  // JPEG (jpg)
	tMap.Store("89504e470d0a1a0a0000", "png")  // PNG (png)
	tMap.Store("47494638396126026f01", "gif")  // GIF (gif)
	tMap.Store("49492a00227105008037", "tif")  // TIFF (tif)
	tMap.Store("424d228c010000000000", "bmp")  // 16 color bitmap (bmp)
	tMap.Store("424d8240090000000000", "bmp")  // 24 color bitmap (bmp)
	tMap.Store("424d8e1b030000000000", "bmp")  // 256 color bitmap (bmp)
	tMap.Store("41433130313500000000", "dwg")  // CAD (dwg)
	tMap.Store("3c21444f435459504520", "html") // HTML (html)
	tMap.Store("3c68746d6c3e0", "html")        // HTML (html)
	tMap.Store("3c21646f637479706520", "htm")  // HTM (htm)
	tMap.Store("48544d4c207b0d0a0942", "css")  // css
	tMap.Store("696b2e71623d696b2e71", "js")   // js
	tMap.Store("7b5c727466315c616e73", "rtf")  // Rich Text Format (rtf)
	tMap.Store("38425053000100000000", "psd")  // Photoshop (psd)
	tMap.Store("46726f6d3a203d3f6762", "eml")  // Email [Outlook Express 6] (eml)
	tMap.Store("d0cf11e0a1b11ae10000", "doc")  // MS Excel Note: Word, msi and excel file headers are the same
	tMap.Store("d0cf11e0a1b11ae10000", "vsd")  // Visio
	tMap.Store("5374616E64617264204A", "mdb")  // MS Access (mdb)
	tMap.Store("252150532D41646F6265", "ps")
	tMap.Store("255044462d312e350d0a", "pdf")  // Adobe Acrobat (pdf)
	tMap.Store("2e524d46000000120001", "rmvb") // rmvb/rm
	tMap.Store("464c5601050000000900", "flv")  // flv/f4v
	tMap.Store("00000020667479706d70", "mp4")
	tMap.Store("49443303000000002176", "mp3")
	tMap.Store("000001ba210001000180", "mpg") //
	tMap.Store("3026b2758e66cf11a6d9", "wmv") // wmv/asf
	tMap.Store("52494646e27807005741", "wav") // Wave (wav)
	tMap.Store("52494646d07d60074156", "avi")
	tMap.Store("4d546864000000060001", "mid") // MIDI (mid)
	tMap.Store("504b0304140000000800", "zip")
	tMap.Store("526172211a0700cf9073", "rar")
	tMap.Store("235468697320636f6e66", "ini")
	tMap.Store("504b03040a0000000000", "jar")
	tMap.Store("4d5a9000030000000400", "exe")        // executable
	tMap.Store("3c25402070616765206c", "jsp")        // jsp
	tMap.Store("4d616e69666573742d56", "mf")         // MF
	tMap.Store("3c3f786d6c2076657273", "xml")        // xml
	tMap.Store("494e5345525420494e54", "sql")        // xml
	tMap.Store("7061636b616765207765", "java")       // java
	tMap.Store("406563686f206f66660d", "bat")        // bat
	tMap.Store("1f8b0800000000000000", "gz")         // gz
	tMap.Store("6c6f67346a2e726f6f74", "properties") // bat
	tMap.Store("cafebabe0000002e0041", "class")      // bat
	tMap.Store("49545346030000006000", "chm")        // bat
	tMap.Store("04000000010000001300", "mxp")        // bat
	tMap.Store("504b0304140006000800", "docx")       // docx
	tMap.Store("d0cf11e0a1b11ae10000", "wps")        // WPS
	tMap.Store("6431303a637265617465", "torrent")
	tMap.Store("6D6F6F76", "mov")         // Quicktime (mov)
	tMap.Store("FF575043", "wpd")         // WordPerfect (wpd)
	tMap.Store("CFAD12FEC5FD746F", "dbx") // Outlook Express (dbx)
	tMap.Store("2142444E", "pst")         // Outlook (pst)
	tMap.Store("AC9EBD8F", "qdf")         // Quicken (qdf)
	tMap.Store("E3828596", "pwl")         // Windows Password (pwl)
	tMap.Store("2E7261FD", "ram")         // Real Audio (ram)
}

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

// Ext 根据二进制流获取文件类型
func Ext(b []byte) string {
	var t string
	enCodeStr := hex.EncodeToString(b[:10])
	tMap.Range(func(key, value any) bool {
		k := key.(string)
		if strings.HasPrefix(enCodeStr, strings.ToLower(k)) || strings.HasPrefix(k, strings.ToLower(enCodeStr)) {
			t = value.(string)
			return false
		}
		return true
	})
	return t
}

// IsExt 根据二进制流判断文件类型是否是所需要的
func IsExt(b []byte, t string) bool {
	typeSlice := strings.Split(t, ",")
	// 获取文件类型
	fileType := Ext(b)
	for _, item := range typeSlice {
		if strings.ToLower(fileType) == strings.ToLower(item) {
			return true
		}
	}
	return false
}
