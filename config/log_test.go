package config

import (
	"fmt"
	"go-server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)


var (
	testLog = `
　　孙子曰：凡用兵之法，驰车千驷，革车千乘，带甲十万，千里馈粮，内外之费，宾客之用，胶漆之材，车甲之奉，日费千金，然后十万之师举矣。

　　其用战也胜，久则钝兵挫锐，攻城则力屈，久暴师则国用不足。夫钝兵挫锐，屈力殚货，则诸侯乘其弊而起，虽有智者，不能善其后矣。故兵闻拙速，未睹巧之久也。夫兵久而国利者，未之有也。故不尽知用兵之害者，则不能尽知用兵之利也。

　　善用兵者，役不再籍，粮不三载，取用于国，因粮于敌，故军食可足也。国之贫于师者远输，远输则百姓贫；近师者贵卖，贵卖则百姓财竭，财竭则急于丘役。力屈财殚，中原内虚于家，百姓之费，十去其七；公家之费，破军罢马，甲胄矢弩，戟楯蔽橹，丘牛大车，十去其六。

　　故智将务食于敌，食敌一钟，当吾二十钟；萁秆一石，当吾二十石。

　　故杀敌者，怒也；取敌之利者，货也。车战得车十乘以上，赏其先得者，而更其旌旗，车杂而乘之，卒善而养之，是谓胜敌而益强。

　　故兵贵胜，不贵久。故知兵之将，生民之司命，国家安危之主也。
`
)

func TestLog(t *testing.T) {
	write := getLogWriter("/data/logs/test.log")
	core := zapcore.NewTee(zapcore.NewCore(getEncoder(), zapcore.AddSync(write), zap.InfoLevel))

	// 构建日志
	logger := zap.New(core, zap.AddCaller())
	logger.Info("log 初始化成功")
	global.Logger = logger
	for i := 0; i < 10000; i++ {
		global.Logger.Info(testLog+fmt.Sprintf(",i=%d",i))
	}

}


