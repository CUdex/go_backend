package aws_api

type Instance struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Tag  []Tag `json:"tag"`
}

type Tag struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

