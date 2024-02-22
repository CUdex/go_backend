package aws_api

type Instance struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Tag  []Tag `json:"tag"`
}

type TagRequest struct {
    InstanceID string `json:"instanceId"`
    RequestTag       Tag `json:"tag"`
}


type Tag struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

