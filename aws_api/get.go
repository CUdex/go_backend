package aws_api

import (
	"context"
	"net/http"
	"encoding/json"
	"tag-controller/logger"
    "tag-controller/prom"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/ec2"
)

func TagGetHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
        return
    }
    prom.GetRequestCounter.WithLabelValues(r.URL.Path).Inc()

	// instance list와 tag 조회
	instances, err := getInstances()
	
	if err != nil {
		logger.Error("failed Get Instance Info")
		http.Error(w, "failed Get Instance Info", http.StatusMethodNotAllowed)
        return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(instances)
}

func getInstances() ([]Instance, error) {
    // AWS 설정 로드
    cfg, err := config.LoadDefaultConfig(context.TODO())

    if err != nil {
        return nil, http.ErrBodyNotAllowed
    }

    // EC2 클라이언트 생성
    svc := ec2.NewFromConfig(cfg)

    // EC2 인스턴스 조회
    result, err := svc.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	var instances []Instance
    for _, reservation := range result.Reservations {
        for _, inst := range reservation.Instances {
            var insName string
            var tag []Tag

            for _, t := range inst.Tags {
                // 인스턴스 NO_AUTO 관련 태그 저장
				if *t.Key == "NO_AUTO_STOP" || *t.Key == "NO_AUTO_TERMINATE" {
					tag = append(tag, Tag{Key: *t.Key, Value: *t.Value}) 
				}
                // 인스턴스 name tag 저장
				if *t.Key == "Name" {
					insName = *t.Value 
				}
            }
			typeInfo := string(inst.InstanceType)
            instances = append(instances, Instance{
                ID:  *inst.InstanceId,
				Type: typeInfo,
                Name: insName,
                Tag: tag,
            })
        }
    }
    return instances, err
}