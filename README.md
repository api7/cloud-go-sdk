# cloud-go-sdk

The Go SDK for accessing [API7 Cloud API](https://api7.cloud/api)

# Installation

Make sure you're using Go Modules to manage your Go Project, and your Go version needs to be `1.19` or later.

```shell
go mod init
```

Then, import the cloud-go-sdk in your Go program by running:

```go
import (
	"github.com/api7/cloud-go-sdk"
)
```

Alternatively, download the cloud-go-sdk by running:

```shell
go get -u github.com/api7/cloud-go-sdk
```

# Mock

You may want to add unit test cases after you import cloud-go-sdk into your Go program,
for the sake of better integration, cloud-go-sdk uses [mockgen](https://github.com/golang/mock) to generate mock implementations.

e.g., for the `Interface`, we have a `MockInterface` implementation. If you want to mock an Application create error, you can write some codes like:

```go
ctrl := gomock.NewController(t) // var t *testing.T
cloud := NewMockInterface(ctrl)
cloud.EXPECT().CreateApplication(gomock.Any(), &Application{
    ApplicationSpec: ApplicationSpec{
        Name:        "test app",
		Description: "This is a test app",
           Protocols:   []string{ProtocolHTTP},
           PathPrefix:  "/api/v1",
           Hosts:       []string{"app.test.com"},
           Upstreams: []UpstreamAndVersion{
			{
				Upstream: Upstream{
					Scheme: "https",
					LBType: "roundrobin",
					Targets: []UpstreamTarget{
						{
							Host:   "10.0.5.1", 
							Port:   80,
							Weight: 100,
						},
						{
							Host:   "10.0.5.2",
							Port:   80,
							Weight: 100,
						},
					},
				}, 
				Version: "default",
			},
		   }, 
		   DefaultUpstreamVersion: "default", 
		   Active:                 ActiveStatus,
	},
}, &ApplicationCreateOptions{
	Cluster: &Cluster{
		ID: 1,
	},
}).Return(nil, errors.New("mock error"))
```

## Development

Run test cases via running:

```shell
make test
```

Regenerate mock codes by running:

```shell
make mockgen
```
