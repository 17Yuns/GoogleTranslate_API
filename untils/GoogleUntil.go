package untils

import "fmt"

//翻译链接

func ConstructUrl(url string, tk string, message string, be string, to string) string {
	return fmt.Sprintf(`%s/translate_a/single?client=t&sl=%s&tl=%s&hl=%s&dt=bd&dt=ex&dt=ld&dt=md&dt=qc&dt=rw&dt=rm&dt=ss&dt=t&dt=at&ie=UTF-8&oe=UTF-8&source=sel&tk=%s&q=%s`, url, be, to, to, tk, message)
}

//通过下面这个链接构造之后直接请求就是一个音频，直接保存为MP3就行，我给出样例代码
//response, err := http.Get(url)
//if err != nil {
//	fmt.Println("Error making GET request:", err)
//	return
//}
//defer response.Body.Close() // 确保在函数结束时关闭响应体
//
//// 检查响应状态
//if response.StatusCode != http.StatusOK {
//	fmt.Println("Error: received non-200 response status:", response.Status)
//	return
//}
//
//// 创建一个文件来保存 MP3
//outputFile, err := os.Create("output.mp3")
//if err != nil {
//	fmt.Println("Error creating file:", err)
//	return
//}
//defer outputFile.Close() // 确保在函数结束时关闭文件
//
//// 将响应体写入文件
//_, err = io.Copy(outputFile, response.Body)
//if err != nil {
//	fmt.Println("Error saving MP3 to file:", err)
//	return
//}
//
//fmt.Println("MP3 file saved successfully as output.mp3")

func ConstructSpeakUrl(url string, tk string, message string, to string) string {
	return fmt.Sprintf(`%s/translate_tts?ie=UTF-8&client=t&prev=input&q=%s&tl=%s&total=1&idx=0&textlen=4&tk=%s`, url, message, to, tk)
}
