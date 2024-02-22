package aws_api

import (
	"fmt"
	"context"
	"net/http"
	"encoding/json"
	"tag-controller/logger"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func TagAddHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Only PUT method is allowed", http.StatusMethodNotAllowed)
        return
    }

	// request body 값이 정의한 구조체와 다를 경우 Bad Request 반환
	var req TagRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Error decoding request body", http.StatusBadRequest)
        return
    }

	// AUTO_STOP, AUTO_TERMINATE key가 아닌 경우 잘못된 요청으로 처리
	if req.RequestTag.Key == "NO_AUTO_STOP" || req.RequestTag.Key == "NO_AUTO_TERMINATE" {
		
		if err := updateEC2Tags(req); err != nil {
			http.Error(w, fmt.Sprintf("Error updating tags: %s", err), http.StatusInternalServerError)
			return
		}
	
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Tags updated successfully")
		return
	} 
	
	http.Error(w, "Error Bad Request Tag Key", http.StatusBadRequest)
	return   
}

func updateEC2Tags(tagInfo TagRequest) error {
	// AWS 설정 로드
    cfg, err := config.LoadDefaultConfig(context.TODO())

    if err != nil {
		logger.Error("error loading AWS configuration")
        return fmt.Errorf("error loading AWS configuration: %w", err)
    }

    svc := ec2.NewFromConfig(cfg)

    // 태그 형식 변환
    var ec2Tags []types.Tag
	ec2Tags = append(ec2Tags, types.Tag{Key: &tagInfo.RequestTag.Key, Value: &tagInfo.RequestTag.Value})


    // EC2 인스턴스 태그 변경 요청
	logger.Info(fmt.Sprintf("Request updating tags: ID(%s), Key(%s)", tagInfo.InstanceID, tagInfo.RequestTag.Key))
    _, err = svc.CreateTags(context.TODO(), &ec2.CreateTagsInput{
        Resources: []string{tagInfo.InstanceID},
        Tags:      ec2Tags,
    })

    if err != nil {
        return fmt.Errorf("error creating tags: %w", err)
    }

    return nil
}