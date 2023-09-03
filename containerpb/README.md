# containerpb

https://github.com/googleapis/google-cloud-go/issues/8237

`Parent: us-central1`

```
2023/09/02 19:51:34 rpc error: code = PermissionDenied desc = Permission denied on resource project us-central1.
error details: name = ErrorInfo reason = CONSUMER_INVALID domain = googleapis.com metadata = map[consumer:projects/us-central1 service:edgecontainer.googleapis.com]
error details: name = Help desc = Google developer console API key url = https://console.developers.google.com/project/us-central1/apiui/credential
exit status 1
```

`projects/{project}/locations/{location}`

```
2023/09/02 19:53:46 rpc error: code = PermissionDenied desc = Distributed Cloud Edge Container API has not been used in project theviewfromjq before or it is disabled. Enable it by visiting https://console.developers.google.com/apis/api/edgecontainer.googleapis.com/overview?project=theviewfromjq then retry. If you enabled this API recently, wait a few minutes for the action to propagate to our systems and retry.
error details: name = ErrorInfo reason = SERVICE_DISABLED domain = googleapis.com metadata = map[consumer:projects/theviewfromjq service:edgecontainer.googleapis.com]
error details: name = Help desc = Google developers console API activation url = https://console.developers.google.com/apis/api/edgecontainer.googleapis.com/overview?project=theviewfromjq
exit status 1
```


- How is this mismatch possible:
https://github.com/googleapis/google-cloud-go/issues/8237#issuecomment-1629340566

https://pkg.go.dev/cloud.google.com/go/edgecontainer/apiv1/edgecontainerpb#ListClustersRequest
