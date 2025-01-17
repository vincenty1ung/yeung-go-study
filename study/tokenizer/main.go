package main

import (
	"fmt"
	"regexp"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter

	new, _ = gse.New("zh,testdata/test_en_dict3.txt", "alpha")

	text = "你好世界, Hello world, Helloworld."
)

func main() {
	// 加载默认词典
	seg.LoadDict()
	// 加载默认 embed 词典
	// seg.LoadDictEmbed()
	//
	// 加载简体中文词典
	// seg.LoadDict("zh_s")
	// seg.LoadDictEmbed("zh_s")
	//
	// 加载繁体中文词典
	// seg.LoadDict("zh_t")
	//
	// 加载日文词典
	// seg.LoadDict("jp")
	//
	// 载入词典
	// seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	// cut()

	segCut()
}

func cut() {
	hmm := new.Cut(text, true)
	fmt.Println("cut use hmm: ", hmm)

	hmm = new.CutSearch(text, true)
	fmt.Println("cut search use hmm: ", hmm)
	fmt.Println("analyze: ", new.Analyze(hmm, text))

	hmm = new.CutAll(text)
	fmt.Println("cut all: ", hmm)

	reg := regexp.MustCompile(`(\d+年|\d+月|\d+日|[\p{Latin}]+|[\p{Hangul}]+|\d+\.\d+|[a-zA-Z0-9]+)`)
	text1 := `헬로월드 헬로 서울, 2021年09月10日, 3.14`
	hmm = seg.CutDAG(text1, reg)
	fmt.Println("Cut with hmm and regexp: ", hmm, hmm[0], hmm[6])
}

func analyzeAndTrim(cut []string) {
	a := seg.Analyze(cut, "")
	fmt.Println("analyze the segment: ", a)

	cut = seg.Trim(cut)
	fmt.Println("cut all: ", cut)

	fmt.Println(seg.String(text, true))
	fmt.Println(seg.Slice(text, true))
}

func cutPos() {
	po := seg.Pos(text, true)
	fmt.Println("pos: ", po)
	po = seg.TrimPos(po)
	fmt.Println("trim pos: ", po)

	posSeg.WithGse(seg)
	po = posSeg.Cut(text, true)
	fmt.Println("pos: ", po)

	po = posSeg.TrimWithPos(po, "zg")
	fmt.Println("trim pos: ", po)
}

func segCut() {
	// 分词文本
	tb := []byte("请各位高人指点一下，我用Macbook pro直接USB输出到高登GDX2解码耳放一体机。\nMAC OS下不能用foobar. 无奈之下用的audirvana。感觉Audirvana播放音质很好，管理专辑也方便。但有一个小问题：也许是我的本子比较老了算力不足？就是加载高码率文件时，尤其是ISO时，加载的比较慢，播放DSD 256时偶尔还会有卡顿。\n\n用windows本子（Thinkpad X1）在foobar下就不会有任何卡顿问题。但是，又主观上感觉foobar的声音较Audirvana有些“软绵绵”的，不够坚定有力量的感觉，而且foobar管理文件不如Audirvana方便。\n\n\n大家指点一下Mac OS系统下还有什么好的播放软件？或者Windows系统下foobar用什么版本可能更好？\n\n谢谢各位指点。")

	// 处理分词结果
	// fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(string(tb), true))
	// fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(string(tb)))

	segments := seg.Segment(tb)
	// 处理分词结果, 普通模式
	fmt.Println(gse.ToString(segments))
	fmt.Println()

	// segments1 := seg.Segment([]byte(text))
	// 搜索模式
	// fmt.Println(gse.ToString(segments1, true))
}
