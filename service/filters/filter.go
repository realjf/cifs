package filters

import (
	pb "asifs/service/proto/filter"
	"asifs/service/utils"
	"context"
	"log"
)

var (
	StopWordTable *StopWord
	SensitiveTable *SensitiveWord
)

func init() {
	// 加载停用词
	StopWordTable = NewStopWord()
	err := StopWordTable.LoadTable("./filters/stopword_table")
	if err != nil {
		log.Printf("failed to load stop word table: %v", err)
	}

	// 加载敏感词
	SensitiveTable = NewSensitiveWord()
	err = SensitiveTable.LoadTable("./filters/sensitiveword_table")
	if err != nil {
		log.Printf("failed to load sensitive word table: %v", err)
	}
}

type Filter struct {

}

func NewFilter() *Filter {
	return &Filter{}
}

func (fs *Filter) StopWord(ctx context.Context, request *pb.Request) (response *pb.Response, err error) {
	if request.Content == "" {
		return &pb.Response{
			Code: 101,
			Message: "待过滤内容为空",
			Result: &pb.Result{
				OriginContent: request.Content,
				FilteredContent: "",
			},
		}, nil
	}

	// 全角转换半角
	var content string = utils.D2SConvertString(request.Content)

	filteredContent := StopWordTable.Filter(content)

	return &pb.Response{
		Code: 100,
		Message: "停用词过滤成功",
		Result: &pb.Result{
			OriginContent: request.Content,
			FilteredContent: filteredContent,
		},
	}, nil
}


func (fs *Filter) SensitiveWord(ctx context.Context, request *pb.Request) (response *pb.Response, err error) {
	if request.Content == "" {
		return &pb.Response{
			Code: 101,
			Message: "待过滤内容为空",
			Result: &pb.Result{
				OriginContent: request.Content,
				FilteredContent: "",
			},
		}, nil
	}
	// 全角转换半角
	var content string = utils.D2SConvertString(request.Content)

	filteredContent := SensitiveTable.Filter(content)

	return &pb.Response{
		Code: 100,
		Message: "敏感词过滤成功",
		Result: &pb.Result{
			OriginContent: request.Content,
			FilteredContent: filteredContent,
		},
	}, nil
}
