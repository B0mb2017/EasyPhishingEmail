package Plugin

import (
	"fmt"
	"strings"
)

type Bar struct {
	percent  int    //百分比
	cur      int    //当前进度位置
	total    int    //总进度
	rate     string //进度条
	graph    string //显示符号
	accuracy int    //一个 graph 代表 accuracy 个 percent，最小为1(默认值)，最大100
}

//  NewBar
//  @Description: 创建进度条
//  @receiver b
//  @param start 当前值
//  @param total 总进度值
//
func (b *Bar) NewBar(start, total int) {
	b.cur = start
	b.total = total
	if b.graph == "" {
		b.graph = "█"
	}
	if b.accuracy == 0 {
		b.accuracy = 1
	}
	b.percent = b.getPercent()
	for i := 0; i < int(b.percent); i += b.accuracy {
		b.rate += b.graph
	}
}

//  NewBarWithCustomGraph
//  @Description: 自定义进度条图案
//  @receiver b
//  @param start 当前值
//  @param total 总进度值
//  @param graph 使用graph作为输出字符
//
func (b *Bar) NewBarWithCustomGraph(start, total int, graph string) {
	b.graph = graph
	b.NewBar(start, total)
}

//  NewBarWithCustomAccuracy
//  @Description: 自定义进度条精度
//  @receiver b
//  @param start 当前值
//  @param total 总计值
//  @param acc   一个 graph 代表 accuracy 个 percent，最小为1(默认值)，最大100
//
func (b *Bar) NewBarWithCustomAccuracy(start, total, acc int) {
	b.accuracy = acc % 11
	b.NewBar(start, total)
}

//  NewBarWithCustomAccuracy
//  @Description: 自定义进度条精度
//  @receiver b
//  @param start 当前值
//  @param total 总计值
//  @param acc   一个 graph 代表 accuracy 个 percent，最小为1(默认值)，最大100
//  @param graph 自定义图案
//
func (b *Bar) NewBarWithCustomGraphAndAccuracy(start, total, acc int, graph string) {
	b.graph = graph
	b.NewBarWithCustomAccuracy(start, total, acc)
}

//  getPercent
//  @Description: 获取进度条百分比
//  @receiver b
//  @return int
func (b *Bar) getPercent() int {
	return int(float32(b.cur) / float32(b.total) * 100)
}

//  Play
//  @Description: 设置播放进度
//  @receiver b
//  @param cur
//
func (b *Bar) Play(cur int) {
	b.cur = cur
	lastPercent := b.percent
	b.percent = b.getPercent()
	last := (lastPercent + b.accuracy - 1) / b.accuracy
	current := (b.percent + b.accuracy - 1) / b.accuracy
	if current-last > 0 {
		b.rate += strings.Repeat(b.graph, current-last)
		fmt.Printf("\r%4d/%d %3d%% [%-1s]", b.cur, b.total, b.percent, b.rate)
	}
}

//
//  Finish
//  @Description: 结束进度条
//  @receiver b
//
func (b *Bar) Finish() {
	b.graph = ""
	b.accuracy = 0
	b.rate = ""
	fmt.Println()
}
